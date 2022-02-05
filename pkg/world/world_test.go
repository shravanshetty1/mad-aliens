package world

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ParseMap(t *testing.T) {
	tcs := []struct {
		MapName string
		Err     string
	}{
		{
			MapName: "map.txt",
			Err:     "",
		},
		{
			MapName: "invalid.txt",
			Err:     "invalid path",
		},
	}

	path, err := filepath.Abs("../../assets/")
	if err != nil {
		t.Fatal(err)
	}

	for _, tc := range tcs {
		_, err = ParseMap(filepath.Join(path, tc.MapName))
		if tc.Err == "" {
			require.Nil(t, err)
		} else {
			require.NotNil(t, err)
			require.Contains(t, err.Error(), tc.Err)
		}
	}
}

func Test_SpawnAliens(t *testing.T) {
	tcs := []struct {
		World          *World
		NumberOfAliens int
		Err            string
	}{
		{
			World: &World{
				Cities: map[string]*City{
					"York": &City{
						Name:   "York",
						Aliens: make(map[string]*Alien),
					},
					"York2": &City{
						Name:   "York2",
						Aliens: make(map[string]*Alien),
					},
				},
				Aliens: nil,
			},
			NumberOfAliens: 3,
			Err:            "",
		},
		{
			World: &World{
				Cities: nil,
				Aliens: nil,
			},
			NumberOfAliens: 3,
			Err:            "no cities",
		},
	}

	for _, tc := range tcs {
		err := tc.World.SpawnAliens(tc.NumberOfAliens)
		if tc.Err == "" {
			require.Nil(t, err)
		} else {
			require.NotNil(t, err)
			require.Contains(t, err.Error(), tc.Err)
		}
	}
}
