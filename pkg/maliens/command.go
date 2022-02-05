package maliens

import (
	"fmt"

	"github.com/shravanshetty1/mad-aliens/pkg/world"

	"github.com/spf13/cobra"
)

const MAX_ITERATIONS = 10000

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "maliens",
		Short:   "Simulates an alien invasion on a world. Map of the world and number of aliens are the inputs.",
		Example: "maliens ./map.txt --aliens 3",
	}

	var numberOfAliens int
	cmd.Flags().IntVar(&numberOfAliens, "aliens", 0, "number of aliens invading the world")

	cmd.RunE = func(cmd *cobra.Command, args []string) (err error) {
		if len(args) < 1 || args[0] == "" {
			return fmt.Errorf("path to world map is required")
		}

		mapFile := args[0]
		wrld, err := world.ParseMap(mapFile)
		if err != nil {
			return err
		}
		wrld.SpawnAliens(numberOfAliens)

		for i := 0; i < MAX_ITERATIONS; i++ {
			events := wrld.Update()

			for _, e := range events {
				fmt.Println(e)
			}

			if len(wrld.ListCities()) < 1 {
				fmt.Println("Invasion successful - All the cities have been destroyed")
				return nil
			}
		}

		fmt.Println("---")
		fmt.Println("Invasion failed - cities still exist after " + fmt.Sprint(MAX_ITERATIONS) + " iterations, following is the map of the remains")
		fmt.Println("---")
		fmt.Println(wrld.GetMap())
		fmt.Println("---")

		return nil
	}

	return cmd
}
