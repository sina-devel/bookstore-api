package logrus

import (
	"github.com/alecthomas/units"
	"github.com/kianooshaz/bookstore-api/pkg/logger"
	rotators "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"github.com/xhit/go-str2duration/v2"
	"io"
)

type logBundle struct {
	logger *logrus.Logger
}

//New is Constructor of the logrus package
func New(path, pattern, maxAgeStr, rotationTimeStr, rotationSizeStr string) (logger.Logger, error) {
	l := &logBundle{logger: logrus.New()}
	writer, err := getLoggerWriter(path, pattern, maxAgeStr, rotationTimeStr, rotationSizeStr)
	if err != nil {
		return nil, err
	}
	l.logger.SetOutput(writer)
	l.logger.SetFormatter(&logrus.JSONFormatter{})

	return l, nil
}

//getLoggerWriter return io.Writer which can create different
//files with custom names at different time intervals
func getLoggerWriter(path, pattern, maxAgeStr, rotationTimeStr, rotationSizeStr string) (io.Writer, error) {

	maxAge, err := str2duration.ParseDuration(maxAgeStr)
	if err != nil {
		return nil, err
	}

	rotationTime, err := str2duration.ParseDuration(rotationTimeStr)
	if err != nil {
		return nil, err
	}

	rotationSize, err := units.ParseBase2Bytes(rotationSizeStr)
	if err != nil {
		return nil, err
	}

	writer, err := rotators.New(
		path+pattern,
		rotators.WithLinkName(path),
		rotators.WithMaxAge(maxAge),
		rotators.WithRotationTime(rotationTime),
		rotators.WithRotationSize(int64(rotationSize)),
	)
	if err != nil {
		return nil, err
	}

	return writer, nil
}

//Info is logger with level info
func (l *logBundle) Info(field *logger.LogField) {
	l.logger.WithFields(logrus.Fields{
		"section":  field.Section,
		"function": field.Function,
		"params":   field.Params,
	}).Info(field.Message)
}

//Warning is logger with level warning
func (l *logBundle) Warning(field *logger.LogField) {
	l.logger.WithFields(logrus.Fields{
		"section":  field.Section,
		"function": field.Function,
		"params":   field.Params,
	}).Warning(field.Message)
}

//Error is logger with level error
func (l *logBundle) Error(field *logger.LogField) {
	l.logger.WithFields(logrus.Fields{
		"section":  field.Section,
		"function": field.Function,
		"params":   field.Params,
	}).Error(field.Message)
}
