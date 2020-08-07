package core

import (
	"fmt"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

func LogrusJsonFmt() *logrus.JSONFormatter {
	return &logrus.JSONFormatter{
		CallerPrettyfier: LogrusFormatter,
	}
}

func LogrusTextFmt() *logrus.TextFormatter {
	return &logrus.TextFormatter{
		CallerPrettyfier: LogrusFormatter,
	}
}

func LogrusFormatter(f *runtime.Frame) (string, string) {
	filename := path.Base(f.File)
	function := path.Base(f.Function)
	return fmt.Sprintf("%s()", function), fmt.Sprintf("%s:%d", filename, f.Line)
}
