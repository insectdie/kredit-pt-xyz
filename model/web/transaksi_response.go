package web

type TransaksiResponse struct {
	No_kontrak  string `json:"no_kontrak"`
	Nik         string `json:"nik"`
	Otr         int    `json:"otr"`
	Admin_fee   int    `json:"admin_fee"`
	Jml_cicilan int    `json:"jml_cicilan"`
	Jml_bunga   int    `json:"jml_bunga"`
	Nama_asset  string `json:"nama_asset"`
}
