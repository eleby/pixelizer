package main

import (
	log "github.com/sirupsen/logrus"
)

const (
	unset     = "unset"
	pixelizer = "pixelizer"
)

var ( // build info
	version   = unset
	date      = unset
	commit    = unset
	appname   = pixelizer
	goversion = unset
)

func printVersion() {
	log.WithFields(log.Fields{
		"version":    version,
		"build_time": date,
		"commit":     commit,
		"app_name":   appname,
		"go_version": goversion,
	}).Info("Build info")
}
