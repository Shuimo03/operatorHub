package kubernetesutil

import (
	"context"
	"github.com/sirupsen/logrus"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateStatefulSet(namespace string, statefulSet *appsv1.StatefulSet) error {
	client, err := GenerateKubeClient()
	if err != nil {
		logrus.Error("Could not generate kubernetes client")
		return err
	}
	_, err = client.AppsV1().StatefulSets(namespace).Create(context.TODO(), statefulSet, metav1.CreateOptions{})
	if err != nil {
		logrus.Error("Could not Create StatefulSets")
		return err
	}
	return nil
}

func GenStatefulSetTemplate(stsMeta metav1.ObjectMeta, podTemplate corev1.PodTemplateSpec) *appsv1.StatefulSet {
	statefulSet := &appsv1.StatefulSet{
		TypeMeta:   generateMetaInformation(StatefulSetKind, ApiVersionV1),
		ObjectMeta: stsMeta,
		Spec: appsv1.StatefulSetSpec{
			Replicas:    int32Ptr(1),
			ServiceName: "",
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{},
			},
			Template: podTemplate,
		},
	}
	return statefulSet
}

// generateMetaInformation generates the meta information
func generateMetaInformation(resourceKind string, apiVersion string) metav1.TypeMeta {
	return metav1.TypeMeta{
		Kind:       resourceKind,
		APIVersion: apiVersion,
	}
}

func int32Ptr(i int32) *int32 {
	return &i
}
