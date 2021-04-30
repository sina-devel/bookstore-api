package logrus

import (
	"errors"
	"io"
	"path/filepath"

	"github.com/alecthomas/units"
	"github.com/kianooshaz/bookstore-api/pkg/logger"
	rotators "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"github.com/xhit/go-str2duration/v2"
)

var ErrNilOption = errors.New("option can not be nil")

type logBundle struct {
	logger *logrus.Logger
}

type Option struct {
	Path, Pattern, MaxAge, RotationTime, RotationSize string
}

//New is constructor of the logrus package
func New(opt *Option) (logger.Logger, error) {
	if opt == nil {
		return nil, ErrNilOption
	}
	l := &logBundle{logger: logrus.New()}
	writer, err := getLoggerWriter(opt)
	if err != nil {
		return nil, err
	}
	l.logger.SetOutput(writer)
	l.logger.SetFormatter(&logrus.JSONFormatter{})

	return l, nil
}


//getLoggerWriter return io.Writer which can create different
//files with custom names at different time intervals
func getLoggerWriter(opt *Option) (io.Writer, error) {

	maxAge, err := str2duration.ParseDuration(opt.MaxAge)
	if err != nil {
		return nil, err
	}

	rotationTime, err := str2duration.ParseDuration(opt.RotationTime)
	if err != nil {
		return nil, err
	}

	rotationSize, err := units.ParseBase2Bytes(opt.RotationSize)
	if err != nil {
		return nil, err
	}

	return rotators.New(
		filepath.Join(opt.Path, opt.Pattern),
		rotators.WithLinkName(opt.Path),
		rotators.WithMaxAge(maxAge),
		rotators.WithRotationTime(rotationTime),
		rotators.WithRotationSize(int64(rotationSize)),
	)
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
