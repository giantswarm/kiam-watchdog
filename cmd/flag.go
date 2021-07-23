package cmd

import (
	goflag "flag"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	regionFlag        = "region"
	failThresholdFlag = "fail-threshold"
	intervalFlag      = "interval"
)

type flag struct {
	Region        string
	FailThreshold int
	Interval      int
}

func (f *flag) Init(cmd *cobra.Command) {
	// Add command line flags for glog.
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)

	cmd.Flags().StringVar(&f.Region, regionFlag, "eu-west-1", `The AWS region to use for the tests. Defaults to eu-west-1.`)
	cmd.Flags().IntVar(&f.FailThreshold, failThresholdFlag, 5, `How many failed probes in a row to consider kiam unhealthy. Defaults to 5.`)
	cmd.Flags().IntVar(&f.Interval, intervalFlag, 60, `Interval in seconds to wait between tests. Defaults to 60.`)
}

func (f *flag) Validate(cmd *cobra.Command) error {
	var err error

	if f.Interval <= 0 {
		return fmt.Errorf("--%s should be greater than 0", intervalFlag)
	}

	if f.FailThreshold <= 0 {
		return fmt.Errorf("--%s should be greater than 0", intervalFlag)
	}

	// TODO validate AWS region.

	return err
}
