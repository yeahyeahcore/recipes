package main

import (
	"fmt"
	"path"
	"runtime"
	"strconv"

	formatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"

	"github.com/yeahyeahcore/recipes/internal/app"
	"github.com/yeahyeahcore/recipes/internal/models"
	"github.com/yeahyeahcore/recipes/pkg/env"
)

func main() {
	logger := logrus.New()

	logger.SetReportCaller(true)
	logger.SetFormatter(&formatter.Formatter{
		TimestampFormat: "02-01-2006 15:04:05",
		HideKeys:        true,
		NoColors:        false,
		NoFieldsSpace:   true,
		CustomCallerFormatter: func(frame *runtime.Frame) string {
			return fmt.Sprintf(" (%s:%s)", path.Base(frame.File), strconv.Itoa(frame.Line))
		},
	})

	config, err := env.Parse[models.Config](".env")
	if err != nil {
		logger.Fatalf("config read is failed: %s", err)
	}

	app.Run(config, logger)
}
