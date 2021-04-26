package etcdstatus

import (
	"github.com/openshift/library-go/pkg/operator/configobserver"
	"github.com/openshift/library-go/pkg/operator/events"
	"github.com/openshift/library-go/pkg/operator/status"
)

// NewCloudProviderObserver returns a new cloudprovider observer for syncing cloud provider specific
// information to controller-manager and api-server.
func NewEtcdStatusObserver(versionRecorder status.VersionGetter, operatorVersion string) configobserver.ObserveConfigFunc {
	cloudObserver := &etcdStatusObserver{
		versionRecorder: versionRecorder,
		operatorVersion: operatorVersion,
	}
	return cloudObserver.ObserveEtcdStatus
}

type etcdStatusObserver struct {
	versionRecorder status.VersionGetter
	operatorVersion string
}

// Observe etcd status observes the version status of the etcd instances.
func (c *etcdStatusObserver) ObserveEtcdStatus(genericListers configobserver.Listers, recorder events.Recorder, currentConfig map[string]interface{}) (map[string]interface{}, []error) {
	observedConfig := map[string]interface{}{}
	// TODO: check if all the etcd operands are available
	allEtcdsUpdated := true
	if allEtcdsUpdated {
		c.setVersion()
	}
	return observedConfig, nil
}

func (c *etcdStatusObserver) setVersion() {
	if c.versionRecorder.GetVersions()["operator"] != c.operatorVersion {
		// Set current version
		c.versionRecorder.SetVersion("operator", c.operatorVersion)
	}
}
