package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Pembayaran struct {
	MavisModel
	HutangId        uint      `json:"hutang_id"`
	Hutang          *Hutang   `json:"hutang"`
	TglTransaksi    time.Time `json:"tgl_transaksi"`
	TotalDibayar    int64     `json:"total_dibayar"`
	BuktiPembayaran string    `json:"bukti_pembayaran"`
	Status          string    `json:"status"`
}

func (p *Pembayaran) BeforeCreate(tx *gorm.DB) (err error) {
	var hutang Hutang
	if err := tx.First(&hutang, p.HutangId).Error; err != nil {
		return err
	}

	hutang.SisaTagihan = hutang.SisaTagihan - p.TotalDibayar
	hutang.TotalDibayar += p.TotalDibayar
	if hutang.SisaTagihan < 0 {
		return errors.New("jumlah hutang melebihi nilai maksimal")
	}
	if hutang.SisaTagihan == 0 {
		hutang.Status = "paid"
	}

	errData := tx.Save(&hutang).Error
	if errData != nil {
		return errData
	}

	p.Status = "paid"
	return nil
}
