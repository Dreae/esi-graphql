package universe

import (
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/dreae/esi-graphql/resolvers/dogma"
)

type EVEType struct {
	TypeID          int32                      `json:"type_id"`
	Name            string                     `json:"name"`
	Description     string                     `json:"description"`
	Published       bool                       `json:"published"`
	GroupID         int32                      `json:"group_id"`
	Radius          float64                    `json:"radius"`
	Volume          float64                    `json:"volume"`
	Capacity        float64                    `json:"capacity"`
	PortionSize     int32                      `json:"portion_size"`
	Mass            float64                    `json:"mass"`
	GraphicID       int32                      `json:"graphic_id"`
	DogmaAttributes []dogma.DogmaAttributeNode `json:"dogma_attributes"`
}

type EVETypeResolver struct {
	EveType *EVEType
}

func (t *EVETypeResolver) TypeID() *int32 {
	return &t.EveType.TypeID
}

func (t *EVETypeResolver) Name() *string {
	return &t.EveType.Name
}

func (t *EVETypeResolver) Description() *string {
	return &t.EveType.Description
}

func (t *EVETypeResolver) Published() *bool {
	return &t.EveType.Published
}

func (t *EVETypeResolver) GroupID() *int32 {
	return &t.EveType.GroupID
}

func (t *EVETypeResolver) Radius() *float64 {
	return &t.EveType.Radius
}

func (t *EVETypeResolver) Volume() *float64 {
	return &t.EveType.Volume
}

func (t *EVETypeResolver) Capacity() *float64 {
	return &t.EveType.Capacity
}

func (t *EVETypeResolver) PortionSize() *int32 {
	return &t.EveType.PortionSize
}

func (t *EVETypeResolver) Mass() *float64 {
	return &t.EveType.Mass
}

func (t *EVETypeResolver) GraphicID() *int32 {
	return &t.EveType.GraphicID
}

func (t *EVETypeResolver) DogmaAttributes() *[]*dogma.DogmaAttributeNodeResolver {
	var nodes []*dogma.DogmaAttributeNodeResolver
	for _, attributeNode := range t.EveType.DogmaAttributes {
		nodes = append(nodes, &dogma.DogmaAttributeNodeResolver{&dogma.DogmaAttributeNode{
			attributeNode.AttributeID,
			attributeNode.Value,
		}})
	}

	return &nodes
}

func GetEVEType(typeID int32) (*EVETypeResolver, error) {
	var type_ EVEType
	resp, err := http.Get(fmt.Sprintf("https://esi.tech.ccp.is/latest/universe/types/%d/?datasource=tranquility&language=en-us", typeID))
	if err != nil {
		return &EVETypeResolver{&type_}, err
	}

	json.NewDecoder(resp.Body).Decode(&type_)

	return &EVETypeResolver{&type_}, nil
}
