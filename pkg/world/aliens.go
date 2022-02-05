package world

import "math/rand"

type Alien struct {
	Name     string
	Location *City
}

func (a *Alien) Move() string {
	prevCity := a.Location
	delete(prevCity.Aliens, a.Name)

	connectedCities := prevCity.ListConnectedCities()
	if len(connectedCities) < 1 {
		return ""
	}
	currCity := connectedCities[rand.Intn(len(connectedCities))]

	a.Location = currCity
	currCity.Aliens[a.Name] = a

	return "Alien " + a.Name + " moved from " + prevCity.Name + " to " + currCity.Name
}
