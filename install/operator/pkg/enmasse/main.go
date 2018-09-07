package enmasse

import (
	"encoding/json"

	"github.com/operator-framework/operator-sdk/pkg/sdk"
	glog "github.com/sirupsen/logrus"
	enmasse "github.com/syndesisio/syndesis/install/operator/pkg/apis/enmasse/v1alpha1"
	syndesis "github.com/syndesisio/syndesis/install/operator/pkg/apis/syndesis/v1alpha1"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Reconcile the state
func ReconcileConfigmap(configmap *v1.ConfigMap, deleted bool) error {
	// TODO: Get configmap of enmasse standard
	if val, ok := configmap.ObjectMeta.Labels["type"]; ok && val == "address-space" {
		return reconcileConfigMap(configmap, deleted)
	}
	return nil
}

func reconcileConfigMap(configmap *v1.ConfigMap, deleted bool) error {
	addressSpace, err := getAddressSpaceObj(configmap)
	if err != nil {
		glog.Errorf("failed to get enmasse address space object %+v", err)
		return err
	}

	// TODO: If not there, create it. If exist, return nil
	err = sdk.Get(getConnectionCRObj(addressSpace))

	// Create Connection CR
	err = sdk.Create(addressSpace)
	if err != nil {
		glog.Errorf("failed to create address space %+v", err)
		return err
	}

	return nil
}

func getAddressSpaceObj(configmap *v1.ConfigMap) (*enmasse.AddressSpace, error) {
	var addressSpace enmasse.AddressSpace
	if err := json.Unmarshal([]byte(configmap.Data["config.json"]), &addressSpace); err != nil {
		glog.Errorf("Failed to unmarshal the configmap data into an AddressSpace object %+v", err)
		return nil, err
	}
	return &addressSpace, nil
}

func getConnectionCRObj(addressSpace *enmasse.AddressSpace) *syndesis.Connection {
	var messagingEndpoint enmasse.EndpointStatus
	var amqpPort string

	for _, v := range addressSpace.Status.EndpointStatuses {
		if v.Name == "messaging" {
			messagingEndpoint = v
		}
	}
	for _, v := range messagingEndpoint.ServicePorts {
		if v.Name == "amqp" {
			amqpPort = string(v.Port)
		}
	}
	amqpURL := "amqp://" + messagingEndpoint.ServiceHost + ":" + amqpPort + "?amqp.saslMechanisms=PLAIN"

	return &syndesis.Connection{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "connection-" + addressSpace.Name,
			Namespace: addressSpace.Namespace,
		},
		Spec: syndesis.ConnectionSpec{
			URL: amqpURL,
		},
		Status: syndesis.ConnectionStatus{
			Ready: false,
			Phase: "PendingAuth",
		},
	}
}
