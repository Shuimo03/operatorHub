package kubernetesutil

import (
	v1 "cola.redis/api/v1"
	"context"
	"github.com/sirupsen/logrus"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateStatefulSet(cr *v1.Standalone) error {
	client, err := GenerateKubeClient()
	if err != nil {
		logrus.Error("Could not generate kubernetes client")
		return err
	}
	label := buildLabel(cr.Kind, cr.Name)
	statefulSet := builderStatefulSet(cr.Name, cr.Namespace, cr.Spec.KubernetesConfig.Service, label, cr.Name, cr.Spec.KubernetesConfig.Image)
	_, err = client.AppsV1().StatefulSets(cr.Namespace).Create(context.TODO(), statefulSet, metav1.CreateOptions{})
	if err != nil {
		logrus.Error("Could not Create StatefulSets")
		return err
	}
	return nil
}

func builderStatefulSet(name, namespace, svc string, label map[string]string, containerName, image string) *appsv1.StatefulSet {
	statefulSet := &appsv1.StatefulSet{
		TypeMeta: generateMetaInformation("StatefulSet", "apps/v1"),
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: appsv1.StatefulSetSpec{
			Replicas: int32Ptr(1),
			//ServiceName: svc,
			Selector: &metav1.LabelSelector{
				MatchLabels: label,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: label,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  containerName,
							Image: image,
						},
					},
				},
			},
		},
	}

	return statefulSet
}

func int32Ptr(i int32) *int32 {
	return &i
}
