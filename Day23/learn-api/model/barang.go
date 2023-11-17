package model

import "gorm.io/gorm"

type BarangModel struct {
	gorm.Model
	NamaBarang string
	Stok       int
	Harga      int
	UserID     uint
}

type BarangQuery struct {
	DB *gorm.DB
}

func (bq *BarangQuery) AddBarang(newBarang BarangModel) (BarangModel, error) {
	if err := bq.DB.Create(&newBarang).Error; err != nil {
		return BarangModel{}, err
	}

	return newBarang, nil
}

func (bq *BarangQuery) GetBarang() ([]BarangModel, error) {
	var barang []BarangModel
	if err := bq.DB.Find(&barang).Error; err != nil {
		return nil, err
	}

	return barang, nil
}

func (bq *BarangQuery) UpdateBarang(updatedBarang BarangModel) (BarangModel, error) {
	var barang BarangModel
	if err := bq.DB.First(&barang, updatedBarang.ID).Error; err != nil {
		return BarangModel{}, err
	}

	barang.NamaBarang = updatedBarang.NamaBarang
	barang.Harga = updatedBarang.Harga
	barang.Stok = updatedBarang.Stok

	if err := bq.DB.Save(&barang).Error; err != nil {
		return BarangModel{}, err
	}

	return barang, nil
}

func (bq *BarangQuery) DeleteBarang(id int) error {
	var barang BarangModel
	if err := bq.DB.First(&barang, id).Error; err != nil {
		return err
	}

	if err := bq.DB.Delete(&barang).Error; err != nil {
		return err
	}

	return nil
}