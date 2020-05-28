package common

import (
	"fmt"
	"strings"
)

type PartitionInfo struct {
	topic           string
	partition       int
	leader          *Node
	replicas        []*Node
	inSyncReplicas  []*Node
	offlineReplicas []*Node
}

func NewPartitionInfo(
	topic string,
	partition int,
	leader *Node,
	replicas []*Node,
	inSyncReplicas []*Node,
	offlineReplicas []*Node,
) *PartitionInfo {
	p := &PartitionInfo{}
	p.topic = topic
	p.partition = partition
	p.leader = leader
	p.replicas = replicas
	p.inSyncReplicas = inSyncReplicas
	p.offlineReplicas = offlineReplicas
	return p
}

func (p *PartitionInfo) Topic() string {
	return p.topic
}

func (p *PartitionInfo) Partition() int {
	return p.partition
}

// if nil, no leader
func (p *PartitionInfo) Leader() *Node {
	return p.leader
}

func (p *PartitionInfo) nilSafeLeaderIdString() string {
	if p == nil || p.leader == nil {
		return "none"
	}
	return p.leader.idString
}

func (p *PartitionInfo) Replicas() []*Node {
	return p.replicas
}

func (p *PartitionInfo) InSyncReplicas() []*Node {
	return p.inSyncReplicas
}

func (p *PartitionInfo) OfflineReplicas() []*Node {
	return p.offlineReplicas
}

func (p *PartitionInfo) String() string {
	return fmt.Sprintf(
		"Partition(topic = %s, partition = %d, leader = %s, replicas = %s, isr = %s, offlineReplicas = %s)",
		p.topic,
		p.partition,
		p.nilSafeLeaderIdString(),
		p.formatNodeIds(p.replicas),
		p.formatNodeIds(p.inSyncReplicas),
		p.formatNodeIds(p.offlineReplicas))
}

func (p *PartitionInfo) formatNodeIds(nodes []*Node) string {
	if nodes == nil {
		return "[]"
	}
	var ids []string
	for _, n := range nodes {
		ids = append(ids, n.idString)
	}
	return "[" + strings.Join(ids, ",") + "]"
}
