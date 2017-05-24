package resolvers

import (
	"encoding/json"

	"github.com/dreae/esi-graphql/resolvers/http"
)

type DogmaAttributeResolver struct {
	Attribute *DogmaAttribute
}

// AttributeID returns the ID of the dogma attribute
func (a *DogmaAttributeResolver) AttributeID() *int32 {
	return &a.Attribute.AttributeID
}

// DefaultValue returns the dogma attribute's default value
func (a *DogmaAttributeResolver) DefaultValue() *int32 {
	return &a.Attribute.DefaultValue
}

// Description returns the description of the dogma attribute
func (a *DogmaAttributeResolver) Description() *string {
	return &a.Attribute.Description
}

// DisplayName returns the display name of the dogma attribute
func (a *DogmaAttributeResolver) DisplayName() *string {
	return &a.Attribute.DisplayName
}

// HighIsGood returns if higher is better for this dogma attribute
func (a *DogmaAttributeResolver) HighIsGood() *bool {
	return &a.Attribute.HighIsGood
}

// IconID returns the ID of the attribute's icon
func (a *DogmaAttributeResolver) IconID() *int32 {
	return &a.Attribute.IconID
}

// Name returns the attribute's name
func (a *DogmaAttributeResolver) Name() *string {
	return &a.Attribute.Name
}

// Published returns if the attribute is published
func (a *DogmaAttributeResolver) Published() *bool {
	return &a.Attribute.Published
}

// Stackable returns if the attribute is stackable
func (a *DogmaAttributeResolver) Stackable() *bool {
	return &a.Attribute.Stackable
}

// UnitID returns the attribute's unit ID
func (a *DogmaAttributeResolver) UnitID() *int32 {
	return &a.Attribute.UnitID
}

// DogmaAttribute holds the details for an EVE dogma attribute
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

// AttributeID returns the ID of the attribute associated with this node
func (a *DogmaAttributeNodeResolver) AttributeID() *int32 {
	return &a.Node.AttributeID
}

// Value returns the value of the attribute associated with this node
func (a *DogmaAttributeNodeResolver) Value() *float64 {
	return &a.Node.Value
}

// Attribute serves as a pointer to fetch the details of the attribute
// associated with this node
func (a *DogmaAttributeNodeResolver) Attribute() (*DogmaAttributeResolver, error) {
	return GetDogmaAttributeResolver(a.Node.AttributeID)
}

type DogmaAttributeNode struct {
	AttributeID int32   `json:"attribute_id"`
	Value       float64 `json:"value"`
}

// GetDogmaAttributeResolver gets the details of a dogma attribute by ID
func GetDogmaAttributeResolver(attributeID int32) (*DogmaAttributeResolver, error) {
	var attribute DogmaAttribute
	resp, err := http.MakeRequest("dogma/attributes/%d/", attributeID)
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&attribute); err != nil {
		return nil, err
	}

	return &DogmaAttributeResolver{&attribute}, nil
}

// GetDogmaList gets all dogma attributes as reported by ESI
func GetDogmaList() (*[]*DogmaAttributeNodeResolver, error) {
	var attributes []int32
	resp, err := http.MakeRequest("dogma/attributes/")
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&attributes); err != nil {
		return nil, err
	}

	var resolvers []*DogmaAttributeNodeResolver
	for _, attribute := range attributes {
		resolvers = append(resolvers, &DogmaAttributeNodeResolver{&DogmaAttributeNode{attribute, 0}})
	}

	return &resolvers, nil
}
