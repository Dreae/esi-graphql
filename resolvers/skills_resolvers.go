package resolvers

import (
	"encoding/json"

	"github.com/dreae/esi-graphql/resolvers/http"
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

type SkillQueueSkill struct {
	FinishDate      string `json:"finish_date"`
	FinishedLevel   int32  `json:"finished_level"`
	LevelEndSP      int32  `json:"level_end_sp"`
	LevelStartSP    int32  `json:"level_start_sp"`
	QueuePosition   int32  `json:"queue_position"`
	SkillID         int32  `json:"skill_id"`
	StartDate       string `json:"start_date"`
	TrainingStartSP int32  `json:"training_start_sp"`
}

type CharacterSkillResolver struct {
	skill *CharacterSkill
}

type CharacterSkillsResolver struct {
	skills *CharacterSkills
}

type SkillQueueResolver struct {
	skill *SkillQueueSkill
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
	for idx, _ := range skills.skills.Skills {
		resolvers = append(resolvers, &CharacterSkillResolver{&skills.skills.Skills[idx]})
	}

	return &resolvers
}

func (sq *SkillQueueResolver) FinishDate() *string {
	return &sq.skill.FinishDate
}

func (sq *SkillQueueResolver) FinishedLevel() *int32 {
	return &sq.skill.FinishedLevel
}

func (sq *SkillQueueResolver) LevelEndSP() *int32 {
	return &sq.skill.LevelEndSP
}

func (sq *SkillQueueResolver) LevelStartSP() *int32 {
	return &sq.skill.LevelStartSP
}

func (sq *SkillQueueResolver) QueuePosition() *int32 {
	return &sq.skill.QueuePosition
}

func (sq *SkillQueueResolver) SkillID() *int32 {
	return &sq.skill.SkillID
}

func (sq *SkillQueueResolver) StartDate() *string {
	return &sq.skill.StartDate
}

func (sq *SkillQueueResolver) TrainingStartSP() *int32 {
	return &sq.skill.TrainingStartSP
}

func (sq *SkillQueueResolver) Type() (*EVETypeResolver, error) {
	return GetEVEType(sq.skill.SkillID)
}

func (skills *CharacterSkillsResolver) TotalSP() *int32 {
	return &skills.skills.TotalSP
}

func GetSkillsForCharID(auth string, charID int32) (*CharacterSkillsResolver, error) {
	var skills CharacterSkills
	resp, err := http.MakeAuthorizedRequest(auth, "characters/%d/skills/", charID)
	if err != nil {
		return &CharacterSkillsResolver{&skills}, err
	}

	json.NewDecoder(resp.Body).Decode(&skills)

	return &CharacterSkillsResolver{&skills}, nil
}

func GetSkillQueueForCharID(auth string, charID int32) (*[]*SkillQueueResolver, error) {
	var skillQueue []SkillQueueSkill

	resp, err := http.MakeAuthorizedRequest(auth, "characters/%d/skillqueue/", charID)
	if err != nil {
		return nil, err
	}

	json.NewDecoder(resp.Body).Decode(&skillQueue)

	var resolvers []*SkillQueueResolver
	for idx, _ := range skillQueue {
		resolvers = append(resolvers, &SkillQueueResolver{&skillQueue[idx]})
	}

	return &resolvers, nil
}
