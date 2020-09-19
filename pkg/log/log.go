package log

import (
log	"github.com/sirupsen/logrus"
	//"log"
	"os"
)

//var log = logrus.New()

func init() {

	
	log.SetFormatter(&log.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)
	log.SetReportCaller(true)
	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}
