package web

type KonsumenResponse struct {
	Nik           string `json:"nik"`
	Full_name     string `json:"full_name"`
	Legal_name    string `json:"legal_name"`
	Tempat_lahir  string `json:"tempat_lahir"`
	Tanggal_lahir string `json:"tanggal_lahir"`
	Gaji          int    `json:"gaji"`
	Foto_ktp      string `json:"foto_ktp"`
	Foto_selfie   string `json:"foto_selfie"`
}
