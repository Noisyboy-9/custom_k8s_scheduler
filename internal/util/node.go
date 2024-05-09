package util

import (
	"github.com/noisyboy-9/random-k8s-scheduler/internal/model"
	"k8s.io/api/core/v1"
	"k8s.io/utils/strings/slices"
)

var edgeNodes = []string{"uq7j5k991-01", "uq7g5w631-01", "uq7p7x251-01"}
var masterNodeName = "uq7g5t611-01"

func IsNodeOnEdge(nodeKubernetesObject *v1.Node) bool {
	n := nodeKubernetesObject.GetName()
	return slices.Contains(edgeNodes, n)
}

func IsMasterNode(node *model.Node) bool {
	return node.Name == masterNodeName
}
