package bubble

import (
	"fmt"

	"github.com/spf13/cobra"

	"goalgo/cmd/cobra/sort/children"
	"goalgo/internal/sort"
)

func New() *cobra.Command {
	var (
		sortFn      func([]int)
		sortVersion int
	)

	cmd := &cobra.Command{
		Use:   "bubble",
		Short: "Sorts input array using bubble sort algorithm.",
		Long:  "Sorts input array using bubble sort algorithm.",
		RunE: children.RunE(func(ints []int) error {
			switch sortVersion {
			case 0:
				sortFn = sort.Bubble[int]
			case 1:
				sortFn = sort.Bubble1[int]
			default:
				return fmt.Errorf("unsupported bubble sort version [%d]", sortVersion)
			}

			sortFn(ints)
			fmt.Println(ints)
			return nil
		}),
	}

	cmd.Flags().IntVarP(
		&sortVersion,
		"version", "v",
		0,
		`set bubble sort version:
	0 - simple,
	1 - optimized`)

	return cmd
}
