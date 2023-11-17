package barang

import (
	"19api/model"
	"19api/utils/jwt"
	"net/http"
	"strconv"
	"strings"

	gojwt "github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type BarangController struct {
	Model model.BarangQuery
}

func (bc *BarangController) AddBarang() echo.HandlerFunc {
	return func(c echo.Context) error {
		userid, err := jwt.ExtractToken(c.Get("user").(*gojwt.Token))
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]any{
					"message": "Tidak Memiliki Kuasa Untuk Mengakses",
				})
		}

		var input = new(BarangRequest)
		if err := c.Bind(input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "input yang diberikan tidak sesuai",
			})
		}

		var inputProcess = new(model.BarangModel)
		inputProcess.NamaBarang = input.NamaBarang
		inputProcess.Harga = input.Harga
		inputProcess.Stok = input.Stok
		inputProcess.UserID = userid

		result, err := bc.Model.AddBarang(*inputProcess)

		if err != nil {
			c.Logger().Error("ERROR Register, explain:", err.Error())
			if strings.Contains(err.Error(), "duplicate") {
				return c.JSON(http.StatusBadRequest, map[string]any{
					"message": "data yang diinputkan sudah terdaftar ada sistem",
				})
			}
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "terjadi permasalahan ketika memproses data",
			})
		}

		var response = new(BarangResponse)
		response.NamaBarang = result.NamaBarang
		response.Harga = result.Harga
		response.Stok = result.Stok
		response.ID = result.ID

		return c.JSON(http.StatusCreated, map[string]any{
			"message": "success create data",
			"data":    response,
		})
	}
}

func (bc *BarangController) GetBarang() echo.HandlerFunc {
	return func(c echo.Context) error {
		barang, err := bc.Model.GetBarang()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "terjadi permasalahan ketika memproses data",
			})
		}

		return c.JSON(http.StatusOK, map[string]any{
			"message": "success get data",
			"data":    barang,
		})
	}
}

func (bc *BarangController) UpdateBarang() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		var input = new(BarangRequest)
		if err := c.Bind(input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "input yang diberikan tidak sesuai",
			})
		}

		var inputProcess = new(model.BarangModel)
		inputProcess.ID = uint(id)
		inputProcess.NamaBarang = input.NamaBarang
		inputProcess.Harga = input.Harga
		inputProcess.Stok = input.Stok

		result, err := bc.Model.UpdateBarang(*inputProcess)

		if err != nil {
			c.Logger().Error("ERROR Update Barang, explain:", err.Error())
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusBadRequest, map[string]any{
					"message": "data yang diinputkan tidak ditemukan",
				})
			}
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "terjadi permasalahan ketika memproses data",
			})
		}

		var response = new(BarangResponse)
		response.NamaBarang = result.NamaBarang
		response.Harga = result.Harga
		response.Stok = result.Stok
		response.ID = result.ID

		return c.JSON(http.StatusOK, map[string]any{
			"message": "success update data",
			"data":    response,
		})
	}
}

func (bc *BarangController) DeleteBarang() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		err := bc.Model.DeleteBarang(id)

		if err != nil {
			c.Logger().Error("ERROR Delete Barang, explain:", err.Error())
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusBadRequest, map[string]any{
					"message": "data yang diinputkan tidak ditemukan",
				})
			}
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "terjadi permasalahan ketika memproses data",
			})
		}

		return c.JSON(http.StatusOK, map[string]any{
			"message": "success delete data",
		})
	}
}