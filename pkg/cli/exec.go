package cli

import (
	"context"

	"github.com/spf13/cobra"
)

type Executor interface {
	Execute(ctx context.Context) error
	Cobra() *cobra.Command
}

type Exec struct {
	r *Command
}

func (ex *Exec) Execute(ctx context.Context) error {
	return ex.r.Cobra.ExecuteContext(ctx)
}

func (ex *Exec) Cobra() *cobra.Command {
	return ex.r.Cobra
}