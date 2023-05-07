package domain

type Transaksi struct {
	No_kontrak  string
	Nik         string
	Otr         int
	Admin_fee   int
	Jml_cicilan int
	Jml_bunga   int
	Nama_asset  string
}
