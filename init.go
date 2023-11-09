package main

import (
	"github.com/labstack/gommon/log"
	"github.com/qiuzhanghua/common/util"
	"os"
	"strings"
)

func init() {
	log.SetPrefix("tz")
	level := os.Getenv("LOGGING_LEVEL")
	if level == "" {
		level = "off"
	}
	format := strings.ToLower(os.Getenv("LOGGING_FORMAT"))
	if format != "json" {
		log.SetHeader(`${time_rfc3339_nano}, ${prefix}, ${level} ${short_file}(${line})`)
	}
	log.SetLevel(util.LevelOf(level))
	log.SetOutput(os.Stdout)
}
