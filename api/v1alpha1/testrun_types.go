package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type TestRunSpec struct {
	Service     string `json:"service"`
	Version     string `json:"version"`
	CallbackURL string `json:"callbackURL,omitempty"` // omitempty? if not tied to a pipe but only running on cron
}

type TestRunStatus struct {
	Phase   string `json:"phase"`
	Promote bool   `json:"promote"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type TestRun struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TestRunSpec   `json:"spec,omitempty"`
	Status            TestRunStatus `json:"status,omitempty"`
}
