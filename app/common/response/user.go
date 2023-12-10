package response

type UserRes struct {
	ID        uint     `json:"id"`
	Nickname  string   `json:"nickname"`
	AvatarUrl string   `json:"avatar_url"`
	Role      []string `json:"role"`
}
