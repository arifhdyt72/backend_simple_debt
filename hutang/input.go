package hutang

type HutangInput struct {
	ID             int    `json:"id"`
	CustomerId     int    `json:"customer_id"`
	MasterHutangId int    `json:"master_hutang_id"`
	TglTransaksi   string `json:"tgl_transaksi"`
	JumlahHutang   int64  `json:"jumlah_hutang"`
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
