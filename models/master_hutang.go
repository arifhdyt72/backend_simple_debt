package models

type MasterHutang struct {
	MavisModel
	NamaHutang     string `json:"nama_hutang"`
	JumlahMaksimal int64  `json:"jumlah_maksimal"`
	JatuhTempo     int    `json:"jatuh_tempo"`
}
