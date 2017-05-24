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

type CharacterResult int32

type CorporationResult int32

type AllianceResult int32

type EVETypeResult int32

type SolarSystemResult int32

type SearchResultsResolver struct {
	r *searchResults
}

func (r *SearchResultsResolver) Agents() *[]*int32 {
	var ids []*int32
	for idx, _ := range r.r.Agents {
		ids = append(ids, &r.r.Agents[idx])
	}

	return &ids
}

func (r *SearchResultsResolver) Alliances() *[]*AllianceResult {
	var allianceResults []*AllianceResult

	for idx, _ := range r.r.Alliances {
		allianceResults = append(allianceResults, &r.r.Alliances[idx])
	}

	return &allianceResults
}

func (r *SearchResultsResolver) Characters() *[]*CharacterResult {
	var characterResults []*CharacterResult

	for idx, _ := range r.r.Characters {
		characterResults = append(characterResults, &r.r.Characters[idx])
	}

	return &characterResults
}

func (r *SearchResultsResolver) Constellations() *[]*int32 {
	var ids []*int32
	for idx, _ := range r.r.Constellations {
		ids = append(ids, &r.r.Constellations[idx])
	}

	return &ids
}

func (r *SearchResultsResolver) Corporations() *[]*CorporationResult {
	var corporationResults []*CorporationResult

	for idx, _ := range r.r.Corporations {
		corporationResults = append(corporationResults, &r.r.Corporations[idx])
	}

	return &corporationResults
}

func (r *SearchResultsResolver) Factions() *[]*int32 {
	var ids []*int32
	for idx, _ := range r.r.Factions {
		ids = append(ids, &r.r.Factions[idx])
	}

	return &ids
}

func (r *SearchResultsResolver) InventoryTypes() *[]*EVETypeResult {
	var typeResults []*EVETypeResult

	for idx, _ := range r.r.Types {
		typeResults = append(typeResults, &r.r.Types[idx])
	}

	return &typeResults
}

func (r *SearchResultsResolver) Regions() *[]*int32 {
	var ids []*int32
	for idx, _ := range r.r.Regions {
		ids = append(ids, &r.r.Regions[idx])
	}

	return &ids
}

func (r *SearchResultsResolver) SolarSystems() *[]*SolarSystemResult {
	var ids []*SolarSystemResult
	for idx, _ := range r.r.SolarSystems {
		ids = append(ids, &r.r.SolarSystems[idx])
	}

	return &ids
}

func (r *SearchResultsResolver) Stations() *[]*int32 {
	var ids []*int32
	for idx, _ := range r.r.Stations {
		ids = append(ids, &r.r.Stations[idx])
	}

	return &ids
}

func (r *SearchResultsResolver) Wormholes() *[]*int32 {
	var ids []*int32
	for idx, _ := range r.r.Wormholes {
		ids = append(ids, &r.r.Wormholes[idx])
	}

	return &ids
}

func (r *CharacterResult) CharacterID() *int32 {
	return (*int32)(r)
}

func (r *CharacterResult) Character() (*CharacterResolver, error) {
	return GetCharacterByID(*(*int32)(r))
}

func (r *CorporationResult) CorporationID() *int32 {
	return (*int32)(r)
}

func (r *CorporationResult) Corporation() (*CorporationResolver, error) {
	return GetCorpByID(*(*int32)(r))
}

func (r *AllianceResult) AllianceID() *int32 {
	return (*int32)(r)
}

func (r *AllianceResult) Alliance() (*AllianceResolver, error) {
	return GetAllianceByID(*(*int32)(r))
}

func (r *EVETypeResult) TypeID() *int32 {
	return (*int32)(r)
}

func (r *EVETypeResult) Type() (*EVETypeResolver, error) {
	return GetEVEType(*(*int32)(r))
}

func (r *SolarSystemResult) SystemID() *int32 {
	return (*int32)(r)
}

func (r *SolarSystemResult) System() (*SolarSystemResolver, error) {
	return GetSolarSystemByID(*(*int32)(r))
}

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
	json.NewDecoder(resp.Body).Decode(&results)

	return &SearchResultsResolver{&results}, nil
}
