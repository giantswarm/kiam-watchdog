package cmd

import (
	"context"
	"io"
	"time"

	"github.com/giantswarm/k8sclient/v4/pkg/k8sclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/giantswarm/kiam-watchdog/pkg/awsprober"
	"github.com/giantswarm/kiam-watchdog/pkg/kiamagentrestarter"
)

type runner struct {
	flag   *flag
	logger micrologger.Logger
	stdout io.Writer
	stderr io.Writer
}

func (r *runner) Run(cmd *cobra.Command, args []string) error {
	ctx := context.Background()

	err := r.flag.Validate(cmd)
	if err != nil {
		return err
	}

	err = r.run(ctx, cmd, args)
	if err != nil {
		return err
	}

	return nil
}

func (r *runner) run(ctx context.Context, cmd *cobra.Command, args []string) error {
	var err error
	errors := 0

	r.logger.Debugf(ctx, "Kiam namespace: %q", r.flag.KiamNamespace)
	r.logger.Debugf(ctx, "Kiam label selector: %q", r.flag.KiamLabelSelector)
	r.logger.Debugf(ctx, "Node name: %q", r.flag.NodeName)
	r.logger.Debugf(ctx, "Interval: %d", r.flag.Interval)
	r.logger.Debugf(ctx, "Fail Threshould: %d", r.flag.FailThreshold)
	r.logger.Debugf(ctx, "AWS region: %q", r.flag.Region)

	var prober awsprober.Interface
	{
		prober, err = awsprober.NewRoute53(awsprober.Route53Config{
			Logger: r.logger,
			Region: r.flag.Region,
		})
		if err != nil {
			return microerror.Mask(err)
		}
	}

	var k8sClient kubernetes.Interface
	{
		config, err := rest.InClusterConfig()
		if err != nil {
			return microerror.Mask(err)
		}

		c := k8sclient.ClientsConfig{
			Logger:        r.logger,
			SchemeBuilder: k8sclient.SchemeBuilder{},
			RestConfig:    config,
		}

		clients, err := k8sclient.NewClients(c)
		if err != nil {
			return microerror.Mask(err)
		}

		k8sClient = clients.K8sClient()
	}

	restarter, err := kiamagentrestarter.NewRestarter(kiamagentrestarter.Config{
		Logger:        r.logger,
		K8sClient:     k8sClient,
		Namespace:     r.flag.KiamNamespace,
		LabelSelector: r.flag.KiamLabelSelector,
		NodeName:      r.flag.NodeName,
	})
	if err != nil {
		return microerror.Mask(err)
	}

	for {
		if prober.Probe(ctx) {
			errors = 0
			r.logger.Debugf(ctx, "Probe successful")
		} else {
			errors += 1
			r.logger.Debugf(ctx, "Probe failed (number of failed probes: %d)", errors)
			if errors >= r.flag.FailThreshold {
				r.logger.Debugf(ctx, "Reached threshold of %d errors in a row", r.flag.FailThreshold)
				err = restarter.RestartKiamAgent(ctx)
				if err != nil {
					r.logger.Errorf(ctx, err, "Error restarting kiam agent")
					// Next loop.
					continue
				}

				r.logger.Debugf(ctx, "Kiam agent restarted")

				// Kiam restarted, keep probing.
				errors = 0
			}
		}

		time.Sleep(time.Second * time.Duration(r.flag.Interval))
	}
}
