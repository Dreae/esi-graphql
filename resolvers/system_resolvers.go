package resolvers

import "github.com/dreae/esi-graphql/resolvers/http"
import "encoding/json"

// SystemStargate serves as a wrapper type for a stargate
// in a solar system
type SystemStargate int32

// PlanetMoon serves as a wrapper type for a moon around
// a planet
type PlanetMoon int32

// SolarSystem holds the details about a solar system in EVE
type SolarSystem struct {
	ConstellationID int32            `json:"constellation_id"`
	Name            string           `json:"name"`
	Planets         []SystemPlanet   `json:"planets"`
	Position        EVEPosition      `json:"position"`
	SecurityClass   string           `json:"security_class"`
	SecurityStatus  float64          `json:"security_status"`
	Stargates       []SystemStargate `json:"stargates"`
	SystemID        int32            `json:"system_id"`
}

type SolarSystemResolver struct {
	system *SolarSystem
}

// SystemPlanet associates a planet to a solar system
type SystemPlanet struct {
	Moons    []PlanetMoon `json:"moons"`
	PlanetID int32        `json:"planet_id"`
}

type SystemPlanetResolver struct {
	planet *SystemPlanet
}

// StargateID returns the ID of this stargate
func (stargate *SystemStargate) StargateID() *int32 {
	return (*int32)(stargate)
}

// MoonID returns the ID of this moon
func (moon *PlanetMoon) MoonID() *int32 {
	return (*int32)(moon)
}

// ConstellationID returns the ID of the constellation this solar system is in
func (r *SolarSystemResolver) ConstellationID() *int32 {
	return &r.system.ConstellationID
}

// Name returns the name of this solar system
func (r *SolarSystemResolver) Name() *string {
	return &r.system.Name
}

// Planets returns an array representing the planets in this solar system
func (r *SolarSystemResolver) Planets() *[]*SystemPlanetResolver {
	var resolvers []*SystemPlanetResolver

	for idx := range r.system.Planets {
		resolvers = append(resolvers, &SystemPlanetResolver{&r.system.Planets[idx]})
	}

	return &resolvers
}

// Position returns this solar system's position in space
func (r *SolarSystemResolver) Position() *PositionResolver {
	return &PositionResolver{&r.system.Position}
}

// SecurityClass returns the class of this solar system
func (r *SolarSystemResolver) SecurityClass() *string {
	return &r.system.SecurityClass
}

// SecurityStatus returns the security status of this solar system
func (r *SolarSystemResolver) SecurityStatus() *float64 {
	return &r.system.SecurityStatus
}

// Stargates returns an array representing the stargates in this system
func (r *SolarSystemResolver) Stargates() *[]*SystemStargate {
	var stargates []*SystemStargate

	for idx := range r.system.Stargates {
		stargates = append(stargates, &r.system.Stargates[idx])
	}

	return &stargates
}

// SystemID returns the ID of this solar system
func (r *SolarSystemResolver) SystemID() *int32 {
	return &r.system.SystemID
}

// PlanetID returns the ID of this planet
func (r *SystemPlanetResolver) PlanetID() *int32 {
	return &r.planet.PlanetID
}

// Moons returns an array representing the moons around this planet
func (r *SystemPlanetResolver) Moons() *[]*PlanetMoon {
	var moons []*PlanetMoon

	for idx := range r.planet.Moons {
		moons = append(moons, &r.planet.Moons[idx])
	}

	return &moons
}

// GetSolarSystemByID queries ESI for a solar system by a given ID
func GetSolarSystemByID(systemID int32) (*SolarSystemResolver, error) {
	resp, err := http.MakeRequest("universe/systems/%d/", systemID)
	if err != nil {
		return nil, err
	}

	var system SolarSystem
	if err := json.NewDecoder(resp.Body).Decode(&system); err != nil {
		return nil, err
	}

	return &SolarSystemResolver{&system}, nil
}
