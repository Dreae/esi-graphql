package resolvers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CharacterSkill struct {
	CurrentSkillLevel  int32 `json:"current_skill_level"`
	SkillID            int32 `json:"skill_id"`
	SkillpointsInSkill int32 `json:"skillpoints_in_skill"`
}

type CharacterSkills struct {
	Skills  []CharacterSkill `json:"skills"`
	TotalSP int32            `json:"total_sp"`
}

type CharacterSkillResolver struct {
	skill *CharacterSkill
}

type CharacterSkillsResolver struct {
	skills *CharacterSkills
}

func (skill *CharacterSkillResolver) CurrentSkillLevel() *int32 {
	return &skill.skill.CurrentSkillLevel
}

func (skill *CharacterSkillResolver) SkillID() *int32 {
	return &skill.skill.SkillID
}

func (skill *CharacterSkillResolver) SkillpointsInSkill() *int32 {
	return &skill.skill.SkillpointsInSkill
}

func (skill *CharacterSkillResolver) Type() (*EVETypeResolver, error) {
	return GetEVEType(*skill.SkillID())
}

func (skills *CharacterSkillsResolver) Skills() *[]*CharacterSkillResolver {
	var resolvers []*CharacterSkillResolver
	for _, skill := range skills.skills.Skills {
		resolvers = append(resolvers, &CharacterSkillResolver{
			&CharacterSkill{
				skill.CurrentSkillLevel,
				skill.SkillID,
				skill.SkillpointsInSkill,
			},
		})
	}

	return &resolvers
}

func (skills *CharacterSkillsResolver) TotalSP() *int32 {
	return &skills.skills.TotalSP
}

func GetSkillsForCharID(auth string, charID int32) (*CharacterSkillsResolver, error) {
	var skills CharacterSkills
	req, _ := http.NewRequest("GET", fmt.Sprintf("https://esi.tech.ccp.is/latest/characters/%d/skills/?datasource=tranquility", charID), nil)
	req.Header.Set("Authorization", auth)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return &CharacterSkillsResolver{&skills}, err
	}

	json.NewDecoder(resp.Body).Decode(&skills)

	return &CharacterSkillsResolver{&skills}, nil
}
