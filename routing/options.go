package routing

import commonConfig "common/config"

// Options contains options for the route handlers.
type Options struct {
	config *commonConfig.Config
}

// Option is a function to set options.
type Option func(*Options)

// Config sets the application configuration.
func Config(config *commonConfig.Config) Option {
	return func(options *Options) {
		options.config = config
	}
}

func newOptions(opts ...Option) *Options {
	config, _ := commonConfig.GetConfig()
	options := &Options{
		config: config,
	}
	for _, opt := range opts {
		opt(options)
	}
	return options
}
