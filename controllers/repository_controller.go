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

package controllers

import (
	"context"
	"time"

	"golang.org/x/oauth2"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	repositoriesv1alpha1 "github.com/stone-payments/stone-sreplatform-challenge/api/v1alpha1"
	github "github.com/stone-payments/stone-sreplatform-challenge/client"
)

// RepositoryReconciler reconciles a Repository object
type RepositoryReconciler struct {
	client.Client
	Scheme   *runtime.Scheme
	Recorder record.EventRecorder
}

//+kubebuilder:rbac:groups=repositories.platform.buy4.io,resources=repositories,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=repositories.platform.buy4.io,resources=repositories/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=repositories.platform.buy4.io,resources=repositories/finalizers,verbs=update
func (r *RepositoryReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	cr := &repositoriesv1alpha1.Repository{}
	if err := r.Get(ctx, req.NamespacedName, cr); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	repositories := newClient("MYF4K3P4T")

	if cr.IsBeingDeleted() {
		if cr.HasFinalizer(repositoriesv1alpha1.RepositoryFinalizerName) {
			if err := repositories.Delete(ctx, cr.Spec.Owner, cr.Spec.Name); err != nil {
				return ctrl.Result{}, err
			}

			cr.RemoveFinalizer(repositoriesv1alpha1.RepositoryFinalizerName)
			if err := r.Update(ctx, cr); err != nil {
				return ctrl.Result{}, err
			}
		}
		return ctrl.Result{}, nil
	}

	if !cr.HasFinalizer(repositoriesv1alpha1.RepositoryFinalizerName) {
		cr.AddFinalizer(repositoriesv1alpha1.RepositoryFinalizerName)
		if err := r.Update(ctx, cr); err != nil {
			return ctrl.Result{}, err
		}
	}

	repo := generateRepository(cr.Spec)
	repositories.Create(ctx, repo)
	r.Recorder.Event(cr, "Normal", "Created", "External resource successfully created")

	return ctrl.Result{RequeueAfter: 1 * time.Minute}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *RepositoryReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&repositoriesv1alpha1.Repository{}).
		Complete(r)
}

func newClient(token string) *github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc)
}

func generateRepository(internal repositoriesv1alpha1.RepositorySpec) *github.Repository {
	return &github.Repository{
		Name:            &internal.Name,
		Owner:           &internal.Owner,
		Private:         boolPtr(false),
		HasIssues:       boolPtr(true),
		AutoInit:        boolPtr(true),
		LicenseTemplate: stringPtr("apache-2.0"),
	}
}

func boolPtr(b bool) *bool { return &b }

func stringPtr(s string) *string { return &s }
