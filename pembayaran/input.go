package pembayaran

type PembayaranInput struct {
	HutangId     int    `form:"hutang_id" binding:"required"`
	TotalDibayar int64  `form:"total_dibayar" binding:"required"`
	TglTransaksi string `form:"tgl_transaksi" binding:"required"`
}
