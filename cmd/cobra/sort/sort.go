package sort

import (
	"github.com/spf13/cobra"

	"goalgo/cmd/cobra/sort/children/insertion"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sort",
		Short: "Sorts input array according to specified algorithm.",
		Long:  "Sorts input array according to specified algorithm.",
	}

	cmd.AddCommand(insertion.New())

	return cmd
}
