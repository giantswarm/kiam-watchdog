package awsprober

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
)

type STSConfig struct {
	Logger micrologger.Logger
	Region string
}

type STS struct {
	logger micrologger.Logger
	region string
}

func NewSTS(config STSConfig) (*STS, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	return &STS{
		logger: config.Logger,
		region: config.Region,
	}, nil
}

func (r *STS) Probe(ctx context.Context) bool {
	sess, err := session.NewSessionWithOptions(session.Options{})
	if err != nil {
		r.logger.Errorf(ctx, err, "Error during AWS session setup")
		return false
	}

	client := sts.New(sess)

	_, err = client.GetCallerIdentity(nil)
	if err != nil {
		r.logger.Errorf(ctx, err, "Error calling sts.GetCallerIdentity")
		return false
	}

	return true
}
