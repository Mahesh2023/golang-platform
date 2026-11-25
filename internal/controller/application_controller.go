package controller

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	platformv1alpha1 "github.com/example/golang-platform/api/v1alpha1"
)

type ApplicationReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=platform.platform.io,resources=applications;applications/status,verbs=get;list;watch;update;patch
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups="",resources=namespaces,verbs=get;list;watch

func (r *ApplicationReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	app := &platformv1alpha1.Application{}
	if err := r.Get(ctx, req.NamespacedName, app); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	ns := &corev1.Namespace{}
	if err := r.Get(ctx, types.NamespacedName{Name: app.Namespace}, ns); err != nil {
		return ctrl.Result{}, err
	}
	tenant, ok := ns.Labels["platform.io/tenant"]
	if !ok {
		return ctrl.Result{}, nil
	}

	dep := &appsv1.Deployment{ObjectMeta: metav1.ObjectMeta{Name: app.Name, Namespace: app.Namespace}}
	_, err := controllerutil.CreateOrUpdate(ctx, r.Client, dep, func() error {
		dep.Labels = map[string]string{"app": app.Name, "platform.io/tenant": tenant}
		dep.Spec.Replicas = app.Spec.Replicas
		dep.Spec.Selector = &metav1.LabelSelector{MatchLabels: map[string]string{"app": app.Name}}
		dep.Spec.Template.Labels = dep.Spec.Selector.MatchLabels
		dep.Spec.Template.Spec.Containers = []corev1.Container{{Name: "app", Image: app.Spec.Image}}
		return controllerutil.SetControllerReference(app, dep, r.Scheme)
	})
	return ctrl.Result{}, err
}

func (r *ApplicationReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&platformv1alpha1.Application{}).
		Owns(&appsv1.Deployment{}).
		Complete(r)
}
