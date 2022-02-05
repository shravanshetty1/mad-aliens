package world

type City struct {
	Name                   string
	ConnectedCityDirection map[string]string
	ConnectedCities        map[string]*City
	Aliens                 map[string]*Alien
}

func (c *City) ListAliens() []*Alien {
	var aliens []*Alien
	for _, alien := range c.Aliens {
		aliens = append(aliens, alien)
	}

	return aliens
}

func (c *City) ListConnectedCities() []*City {
	var connectedCities []*City
	for _, connectedCity := range c.ConnectedCities {
		connectedCities = append(connectedCities, connectedCity)
	}

	return connectedCities
}
