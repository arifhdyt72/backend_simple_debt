package user

type UserInput struct {
	ID       int    `json:"id"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserUpdate struct {
	ID       int    `json:"id"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password"`
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type InputId struct {
	ID int `uri:"id" binding:"required"`
}

type DatatableInput struct {
	Filters   string `json:"filters"`
	Page      int    `json:"page"`
	First     int    `json:"first"`
	Rows      int    `json:"rows"`
	PageCount int    `json:"pageCount"`
	SortField string `json:"sortField"`
	SortOrder string `json:"sortOrder"`
}
