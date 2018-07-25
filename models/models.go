package models

type Order struct {
	IsProductValid bool `json:"isProductValid"`
}

// This struct defines the engine's response type.
type ResponseData struct {
	Status int
	Message string
}