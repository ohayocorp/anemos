package apply

import (
	"time"
)

type Options struct {
	DocumentGroups   []string
	SkipConfirmation bool
	ForceConflicts   bool
	Timeout          time.Duration
}

func NewOptions() *Options {
	return &Options{}
}
