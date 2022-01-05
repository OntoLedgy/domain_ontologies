package dto

type EditReleaseGroup struct {
	Edit         int `db:"edit"`
	ReleaseGroup int `db:"release_group"`
}
