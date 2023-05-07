package web

type TransaksiCreateRequest struct {
	// No_kontrak  string `validate:"required" json:"no_kontrak"`
	Nik         string `validate:"required,min=16,max=16" json:"nik"`
	Otr         int    `validate:"required" json:"otr"`
	Admin_fee   int    `validate:"required" json:"admin_fee"`
	Jml_cicilan int    `validate:"required" json:"jml_cicilan"`
	Jml_bunga   int    `validate:"required" json:"jml_bunga"`
	Nama_asset  string `validate:"required" json:"nama_asset"`
}
