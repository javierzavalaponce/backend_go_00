package logging_service

import (
	//re_services "artemisa/src/core/domain/services"
	core_services "github.com/DEINSI-DEVELOP/test_backend_go.git/src/core/domain/services"

	"github.com/sirupsen/logrus"
)

type LoggingServiceLogrus struct {
	logger *logrus.Logger
}

func NewLoggingServiceLogrus() core_services.LoggingService {
	logger := logrus.New()

	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logger.SetLevel(logrus.DebugLevel)

	return &LoggingServiceLogrus{
		logger: logger,
	}
}

func (l *LoggingServiceLogrus) Info(message string, fields map[string]interface{}) {
	l.logger.WithFields(logrus.Fields(fields)).Info(message)
}

func (l *LoggingServiceLogrus) Warn(message string, fields map[string]interface{}) {
	l.logger.WithFields(logrus.Fields(fields)).Warn(message)
}

func (l *LoggingServiceLogrus) Error(message string, fields map[string]interface{}) {
	l.logger.WithFields(logrus.Fields(fields)).Error(message)
}

func (l *LoggingServiceLogrus) Debug(message string, fields map[string]interface{}) {
	l.logger.WithFields(logrus.Fields(fields)).Debug(message)
}

func (l *LoggingServiceLogrus) Write(p []byte) (n int, err error) {
	l.logger.Debug(string(p))
	return len(p), nil
}
