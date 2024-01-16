package ds

import (
	"github.com/drakenchef/RIP/internal/app/role"
	"time"
)

type LoginReq struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginResp struct {
	ExpiresIn   time.Duration `json:"expires_in"`
	AccessToken string        `json:"access_token"`
	TokenType   string        `json:"token_type"`
	Role        role.Role     `json:"role"`
	Username    string        `json:"userName"`
	UserId      int           `json:"userid"`
}

type RegisterReq struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	UserName string `json:"user_name"`
}

// forcommit
type RegisterResp struct {
	Ok bool `json:"ok"`
}

// MARK: - Swagger

type LoginSwaggerResp struct {
	ExpiresIn   string `json:"expires_in"`
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}
