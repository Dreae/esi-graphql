package resolvers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Alliance struct {
	Name        string `json:"alliance_name"`
	DateFounded string `json:"date_founded"`
	ExecutorID  int32  `json:"executor_corp"`
	Ticker      string `json:"ticker"`
}

type AllianceResolver struct {
	alliance *Alliance
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

func GetAllianceByID(allianceID int32) (*AllianceResolver, error) {
	var alliance Alliance
	resp, err := http.Get(fmt.Sprintf("https://esi.tech.ccp.is/latest/alliances/%d/?datasource=tranquility", allianceID))
	if err != nil {
		return &AllianceResolver{&alliance}, err
	}

	json.NewDecoder(resp.Body).Decode(&alliance)

	return &AllianceResolver{&alliance}, nil
}
