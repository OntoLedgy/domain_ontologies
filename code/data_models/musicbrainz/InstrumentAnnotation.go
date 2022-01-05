package dto

type InstrumentAnnotation struct {
	Instrument int `db:"instrument"`
	Annotation int `db:"annotation"`
}
