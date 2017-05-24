package resolvers

import (
	"encoding/json"

	"github.com/dreae/esi-graphql/resolvers/http"
)

// Alliance holds the details of an EVE alliance
type Alliance struct {
	AllianceID  int32
	Name        string `json:"alliance_name"`
	DateFounded string `json:"date_founded"`
	ExecutorID  int32  `json:"executor_corp"`
	Ticker      string `json:"ticker"`
}

// AllianceIcons stores the icon URLs for an alliance
type AllianceIcons struct {
	LargeIcon string `json:"px128x128"`
	SmallIcon string `json:"px64x64"`
}

type AllianceMemberResolver struct {
	corpID int32
}

type AllianceResolver struct {
	alliance *Alliance
}

type AllianceIconResolver struct {
	icons *AllianceIcons
}

// Name returns the name of the alliance
func (a *AllianceResolver) Name() *string {
	return &a.alliance.Name
}

// DateFounded returns the date an alliance was founded
func (a *AllianceResolver) DateFounded() *string {
	return &a.alliance.DateFounded
}

// ExecutorID returns the ID of the alliance executor
func (a *AllianceResolver) ExecutorID() *int32 {
	return &a.alliance.ExecutorID
}

// Ticker returns the alliance ticker string
func (a *AllianceResolver) Ticker() *string {
	return &a.alliance.Ticker
}

// Executor serves as a pointer to get more detailed information on the
// alliance executor corp
func (a *AllianceResolver) Executor() (*CorporationResolver, error) {
	return GetCorpByID(a.alliance.ExecutorID)
}

// Members returns an array representing the alliance member corporations
func (a *AllianceResolver) Members() (*[]*AllianceMemberResolver, error) {
	return GetAllianceMembers(a.alliance.AllianceID)
}

// Icons returns the alliance's icon structure
func (a *AllianceResolver) Icons() (*AllianceIconResolver, error) {
	return GetAllianceIcons(a.alliance.AllianceID)
}

// CorporationID returns the ID of this alliance member
func (m *AllianceMemberResolver) CorporationID() *int32 {
	return &m.corpID
}

// Corporation serves has a pointer to get more informartion on this corporation
func (m *AllianceMemberResolver) Corporation() (*CorporationResolver, error) {
	return GetCorpByID(m.corpID)
}

// LargeIcon returns the URL to the alliance's high resolution icon
func (i *AllianceIconResolver) LargeIcon() *string {
	return &i.icons.LargeIcon
}

// SmallIcon returns the URL to the alliance's low resolution icon
func (i *AllianceIconResolver) SmallIcon() *string {
	return &i.icons.SmallIcon
}

// GetAllianceByID fetches an alliance by the given alliance ID
func GetAllianceByID(allianceID int32) (*AllianceResolver, error) {
	var alliance Alliance
	resp, err := http.MakeRequest("alliances/%d/", allianceID)
	if err != nil {
		return &AllianceResolver{&alliance}, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&alliance); err != nil {
		return nil, err
	}

	alliance.AllianceID = allianceID

	return &AllianceResolver{&alliance}, nil
}

// GetAllianceMembers returns the array of all member corporations for an alliance
func GetAllianceMembers(allianceID int32) (*[]*AllianceMemberResolver, error) {
	var memberIDs []int32
	var resolvers []*AllianceMemberResolver
	resp, err := http.MakeRequest("alliances/%d/corporations/", allianceID)
	if err != nil {
		return &resolvers, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&memberIDs); err != nil {
		return nil, err
	}

	for _, memberID := range memberIDs {
		resolvers = append(resolvers, &AllianceMemberResolver{memberID})
	}

	return &resolvers, nil
}

// GetAllianceIcons returns the alliance icon struct for a give alliance
func GetAllianceIcons(allianceID int32) (*AllianceIconResolver, error) {
	var icons AllianceIcons
	resp, err := http.MakeRequest("alliances/%d/icons/", allianceID)
	if err != nil {
		return &AllianceIconResolver{&icons}, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&icons); err != nil {
		return nil, err
	}

	return &AllianceIconResolver{&icons}, nil
}
