package service

import (
	"math/rand"

	"github.com/noisyboy-9/random-k8s-scheduler/internal/config"
	"github.com/noisyboy-9/random-k8s-scheduler/internal/log"
	"github.com/noisyboy-9/random-k8s-scheduler/internal/model"
)

type scheduler struct {
	Name      string
	Namespace string
}

var Scheduler *scheduler

func NewScheduler() {
	Scheduler = &scheduler{
		Name:      config.Scheduler.Name,
		Namespace: config.Scheduler.Namespace,
	}

	log.App.Infof("scheduler has been initialized in namespace: %v, with name: %v", config.Scheduler.Namespace, config.Scheduler.Name)
}

func (scheduler *scheduler) FindNodeForBinding(pod *model.Pod, nodes []*model.Node) (node *model.Node, err error) {
	// filtering step
	eligibleNodes := make([]*model.Node, 0)
	for _, node := range nodes {
		if node.HasEnoughResourcesForPod(pod) {
			eligibleNodes = append(eligibleNodes, node)
		}
	}

	// select random node
	return eligibleNodes[rand.Intn(len(eligibleNodes))], nil
}
