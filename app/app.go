package app

import (
    "github.com/sirupsen/logrus"
    "os"
    "strings"
    "time"
)

var (
    Name        = "Sturgeon"
    Description = "Tools for check resources and health status for Kubernetes."
)

func init() {
    // init default logger
    logrus.SetLevel(logrus.InfoLevel)
    logrus.SetFormatter(&logrus.TextFormatter{TimestampFormat: time.RFC3339})
    enableDebug := strings.ToLower(os.Getenv(strings.ToUpper(Name) + "_ENABLE_DEBUG"))
    if enableDebug == "yes" || enableDebug == "y" {
        logrus.SetLevel(logrus.DebugLevel)
        logrus.SetReportCaller(true)
    }
}
