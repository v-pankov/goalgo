package bubble

import (
	"fmt"

	"github.com/spf13/cobra"

	"goalgo/cmd/cobra/sort/children"
	"goalgo/internal/sort"
)

func New() *cobra.Command {
	return &cobra.Command{
		Use:   "bubble",
		Short: "Sorts input array using bubble sort algorithm.",
		Long:  "Sorts input array using bubble sort algorithm.",
		RunE: children.RunE(func(ints []int) error {
			sort.Bubble(ints)
			fmt.Println(ints)
			return nil
		}),
	}
}
