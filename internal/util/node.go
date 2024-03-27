package util

import (
	"errors"
	"fmt"
	"github.com/noisyboy-9/random-k8s-scheduler/internal/model"
)

func FindNodeByName(nodes []*model.Node, name string) (*model.Node, error) {
	for _, node := range nodes {
		if node.Name() == name {
			return node, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("can't find node with name: %s in list of nodes", name))
}