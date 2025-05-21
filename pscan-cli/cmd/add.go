/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/itsjayeshrathi/pscan-cli/scan"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <host1>...<hostn>",
	Short: "Add host(s) in the list",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		hostFile, err := cmd.Flags().GetString("hosts-file")
		if err != nil {
			return err
		}
		return addAction(os.Stdout, hostFile, args)
	},
	Aliases:      []string{"a"},
	Args:         cobra.MinimumNArgs(1),
	SilenceUsage: true,
}

func addAction(out io.Writer, hostFile string, args []string) error {
	hl := &scan.HostsList{}

	if err := hl.Load(hostFile); err != nil {
		return err
	}
	for _, h := range args {

		if err := hl.Add(h); err != nil {
			return err
		}
		fmt.Fprintln(out, "Added host:", h)
	}
	return hl.Save(hostFile)
}

func init() {
	hostsCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
