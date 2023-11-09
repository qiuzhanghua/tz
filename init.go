package main

import (
	"github.com/qiuzhanghua/common/log"
	"os"
)

func init() {
	level := os.Getenv("LOGGING_LEVEL")
	if len(level) == 0 {
		level = "info"
	}
	log.SetGlobalLevel(level)
}
