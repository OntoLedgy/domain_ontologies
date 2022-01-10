package dto

type EditorCollectionCollaborator struct {
	Collection int `db:"collection"`
	Editor     int `db:"editor"`
}
