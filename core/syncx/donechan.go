package syncx

import (
	"sync"

	"github.com/zeromicro/go-zero/core/lang"
)

type DoneChan struct {
	done chan lang.PlaceholderType
	once sync.Once
}

func NewDoneChan() *DoneChan { _ = "STUB: not implemented"; return nil }

func (dc *DoneChan) Close() { _ = "STUB: not implemented"; return }

func (dc *DoneChan) Done() chan lang.PlaceholderType { _ = "STUB: not implemented"; return nil }
