package web

type UserRegisterInput struct {
	LastName  string `json:"last_name" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
	Address   string `json:"address" binding:"required"`
}

type UserLoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Address  string `json:"address" binding:"required"`
}

type UserResponse struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Avatar    string `json:"avatar"`
	Token     string `json:"token"`
	Balance   int    `json:"balance"`
	Rank      int    `json:"rank"`
}

type UserRankResponse struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Avatar    string `json:"avatar"`
	Balance   int    `json:"balance"`
	Rank      int    `json:"rank"`
}

type UserDetailResponse struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Avatar    string `json:"avatar"`
	Balance   int    `json:"balance"`
	Rank      int    `json:"rank"`
	LastRank  int
}
