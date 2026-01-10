package logging_service

import (
	//re_services "artemisa/src/core/domain/services"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"

	core_services "github.com/DEINSI-DEVELOP/test_backend_go.git/src/core/domain/services"
)

type LoggingServiceGolog struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
	debugLogger *log.Logger
	warnLogger  *log.Logger
}

func NewLoggingServiceGolog() core_services.LoggingService {
	return &LoggingServiceGolog{
		infoLogger:  log.New(os.Stdout, "INFO:  ", log.Ldate|log.Ltime),
		warnLogger:  log.New(os.Stdout, "WARN:  ", log.Ldate|log.Ltime),
		errorLogger: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime),
		debugLogger: log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime),
	}
}

func (l *LoggingServiceGolog) Info(message string, fields map[string]interface{}) {
	l.infoLogger.Println(l.formatLog(message, fields))
}

func (l *LoggingServiceGolog) Warn(message string, fields map[string]interface{}) {
	l.warnLogger.Println(l.formatLog(message, fields))
}

func (l *LoggingServiceGolog) Error(message string, fields map[string]interface{}) {
	l.errorLogger.Println(l.formatLog(message, fields))
}

func (l *LoggingServiceGolog) Debug(message string, fields map[string]interface{}) {
	l.debugLogger.Println(l.formatLog(message, fields))
}

func (l *LoggingServiceGolog) Write(p []byte) (n int, err error) {
	l.debugLogger.Println(string(p))
	return len(p), nil
}

func (l *LoggingServiceGolog) formatLog(message string, fields map[string]interface{}) string {
	// Get caller information
	callerInfo := getCallerInfo(3) // Adjust depth to get the correct caller
	if len(fields) == 0 {
		return fmt.Sprintf("%s [%s]", message, callerInfo)
	}

	fieldStrings := make([]string, 0, len(fields))
	for key, value := range fields {
		fieldStrings = append(fieldStrings, fmt.Sprintf("%s=%v", key, value))
	}
	return fmt.Sprintf("%s | %s [%s]", message, strings.Join(fieldStrings, ", "), callerInfo)
}

func getCallerInfo(depth int) string {
	// Get the file and line number of the caller
	_, file, line, ok := runtime.Caller(depth)
	if !ok {
		return "unknown:0"
	}

	// Debug: Print the caller details
	fmt.Printf("%s:%d ", file, line)

	// Extract the base file name
	shortFile := file
	if lastSlash := strings.LastIndex(file, "/"); lastSlash >= 0 {
		shortFile = file[lastSlash+1:]
	}
	return fmt.Sprintf("%s:%d", shortFile, line)
}
