package controller

import (
	"errors"
	"fmt"
	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2"
	cek "github.com/irgifauzi/back-bola/module"
	inimodel "github.com/irgifauzi/back-bola/model"
	"github.com/irgifauzi/bp-bola/config"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"

)

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}
//p
// GetAllPemain godoc
// @Summary Get All Data Pemain.
// @Description Mengambil semua data pemain.
// @Tags Pemain
// @Accept json
// @Produce json
// @Success 200 {object} Pemain
// @Router /pemain [get]
func GetAllPemain(c *fiber.Ctx) error {
	ps := cek.GetAllDataPemain(config.Ulbimongoconn, "pemain")
	return c.JSON(ps)
}

// GetPemainID godoc
// @Summary Get By ID Data Pemain.
// @Description Ambil per ID data pemain.
// @Tags Pemain
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200 {object} Pemain
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /pemain/{id} [get]
func GetPemainID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}
	ps, err := cek.GetPemainFromID(objID, config.Ulbimongoconn, "pemain")
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("No data found for id %s", id),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error retrieving data for id %s", id),
		})
	}
	return c.JSON(ps)
}
// InsertDataPemain godoc
// @Summary Insert data pemain.
// @Description Input data pemain.
// @Tags Pemain
// @Accept json
// @Produce json
// @Param request body ReqPemain true "Payload Body [RAW]"
// @Success 200 {object} Pemain
// @Failure 400
// @Failure 500
// @Router /ins [post]
func InsertDataPemain(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var pemain inimodel.Pemain

	// Parse the request body into the pemain struct
	if err := c.BodyParser(&pemain); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Validate required fields
	if pemain.Nama_Pemain == "" || pemain.Tim.Nama_Club == "" || pemain.Tim.Liga == "" || pemain.Tim.Tahun_Berdiri == 0 ||
		pemain.Tim.Stadion == "" || pemain.Tim.Manajer == "" || pemain.Tim.Jumlah_Pemain == 0 || pemain.Tim.Logo == "" ||
		pemain.Posisi == "" || pemain.Tinggi == 0 || pemain.Berat == 0 || pemain.Tanggal_Lahir == "" ||
		pemain.Negara == "" || pemain.No_Punggung == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Fill all the form.",
		})
	}

	insertedID, err := cek.InsertPemain(db, "pemain",
		pemain.Nama_Pemain,
		pemain.Tim,
		pemain.Posisi,
		pemain.Tinggi,
		pemain.Berat,
		pemain.Tanggal_Lahir,
		pemain.Negara,
		pemain.No_Punggung)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}


// UpdateDataPemain godoc
// @Summary Update data Pemain.
// @Description Ubah data Pemain.
// @Tags Pemain
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Param request body ReqPemain true "Payload Body [RAW]"
// @Success 200 {object} Pemain
// @Failure 400
// @Failure 500
// @Router /update/{id} [put]
func UpdateDataPemain(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	var pemain inimodel.Pemain
	if err := c.BodyParser(&pemain); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Call the UpdatePemain function with the parsed ID and the Pemain object
	err = cek.UpdatePemain(db, "pemain",
		objectID,
		pemain.Nama_Pemain,
		pemain.Tim,
		pemain.Posisi,
		pemain.Tinggi,
		pemain.Berat,
		pemain.Tanggal_Lahir,
		pemain.Negara,
		pemain.No_Punggung)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Data pemain successfully updated",
	})
}
// DeletePemainByID godoc
// @Summary Delete data pemain.
// @Description Hapus data pemain.
// @Tags Pemain
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /delete/{id} [delete]
func DeletePemainByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}

	err = cek.DeletePemainByID(objID, config.Ulbimongoconn, "pemain")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error deleting data pemain for id %s", id),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("Data with id %s deleted successfully", id),
	})
}