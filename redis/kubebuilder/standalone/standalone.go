package standalone

import (
	v1 "cola.redis/api/v1"
	"cola.redis/kubernetesutil"
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateStandaloneRedis(cr *v1.Standalone) error {
	stsMeta := metav1.ObjectMeta{
		Name:      cr.Name,
		Namespace: cr.Namespace,
	}

	podTemplate := corev1.PodTemplateSpec{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app": "example-app",
			},
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  cr.Name,
					Image: cr.Spec.KubernetesConfig.Image,
				},
			},
		},
	}

	statefulSetTemplate := kubernetesutil.GenStatefulSetTemplate(stsMeta, podTemplate)
	if err := kubernetesutil.CreateStatefulSet(cr.Namespace, statefulSetTemplate); err != nil {
		logrus.Error(err, "Cannot create standalone statefulset for Redis")
		return err
	}
	return nil
}
