package dtokabupaten

type Kabupaten struct {
	Id 		  int `json:"id_kab"`
	Name 	  string `json:"nama"`
	Populasi  int    `json:"populasi"`
	Tahun     string `json:"tahun"`
	Provinsi  string `json:"provinsi"`
	Kecamatan string `json:"kecamatan"`
}

type DetailKabupatenData struct {
	NamaKabupaten  string       `json:"nama_kabupaten"`
	NamaProvinsi   string       `json:"nama_provinsi"`
	TotalKecamatan int          `json:"total_kecamatan"`
	TotalPenduduk  int          `json:"total_penduduk"`
	ListKecamatan  []Kecamatan  `json:"list_kecamatan"`
}

type Kecamatan struct {
	IDKec        string `json:"id_kec"`
	NamaKecamatan string `json:"nama_kecamatan"`
	Total        int    `json:"total"`
}

