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

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha1 "github.com/rybnico/provider-jet-hcloud/apis/floating/v1alpha1"
	v1alpha1hcloud "github.com/rybnico/provider-jet-hcloud/apis/hcloud/v1alpha1"
	v1alpha1load "github.com/rybnico/provider-jet-hcloud/apis/load/v1alpha1"
	v1alpha1managed "github.com/rybnico/provider-jet-hcloud/apis/managed/v1alpha1"
	v1alpha1network "github.com/rybnico/provider-jet-hcloud/apis/network/v1alpha1"
	v1alpha1placement "github.com/rybnico/provider-jet-hcloud/apis/placement/v1alpha1"
	v1alpha1server "github.com/rybnico/provider-jet-hcloud/apis/server/v1alpha1"
	v1alpha1ssh "github.com/rybnico/provider-jet-hcloud/apis/ssh/v1alpha1"
	v1alpha1uploaded "github.com/rybnico/provider-jet-hcloud/apis/uploaded/v1alpha1"
	v1alpha1apis "github.com/rybnico/provider-jet-hcloud/apis/v1alpha1"
	v1alpha1volume "github.com/rybnico/provider-jet-hcloud/apis/volume/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1alpha1hcloud.SchemeBuilder.AddToScheme,
		v1alpha1load.SchemeBuilder.AddToScheme,
		v1alpha1managed.SchemeBuilder.AddToScheme,
		v1alpha1network.SchemeBuilder.AddToScheme,
		v1alpha1placement.SchemeBuilder.AddToScheme,
		v1alpha1server.SchemeBuilder.AddToScheme,
		v1alpha1ssh.SchemeBuilder.AddToScheme,
		v1alpha1uploaded.SchemeBuilder.AddToScheme,
		v1alpha1apis.SchemeBuilder.AddToScheme,
		v1alpha1volume.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
