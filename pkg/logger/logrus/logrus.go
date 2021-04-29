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

func (l *logBundle) Info(field *logger.LogField) {
	l.logger.WithFields(logrus.Fields{
		"section":  field.Section,
		"function": field.Function,
		"params":   field.Params,
	}).Info(field.Message)
}

func (l *logBundle) Warning(field *logger.LogField) {
	l.logger.WithFields(logrus.Fields{
		"section":  field.Section,
		"function": field.Function,
		"params":   field.Params,
	}).Warning(field.Message)
}

func (l *logBundle) Error(field *logger.LogField) {
	l.logger.WithFields(logrus.Fields{
		"section":  field.Section,
		"function": field.Function,
		"params":   field.Params,
	}).Error(field.Message)
}
