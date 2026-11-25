package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type ResourceQuotaSpec struct {
	CPU    string `json:"cpu"`
	Memory string `json:"memory"`
	// +kubebuilder:validation:Minimum=1
	Pods int64 `json:"pods"`
}

type TenantSpec struct {
	// +kubebuilder:validation:MinLength=1
	DisplayName   string            `json:"displayName"`
	ResourceQuota ResourceQuotaSpec `json:"resourceQuota"`
}

type TenantStatus struct {
	Phase string `json:"phase,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Tenant struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TenantSpec   `json:"spec,omitempty"`
	Status TenantStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
type TenantList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Tenant `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Tenant{}, &TenantList{})
}
