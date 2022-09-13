package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// A Transport is the main product in the store.
// It is used to describe the animals available in the store.
//
// swagger:model transport
type Transport struct {
	// The id of the transport.
	//
	// required: true
	Id primitive.ObjectID `json:"id" bson:"_id"`

	// The name of the pet.
	//
	// required: true
	// pattern: \w[\w-]+
	// minimum length: 3
	// maximum length: 50
	Name *string `json:"name,omitempty" validate:"required"`

	// Description
	Description *string `json:"description,omitempty" validate:"required"`

	//Modalirt
	Modality *string `json:"modality,omitempty" validate:"required"`
}
