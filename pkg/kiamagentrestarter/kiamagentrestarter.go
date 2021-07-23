package kiamagentrestarter

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
)

type Config struct {
	Logger micrologger.Logger
}

type Restarter struct {
	logger micrologger.Logger
}

func NewRestarter(config Config) (*Restarter, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	return &Restarter{
		logger: config.Logger,
	}, nil
}

func (r *Restarter) RestartKiamAgent() error {
	return nil
}
