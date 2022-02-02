package awsprober

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
)

type STSConfig struct {
	Logger       micrologger.Logger
	Region       string
	ExpectedRole string
}

type STS struct {
	logger       micrologger.Logger
	region       string
	expectedRole string
}

func NewSTS(config STSConfig) (*STS, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}
	if config.ExpectedRole == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.ExpectedRole must not be empty", config)
	}

	return &STS{
		logger:       config.Logger,
		region:       config.Region,
		expectedRole: config.ExpectedRole,
	}, nil
}

func (r *STS) Probe(ctx context.Context) bool {
	sess, err := session.NewSessionWithOptions(session.Options{})
	if err != nil {
		r.logger.Errorf(ctx, err, "Error during AWS session setup")
		return false
	}

	client := sts.New(sess)

	identity, err := client.GetCallerIdentity(&sts.GetCallerIdentityInput{})
	if err != nil {
		r.logger.Errorf(ctx, err, "Error calling sts.GetCallerIdentity")
		return false
	}

	if identity.UserId == nil {
		r.logger.Errorf(ctx, err, "sts.GetCallerIdentity returned nil userId")
		return false
	}

	if *identity.UserId != r.expectedRole {
		r.logger.Errorf(ctx, err, "Expected to have assumed role %q, but sts.GetCallerIdentity gave us %q", r.expectedRole, *identity.UserId)
		return false
	}

	return true
}
