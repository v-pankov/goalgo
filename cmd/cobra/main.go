package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"goalgo/cmd/cobra/sort"
)

func main() {
	root := &cobra.Command{
		Use:   "goalgo",
		Short: "execute algorithm",
		Long:  "execute algorithm",
	}

	root.AddCommand(sort.New())

	if err := root.Execute(); err != nil {
		fmt.Println(err)
	}
}
