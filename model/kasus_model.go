package model

type ReqKasusProvinsi struct{
	KodeProvinsi string `form:"kode_provinsi" json:"kode_provinsi" binding:"required"`
}
type Respkasusprovinsi struct{
	KodeProvinsi string `json:"kode_provinsi"`
	NamaProvinsi string `json:"nama_provinsi"`
	KasusPositif string `json:"kasus_positif"`
	KasusSembuh string `json:"kasus_sembuh"`
	KasusMeninggal string `json:"kasus_meninggal"`
	Grafik []RespGrafik `json:"grafik,omitempty"`
}
type RespGrafik struct{
	KasusPositif string `json:"kasus_positif"`
	Tanggal string `json:"tanggal"`
}