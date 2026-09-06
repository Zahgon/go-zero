package logx

import (
	"errors"

	"github.com/zeromicro/go-zero/core/syncx"
)

const (
	DebugLevel uint32 = iota

	InfoLevel

	ErrorLevel

	SevereLevel

	disableLevel = 0xff
)

const (
	jsonEncodingType = iota
	plainEncodingType
)

const (
	plainEncoding    = "plain"
	plainEncodingSep = '\t'
	sizeRotationRule = "size"

	accessFilename = "access.log"
	errorFilename  = "error.log"
	severeFilename = "severe.log"
	slowFilename   = "slow.log"
	statFilename   = "stat.log"

	fileMode   = "file"
	volumeMode = "volume"

	levelAlert  = "alert"
	levelInfo   = "info"
	levelError  = "error"
	levelSevere = "severe"
	levelFatal  = "fatal"
	levelSlow   = "slow"
	levelStat   = "stat"
	levelDebug  = "debug"

	backupFileDelimiter = "-"
	nilAngleString      = "<nil>"
	flags               = 0x0
)

const (
	defaultCallerKey    = "caller"
	defaultContentKey   = "content"
	defaultDurationKey  = "duration"
	defaultLevelKey     = "level"
	defaultSpanKey      = "span"
	defaultTimestampKey = "@timestamp"
	defaultTraceKey     = "trace"
	defaultTruncatedKey = "truncated"
)

var (
	ErrLogPathNotSet = errors.New("log path must be set")

	ErrLogServiceNameNotSet = errors.New("log service name must be set")

	ExitOnFatal = syncx.ForAtomicBool(true)

	truncatedField = Field(truncatedKey, true)
)

var (
	callerKey    = defaultCallerKey
	contentKey   = defaultContentKey
	durationKey  = defaultDurationKey
	levelKey     = defaultLevelKey
	spanKey      = defaultSpanKey
	timestampKey = defaultTimestampKey
	traceKey     = defaultTraceKey
	truncatedKey = defaultTruncatedKey
)
