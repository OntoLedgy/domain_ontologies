package dto

type EditorWatchReleaseStatus struct {
	Editor        int `db:"editor"`
	ReleaseStatus int `db:"release_status"`
}
