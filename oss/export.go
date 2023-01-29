package oss

import (
	"github.com/reedchan7/aliyungo/util"
)

var originalStrategy = attempts

func SetAttemptStrategy(s *util.AttemptStrategy) {
	if s == nil {
		attempts = originalStrategy
	} else {
		attempts = *s
	}
}

func SetListPartsMax(n int) {
	listPartsMax = n
}

func SetListMultiMax(n int) {
	listMultiMax = n
}
