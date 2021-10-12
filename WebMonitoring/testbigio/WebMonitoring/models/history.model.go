package models

import (
	"net/http"
	"onboarding/helpers"

	"gorm.io/gorm"
)

type Raport struct {
	gorm.Model
	IDSiswa uint   `gorm:"column:siswa_id" json:"id_siswa"`
	Nama    string `gorm:"column:nama" json:"nama"`
	Mapel   string `gorm:"column:mapel" json:"mapel"`
	Nilai   string `gorm:"column:nilai" json:"nilai"`
	User    User   `gorm:"foreignKey:IDSiswa"`
}

func (raport *Raport) IsiRaport(conn *gorm.DB, idSiswa, idGuru int) (helpers.MessageResponse, bool) {
	var user User
	if err := conn.Debug().First(&user, idGuru).Error; err != nil {
		return *helpers.MessageResponses(false, http.StatusUnprocessableEntity, "User tidak ditemukan"), false
	}
	if user.RoleID == 1 {
		addNilai := Raport{
			IDSiswa: raport.IDSiswa,
			Nama:    raport.Nama,
			Mapel:   raport.Mapel,
			Nilai:   raport.Nilai,
		}
		err := conn.Debug().Create(&addNilai).Error
		if err != nil {
			return *helpers.MessageResponses(false, http.StatusBadRequest, "Cannot add User"), false
		}
		return *helpers.MessageResponses(true, 200, "Succesfully"), true
	}
	return *helpers.MessageResponses(false, http.StatusBadRequest, "Cannot add User"), false
}

func (raport *Raport) GetRaport(conn *gorm.DB, idSiswa int) (bool, []Raport) {
	var user User
	var dataRaport []Raport
	if err := conn.Debug().First(&user, idSiswa).Error; err != nil {
		return false, nil
	}
	if user.RoleID == 2 {
		if err := conn.Debug().Find(&dataRaport, idSiswa).Error; err != nil {
			return false, nil
		}
		return true, dataRaport
	}
	return false, nil
}
