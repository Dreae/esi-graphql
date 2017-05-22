package resolvers

type searchResults struct {
	Agents         []int32 `json:"agent"`
	Alliances      []int32 `json:"alliance"`
	Characters     []int32 `json:"character"`
	Constellations []int32 `json:"constellation"`
	Corporations   []int32 `json:"corporation"`
	Factions       []int32 `json:"faction"`
	Types          []int32 `json:"inventorytype"`
	Regions        []int32 `json:"region"`
	SolarSystems   []int32 `json:"solarsystem"`
	Stations       []int32 `json:"station"`
	Wormholes      []int32 `json:"wormhole"`
}

type CharacterResult struct {
	id int32
}

type CorporationResult struct {
	id int32
}

type AllianceResult struct {
	id int32
}

type EVETypeResult struct {
	id int32
}

func (r *CharacterResult) CharacterID() *int32 {
	return &r.id
}

func (r *CharacterResult) Character() (*CharacterResolver, error) {
	return GetCharacterByID(r.id)
}

func (r *CorporationResult) CorporationID() *int32 {
	return &r.id
}

func (r *CorporationResult) Corporation() (*CorporationResolver, error) {
	return GetCorpByID(r.id)
}

func (r *AllianceResult) AllianceID() *int32 {
	return &r.id
}

func (r *AllianceResult) Alliance() (*AllianceResolver, error) {
	return GetAllianceByID(r.id)
}

func (r *EVETypeResult) TypeID() *int32 {
	return &r.id
}

func (r *EVETypeResult) Type() (*EVETypeResolver, error) {
	return GetEVEType(r.id)
}
