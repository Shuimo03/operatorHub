package kubernetesutil

import (
	redisv1 "cola.redis/api/v1"
	"github.com/go-logr/logr"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	StandaloneFinalizer string = "standaloneFinalizer"
)

func HandleStandaloneFinalizer(cli client.Client, k8sClient kubernetes.Interface, logger logr.Logger, cr *redisv1.Standalone) error {
	return nil
}
