/*
Copyright 2021.

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

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	managementv1alpha1 "github.com/project-flotta/flotta-operator/api/v1alpha1"
	"github.com/project-flotta/flotta-operator/internal/common/repository/edgeautoconfig"
	"github.com/project-flotta/flotta-operator/internal/common/repository/edgedevice"
)

const (
	EdgeAutoConfigFinalizer = "edge-auto-config-finalizer"
)

// EdgeAutoConfigReconciler reconciles a EdgeAutoConfig object
type EdgeAutoConfigReconciler struct {
	client.Client
	Scheme                   *runtime.Scheme
	EdgeAutoConfigRepository edgeautoconfig.Repository
	EdgeDeviceRepository     edgedevice.Repository
	MaxConcurrentReconciles  int
	AutoApproval             bool
}

//+kubebuilder:rbac:groups=management.project-flotta.io,resources=edgeautoconfigs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=management.project-flotta.io,resources=edgeautoconfigs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=management.project-flotta.io,resources=edgeautoconfigs/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the EdgeAutoConfig object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.2/pkg/reconcile
func (r *EdgeAutoConfigReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here
	logger := log.FromContext(ctx)
	logger.Info("Reconciling", "edgeautoconfig", req)

	edgeautocfg, err := r.EdgeAutoConfigRepository.Read(ctx, req.Name, req.Namespace)
	if err != nil {
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{Requeue: true}, err
	}

	if edgeautocfg.DeletionTimestamp != nil {
		logger.Info("Reconciling", "edgeautoconfig delete")
		// if err := r.removeRelatedEDSR(ctx, edgeDevice); err != nil {
		// 	return ctrl.Result{Requeue: true}, err
		// }
		// return ctrl.Result{}, r.removeFinalizer(ctx, edgeDevice)
	}

	// if !r.ObcAutoCreate && !storage.ShouldCreateOBC(edgeDevice.Spec.Storage) {
	// 	return ctrl.Result{}, nil
	// }

	edgeautocfgcpy := edgeautocfg.DeepCopy()
	//get devices which do not have autoconfig workloads set
	edgedevicesstatus := edgeautocfgcpy.Status.EdgeDevices
	logger.Info("device status", edgedevicesstatus[0].Name)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *EdgeAutoConfigReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&managementv1alpha1.EdgeAutoConfig{}).
		WithEventFilter(predicate.GenerationChangedPredicate{}).
		WithOptions(controller.Options{MaxConcurrentReconciles: r.MaxConcurrentReconciles}).
		Complete(r)
}