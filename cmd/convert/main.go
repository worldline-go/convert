package main

import (
	"context"
	"os"
	"sync"

	"github.com/spf13/cobra"
	"github.com/worldline-go/initializer"
	"github.com/worldline-go/logz"

	"github.com/worldline-go/convert/internal/config"
	"github.com/worldline-go/convert/internal/convert"
)

var (
	version = "v0.0.0"
	commit  = "-"
	date    = "-"
)

func main() {
	initializer.Init(
		runCommand,
		initializer.WithInitLog(false),
		initializer.WithOptionsLogz(logz.WithCaller(false)),
	)
}

var rootCmd = &cobra.Command{
	Use:           "convert",
	Short:         "file convertor",
	Long:          "convert from excel to text format",
	SilenceUsage:  true,
	SilenceErrors: true,
	Example:       "convert -i input.xlsx -o output.txt",
	RunE: func(cmd *cobra.Command, args []string) error {
		configFile := os.Getenv("CONFIG_FILE")
		if configFile != "" {
			if err := config.Load(configFile); err != nil {
				return err
			}
		}

		if values.Input != "" {
			config.AppConfig.Input = values.Input
		}

		return run(cmd.Context())
	},
}

var values = struct {
	Input string `cfg:"input"`
}{
	Input: "",
}

func init() {
	rootCmd.Flags().StringVarP(&values.Input, "input", "i", values.Input, "input file")
}

func runCommand(ctx context.Context, _ *sync.WaitGroup) error {
	rootCmd.Version = version
	rootCmd.Long += "\nversion: " + version + " commit: " + commit + " buildDate:" + date

	return rootCmd.ExecuteContext(ctx)
}

func run(ctx context.Context) error {
	if err := convert.Convert(&config.AppConfig); err != nil {
		return err
	}

	return nil
}
