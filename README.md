# golang-platform

Kubernetes platform operator in Go. CRDs are generated from Go types with controller-gen.

Day 1 status: Tenant API types only (scaffolded by hand; no cluster steps run yet).

On a machine with Go, Docker, kind and kubebuilder:

    kubebuilder init --domain platform.io --repo github.com/example/golang-platform
    kubebuilder create api --group platform --version v1alpha1 --kind Tenant
    # replace api/v1alpha1/tenant_types.go with this repo's version
    make manifests generate install
    kubectl apply -f examples/tenant-payments.yaml
