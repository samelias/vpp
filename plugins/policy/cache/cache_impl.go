package cache

import (
	"github.com/ligato/cn-infra/datasync"
	"github.com/ligato/cn-infra/logging"

	nsmodel "github.com/contiv/vpp/plugins/ksr/model/namespace"
	podmodel "github.com/contiv/vpp/plugins/ksr/model/pod"
	policymodel "github.com/contiv/vpp/plugins/ksr/model/policy"
)

// PolicyCache s used for a in-memory storage of K8s State data with fast
// lookups using idxmap-s.
// The cache processes K8s State data updates and RESYNC events through Update()
// and Resync() APIs, respectively.
// The cache allows to get notified about changes via convenient callbacks.
// A watcher needs to implement the interface PolicyCacheWatcher and subscribe
// for watching using Watch() API.
// The cache provides various fast lookup methods (e.g. by the label selector).
type PolicyCache struct {
	Deps
}

// Deps lists dependencies of PolicyCache.
type Deps struct {
	Log logging.Logger
}

// Update processes a datasync change event associated with K8s State data.
// The change is applied into the cache and all subscribed watchers are
// notified.
// The function will forward any error returned by a watcher.
func (pc *PolicyCache) Update(dataChngEv datasync.ChangeEvent) error {
	return nil
}

// Resync processes a datasync resync event associated with K8s State data.
// The cache content is full replaced with the received data and all
// subscribed watchers are notified.
// The function will forward any error returned by a watcher.
func (pc *PolicyCache) Resync(resyncEv datasync.ResyncEvent) error {
	return nil
}

// Watch subscribes a new watcher.
func (pc *PolicyCache) Watch(watcher *PolicyCacheWatcher) {
}

// LookupPod returns data of a given Pod.
func (pc *PolicyCache) LookupPod(pod podmodel.ID) (found bool, data *podmodel.Pod) {
	return false, nil
}

// LookupPodsByLabelSelector evaluates label selector (expression and/or match
// labels) and returns IDs of matching pods.
func (pc *PolicyCache) LookupPodsByLabelSelector(podLabelSelector *policymodel.Policy_LabelSelector) (pods []podmodel.ID) {
	return nil
}

// LookupPodsByNamespace returns IDs of all pods inside a given namespace.
func (pc *PolicyCache) LookupPodsByNamespace(namespace nsmodel.ID) (pods []podmodel.ID) {
	return nil
}

// ListAllPods returns IDs of all known pods.
func (pc *PolicyCache) ListAllPods() (pods []podmodel.ID) {
	return nil
}

// LookupPolicy returns data of a given Policy.
func (pc *PolicyCache) LookupPolicy(policy podmodel.ID) (found bool, data *policymodel.Policy) {
	return false, nil
}

// LookupPoliciesByPod returns IDs of all policies assigned to a given pod.
func (pc *PolicyCache) LookupPoliciesByPod(pod podmodel.ID) (policies []policymodel.ID) {
	return nil
}

// ListAllPolicies returns IDs of all policies.
func (pc *PolicyCache) ListAllPolicies() (policies []policymodel.ID) {
	return nil
}

// LookupNamespace returns data of a given namespace.
func (pc *PolicyCache) LookupNamespace(namespace nsmodel.ID) (found bool, data *nsmodel.Namespace) {
	return false, nil
}

// LookupNamespacesByLabelSelector evaluates label selector (expression
// and/or match labels) and returns IDs of matching namespaces.
func (pc *PolicyCache) LookupNamespacesByLabelSelector(nsLabelSelector *policymodel.Policy_LabelSelector) (namespaces []nsmodel.ID) {
	return nil
}

// ListAllNamespaces returns IDs of all known namespaces.
func (pc *PolicyCache) ListAllNamespaces() (namespaces []nsmodel.ID) {
	return nil
}