/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type BalancerTargetObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type BalancerTargetParameters struct {

	// +kubebuilder:validation:Optional
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// +kubebuilder:validation:Optional
	LabelSelector *string `json:"labelSelector,omitempty" tf:"label_selector,omitempty"`

	// +kubebuilder:validation:Required
	LoadBalancerID *int64 `json:"loadBalancerId" tf:"load_balancer_id,omitempty"`

	// +kubebuilder:validation:Optional
	ServerID *int64 `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	UsePrivateIP *bool `json:"usePrivateIp,omitempty" tf:"use_private_ip,omitempty"`
}

// BalancerTargetSpec defines the desired state of BalancerTarget
type BalancerTargetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BalancerTargetParameters `json:"forProvider"`
}

// BalancerTargetStatus defines the observed state of BalancerTarget.
type BalancerTargetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BalancerTargetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BalancerTarget is the Schema for the BalancerTargets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,hcloudjet}
type BalancerTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BalancerTargetSpec   `json:"spec"`
	Status            BalancerTargetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BalancerTargetList contains a list of BalancerTargets
type BalancerTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BalancerTarget `json:"items"`
}

// Repository type metadata.
var (
	BalancerTarget_Kind             = "BalancerTarget"
	BalancerTarget_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BalancerTarget_Kind}.String()
	BalancerTarget_KindAPIVersion   = BalancerTarget_Kind + "." + CRDGroupVersion.String()
	BalancerTarget_GroupVersionKind = CRDGroupVersion.WithKind(BalancerTarget_Kind)
)

func init() {
	SchemeBuilder.Register(&BalancerTarget{}, &BalancerTargetList{})
}
