/*
Copyright 2022.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RepositorySpec defines the desired state of Repository
type RepositorySpec struct {
	// The name of the Repository.
	Name string `json:"name"`

	// The owner of the Repository. This field can be either
	// an user or an organization.
	Owner string `json:"owner"`

	// The Repository configuration type.
	Type string `json:"type"`

	// The Secret reference that contains the credentials.
	CredentialsRef SecretKeyReference `json:"credentialsRef"`
}

// RepositoryStatus defines the observed state of Repository
type RepositoryStatus struct{}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Repository is the Schema for the repositories API
type Repository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RepositorySpec   `json:"spec,omitempty"`
	Status RepositoryStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// RepositoryList contains a list of Repository
type RepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Repository `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Repository{}, &RepositoryList{})
}

type SecretKeyReference struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

const RepositoryFinalizerName string = "repositories.platform.buy4.io/finalizer"

func (r *Repository) IsBeingDeleted() bool {
	return !r.ObjectMeta.DeletionTimestamp.IsZero()
}

func (r *Repository) HasFinalizer(finalizerName string) bool {
	return containsString(r.ObjectMeta.Finalizers, finalizerName)
}

func (r *Repository) AddFinalizer(finalizerName string) {
	r.ObjectMeta.Finalizers = append(r.ObjectMeta.Finalizers, finalizerName)
}

func (r *Repository) RemoveFinalizer(finalizerName string) {
	r.ObjectMeta.Finalizers = removeString(r.ObjectMeta.Finalizers, finalizerName)
}

func containsString(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

func removeString(slice []string, s string) (result []string) {
	for _, item := range slice {
		if item == s {
			continue
		}
		result = append(result, item)
	}
	return
}
