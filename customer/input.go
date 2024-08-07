package customer

type CustomerInput struct {
	ID          int    `json:"id"`
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Address     string `json:"address" binding:"required"`
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
