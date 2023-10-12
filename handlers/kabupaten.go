package handlers

import (
	dtokabupaten "TestBackEnd/dto/kabupaten"
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

var listKabupaten []dtokabupaten.Kabupaten
var listKecamatan []dtokabupaten.Kecamatan

func InitDataKabupaten() {
	dataKabupaten, err := ioutil.ReadFile("agregatpenduduk2022.json")
	if err != nil {
		log.Fatal("file JSON gagal terbaca:", err)
	}

	var data struct {
		Data [][]interface{} `json:"data"`
	}
	err = json.Unmarshal(dataKabupaten, &data)
	if err != nil {
		log.Fatalf("data JSON gagal terurai: %v", err)
	}

	for _, row := range data.Data {
		if len(row) == 6 {
			idFloat, _ := row[0].(float64)
			id := int(idFloat)
			tahun, _ := row[1].(string)
			provinsi, _ := row[2].(string)
			nama, _ := row[3].(string)
			kecamatan, _ := row[4].(string)
			populasiFloat, _ := row[5].(float64)
			populasi := int(populasiFloat)

			kabupaten := dtokabupaten.Kabupaten{
				Id:        id,
				Tahun:     tahun,
				Provinsi:  provinsi,
				Name:      nama,
				Kecamatan: kecamatan,
				Populasi:  populasi,
			}
			listKabupaten = append(listKabupaten, kabupaten)
		}
	}
}

func ListKabupaten(c *fiber.Ctx) error {
	search := c.Query("search")
	page := c.Query("page")

	limitListData := 10
	currentPage := 1

	if page != "" {
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status": false,
				"message": "Invalid page parameter",
			})
		}
		currentPage = pageInt
	}

	startIndex := (currentPage - 1) * limitListData
	endIndex := startIndex + limitListData
	var filteredData []dtokabupaten.Kabupaten

	if search != "" {
		for _, kabupaten := range listKabupaten {
			if strings.Contains(strings.ToLower(kabupaten.Name), strings.ToLower(search)) {
				filteredData = append(filteredData, kabupaten)
			}
		}
	} else {
		filteredData = listKabupaten
	}

	totalPages := int(math.Ceil(float64(len(filteredData)) / float64(limitListData)))
	nextPage := currentPage < totalPages

	response := dtokabupaten.Response{
		Status:      true,
		Data:        filteredData[startIndex:endIndex],
		Total:       len(filteredData),
		Search:      search,
		Limit:       limitListData,
		CurrentPage: currentPage,
		TotalPage:   totalPages,
		Next:        nextPage,
	}

	return c.JSON(response)
}

func DetailKabupatenData(c *fiber.Ctx) error {
	response := dtokabupaten.DetailResponse{
		Status: true,
		Data: dtokabupaten.DetailKabupatenData{},
	}

	var inputData struct {
		IDKab string `json:"id_kab"`
	}
	
	if err := c.BodyParser(&inputData); err != nil {
		response.Status = false
		response.Message = "Input JSON Gagal"
		return c.JSON(response)
	}
	
	var kabupaten dtokabupaten.Kabupaten
	found := false
	idKab, err := strconv.Atoi(strings.Split(inputData.IDKab, "-")[0])
	if err != nil {
		response.Status = false
		response.Message = "Format ID Kabupaten tidak valid"
		return c.JSON(response)
	}
	for _, k := range listKabupaten {
		if k.Id == idKab {
			kabupaten = k
			found = true
			break
		}
	}

	if !found {
		response.Status = false
		response.Message = "Kabupaten tidak ditemukan"
		return c.JSON(response)
	}


	detailKabupaten := dtokabupaten.DetailKabupatenData{
		NamaKabupaten: kabupaten.Name,
		NamaProvinsi:  kabupaten.Provinsi,
		TotalKecamatan: 0,
		TotalPenduduk:  0,
		ListKecamatan:  []dtokabupaten.Kecamatan{},
	}

		for _, kecamatan := range listKecamatan {
			if strings.HasPrefix(kecamatan.IDKec, inputData.IDKab+"-") {
				detailKabupaten.TotalKecamatan++
				detailKabupaten.TotalPenduduk += kecamatan.Total
				detailKabupaten.ListKecamatan = append(detailKabupaten.ListKecamatan, dtokabupaten.Kecamatan{
					IDKec:         kecamatan.IDKec,
					NamaKecamatan: kecamatan.NamaKecamatan,
					Total:         kecamatan.Total,
				})
			}
		}		
	

	sort.Slice(detailKabupaten.ListKecamatan, func(i, j int) bool {
		return detailKabupaten.ListKecamatan[i].Total > detailKabupaten.ListKecamatan[j].Total
	})

	response.Data = detailKabupaten

	return c.JSON(response)
}