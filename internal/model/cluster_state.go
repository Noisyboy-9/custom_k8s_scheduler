package model

import (
	"errors"
	"fmt"
	"k8s.io/apimachinery/pkg/types"
	"sync"
)

type ClusterState struct {
	mux           *sync.RWMutex
	nodes         map[types.UID]*Node
	isNodesSynced bool
	pods          map[types.UID]*Pod
	placement     map[*Pod]*Node
}

func NewClusterState() *ClusterState {
	return &ClusterState{
		mux:       new(sync.RWMutex),
		nodes:     make(map[types.UID]*Node),
		pods:      make(map[types.UID]*Pod),
		placement: make(map[*Pod]*Node),
	}
}

func (s *ClusterState) IsNodesSynced() bool {
	s.mux.RLock()
	defer s.mux.RUnlock()
	return s.isNodesSynced
}

func (s *ClusterState) SetIsNodesSynced(isNodesSynced bool) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.isNodesSynced = isNodesSynced
}

func (s *ClusterState) AddPod(p *Pod) {
	s.pods[p.ID] = p
}

func (s *ClusterState) RemovePod(p *Pod) {
	delete(s.pods, p.ID)
	delete(s.placement, p)
}
func (s *ClusterState) RemovePodByID(id types.UID) {
	delete(s.pods, id)
}

func (s *ClusterState) EditPodWithUID(id types.UID, edited *Pod) error {
	if _, ok := s.pods[id]; !ok {
		return errors.New("pod with given uid doesn't exist")
	}

	s.pods[id] = edited
	return nil
}

func (s *ClusterState) GetPodByUID(id types.UID) (pod *Pod, exists bool) {
	pod, exists = s.pods[id]
	return
}

func (s *ClusterState) AddNode(n *Node) {
	s.nodes[n.ID] = n
}

func (s *ClusterState) RemoveNode(n *Node) {
	delete(s.nodes, n.ID)
}

func (s *ClusterState) RemoveNodeByID(id types.UID) {
	delete(s.nodes, id)
}

func (s *ClusterState) EditNodeWithUID(id types.UID, edited *Node) error {
	if _, ok := s.nodes[id]; !ok {
		return errors.New("node with given id doesn't exist")
	}

	s.nodes[id] = edited
	return nil
}
func (s *ClusterState) GetNodeWithID(id types.UID) (node *Node, exists bool) {
	node, exists = s.nodes[id]
	return
}

func (s *ClusterState) Nodes() []*Node {
	nodes := make([]*Node, 0, len(s.nodes))
	i := 0
	for _, n := range s.nodes {
		nodes = append(nodes, n)
		i++
	}

	return nodes
}

func (s *ClusterState) Pods() []*Pod {
	pods := make([]*Pod, 0, len(s.pods))
	i := 0
	for _, p := range s.pods {
		pods = append(pods, p)
		i++
	}

	return pods
}

func (s *ClusterState) AllocateResources(node *Node, pod *Pod) {
	node.allocateCpu(*pod.Cores)
	node.allocateMemory(*pod.Memory)
}

func (s *ClusterState) DeAllocateResources(node *Node, pod *Pod) {
	node.deAllocateCpu(*pod.Cores)
	node.deallocateMemory(*pod.Memory)
}

func (s *ClusterState) SaveSelectedNodeForPod(selectedNode *Node, pod *Pod) {
	s.placement[pod] = selectedNode
}

func (s *ClusterState) GetSelectedNodeForPod(pod *Pod) (n *Node, exist bool) {
	n, exist = s.placement[pod]
	return
}

func (s *ClusterState) FindNodeByName(name string) (n *Node, err error) {
	for _, node := range s.nodes {
		if node.Name == name {
			return node, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("node with name: %s can't be found", name))
}
