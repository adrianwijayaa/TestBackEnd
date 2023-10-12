package dtokabupaten

type Response struct {
	Status      bool      `json:"status"`
	Data        []Kabupaten `json:"data"`
	Total       int       `json:"total"`
	Search      string    `json:"search"`
	Limit       int       `json:"limit"`
	CurrentPage int       `json:"current_page"`
	TotalPage   int       `json:"total_page"`
	Next        bool      `json:"next"`
}

type DetailResponse struct {
	Status  bool              `json:"status"`
	Data    DetailKabupatenData `json:"data"`
	Message string            `json:"message"`
}