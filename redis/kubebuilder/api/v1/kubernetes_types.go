package v1

import corev1 "k8s.io/api/core/v1"

type KubernetesConfig struct {
	Image           string                       `json:"image"`
	ImagePullPolicy corev1.PullPolicy            `json:"imagePullPolicy,omitempty"`
	Storage         Storage                      `json:"storage"`
	Resources       *corev1.ResourceRequirements `json:"resources,omitempty"`
	Service         string                       `json:"service,omitempty"`
}
type Storage struct {
	VolumeClaimTemplate corev1.PersistentVolumeClaim `json:"volumeClaimTemplate,omitempty"`
	VolumeMount         AdditionalVolume             `json:"volumeMount,omitempty"`
}
type AdditionalVolume struct {
	Volume    []corev1.Volume      `json:"volume,omitempty"`
	MountPath []corev1.VolumeMount `json:"mountPath,omitempty"`
}
