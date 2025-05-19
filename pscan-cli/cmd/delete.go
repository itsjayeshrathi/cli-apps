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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete <host1>...<hostn>",
	Short: "Remove host(s) in the list",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		hostFile, err := cmd.Flags().GetString("host-file")
		if err != nil {
			return err
		}
		return deleteAction(os.Stdout, hostFile, args)
	},
	Aliases:      []string{"d"},
	SilenceUsage: true,
	Args:         cobra.MinimumNArgs(1),
}

func deleteAction(out io.Writer, hostFile string, args []string) error {
	hl := &scan.HostsList{}

	if err := hl.Load(hostFile); err != nil {
		return err
	}
	for _, h := range hl.Hosts {
		if err := hl.Remove(h); err != nil {
			return err
		}
		fmt.Println(out, "Deleted host:", h)
	}
	return hl.Save(hostFile)
}

func init() {
	hostsCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
