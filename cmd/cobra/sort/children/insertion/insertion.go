package insertion

import (
	"fmt"

	"github.com/spf13/cobra"

	"goalgo/cmd/cobra/sort/children"
	"goalgo/internal/sort"
)

func New() *cobra.Command {
	return &cobra.Command{
		Use:     "insertion",
		Aliases: []string{"insert"},
		Short:   "Sorts input array using insertion sort algorithm.",
		Long:    "Sorts input array using insertion sort algorithm.",
		RunE: children.RunE(func(ints []int) error {
			sort.Insertion(ints)
			fmt.Println(ints)
			return nil
		}),
	}
}
