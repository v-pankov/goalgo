package sort

import (
	"github.com/spf13/cobra"

	"goalgo/cmd/cobra/sort/children/bubble"
	"goalgo/cmd/cobra/sort/children/insertion"
	"goalgo/cmd/cobra/sort/children/selection"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sort",
		Short: "Sorts input array according to specified algorithm.",
		Long:  "Sorts input array according to specified algorithm.",
	}

	cmd.AddCommand(insertion.New())
	cmd.AddCommand(selection.New())
	cmd.AddCommand(bubble.New())

	return cmd
}
