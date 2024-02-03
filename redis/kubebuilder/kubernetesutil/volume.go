package kubernetesutil

import (
	v1 "cola.redis/api/v1"
	"context"
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 根据sc和pvc来自动创建pvc
func CreateVolume(cr *v1.Standalone) error {
	client, err := GenerateKubeClient()
	if err != nil {
		logrus.Error("Could not generate kubernetes client")
		return err
	}
	volumeInfo := cr.Spec.KubernetesConfig.Storage
	pvc := buildPVCTemplate(cr.Name, cr.Namespace, volumeInfo.VolumeClaimTemplate.Spec)
	_, err = client.CoreV1().PersistentVolumeClaims(cr.Namespace).Create(context.TODO(), pvc, metav1.CreateOptions{})
	if err != nil {
		logrus.Error("Could Create PVC:", err)
	}
	return nil
}

func buildPVCTemplate(volumeName, namespace string, storageSpec corev1.PersistentVolumeClaimSpec) *corev1.PersistentVolumeClaim {
	pvc := &corev1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name:      volumeName,
			Namespace: namespace,
		},
		Spec: storageSpec,
	}
	return pvc
}
