package awsprober

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/route53"
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
	client := route53.New(route53.Options{
		Region: "eu-west-1",
	})

	_, err := client.ListHostedZones(context.Background(), nil)
	if err != nil {
		r.logger.Errorf(ctx, err, "Error listing route53 hosted zones")
		return false
	}

	return true
}
