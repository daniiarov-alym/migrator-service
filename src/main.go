package main

import (
	"github.com/daniiarov-alym/migrator-service/src/config"
	"github.com/daniiarov-alym/migrator-service/src/db"
	logger "github.com/sirupsen/logrus"
	"os"
	"context"
)

func main() {
	ctx := context.Background()
	logger.SetOutput(os.Stdout)
	logger.SetReportCaller(true)
	logger.SetLevel(logger.TraceLevel)
	config.Init()
	db.Run(ctx)
}