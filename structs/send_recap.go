package structs

type SendRecapPerson struct {
	Type        string
	Destination string
	Success     bool
	Error       error
}
