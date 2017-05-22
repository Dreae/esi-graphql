package resolvers

import (
	"encoding/json"

	"github.com/dreae/esi-graphql/resolvers/http"
)

type Alliance struct {
	AllianceID  int32
	Name        string `json:"alliance_name"`
	DateFounded string `json:"date_founded"`
	ExecutorID  int32  `json:"executor_corp"`
	Ticker      string `json:"ticker"`
}

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

func (a *AllianceResolver) Name() *string {
	return &a.alliance.Name
}

func (a *AllianceResolver) DateFounded() *string {
	return &a.alliance.DateFounded
}

func (a *AllianceResolver) ExecutorID() *int32 {
	return &a.alliance.ExecutorID
}

func (a *AllianceResolver) Ticker() *string {
	return &a.alliance.Ticker
}

func (a *AllianceResolver) Executor() (*CorporationResolver, error) {
	return GetCorpByID(a.alliance.ExecutorID)
}

func (a *AllianceResolver) Members() (*[]*AllianceMemberResolver, error) {
	return GetAllianceMembers(a.alliance.AllianceID)
}

func (a *AllianceResolver) Icons() (*AllianceIconResolver, error) {
	return GetAllianceIcons(a.alliance.AllianceID)
}

func (m *AllianceMemberResolver) CorporationID() *int32 {
	return &m.corpID
}

func (m *AllianceMemberResolver) Corporation() (*CorporationResolver, error) {
	return GetCorpByID(m.corpID)
}

func (i *AllianceIconResolver) LargeIcon() *string {
	return &i.icons.LargeIcon
}

func (i *AllianceIconResolver) SmallIcon() *string {
	return &i.icons.SmallIcon
}

func GetAllianceByID(allianceID int32) (*AllianceResolver, error) {
	var alliance Alliance
	resp, err := http.MakeRequest("alliances/%d/", allianceID)
	if err != nil {
		return &AllianceResolver{&alliance}, err
	}

	json.NewDecoder(resp.Body).Decode(&alliance)
	alliance.AllianceID = allianceID

	return &AllianceResolver{&alliance}, nil
}

func GetAllianceMembers(allianceID int32) (*[]*AllianceMemberResolver, error) {
	var memberIDs []int32
	var resolvers []*AllianceMemberResolver
	resp, err := http.MakeRequest("alliances/%d/corporations/", allianceID)
	if err != nil {
		return &resolvers, err
	}

	json.NewDecoder(resp.Body).Decode(&memberIDs)
	for _, memberID := range memberIDs {
		resolvers = append(resolvers, &AllianceMemberResolver{memberID})
	}

	return &resolvers, nil
}

func GetAllianceIcons(allianceID int32) (*AllianceIconResolver, error) {
	var icons AllianceIcons
	resp, err := http.MakeRequest("alliances/%d/icons/", allianceID)
	if err != nil {
		return &AllianceIconResolver{&icons}, err
	}

	json.NewDecoder(resp.Body).Decode(&icons)

	return &AllianceIconResolver{&icons}, nil
}
