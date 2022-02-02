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
	probeModeFlag     = "probe-mode"
	roleNameFlag      = "role-name"

	kiamNamespaceFlag     = "kiam-namespace"
	kiamLabelSelectorFlag = "kiam-label-selector"
	nodeNameFlag          = "node-name"

	probeModeRoute53 = "route53"
	probeModeSTS     = "sts"
)

type flag struct {
	Region        string
	FailThreshold int
	Interval      int
	ProbeMode     string
	RoleName      string

	KiamNamespace     string
	KiamLabelSelector string
	NodeName          string
}

func (f *flag) Init(cmd *cobra.Command) {
	// Add command line flags for glog.
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)

	cmd.Flags().StringVar(&f.Region, regionFlag, "", `The AWS region to use for the tests.`)
	cmd.Flags().IntVar(&f.FailThreshold, failThresholdFlag, 5, `How many failed probes in a row to consider kiam unhealthy. Defaults to 5.`)
	cmd.Flags().IntVar(&f.Interval, intervalFlag, 60, `Interval in seconds to wait between tests. Defaults to 60.`)
	cmd.Flags().StringVar(&f.ProbeMode, probeModeFlag, "sts", `What AWS API to use to check if kiam is working. Either "route53" or "sts". Defaults to "sts"`)
	cmd.Flags().StringVar(&f.RoleName, roleNameFlag, "", `Name of the IAM role that kiam should make this pod assume. Used by the 'sts' probe mode.`)

	cmd.Flags().StringVar(&f.KiamNamespace, kiamNamespaceFlag, "kube-system", `The namespace where kiam agent pods are running. Defaults to 'kube-system'`)
	cmd.Flags().StringVar(&f.KiamLabelSelector, kiamLabelSelectorFlag, "component=kiam-agent", `The label selector to select kiam pods. Defaults to 'component=kiam-agent'`)
	cmd.Flags().StringVar(&f.NodeName, nodeNameFlag, "", `The node where the watchdog is running (to only terminate current node's kiam agent pod).`)
}

func (f *flag) Validate(cmd *cobra.Command) error {
	var err error

	if f.Region == "" {
		return fmt.Errorf("--%s can't be empty", regionFlag)
	}

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

	if f.ProbeMode != probeModeSTS && f.ProbeMode != probeModeRoute53 {
		return fmt.Errorf("--%s should be either %q or %q", probeModeFlag, probeModeSTS, probeModeRoute53)
	}

	if f.ProbeMode == probeModeSTS && f.RoleName == "" {
		return fmt.Errorf("--%s can't be empty when --%s is %q", roleNameFlag, probeModeFlag, probeModeSTS)
	}

	// TODO validate AWS region.

	return err
}
