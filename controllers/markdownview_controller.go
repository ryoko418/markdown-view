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

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	//	"github.com/operator-framework/operator-lib/handler"
	//	"k8s.io/apimachinery/pkg/runtime/schema"
	//	"sigs.k8s.io/controller-runtime/pkg/source"

	viewv1 "github.com/ryoko418/markdown-view/api/v1"
)

// MarkdownViewReconciler reconciles a MarkdownView object
type MarkdownViewReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=view.indigo.co.jp,resources=markdownviews,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=view.indigo.co.jp,resources=markdownviews/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=view.indigo.co.jp,resources=markdownviews/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the MarkdownView object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.12.2/pkg/reconcile
func (r *MarkdownViewReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	//_ = log.FromContext(ctx)

	log := log.FromContext(ctx)
	log.Info("ReplicaSet")
	// TODO(user): your logic here

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *MarkdownViewReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&viewv1.MarkdownView{}).
		//	        Watches(
		//			&source.Kind{Type: &viewv1.MarkdownView{}},
		//			&handler.EnqueueRequestForAnnotation{schema.GroupKind{Group: "apps", Kind: "ReplicaSet"}},
		//		).
		Complete(r)
}
