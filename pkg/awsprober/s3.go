package awsprober

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
)

type S3Config struct {
	Logger micrologger.Logger
}

type S3 struct {
	logger micrologger.Logger
}

func NewS3(config S3Config) (*S3, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	return &S3{
		logger: config.Logger,
	}, nil
}

func (r *S3) Probe(ctx context.Context) bool {
	sess, err := session.NewSessionWithOptions(session.Options{})
	if err != nil {
		r.logger.Errorf(ctx, err, "Error during AWS session setup")
		return false
	}

	client := s3.New(sess)

	_, err = client.ListBuckets(nil)
	if err != nil {
		r.logger.Errorf(ctx, err, "Error listing s3 buckets")
		return false
	}

	return true
}
