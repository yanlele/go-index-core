package request

type UpdateRequest struct {
	Name string `form:"name" json:"name" binding:"required"`
}

type CreateRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Name     string `form:"name" json:"name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
