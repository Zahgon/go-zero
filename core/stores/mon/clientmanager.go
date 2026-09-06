package mon

import (
	"github.com/zeromicro/go-zero/core/syncx"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var clientManager = syncx.NewResourceManager()

type ClosableClient struct {
	*mongo.Client
}

func (cs *ClosableClient) Close() error { _ = "STUB: not implemented"; return nil }

func Inject(key string, client *mongo.Client) { _ = "STUB: not implemented"; return }

func getClient(url string, opts ...Option) (*mongo.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
