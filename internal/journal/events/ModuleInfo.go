package events

import (
	"time"
)

// ModuleInfo event structure
type evModuleInfo struct {
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// ModuleInfo event handler
func ModuleInfo(e interface{}) {
    // ev := e.(evModuleInfo)
}

