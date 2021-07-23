package kiamagentrestarter

import "github.com/giantswarm/microerror"

var invalidConfigError = &microerror.Error{
	Kind: "invalidConfig",
}

var invalidLabelSelectorError = &microerror.Error{
	Kind: "invalidLabelSelector",
}
