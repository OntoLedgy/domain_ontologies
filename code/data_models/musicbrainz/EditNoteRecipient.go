package dto

type EditNoteRecipient struct {
	Recipient int `db:"recipient"`
	EditNote  int `db:"edit_note"`
}
