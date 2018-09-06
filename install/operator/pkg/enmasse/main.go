package enmasse

import (
	"encoding/json"

	glog "github.com/sirupsen/logrus"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/syndesisio/syndesis/install/operator/dep-cache/sources/https---github.com-operator--framework-operator--sdk/pkg/sdk"
	enmasse "github.com/syndesisio/syndesis/install/operator/pkg/apis/enmasse/v1alpha1"
)

// Reconcile the state
func ReconcileConfigmap(configmap *v1.ConfigMap, deleted bool) error {
	// TODO: Get configmap of enmasse standard
	if val, ok := configmap.ObjectMeta.Labels["type"]; ok && val == "address-space" {
		return reconcileAddressSpace(configmap, deleted)
	}
	return nil
}

func reconcileAddressSpace(configmap *v1.ConfigMap, deleted bool) error {
	var addressSpace enmasse.AddressSpace
	if err := json.Unmarshal([]byte(configmap.Data["config.json"]), &addressSpace); err != nil {
		glog.Errorf("Failed to unmarshal the configmap data into an AddressSpace object %+v", err)
		return nil
	}

	// TODO:Check if connection cr already exist

	createKeycloakUser()
	createConnectionCR()
	return nil
}

func createKeycloakUser() error {
	// TODO:Get keycloak credentials
	keycloakCredSecret := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{

		}
	}
	sdk.Get()
	// TODO:Get authenticated client + token
	// TODO:Create user
	return nil
}

func createConnectionCR() error {
	// CR needs to have username, password and url
	return nil
}

func getConnectionCR() error {
	return nil
}
