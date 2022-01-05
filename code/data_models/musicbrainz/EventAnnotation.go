package dto

type EventAnnotation struct {
	Event      int `db:"event"`
	Annotation int `db:"annotation"`
}
