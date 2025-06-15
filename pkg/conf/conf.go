package conf

import (
	"sync"

	"github.com/spf13/viper"
)

// VConf is a configurator base on Viper.
type VConf struct {
	v  *viper.Viper
	rw sync.RWMutex
}
