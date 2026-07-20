package redis

import (
	red "github.com/redis/go-redis/v9"
)

type ClosableNode interface {
	RedisNode
	Close()
}

func CreateBlockingNode(r *Redis) (ClosableNode, error) {
	_ = "STUB: not implemented"
	return *new(ClosableNode), nil
}

type (
	clientBridge struct {
		*red.Client
	}

	clusterBridge struct {
		*red.ClusterClient
	}
)

func (bridge *clientBridge) Close() { _ = "STUB: not implemented"; return }

func (bridge *clusterBridge) Close() { _ = "STUB: not implemented"; return }
