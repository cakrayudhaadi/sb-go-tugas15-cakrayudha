package services

import (
	"net/http"
	"strconv"
	"tugas15/commons"
	"tugas15/databases/connection"
	"tugas15/repositories"
	"tugas15/structs"

	"github.com/gin-gonic/gin"
)

func CreateBioskop(ctx *gin.Context) {
	var newBioskop structs.Bioskop

	if err := ctx.ShouldBindJSON(&newBioskop); err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, "Parameter yang dimasukkan salah")
		return
	}
	if commons.IsStringEmpty(newBioskop.Nama) {
		commons.ResponseError(ctx, http.StatusBadRequest, "Parameter nama harus diisi")
		return
	}
	if commons.IsStringEmpty(newBioskop.Lokasi) {
		commons.ResponseError(ctx, http.StatusBadRequest, "Parameter lokasi harus diisi")
		return
	}

	err := repositories.CreateBioskop(connection.DBConnection, newBioskop)
	if err != nil {
		commons.ResponseError(ctx, http.StatusInternalServerError, "Data bioskop gagal dibuat")
	} else {
		commons.ResponseSuccessWithoutData(ctx, http.StatusCreated, "Data bioskop berhasil dibuat")
	}
}

func GetAllBioskop(ctx *gin.Context) {
	bioskops, err := repositories.GetAllBioskop(connection.DBConnection)
	if err != nil {
		commons.ResponseError(ctx, http.StatusInternalServerError, "Data bioskop gagal diambil")
	} else if len(bioskops) == 0 {
		commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "Data bioskop kosong")
	} else {
		commons.ResponseSuccessWithData(ctx, http.StatusOK, "Data bioskop berhasil diambil", bioskops)
	}
}

func GetBioskop(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	bioskop, err := repositories.GetBioskop(connection.DBConnection, id)
	if bioskop.ID == -1 {
		commons.ResponseSuccessWithoutData(ctx, http.StatusOK, "Data bioskop tidak ada")
	} else if err != nil {
		commons.ResponseError(ctx, http.StatusInternalServerError, "Data bioskop gagal diambil")
	} else {
		commons.ResponseSuccessWithData(ctx, http.StatusOK, "Data bioskop berhasil diambil", bioskop)
	}
}

func UpdateBioskop(ctx *gin.Context) {
	var newBioskop structs.Bioskop
	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := ctx.ShouldBindJSON(&newBioskop); err != nil {
		commons.ResponseError(ctx, http.StatusBadRequest, "Parameter yang dimasukkan salah")
		return
	}
	if commons.IsStringEmpty(newBioskop.Nama) {
		commons.ResponseError(ctx, http.StatusBadRequest, "Parameter nama harus diisi")
		return
	}
	if commons.IsStringEmpty(newBioskop.Lokasi) {
		commons.ResponseError(ctx, http.StatusBadRequest, "Parameter lokasi harus diisi")
		return
	}

	newBioskop.ID = id

	err := repositories.UpdateBioskop(connection.DBConnection, newBioskop)
	if err != nil {
		commons.ResponseError(ctx, http.StatusInternalServerError, "Data bioskop gagal diubah")
	} else {
		commons.ResponseSuccessWithoutData(ctx, http.StatusCreated, "Data bioskop berhasil diubah")
	}
}

func DeleteBioskop(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := repositories.DeleteBioskop(connection.DBConnection, id)
	if err != nil {
		commons.ResponseError(ctx, http.StatusInternalServerError, "Data bioskop gagal dihapus")
	} else {
		commons.ResponseSuccessWithoutData(ctx, http.StatusCreated, "Data bioskop berhasil dihapus")
	}
}
