package app

import (
	"bifromq_engine/pkg/db"
	"bifromq_engine/pkg/logs"
	"bifromq_engine/pkg/routes"
	"bifromq_engine/pkg/server"
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

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
	//staticFp, _ := fs.Sub(staticFiles, "ui")
	//gin.StaticFS("/ui", http.FS(staticFp))
	routes.InitRoutes(gin)

	// init cordinator
	coordinator, err2 := server.InitCoordinator(option.CooridnatorPort)
	if err2 != nil {
		return err2
	}

	go coordinator.Run(ctx.Done())

	gin.Run(fmt.Sprintf(":%d", option.APIPort))
	return nil
}
