package dto

type LinkAttributeTextValue struct {
	Link          int    `db:"link"`
	AttributeType int    `db:"attribute_type"`
	TextValue     string `db:"text_value"`
}
