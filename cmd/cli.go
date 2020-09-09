package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var (
	persistFlag bool
	PrintFlag   int
	rootCmd     = &cobra.Command{
		Use:          "gifm",
		Short:        "Hello, Gophers!",
		SilenceUsage: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, World!")

			fmt.Println(persistFlag)
		},
	}

	printTimeCmd = &cobra.Command{
		Use: "curtime",
		RunE: func(cmd *cobra.Command, args []string) error {
			now := time.Now()

			prettyTime := now.Format(time.RubyDate)

			cmd.Println("Hey Gophers! The current time is", prettyTime)

			fmt.Println("times", PrintFlag)

			return nil
		},
	}

	echoCmd = &cobra.Command{
		Use:   "echo [string to echo]",
		Short: "prints give string to stdout",
		Args:  cobra.MinimumNArgs(1), // minimum you must pass arguments 1
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Echo", strings.Join(args, " "))
		},
	}
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&persistFlag, "persistFlag", "p", true, "a persistent root flag")
	rootCmd.MarkFlagRequired("persistFlag")

	printTimeCmd.Flags().IntVarP(&PrintFlag, "times", "t", 1, "number of times")

	rootCmd.AddCommand(printTimeCmd) // add cmd sub command from printTimeCmd
	rootCmd.AddCommand(echoCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
