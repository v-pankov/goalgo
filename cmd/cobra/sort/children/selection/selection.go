package selection

import (
	"fmt"

	"github.com/spf13/cobra"

	"goalgo/cmd/cobra/sort/children"
	"goalgo/internal/sort"
)

func New() *cobra.Command {
	return &cobra.Command{
		Use:     "selection",
		Aliases: []string{"select"},
		Short:   "Sorts input array using selection sort algorithm.",
		Long:    "Sorts input array using selection sort algorithm.",
		RunE: children.RunE(func(ints []int) error {
			sort.Selection(ints)
			fmt.Println(ints)
			return nil
		}),
	}
}
