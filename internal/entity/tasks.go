package entity

type Task struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Team        *Teams `json:"responsible_team"`
	Status      string `json:"status"`
	AssignedTo  *User  `json:"assigned_to"`
	CreatedBy   *User  `json:"created_by"`
}
