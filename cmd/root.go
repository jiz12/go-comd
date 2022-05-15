package cmd

import (
	"fmt"
	"github.com/jiz12/go-comd/core/logger"
	"github.com/jiz12/go-comd/gen"
	"github.com/spf13/cobra"
	"os"
	"runtime"
)

var rootCmd = &cobra.Command{
	Use:   "go-comd",
	Short: "A cli tool to generate template code",
	Long:  "A cli tool to generate code",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logger.Logger.Error(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Version = fmt.Sprintf("%s %s/%s", "1.1.0",
		runtime.GOOS, runtime.GOARCH)
	rootCmd.AddCommand(gen.Cmd)
}
