package web

type KonsumenCreateRequest struct {
	Nik           string `validate:"required,min=16,max=16" json:"nik"`
	Full_name     string `validate:"required,min=1,max=60" json:"full_name"`
	Legal_name    string `validate:"required,min=1,max=60" json:"legal_name"`
	Tempat_lahir  string `validate:"required,min=1,max=30" json:"tempat_lahir"`
	Tanggal_lahir string `validate:"required" json:"tanggal_lahir"`
	Gaji          int    `validate:"required" json:"gaji"`
	Foto_ktp      string `validate:"required,min=1,max=255" json:"foto_ktp"`
	Foto_selfie   string `validate:"required,min=1,max=255" json:"foto_selfie"`
}
