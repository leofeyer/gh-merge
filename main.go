package main

import (
	"fmt"
	"os"

	"github.com/leofeyer/gh-merge/api"
	"github.com/spf13/cobra"
)

func rootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gh merge <number>",
		Short: "Merge a pull request",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			number, err := api.GetNumber(args)
			if err != nil {
				return err
			}

			info, err := api.GetInfo(number)
			if err != nil {
				return err
			}

			if err := api.MergePr(number, info); err != nil {
				return err
			}

			if err := api.ThankAuthor(number, info); err != nil {
				fmt.Fprintf(os.Stderr, "Warning: failed to thank author: %s\n", err)
			}

			return nil
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	return cmd
}

func main() {
	if err := rootCmd().Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
