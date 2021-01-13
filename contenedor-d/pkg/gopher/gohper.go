package gopher

// Gopher defines the properties of a gopher to be listed
type Gopher struct {
	ID    string `json:"ID"`
	Name  string `json:"name,omitempty"`
	Image string `json:"image,omitempty"`
	Age   int    `json:"age,omitempty"`
}
