package schemas

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	GitID string `json:"git_id"`
}

type LoginBonusRequest struct {
	ID    string `json:"id"`
	GitID string `json:"git_id"`
}
