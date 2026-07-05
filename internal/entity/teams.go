package entity

type Teams struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type TeamMembers struct {
	Id   int    `json:"id"`
	Team *Teams `json:"team"`
	User *User  `json:"user"`
	Role string `json:"role"`
}
