package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Hutang struct {
	MavisModel
	CustomerId     uint         `json:"customer_id"`
	Customer       Customer     `json:"customer"`
	MasterHutangId uint         `json:"master_hutang_id"`
	MasterHutang   MasterHutang `json:"master_hutang"`
	TglTransaksi   time.Time    `json:"tgl_transaksi"`
	TglJatuhTempo  time.Time    `json:"tgl_jatuh_tempo"`
	JumlahHutang   int64        `json:"jumlah_hutang"`
	TotalDibayar   int64        `json:"total_dibayar"`
	SisaTagihan    int64        `json:"sisa_tagihan"`
	Status         string       `json:"status"`
	Pembayaran     []Pembayaran `json:"pembayaran" gorm:"foreignKey:HutangId"`
}

func (h *Hutang) BeforeCreate(tx *gorm.DB) (err error) {
	var master MasterHutang
	if err := tx.First(&master, h.MasterHutangId).Error; err != nil {
		return err
	}

	if h.JumlahHutang > master.JumlahMaksimal {
		return errors.New("jumlah hutang melebihi nilai maksimal")
	}

	h.TglJatuhTempo = h.TglTransaksi.AddDate(0, 0, master.JatuhTempo)
	if time.Now().Unix() > h.TglJatuhTempo.Unix() {
		h.Status = "overdue"
	} else {
		h.Status = "unpaid"
	}

	h.SisaTagihan = h.JumlahHutang

	return nil
}

func (h *Hutang) BeforeSave(tx *gorm.DB) (err error) {
	var master MasterHutang
	if err := tx.First(&master, h.MasterHutangId).Error; err != nil {
		return err
	}

	if h.JumlahHutang > master.JumlahMaksimal {
		return errors.New("jumlah hutang melebihi nilai maksimal")
	}

	return nil
}

func (h *Hutang) BeforeDelete(tx *gorm.DB) (err error) {
	if h.TotalDibayar > 0 {
		return errors.New("tidak bisa melakukan delete, sudah ada pembayaran")
	}

	return nil
}
