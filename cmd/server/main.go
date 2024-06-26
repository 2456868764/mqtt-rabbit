package main

import (
	"bifromq_engine/cmd/server/app"
	"fmt"
	"os"

	"bifromq_engine/pkg/logs"
	"bifromq_engine/pkg/signals"
)

//var (
//	//go:embed ui/*
//	staticFiles embed.FS
//)

func main() {
	logs.InitLogger()
	ctx := signals.SetupSignalHandler()
	if err := app.NewAppCommand(ctx).Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
