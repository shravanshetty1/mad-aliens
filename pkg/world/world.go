package world

import (
	"io/ioutil"
	"strings"
)

type Alien struct {
	Name     string
	Location *City
}

type City struct {
	Name            string
	ConnectedCities map[string]*City
	Aliens          map[string]*Alien
}

type World struct {
	Cities map[string]*City
	Aliens map[string]*Alien
}

func (w *World) Update() []string {
	return nil
}

func (w *World) GetCities() []string {
	return nil
}

func (w *World) GetMap() string {
	return ""
}

func (w *World) SpawnAliens(aliens int) {
}

func ParseMap(pathToMap string) (*World, error) {
	content, err := ioutil.ReadFile(pathToMap)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	cityToAdjCity := make(map[string]*City)
	for _, line := range lines {
		lineParts := strings.Split(line, " ")
		city := lineParts[0]

		var adjCities []string
		for _, paths := range lineParts[0:] {
			connectedCity := strings.Split(paths, "=")[1]
			adjCities = append(adjCities, connectedCity)
		}

		cityToAdjCity[city] = adjCities
	}

	return nil, nil
}
