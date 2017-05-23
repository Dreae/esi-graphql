package resolvers

import "github.com/dreae/esi-graphql/resolvers/http"
import "encoding/json"

type SystemStargate int32
type PlanetMoon int32

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

type SystemPlanet struct {
	Moons    []PlanetMoon `json:"moons"`
	PlanetID int32        `json:"planet_id"`
}

type SystemPlanetResolver struct {
	planet *SystemPlanet
}

func (stargate *SystemStargate) StargateID() *int32 {
	return (*int32)(stargate)
}

func (moon *PlanetMoon) MoonID() *int32 {
	return (*int32)(moon)
}

func (r *SolarSystemResolver) ConstellationID() *int32 {
	return &r.system.ConstellationID
}

func (r *SolarSystemResolver) Name() *string {
	return &r.system.Name
}

func (r *SolarSystemResolver) Planets() *[]*SystemPlanetResolver {
	var resolvers []*SystemPlanetResolver

	for idx, _ := range r.system.Planets {
		resolvers = append(resolvers, &SystemPlanetResolver{&r.system.Planets[idx]})
	}

	return &resolvers
}

func (r *SolarSystemResolver) Position() *PositionResolver {
	return &PositionResolver{&r.system.Position}
}

func (r *SolarSystemResolver) SecurityClass() *string {
	return &r.system.SecurityClass
}

func (r *SolarSystemResolver) SecurityStatus() *float64 {
	return &r.system.SecurityStatus
}

func (r *SolarSystemResolver) Stargates() *[]*SystemStargate {
	var stargates []*SystemStargate

	for idx, _ := range r.system.Stargates {
		stargates = append(stargates, &r.system.Stargates[idx])
	}

	return &stargates
}

func (r *SolarSystemResolver) SystemID() *int32 {
	return &r.system.SystemID
}

func (r *SystemPlanetResolver) PlanetID() *int32 {
	return &r.planet.PlanetID
}

func (r *SystemPlanetResolver) Moons() *[]*PlanetMoon {
	var moons []*PlanetMoon

	for idx, _ := range r.planet.Moons {
		moons = append(moons, &r.planet.Moons[idx])
	}

	return &moons
}

func GetSolarSystemByID(systemID int32) (*SolarSystemResolver, error) {
	resp, err := http.MakeRequest("universe/systems/%d/", systemID)
	if err != nil {
		return nil, err
	}

	var system SolarSystem
	json.NewDecoder(resp.Body).Decode(&system)

	return &SolarSystemResolver{&system}, nil
}
