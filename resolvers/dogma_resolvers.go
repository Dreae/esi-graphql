package resolvers

import (
	"encoding/json"

	"github.com/dreae/esi-graphql/cache"
	"github.com/dreae/esi-graphql/resolvers/http"
)

type DogmaAttributeResolver struct {
	Attribute *DogmaAttribute
}

func (a *DogmaAttributeResolver) AttributeID() *int32 {
	return &a.Attribute.AttributeID
}

func (a *DogmaAttributeResolver) DefaultValue() *int32 {
	return &a.Attribute.DefaultValue
}

func (a *DogmaAttributeResolver) Description() *string {
	return &a.Attribute.Description
}

func (a *DogmaAttributeResolver) DisplayName() *string {
	return &a.Attribute.DisplayName
}

func (a *DogmaAttributeResolver) HighIsGood() *bool {
	return &a.Attribute.HighIsGood
}

func (a *DogmaAttributeResolver) IconID() *int32 {
	return &a.Attribute.IconID
}

func (a *DogmaAttributeResolver) Name() *string {
	return &a.Attribute.Name
}

func (a *DogmaAttributeResolver) Published() *bool {
	return &a.Attribute.Published
}

func (a *DogmaAttributeResolver) Stackable() *bool {
	return &a.Attribute.Stackable
}

func (a *DogmaAttributeResolver) UnitID() *int32 {
	return &a.Attribute.UnitID
}

type DogmaAttribute struct {
	AttributeID  int32  `json:"attribute_id"`
	DefaultValue int32  `json:"default_value"`
	Description  string `json:"description"`
	DisplayName  string `json:"display_name"`
	HighIsGood   bool   `json:"high_is_good"`
	IconID       int32  `json:"icon_id"`
	Name         string `json:"name"`
	Published    bool   `json:"published"`
	Stackable    bool   `json:"stackable"`
	UnitID       int32  `json:"unit_id"`
}

type DogmaAttributeNodeResolver struct {
	Node *DogmaAttributeNode
}

func (a *DogmaAttributeNodeResolver) AttributeID() *int32 {
	return &a.Node.AttributeID
}

func (a *DogmaAttributeNodeResolver) Value() *float64 {
	return &a.Node.Value
}

func (a *DogmaAttributeNodeResolver) Attribute() (*DogmaAttributeResolver, error) {
	return GetDogmaAttributeResolver(a.Node.AttributeID)
}

type DogmaAttributeNode struct {
	AttributeID int32   `json:"attribute_id"`
	Value       float64 `json:"value"`
}

var dogmaCache = cache.New(3600)

func GetDogmaAttributeResolver(attributeID int32) (*DogmaAttributeResolver, error) {
	if item, ok := dogmaCache.Get(attributeID); ok {
		attribute := item.(DogmaAttribute)
		return &DogmaAttributeResolver{&attribute}, nil
	}

	var attribute DogmaAttribute
	resp, err := http.MakeRequest("dogma/attributes/%d/", attributeID)
	if err != nil {
		return &DogmaAttributeResolver{&attribute}, err
	}

	json.NewDecoder(resp.Body).Decode(&attribute)

	dogmaCache.Set(attributeID, attribute)

	return &DogmaAttributeResolver{&attribute}, nil
}

func GetDogmaList() (*[]*DogmaAttributeNodeResolver, error) {
	var attributes []int32
	resp, err := http.MakeRequest("dogma/attributes/")
	if err != nil {
		return nil, err
	}

	json.NewDecoder(resp.Body).Decode(&attributes)

	var resolvers []*DogmaAttributeNodeResolver
	for _, attribute := range attributes {
		resolvers = append(resolvers, &DogmaAttributeNodeResolver{&DogmaAttributeNode{attribute, 0}})
	}

	return &resolvers, nil
}
