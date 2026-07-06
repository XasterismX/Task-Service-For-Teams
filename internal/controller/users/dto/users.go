package dto

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
type CreateResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type GetUserRequest struct {
	ID string `json:"id"`
}
type GetUserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UpdateUserRequest struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UpdateUserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type DeleteUserRequest struct {
	ID string `json:"id"`
}
type DeleteUserResponse struct {
	IsDeleted bool `json:"is_deleted"`
}
type GetAllUsersRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}
type GetAllUsersResponse struct {
	Users []GetUserResponse `json:"users"`
}
