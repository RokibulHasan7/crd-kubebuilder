/*
Copyright 2023.

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
	"fmt"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	apiserverv1 "my.crd/crd-kubebuilder/api/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"time"
)

var logger = log.Log.WithName("rokibulhasan_controller")

// RokibulHasanReconciler reconciles a RokibulHasan object
type RokibulHasanReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=apiserver.my.crd,resources=rokibulhasans,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=apiserver.my.crd,resources=rokibulhasans/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=apiserver.my.crd,resources=rokibulhasans/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the RokibulHasan object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.12.2/pkg/reconcile
func (r *RokibulHasanReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := logger.WithValues("Request.Namespace", req.Namespace, "Request.Name", req.Name)

	// TODO(user): your logic here
	log.Info("Reconcile Started.")
	rokibul := &apiserverv1.RokibulHasan{}
	err := r.Get(ctx, req.NamespacedName, rokibul)
	if err != nil {
		return ctrl.Result{}, err
	}

	startTime := rokibul.Spec.StartTime
	endTime := rokibul.Spec.EndTime
	replicas := rokibul.Spec.Replicas
	log.Info(fmt.Sprintf("Current Replicas: %v\n", replicas))
	currenttime := time.Now()
	log.Info(fmt.Sprintf("current time in hour : %v\n", currenttime))
	currentTime := time.Now().UTC().Hour()
	log.Info(fmt.Sprintf("current time in hour : %d\n", currentTime))
	if currentTime >= startTime && currentTime < endTime {
		if err = scaleDeployment(rokibul, r, ctx, int32(replicas)); err != nil {
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{RequeueAfter: time.Duration(30 * time.Second)}, nil
}

func scaleDeployment(rokibul *apiserverv1.RokibulHasan, r *RokibulHasanReconciler, ctx context.Context, replicas int32) error {
	for _, deploy := range rokibul.Spec.Deployments {
		deployment := &v1.Deployment{}
		err := r.Get(ctx, types.NamespacedName{
			Namespace: deploy.Namespace,
			Name:      deploy.Name,
		}, deployment)
		if err != nil {
			return err
		}

		if deployment.Spec.Replicas != &replicas {
			deployment.Spec.Replicas = &replicas
			err := r.Update(ctx, deployment)
			if err != nil {
				rokibul.Status.Status = apiserverv1.FAILED
				return err
			}
			rokibul.Status.Status = apiserverv1.SUCCESS
			err = r.Status().Update(ctx, rokibul)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *RokibulHasanReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&apiserverv1.RokibulHasan{}).
		Complete(r)
}
