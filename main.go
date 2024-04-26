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

			err = api.MergePr(number)
			if err != nil {
				return err
			}

			err = api.ThankAuthor(number)
			if err != nil {
				return err
			}

			return nil
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	return cmd
}

func main() {
	err := rootCmd().Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
