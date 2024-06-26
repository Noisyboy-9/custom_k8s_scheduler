package scheduler

import (
	"github.com/noisyboy-9/random-k8s-scheduler/internal/config"
	"github.com/noisyboy-9/random-k8s-scheduler/internal/enum"
	"github.com/noisyboy-9/random-k8s-scheduler/internal/log"
	"github.com/noisyboy-9/random-k8s-scheduler/internal/model"
	cloudFirst "github.com/noisyboy-9/random-k8s-scheduler/internal/scheduler/cloud-first"
	edgeFirst "github.com/noisyboy-9/random-k8s-scheduler/internal/scheduler/edge-first"
	"github.com/noisyboy-9/random-k8s-scheduler/internal/scheduler/random"
)

type Scheduler interface {
	Run(pod *model.Pod, nodes []*model.Node) (node *model.Node, err error)
}

var S Scheduler

func NewScheduler() {
	switch enum.AlgorithmName(config.Scheduler.Algorithm) {
	case enum.RandomScheduler:
		S = newRandomScheduler()
		log.App.Info("random scheduler created successfully")
	case enum.SmallestFittingEdgeNodeScheduler:
		S = newSmallestFittingEdgeNodeScheduler()
		log.App.Info("smallest-fitting-edge-node scheduler created successfully")
	case enum.BiggestFittingEdgeNodeScheduler:
		S = newBiggestFittingEdgeNodeScheduler()
		log.App.Info("biggest-fitting-edge-node scheduler created successfully")
	case enum.CloudFirstNodeScheduler:
		S = newCloudFirstScheduler()
		log.App.Info("cloud-first scheduler created successfully")
	default:
		log.App.Panic("not known scheduler type")
	}
}

func newRandomScheduler() Scheduler {
	return &random.RandScheduler{}
}

func newSmallestFittingEdgeNodeScheduler() Scheduler {
	return &edgeFirst.SmallestFittingEdgeNodeScheduler{}
}

func newBiggestFittingEdgeNodeScheduler() Scheduler {
	return &edgeFirst.BiggestFittingEdgeNodeScheduler{}
}

func newCloudFirstScheduler() Scheduler {
	return &cloudFirst.CloudFirstScheduler{}
}
