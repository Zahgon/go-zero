package discov

const (
	_ = iota
	indexOfId
)

const timeToLive int64 = 10

var TimeToLive = timeToLive

func extract(etcdKey string, index int) (string, bool) { _ = "STUB: not implemented"; return "", false }

func extractId(etcdKey string) (string, bool) { _ = "STUB: not implemented"; return "", false }

func makeEtcdKey(key string, id int64) string { _ = "STUB: not implemented"; return "" }
