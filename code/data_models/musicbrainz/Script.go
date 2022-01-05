package dto

type Script struct {
	ID        int    `db:"id"`
	IsoCode   string `db:"iso_code"`
	IsoNumber string `db:"iso_number"`
	Name      string `db:"name"`
	Frequency int    `db:"frequency"`
}
