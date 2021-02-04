package types

// Common a common type for all project
type Common struct {
	Name string `json:"name"`
}

// GenericResponse generic response
type GenericResponse struct {
	Generic string `json:"generic"`
}

// Person a person
type Person struct {
	Age int `json:"age"`
}
