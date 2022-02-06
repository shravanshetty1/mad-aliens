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

//
//func Test_Update(t *testing.T) {
//	const seed = 2
//
//	city1 := &City{
//		Name:                   "York1",
//		ConnectedCities:        make(map[string]*City),
//		ConnectedCityDirection: make(map[string]string),
//		Aliens:                 make(map[string]*Alien),
//	}
//	city2 := &City{
//		Name:                   "York2",
//		ConnectedCities:        make(map[string]*City),
//		ConnectedCityDirection: make(map[string]string),
//		Aliens:                 make(map[string]*Alien),
//	}
//	city3 := &City{
//		Name:                   "York3",
//		ConnectedCities:        make(map[string]*City),
//		ConnectedCityDirection: make(map[string]string),
//		Aliens:                 make(map[string]*Alien),
//	}
//	city1.ConnectedCities[city2.Name] = city2
//	city1.ConnectedCities[city3.Name] = city3
//	city2.ConnectedCities[city3.Name] = city1
//	city2.ConnectedCities[city3.Name] = city3
//	city3.ConnectedCities[city3.Name] = city1
//	city3.ConnectedCities[city3.Name] = city2
//
//	wrld := &World{
//		Cities: map[string]*City{
//			city1.Name: city1,
//			city2.Name: city2,
//			city3.Name: city3,
//		},
//	}
//
//	rand.Seed(seed)
//	err := wrld.SpawnAliens(2)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	wrld.Update()
//
//	require.Equal(t, wrld.Aliens["alien1"].Location.Name, city3.Name)
//	require.Equal(t, wrld.Aliens["alien2"].Location.Name, city2.Name)
//}
