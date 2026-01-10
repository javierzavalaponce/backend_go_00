package logging_service

import (
	"go.uber.org/zap"
)

type LoggingServiceZap struct {
	logger *zap.Logger
}

func NewLoggingServiceZap() (*LoggingServiceZap, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	return &LoggingServiceZap{logger: logger}, nil
}

func (l *LoggingServiceZap) Info(message string, fields map[string]interface{}) {
	l.logger.With(l.convertFields(fields)...).Info(message)
}

func (l *LoggingServiceZap) Warn(message string, fields map[string]interface{}) {
	l.logger.With(l.convertFields(fields)...).Warn(message)
}

func (l *LoggingServiceZap) Error(message string, fields map[string]interface{}) {
	l.logger.With(l.convertFields(fields)...).Error(message)
}

func (l *LoggingServiceZap) Debug(message string, fields map[string]interface{}) {
	l.logger.With(l.convertFields(fields)...).Debug(message)
}

func (l *LoggingServiceZap) Write(p []byte) (n int, err error) {
	l.logger.Debug(string(p))
	return len(p), nil
}

func (l *LoggingServiceZap) convertFields(fields map[string]interface{}) []zap.Field {
	zapFields := make([]zap.Field, 0, len(fields))
	for key, value := range fields {
		zapFields = append(zapFields, zap.Any(key, value))
	}
	return zapFields
}
