package main

import (
	"context"
	"embed"
	"flag"
	"fmt"
	"os"

	"bifromq_engine/pkg/db"
	"bifromq_engine/pkg/logs"
	"bifromq_engine/pkg/routes"
	"bifromq_engine/pkg/server"
	"bifromq_engine/pkg/signals"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

//go:embed ui/dist/*
var buildFS embed.FS

func main() {
	logs.InitLogger()
	ctx := signals.SetupSignalHandler()
	if err := NewAppCommand(ctx).Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func NewAppCommand(ctx context.Context) *cobra.Command {
	option := server.NewServerConfig()
	cmd := &cobra.Command{
		Use:  "serve",
		Long: `serve`,
		Run: func(cmd *cobra.Command, args []string) {
			logs.Infof("run with option:%+v", option)
			if err := Run(ctx, option); err != nil {
				logs.Fatal(err)
			}
		},
	}
	cmd.Flags().AddGoFlagSet(flag.CommandLine)
	option.AddFlags(cmd.Flags())
	return cmd
}

func Run(ctx context.Context, option *server.ServerConfig) error {
	// Init DB
	err := db.InitDB(option.DNS)
	if err != nil {
		return err
	}
	// Init Routes
	gin := gin.Default()
	routes.InitRoutes(gin, buildFS)

	// init cordinator
	coordinator, err2 := server.InitCoordinator(option.CooridnatorPort)
	if err2 != nil {
		return err2
	}

	go coordinator.Run(ctx.Done())

	gin.Run(fmt.Sprintf(":%d", option.APIPort))
	return nil
}
