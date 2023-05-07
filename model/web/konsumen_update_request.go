package web

type KonsumenUpdateRequest struct {
	Nik         int    `validate:"required"`
	Gaji        int    `validate:"required" json:"gaji"`
	Foto_selfie string `validate:"required,max=255,min=1" json:"foto_selfie"`
}
