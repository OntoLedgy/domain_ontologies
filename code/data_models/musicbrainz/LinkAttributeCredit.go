package dto

type LinkAttributeCredit struct {
	Link          int    `db:"link"`
	AttributeType int    `db:"attribute_type"`
	CreditedAs    string `db:"credited_as"`
}
