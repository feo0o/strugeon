package cmd

import (
    "github.com/feo0o/sturgeon/app"
    "github.com/sirupsen/logrus"
    "github.com/spf13/cobra"
    "strings"
)

var rootCmd = &cobra.Command{
    Use:   strings.ToLower(app.Name),
    Short: app.Description,
    Run: func(c *cobra.Command, args []string) {
        // todo
    },
}

func init() {
    // todo
}

func Exec() {
    if err := rootCmd.Execute(); err != nil {
        logrus.Fatal(err)
    }
}
