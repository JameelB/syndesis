package enmasse

import (
	"os"

	"github.com/operator-framework/operator-sdk/pkg/sdk"
	glog "github.com/sirupsen/logrus"
	enmasse "github.com/syndesisio/syndesis/install/operator/pkg/apis/enmasse/v1alpha1"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Reconcile the state
func ReconcileConnection(addressSpace *enmasse.AddressSpace, deleted bool) error {
	glog.Infoln("Reconciling connection...")
	return nil
}

func getKeycloakAuthToken(adminUsername, adminPassword string) error {
	return nil
}

func createKeycloakUser(adminUsername, adminPassword string) error {
	glog.Infof("Creating keycloak user: %s - %s", adminUsername, adminPassword)
	// TODO:Get auth token
	// TODO:Create user
	return nil
}

func getKeycloakAdminCredentials() (string, string, error) {
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
