package masterhutang

import (
	"backend_test_debt/initializer"
	"backend_test_debt/models"
)

type Repository interface {
	Create(data models.MasterHutang) (models.MasterHutang, error)
	Update(data models.MasterHutang) (models.MasterHutang, error)
	DeleteData(id int) (models.MasterHutang, error)
	GetAllData(input DatatableInput) ([]models.MasterHutang, error)
	GetCountData(input DatatableInput) (int, error)
}

type repository struct {
}

func NewRepository() *repository {
	return &repository{}
}

func (r *repository) Create(data models.MasterHutang) (models.MasterHutang, error) {
	err := initializer.DB.Create(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) Update(data models.MasterHutang) (models.MasterHutang, error) {
	err := initializer.DB.Save(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) DeleteData(id int) (models.MasterHutang, error) {
	var masterHutang models.MasterHutang
	err := initializer.DB.First(&masterHutang, id).Error
	if err != nil {
		return masterHutang, err
	}

	err = initializer.DB.Delete(&masterHutang, masterHutang).Error
	if err != nil {
		return masterHutang, err
	}

	return masterHutang, nil
}

func (r *repository) GetAllData(input DatatableInput) ([]models.MasterHutang, error) {
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

	var data []models.MasterHutang
	err := initializer.DB.Where("nama_hutang LIKE ? OR jumlah_maksimal LIKE ? OR jatuh_tempo LIKE ?",
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
	var data []models.MasterHutang
	err := initializer.DB.Where("nama_hutang LIKE ? OR jumlah_maksimal LIKE ? OR jatuh_tempo LIKE ?",
		search, search, search).Find(&data).Count(&count).Error
	if err != nil {
		return int(count), err
	}

	return int(count), nil
}
