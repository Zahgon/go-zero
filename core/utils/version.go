package utils

import (
	"github.com/zeromicro/go-zero/core/stringx"
)

var replacer = stringx.NewReplacer(map[string]string{
	"V": "",
	"v": "",
	"-": ".",
})

func CompareVersions(v1, op, v2 string) bool { _ = "STUB: not implemented"; return false }

func compare(v1, v2 string) int { _ = "STUB: not implemented"; return 0 }

func strsToInts(strs []string) []int64 { _ = "STUB: not implemented"; return nil }
