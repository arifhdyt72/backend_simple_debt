package hutang

import (
	"backend_test_debt/initializer"
	"backend_test_debt/models"

	"gorm.io/gorm/clause"
)

type Repository interface {
	Create(data models.Hutang) (models.Hutang, error)
	Update(data models.Hutang) (models.Hutang, error)
	DeleteData(id int) (models.Hutang, error)
	GetAllData(input DatatableInput) ([]models.Hutang, error)
	GetCountData(input DatatableInput) (int, error)
	GetAllReport() ([]models.Hutang, error)
}

type repository struct {
}

func NewRepository() *repository {
	return &repository{}
}

func (r *repository) Create(data models.Hutang) (models.Hutang, error) {
	err := initializer.DB.Create(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) Update(data models.Hutang) (models.Hutang, error) {
	err := initializer.DB.Save(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) DeleteData(id int) (models.Hutang, error) {
	var hutang models.Hutang
	err := initializer.DB.First(&hutang, id).Error
	if err != nil {
		return hutang, err
	}

	err = initializer.DB.Delete(&hutang, hutang).Error
	if err != nil {
		return hutang, err
	}

	return hutang, nil
}

func (r *repository) GetAllData(input DatatableInput) ([]models.Hutang, error) {
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

	var data []models.Hutang
	err := initializer.DB.Preload("Customer").Preload("MasterHutang").
		Preload("Pembayaran").Preload(clause.Associations).
		Joins("JOIN customers ON customers.id = hutangs.customer_id").
		Joins("JOIN master_hutangs ON master_hutangs.id = hutangs.master_hutang_id").
		Where("customers.name LIKE ? OR master_hutangs.nama_hutang LIKE ? OR tgl_transaksi LIKE ? OR tgl_jatuh_tempo LIKE ? OR jumlah_hutang LIKE ? OR total_dibayar LIKE ? OR sisa_tagihan LIKE ? OR hutangs.status LIKE ?",
			search, search, search, search, search, search, search, search).Limit(input.Rows).Offset(input.First).
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
	var data []models.Hutang
	err := initializer.DB.Preload("Customer").Preload("MasterHutang").
		Preload("Pembayaran").Preload(clause.Associations).
		Joins("JOIN customers ON customers.id = hutangs.customer_id").
		Joins("JOIN master_hutangs ON master_hutangs.id = hutangs.master_hutang_id").
		Where("customers.name LIKE ? OR master_hutangs.nama_hutang LIKE ? OR tgl_transaksi LIKE ? OR tgl_jatuh_tempo LIKE ? OR jumlah_hutang LIKE ? OR total_dibayar LIKE ? OR sisa_tagihan LIKE ? OR hutangs.status LIKE ?",
			search, search, search, search, search, search, search, search).Find(&data).Count(&count).Error
	if err != nil {
		return int(count), err
	}

	return int(count), nil
}

func (r *repository) GetAllReport() ([]models.Hutang, error) {
	var data []models.Hutang
	err := initializer.DB.Preload("Customer").Preload("MasterHutang").
		Preload("Pembayaran").Preload(clause.Associations).
		Joins("JOIN customers ON customers.id = hutangs.customer_id").
		Joins("JOIN master_hutangs ON master_hutangs.id = hutangs.master_hutang_id").
		Find(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}
