package httpsvr

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

type Server interface {
	ListenAndServe() error
	Shutdown(ctx context.Context) error
}

type ServerPlur struct {
	servers []Server
	eg      *errgroup.Group
}

func NewServerPlur() *ServerPlur {
	group, _ := errgroup.WithContext(context.Background())
	return &ServerPlur{
		eg: group,
	}
}

func (s *ServerPlur) AddServer(server Server) {
	s.servers = append(s.servers, server)
}

func (s *ServerPlur) StartAll() error {
	for _, server := range s.servers {
		s.eg.Go(server.ListenAndServe)
	}

	return s.eg.Wait()
}

func (s *ServerPlur) ShutdownAll(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	for _, server := range s.servers {
		s.eg.Go(func() error {
			return server.Shutdown(ctx)
		})
	}
	return s.eg.Wait()
}

func (s *ServerPlur) RunOrDie(sig ...os.Signal) error {
	if err := s.StartAll(); err != nil {
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, sig...)
	<-quit

	return s.ShutdownAll(5 * time.Second)
}

type HttpConfig struct {
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	TLS          *TLSConfig
}

func (hc *HttpConfig) check() {
	if hc.Addr == "" {
		hc.Addr = ":8080"
	}
	if hc.ReadTimeout == 0 {
		hc.ReadTimeout = 5 * time.Second
	}
	if hc.WriteTimeout == 0 {
		hc.WriteTimeout = 10 * time.Second
	}
}

type TLSConfig struct {
	Cert string
	Key  string
}

type httpServer struct {
	cfg    *HttpConfig
	server *http.Server
}

// ListenAndServe implements Server.
func (hs *httpServer) ListenAndServe() error {
	errChan := make(chan error, 1)
	defer close(errChan)
	go func(errChan chan error) {
		if hs.cfg.TLS != nil {
			if err := hs.server.ListenAndServeTLS(hs.cfg.TLS.Cert, hs.cfg.TLS.Key); err != nil &&
				err != http.ErrServerClosed {
				errChan <- err
			}
		} else {
			if err := hs.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				errChan <- err
			}
		}
	}(errChan)

	if err := <-errChan; err != nil {
		return err
	}

	return hs.wait()
}

// Shutdown implements Server.
func (hs *httpServer) Shutdown(ctx context.Context) error {
	return hs.server.Shutdown(ctx)
}

func (hs *httpServer) wait() error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(quit)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return hs.server.Shutdown(ctx)
}

func NewHttp(cfg *HttpConfig, handler http.Handler) Server {
	cfg.check()
	return &httpServer{
		cfg: cfg,
		server: &http.Server{
			Addr:         cfg.Addr,
			Handler:      handler,
			ReadTimeout:  cfg.ReadTimeout,
			WriteTimeout: cfg.WriteTimeout,
		},
	}
}
