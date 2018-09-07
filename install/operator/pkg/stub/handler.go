package stub

import (
	"context"

	"github.com/operator-framework/operator-sdk/pkg/sdk"
	enmassev1alpha1 "github.com/syndesisio/syndesis/install/operator/pkg/apis/enmasse/v1alpha1"
	api "github.com/syndesisio/syndesis/install/operator/pkg/apis/syndesis/v1alpha1"
	"github.com/syndesisio/syndesis/install/operator/pkg/enmasse"
	"github.com/syndesisio/syndesis/install/operator/pkg/syndesis"

	"k8s.io/api/core/v1"
)

func NewHandler() sdk.Handler {
	return &Handler{}
}

type Handler struct {
	// Fill me
}

func (h *Handler) Handle(ctx context.Context, event sdk.Event) error {
	switch o := event.Object.(type) {
	case *api.Syndesis:
		return syndesis.Reconcile(o, event.Deleted)
	case *v1.ConfigMap:
		return enmasse.ReconcileConfigmap(o, event.Deleted)
	case *enmassev1alpha1.AddressSpace:
		return enmasse.ReconcileConnection(o, event.Deleted)
	}
	return nil
}
