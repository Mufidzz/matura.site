package restful

type User struct {
	AID       string `json:"aid,omitempty"`
	ID        string `json:"id,omitempty"`
	PwordHash string `json:"pwordhash,omitempty"`

	Nama       string `json:"nama,omitempty"`
	POB        string `json:"pob,omitempty"`
	DOB        string `json:"dob,omitempty"`
	AlamatA    string `json:"alamata,omitempty"`
	AlamatB    string `json:"alamatb,omitempty"`
	TahunLulus string    `json:"tahunlulus"`

	Status   string    `json:"status"`
	Instansi string `json:"instansi,omitempty"`

	Facebook string  `json:"fbid"`
	Line     string `json:"line,omitempty"`
	Whatsapp string  `json:"wa"`
}

type Instansi struct {
	Instansi string `json:"instansi"`
}

type Provinsi struct {
	Id string 	`json:"id"`
	Nama string `json:"provinsi"`
}

type Kabupaten struct {
	Id string 	`json:"id"`
	Nama string	`json:"kabupaten"`
}

type Kecamatan struct {
	Id string	`json:"id"`
	Nama string `json:"kecamatan"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []User
	Ins		[]Instansi
	Prov	[]Provinsi
	Kab		[]Kabupaten
	Kec		[]Kecamatan
}