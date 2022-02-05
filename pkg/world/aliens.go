package world

import "math/rand"

type Alien struct {
	Name     string
	Location *City
}

// TODO if alien did not move it should not print an event
func (a *Alien) Move() {
	prevCity := a.Location
	delete(prevCity.Aliens, a.Name)

	connectedCities := prevCity.ListConnectedCities()
	if len(connectedCities) < 1 {
		return
	}
	currCity := connectedCities[rand.Intn(len(connectedCities))]

	a.Location = currCity
	currCity.Aliens[a.Name] = a
}
