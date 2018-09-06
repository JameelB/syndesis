package enmasse

import (
	"strings"

	glog "github.com/sirupsen/logrus"
	"k8s.io/api/core/v1"
)

// Reconcile the state
func ReconcileCredentials(secret *v1.Secret, deleted bool) error {
	if strings.Contains(secret.ObjectMeta.Name, "enmasse") && strings.Contains(secret.ObjectMeta.Name, "credentials") {
		glog.Infof("Reconciling Secret: >>> %+v\n", secret)
	}
	return nil
}
