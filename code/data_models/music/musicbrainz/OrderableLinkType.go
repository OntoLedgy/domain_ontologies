package dto

type OrderableLinkType struct {
	LinkType  int `db:"link_type"`
	Direction int `db:"direction"`
}
