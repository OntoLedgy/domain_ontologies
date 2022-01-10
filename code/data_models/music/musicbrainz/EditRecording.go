package dto

type EditRecording struct {
	Edit      int `db:"edit"`
	Recording int `db:"recording"`
}
