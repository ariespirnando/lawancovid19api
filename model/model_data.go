package model

type Negara struct {
	Nama string `json:"name"`
	Positif string `json:"positif"`
	Sembuh string `json:"sembuh"`
	Meninggal string `json:"meninggal"` 
}

type Provinsi struct {
	Attributes struct { 
		KasusMeni int64  `json:"Kasus_Meni"`
		KasusPosi int64  `json:"Kasus_Posi"`
		KasusSemb int64  `json:"Kasus_Semb"`
		KodeProvi int64  `json:"Kode_Provi"`
		Provinsi  string `json:"Provinsi"`
	} `json:"attributes"`
}

type SaveorUpdatedata struct{
	KodeProvinsi int64
	NamaProvinsi string
	KasusPositif int64
	KasusMeninggal int64
	KasusSehat int64
}