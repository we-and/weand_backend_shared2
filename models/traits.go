package models


type HasIDInterface interface {
	GetId() int
}
type HasSetDeleted interface {
	SetDeleted()
}
