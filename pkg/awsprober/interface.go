package awsprober

import "context"

type Interface interface {
	Probe(ctx context.Context) bool
}
