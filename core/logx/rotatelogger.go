package logx

import (
	"errors"
	"os"
	"sync"
	"time"

	"github.com/zeromicro/go-zero/core/lang"
)

const (
	hoursPerDay     = 24
	bufferSize      = 100
	defaultDirMode  = 0o755
	defaultFileMode = 0o600
	gzipExt         = ".gz"
	megaBytes       = 1 << 20
)

var (
	ErrLogFileClosed = errors.New("error: log file closed")

	fileTimeFormat = time.RFC3339
)

type (
	RotateRule interface {
		BackupFileName() string
		MarkRotated()
		OutdatedFiles() []string
		ShallRotate(size int64) bool
	}

	RotateLogger struct {
		filename string
		backup   string
		fp       *os.File
		channel  chan []byte
		done     chan lang.PlaceholderType
		rule     RotateRule
		compress bool

		waitGroup   sync.WaitGroup
		closeOnce   sync.Once
		currentSize int64
	}

	DailyRotateRule struct {
		rotatedTime string
		filename    string
		delimiter   string
		days        int
		gzip        bool
	}

	SizeLimitRotateRule struct {
		DailyRotateRule
		maxSize    int64
		maxBackups int
	}
)

func DefaultRotateRule(filename, delimiter string, days int, gzip bool) RotateRule {
	_ = "STUB: not implemented"
	return *new(RotateRule)
}

func (r *DailyRotateRule) BackupFileName() string { _ = "STUB: not implemented"; return "" }

func (r *DailyRotateRule) MarkRotated() { _ = "STUB: not implemented"; return }

func (r *DailyRotateRule) OutdatedFiles() []string { _ = "STUB: not implemented"; return nil }

func (r *DailyRotateRule) ShallRotate(_ int64) bool { _ = "STUB: not implemented"; return false }

func NewSizeLimitRotateRule(filename, delimiter string, days, maxSize, maxBackups int, gzip bool) RotateRule {
	_ = "STUB: not implemented"
	return *new(RotateRule)
}

func (r *SizeLimitRotateRule) BackupFileName() string { _ = "STUB: not implemented"; return "" }

func (r *SizeLimitRotateRule) MarkRotated() { _ = "STUB: not implemented"; return }

func (r *SizeLimitRotateRule) OutdatedFiles() []string { _ = "STUB: not implemented"; return nil }

func (r *SizeLimitRotateRule) ShallRotate(size int64) bool { _ = "STUB: not implemented"; return false }

func (r *SizeLimitRotateRule) parseFilename() (prefix, ext string) {
	_ = "STUB: not implemented"
	return "", ""
}

func NewLogger(filename string, rule RotateRule, compress bool) (*RotateLogger, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *RotateLogger) Close() error { _ = "STUB: not implemented"; return nil }

func (l *RotateLogger) Write(data []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (l *RotateLogger) getBackupFilename() string { _ = "STUB: not implemented"; return "" }

func (l *RotateLogger) initialize() error { _ = "STUB: not implemented"; return nil }

func (l *RotateLogger) maybeCompressFile(file string) { _ = "STUB: not implemented"; return }

func (l *RotateLogger) maybeDeleteOutdatedFiles() { _ = "STUB: not implemented"; return }

func (l *RotateLogger) postRotate(file string) { _ = "STUB: not implemented"; return }

func (l *RotateLogger) rotate() error { _ = "STUB: not implemented"; return nil }

func (l *RotateLogger) startWorker() { _ = "STUB: not implemented"; return }

func (l *RotateLogger) write(v []byte) { _ = "STUB: not implemented"; return }

func compressLogFile(file string) { _ = "STUB: not implemented"; return }

func getNowDate() string { _ = "STUB: not implemented"; return "" }

func getNowDateInRFC3339Format() string { _ = "STUB: not implemented"; return "" }

func gzipFile(file string, fsys fileSystem) (err error) { _ = "STUB: not implemented"; return nil }
