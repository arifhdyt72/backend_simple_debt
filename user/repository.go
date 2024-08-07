package user

import (
	"backend_test_debt/initializer"
	"backend_test_debt/models"
)

type Repository interface {
	Create(data models.User) (models.User, error)
	GetAllData(input DatatableInput) ([]models.User, error)
	GetCountData(input DatatableInput) (int, error)
	FindUserByUsernameOrEmail(user string) (models.User, error)
	Update(data models.User) (models.User, error)
	DeleteData(id int) (models.User, error)
}

type repository struct {
}

func NewRepository() *repository {
	return &repository{}
}

func (r *repository) Create(data models.User) (models.User, error) {
	err := initializer.DB.Create(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) FindUserByUsernameOrEmail(user string) (models.User, error) {
	var data models.User
	err := initializer.DB.Where("username = ? OR email = ?", user, user).Find(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) Update(data models.User) (models.User, error) {
	err := initializer.DB.Save(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) DeleteData(id int) (models.User, error) {
	var user models.User
	err := initializer.DB.First(&user, id).Error
	if err != nil {
		return user, err
	}

	err = initializer.DB.Delete(&user, user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) GetAllData(input DatatableInput) ([]models.User, error) {
	if input.Rows == 0 {
		input.Rows = 10
	}

	if input.SortField == "" || input.SortField == "null" {
		input.SortField = "id"
	}

	if input.SortOrder == "1" {
		input.SortOrder = "ASC"
	} else {
		input.SortOrder = "DESC"
	}

	search := "%" + input.Filters + "%"

	var data []models.User
	err := initializer.DB.Where("name LIKE ? OR username LIKE ? OR email LIKE ?",
		search, search, search).Limit(input.Rows).Offset(input.First).
		Order(input.SortField + " " + input.SortOrder).Find(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) GetCountData(input DatatableInput) (int, error) {
	if input.Rows == 0 {
		input.Rows = 10
	}

	if input.SortField == "" || input.SortField == "null" {
		input.SortField = "id"
	}

	if input.SortOrder == "1" {
		input.SortOrder = "ASC"
	} else {
		input.SortOrder = "DESC"
	}

	search := "%" + input.Filters + "%"
	var count int64
	var data []models.User
	err := initializer.DB.Where("name LIKE ? OR username LIKE ? OR email LIKE ?",
		search, search, search).Find(&data).Count(&count).Error
	if err != nil {
		return int(count), err
	}

	return int(count), nil
}
