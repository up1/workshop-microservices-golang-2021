package main

import (
	"io"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func main() {
	log.SetFormatter(&logrus.JSONFormatter{})
	// Write to file
	path := "./sample.log"
	logfile, _ := rotatelogs.New(
		path+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithMaxAge(time.Duration(180)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(60)*time.Second),
	)
	// Set multiple output for log
	mw := io.MultiWriter(os.Stdout, logfile)
	log.SetOutput(mw)

	standardFields := logrus.Fields{
		"hostname": "staging-1",
		"appname":  "foo-app",
		"session":  "1ce3f6v",
	}

	log.WithFields(standardFields).WithFields(
		logrus.Fields{
			"string": "foo",
			"int":    1,
			"float":  1.1,
		}).Info("My first event from Golang")

}
