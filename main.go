package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "testvol",
	Short: "testvol - volume plugin for Podman",
	Long:  `Creates simple directory volumes using the Volume Plugin API for testing volume plugin functionality`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Errorf("Error running volume plugin: %v")
		os.Exit(1)
	}

	os.Exit(0)
}
