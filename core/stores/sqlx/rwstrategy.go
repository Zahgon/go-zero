package sqlx

import "context"

const (
	policyRoundRobin = "round-robin"

	policyRandom = "random"

	readPrimaryMode readWriteMode = "read-primary"

	readReplicaMode readWriteMode = "read-replica"

	writeMode readWriteMode = "write"

	notSpecifiedMode readWriteMode = ""
)

type readWriteModeKey struct{}

func WithReadPrimary(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func WithReadReplica(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func WithWrite(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

type readWriteMode string

func (m readWriteMode) isValid() bool { _ = "STUB: not implemented"; return false }

func getReadWriteMode(ctx context.Context) readWriteMode {
	_ = "STUB: not implemented"
	return *new(readWriteMode)
}

func usePrimary(ctx context.Context) bool { _ = "STUB: not implemented"; return false }
