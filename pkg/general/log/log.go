package log

import "github.com/sirupsen/logrus"

func init() {

	//logrus.Logger{}
	logrus.SetReportCaller(true)
}
