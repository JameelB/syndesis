package enmasse

import (
	"encoding/json"
	"os"

	"github.com/operator-framework/operator-sdk/pkg/sdk"
	glog "github.com/sirupsen/logrus"
	enmasse "github.com/syndesisio/syndesis/install/operator/pkg/apis/enmasse/v1alpha1"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
		return err
	}

	adminUsername, adminPassword, err := getKeycloakAdminCredentials()
	if err != nil {
		glog.Errorf("failed to get admin username and password from keycloak credentials secret %+v", err)
		return err
	}

	// TODO:Check if connection cr already exist

	createKeycloakUser(adminUsername, adminPassword)
	createConnectionCR()
	return nil
}

func createKeycloakUser(adminUsername, adminPassword string) error {
	glog.Infof("Creating keycloak user: %s - %s", adminUsername, adminPassword)
	// TODO:Get authenticated client + token
	// TODO:Create user
	return nil
}

func getKeycloakAdminCredentials() (string, string, error) {
	// TODO:Get keycloak credentials
	kcCreds := &v1.Secret{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Secret",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "keycloak-credentials",
			Namespace: os.Getenv("WATCH_NAMESPACE"),
		},
	}
	err := sdk.Get(kcCreds)
	if err != nil {
		glog.Errorf("failed to get keycloak credentials secret %+v", err)
		return "", "", err
	}

	decodedParams := map[string]string{}
	for k, v := range kcCreds.Data {
		decodedParams[k] = string(v)
	}

	return decodedParams["admin.username"], decodedParams["admin.password"], nil
}

func createConnectionCR() error {
	// CR needs to have username, password and url
	return nil
}

func getConnectionCR() error {
	return nil
}
