package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type SecretKeyRef struct {
	// No namespace field on purpose: always resolved in the Application's own namespace.
	Name string `json:"name"`
	Key  string `json:"key"`
}

type ApplicationSpec struct {
	// +kubebuilder:validation:Pattern=`^registry\.example\.com/.+`
	Image string `json:"image"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=20
	Replicas  *int32        `json:"replicas,omitempty"`
	Port      int32         `json:"port,omitempty"`
	SecretRef *SecretKeyRef `json:"secretRef,omitempty"`
}

type ApplicationStatus struct {
	Phase string `json:"phase,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=app
type Application struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApplicationSpec   `json:"spec,omitempty"`
	Status ApplicationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
type ApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Application `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Application{}, &ApplicationList{})
}
