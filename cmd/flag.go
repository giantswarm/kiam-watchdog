package cmd

import (
	goflag "flag"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	failThresholdFlag = "fail-threshold"
	intervalFlag      = "interval"

	kiamNamespaceFlag     = "kiam-namespace"
	kiamLabelSelectorFlag = "kiam-label-selector"
	nodeNameFlag          = "node-name"
)

type flag struct {
	FailThreshold int
	Interval      int

	KiamNamespace     string
	KiamLabelSelector string
	NodeName          string
}

func (f *flag) Init(cmd *cobra.Command) {
	// Add command line flags for glog.
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)

	cmd.Flags().IntVar(&f.FailThreshold, failThresholdFlag, 5, `How many failed probes in a row to consider kiam unhealthy. Defaults to 5.`)
	cmd.Flags().IntVar(&f.Interval, intervalFlag, 60, `Interval in seconds to wait between tests. Defaults to 60.`)

	cmd.Flags().StringVar(&f.KiamNamespace, kiamNamespaceFlag, "kube-system", `The namespace where kiam agent pods are running. Defaults to 'kube-system'`)
	cmd.Flags().StringVar(&f.KiamLabelSelector, kiamLabelSelectorFlag, "component=kiam-agent", `The label selector to select kiam pods. Defaults to 'component=kiam-agent'`)
	cmd.Flags().StringVar(&f.NodeName, nodeNameFlag, "", `The node where the watchdog is running (to only terminate current node's kiam agent pod).`)
}

func (f *flag) Validate(cmd *cobra.Command) error {
	var err error

	if f.KiamNamespace == "" {
		return fmt.Errorf("--%s can't be empty", kiamNamespaceFlag)
	}

	if f.KiamLabelSelector == "" {
		return fmt.Errorf("--%s can't be empty", kiamLabelSelectorFlag)
	}

	if f.NodeName == "" {
		return fmt.Errorf("--%s can't be empty", nodeNameFlag)
	}

	if f.Interval <= 0 {
		return fmt.Errorf("--%s should be greater than 0", intervalFlag)
	}

	if f.FailThreshold <= 0 {
		return fmt.Errorf("--%s should be greater than 0", intervalFlag)
	}

	return err
}
