package space

// Planet is the default type in our system
type Planet string

var planets = map[Planet]float64{
	"Earth":   1.0,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age will take a value in seconds + a planet name and return the age on that planet
func Age(age float64, planet Planet) float64 {
	if _, ok := planets[planet]; !ok {
		return 0.00
	}

	return age / (60.0 * 60.0 * 24.0 * 365.25 * planets[planet])
}
