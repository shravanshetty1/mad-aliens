package world

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
)

const DestructionThreshold = 1
const AlienPrefix = "alien"

type World struct {
	Cities map[string]*City
	Aliens map[string]*Alien
}

// Update updates the current world state.
// Triggers all aliens to move and then checks if any city has been destroyed.
func (w *World) Update() []string {
	var events []string
	for _, alien := range w.Aliens {
		event := alien.Move()
		if event != "" {
			events = append(events, event)
		}
	}

	for _, city := range w.Cities {
		aliens := city.ListAliens()
		if len(aliens) > DestructionThreshold {
			var alienNames []string
			for _, a := range aliens {
				alienNames = append(alienNames, a.Name)
			}

			events = append(events, "City "+city.Name+" has been destroyed by aliens "+strings.Join(alienNames, ", "))
			w.DestroyCity(city)
		}
	}

	return events
}

func (w *World) DestroyCity(city *City) {
	delete(w.Cities, city.Name)
	for _, alien := range city.Aliens {
		delete(w.Aliens, alien.Name)
	}

	for _, connectedCity := range city.ConnectedCities {
		delete(connectedCity.ConnectedCities, city.Name)
		delete(connectedCity.ConnectedCityDirection, city.Name)
	}
}

func (w *World) ListCities() []*City {
	var cities []*City
	for _, city := range w.Cities {
		cities = append(cities, city)
	}

	return cities
}

func (w *World) GetMap() string {
	var mapContent string
	for cityName, city := range w.Cities {
		lineParts := []string{cityName}
		for connectedCityName, direction := range city.ConnectedCityDirection {
			lineParts = append(lineParts, direction+"="+connectedCityName)
		}
		mapContent += strings.Join(lineParts, " ") + "\n"
	}

	return mapContent
}

func (w *World) SpawnAliens(aliens int) error {
	cities := w.ListCities()
	alienMap := make(map[string]*Alien)
	bound := len(cities)

	if bound < 1 {
		return fmt.Errorf("no cities for aliens to spawn in")
	}

	for i := 1; i < aliens+1; i++ {
		name := AlienPrefix + fmt.Sprint(i)
		city := cities[rand.Intn(bound)]
		alien := &Alien{
			Name:     name,
			Location: city,
		}
		alienMap[name] = alien
		city.Aliens[name] = alien
	}

	w.Aliens = alienMap
	return nil
}

// ParseMap parses given file into world object which is the internal representation
// of the world in the program
func ParseMap(pathToMap string) (*World, error) {
	content, err := ioutil.ReadFile(pathToMap)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	cityToCityPaths := make(map[string][]string)
	cities := make(map[string]*City)
	for _, line := range lines {
		lineParts := strings.Split(line, " ")
		city := lineParts[0]

		if len(lineParts) > 1 {
			cityToCityPaths[city] = lineParts[1:]
		}

		cities[city] = &City{
			Name:                   city,
			ConnectedCities:        make(map[string]*City),
			ConnectedCityDirection: make(map[string]string),
			Aliens:                 make(map[string]*Alien),
		}
	}

	for city, paths := range cityToCityPaths {
		for _, path := range paths {
			arr := strings.Split(path, "=")
			if len(arr) != 2 {
				return nil, fmt.Errorf("invalid path in map, got path - " + path)
			}
			direction := arr[0]
			connectedCity := arr[1]
			cities[city].ConnectedCities[connectedCity] = cities[connectedCity]
			cities[city].ConnectedCityDirection[connectedCity] = direction
		}
	}

	return &World{
		Cities: cities,
	}, nil
}
