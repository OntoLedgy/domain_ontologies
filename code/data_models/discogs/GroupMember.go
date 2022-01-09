package dto

type GroupMember struct {
	GroupArtistID  int    `db:"group_artist_id"`
	MemberArtistID int    `db:"member_artist_id"`
	MemberName     string `db:"member_name"`
}
