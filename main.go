package main

import (
	"fmt"
	"os"

	"github.com/leofeyer/gh-merge/api"
	"github.com/spf13/cobra"
)

type Options struct {
	Auto  bool
	Admin bool
}

func rootCmd() *cobra.Command {
	opts := &Options{}

	cmd := &cobra.Command{
		Use:   "gh merge <number>",
		Short: "Merge a pull request",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			number, err := api.GetNumber(args)
			if err != nil {
				return err
			}

			err = api.MergePr(number, opts.Auto, opts.Admin)
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

	cmd.Flags().BoolVarP(&opts.Auto, "auto", "", false, "enable auto-merging the pull request")
	cmd.Flags().BoolVarP(&opts.Admin, "admin", "", false, "merge the pull request with admin rights")

	return cmd
}

func main() {
	err := rootCmd().Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
