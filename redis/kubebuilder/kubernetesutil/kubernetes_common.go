package kubernetesutil

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func generateMetaInformation(resourceKind string, apiVersion string) metav1.TypeMeta {
	return metav1.TypeMeta{
		Kind:       resourceKind,
		APIVersion: apiVersion,
	}
}
