package resolvers

import (
	"encoding/json"

	"bytes"

	"strings"

	"github.com/dreae/esi-graphql/resolvers/http"
)

type searchResults struct {
	Agents         []int32             `json:"agent"`
	Alliances      []AllianceResult    `json:"alliance"`
	Characters     []CharacterResult   `json:"character"`
	Constellations []int32             `json:"constellation"`
	Corporations   []CorporationResult `json:"corporation"`
	Factions       []int32             `json:"faction"`
	Types          []EVETypeResult     `json:"inventorytype"`
	Regions        []int32             `json:"region"`
	SolarSystems   []SolarSystemResult `json:"solarsystem"`
	Stations       []int32             `json:"station"`
	Wormholes      []int32             `json:"wormhole"`
}

// CharacterResult is a wrapper
// type for fetching more details about a character
type CharacterResult int32

// CorporationResult is a wrapper
// type for fetching more details about a corporation
type CorporationResult int32

// AllianceResult is a wrapper
// type for fetching more details about a alliance
type AllianceResult int32

// EVETypeResult is a wrapper
// type for fetching more details about a Type
type EVETypeResult int32

// SolarSystemResult is a wrapper
// type for fetching more details about a solar system
type SolarSystemResult int32

// SearchResultsResolver is a wrapper
// type for fetching search results
type SearchResultsResolver struct {
	r *searchResults
}

// Agents returns the agents found in the search
func (r *SearchResultsResolver) Agents() *[]*int32 {
	var ids []*int32
	for idx := range r.r.Agents {
		ids = append(ids, &r.r.Agents[idx])
	}

	return &ids
}

// Alliances returns the alliances found in the search
func (r *SearchResultsResolver) Alliances() *[]*AllianceResult {
	var allianceResults []*AllianceResult

	for idx := range r.r.Alliances {
		allianceResults = append(allianceResults, &r.r.Alliances[idx])
	}

	return &allianceResults
}

// Characters returns the characters found in a search
func (r *SearchResultsResolver) Characters() *[]*CharacterResult {
	var characterResults []*CharacterResult

	for idx := range r.r.Characters {
		characterResults = append(characterResults, &r.r.Characters[idx])
	}

	return &characterResults
}

// Constellations returns the constellations found in a search
func (r *SearchResultsResolver) Constellations() *[]*int32 {
	var ids []*int32
	for idx := range r.r.Constellations {
		ids = append(ids, &r.r.Constellations[idx])
	}

	return &ids
}

// Corporations returns the corporations found in a search
func (r *SearchResultsResolver) Corporations() *[]*CorporationResult {
	var corporationResults []*CorporationResult

	for idx := range r.r.Corporations {
		corporationResults = append(corporationResults, &r.r.Corporations[idx])
	}

	return &corporationResults
}

// Factions returns the factions found in a search
func (r *SearchResultsResolver) Factions() *[]*int32 {
	var ids []*int32
	for idx := range r.r.Factions {
		ids = append(ids, &r.r.Factions[idx])
	}

	return &ids
}

// InventoryTypes returns the types found in a search
func (r *SearchResultsResolver) InventoryTypes() *[]*EVETypeResult {
	var typeResults []*EVETypeResult

	for idx := range r.r.Types {
		typeResults = append(typeResults, &r.r.Types[idx])
	}

	return &typeResults
}

// Regions returns the regions found in a search
func (r *SearchResultsResolver) Regions() *[]*int32 {
	var ids []*int32
	for idx := range r.r.Regions {
		ids = append(ids, &r.r.Regions[idx])
	}

	return &ids
}

// SolarSystems returns the solar systems found in a search
func (r *SearchResultsResolver) SolarSystems() *[]*SolarSystemResult {
	var ids []*SolarSystemResult
	for idx := range r.r.SolarSystems {
		ids = append(ids, &r.r.SolarSystems[idx])
	}

	return &ids
}

// Stations returns the stations found in a search
func (r *SearchResultsResolver) Stations() *[]*int32 {
	var ids []*int32
	for idx := range r.r.Stations {
		ids = append(ids, &r.r.Stations[idx])
	}

	return &ids
}

// Wormholes returns the wormholes found in a search
func (r *SearchResultsResolver) Wormholes() *[]*int32 {
	var ids []*int32
	for idx := range r.r.Wormholes {
		ids = append(ids, &r.r.Wormholes[idx])
	}

	return &ids
}

// CharacterID returns the ID of the characer search result
func (r *CharacterResult) CharacterID() *int32 {
	return (*int32)(r)
}

// Character serves as a pointer to fetch more detailed information
// about the character found in the search
func (r *CharacterResult) Character() (*CharacterResolver, error) {
	return GetCharacterByID(*(*int32)(r))
}

// CorporationID returns the ID of the corporation search result
func (r *CorporationResult) CorporationID() *int32 {
	return (*int32)(r)
}

// Corporation serves as a pointer to fetch more detailed information
// about the corporation found in the search
func (r *CorporationResult) Corporation() (*CorporationResolver, error) {
	return GetCorpByID(*(*int32)(r))
}

// AllianceID returns the ID of the alliance search result
func (r *AllianceResult) AllianceID() *int32 {
	return (*int32)(r)
}

// Alliance serves as a pointer to fetch more detailed information
// about the alliance found in the search
func (r *AllianceResult) Alliance() (*AllianceResolver, error) {
	return GetAllianceByID(*(*int32)(r))
}

// TypeID returns the ID of the type search result
func (r *EVETypeResult) TypeID() *int32 {
	return (*int32)(r)
}

// Type serves as a pointer to fetch more detailed information
// about the type found in the search
func (r *EVETypeResult) Type() (*EVETypeResolver, error) {
	return GetEVEType(*(*int32)(r))
}

// SystemID returns the ID of the solar system search result
func (r *SolarSystemResult) SystemID() *int32 {
	return (*int32)(r)
}

// System serves as a pointer to fetch more detailed information
// about the solar system found in the search
func (r *SolarSystemResult) System() (*SolarSystemResolver, error) {
	return GetSolarSystemByID(*(*int32)(r))
}

// DoSearch performs the keyword search with the provided search text,
// and types to perform the search on
func DoSearch(types *[]*string, keyword string) (*SearchResultsResolver, error) {
	vals := make(map[string]string)
	vals["search"] = keyword
	var categories bytes.Buffer
	for idx, val := range *types {
		categories.WriteString(strings.ToLower(*val))
		if idx != len(*types)-1 {
			categories.WriteString(",")
		}
	}
	vals["categories"] = categories.String()

	resp, err := http.MakeQuery("search/", vals)
	if err != nil {
		return nil, err
	}

	var results searchResults
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return nil, err
	}

	return &SearchResultsResolver{&results}, nil
}
