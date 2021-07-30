package awsprober

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
)

type Route53Config struct {
	Logger micrologger.Logger
	Region string
}

type Route53 struct {
	logger micrologger.Logger
	region string
}

func NewRoute53(config Route53Config) (*Route53, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	return &Route53{
		logger: config.Logger,
		region: config.Region,
	}, nil
}

func (r *Route53) Probe(ctx context.Context) bool {
	sess, err := session.NewSessionWithOptions(session.Options{})
	if err != nil {
		r.logger.Errorf(ctx, err, "Error during AWS session setup")
		return false
	}

	client := route53.New(sess)

	_, err = client.ListHostedZones(nil)
	if err != nil {
		r.logger.Errorf(ctx, err, "Error listing route53 hosted zones")
		return false
	}

	return true
}
