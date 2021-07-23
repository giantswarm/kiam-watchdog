package kiamagentrestarter

import (
	"context"
	"fmt"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Config struct {
	Logger    micrologger.Logger
	K8sClient kubernetes.Interface

	Namespace     string
	LabelSelector string
	NodeName      string
}

type Restarter struct {
	logger    micrologger.Logger
	k8sClient kubernetes.Interface

	namespace     string
	labelSelector string
	nodeName      string
}

func NewRestarter(config Config) (*Restarter, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	if config.K8sClient == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.K8sClient must not be empty", config)
	}

	if config.Namespace == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.Namespace must not be empty", config)
	}

	if config.LabelSelector == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.LabelSelector must not be empty", config)
	}

	if config.NodeName == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.NodeName must not be empty", config)
	}

	return &Restarter{
		logger:        config.Logger,
		k8sClient:     config.K8sClient,
		namespace:     config.Namespace,
		labelSelector: config.LabelSelector,
		nodeName:      config.NodeName,
	}, nil
}

func (r *Restarter) RestartKiamAgent(ctx context.Context) error {
	podList, err := r.k8sClient.CoreV1().Pods(r.namespace).List(ctx, v1.ListOptions{
		LabelSelector: r.labelSelector,
		FieldSelector: fmt.Sprintf("spec.nodeName=%s", r.nodeName),
	})
	if err != nil {
		return microerror.Mask(err)
	}

	if len(podList.Items) != 1 {
		return microerror.Maskf(invalidLabelSelectorError, "Expected exactly 1 Pod to be found in namespace %q by label selector %q on node %q, got %d", r.namespace, r.labelSelector, r.nodeName, len(podList.Items))
	}

	pod := podList.Items[0]

	err = r.k8sClient.CoreV1().Pods(pod.Namespace).Delete(ctx, pod.Name, v1.DeleteOptions{})
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}
