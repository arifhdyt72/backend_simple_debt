package masterhutang

type MasterHutangInput struct {
	ID             int    `json:"id"`
	NamaHutang     string `json:"nama_hutang" binding:"required"`
	JumlahMaksimal int64  `json:"jumlah_maksimal" binding:"required"`
	JatuhTempo     int    `json:"jatuh_tempo" binding:"required"`
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
