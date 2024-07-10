package controller

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Club struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" example:"12345"`
	Nama_Club     string             `bson:"nama_club,omitempty" json:"nama_club,omitempty" example:"Real Madrid"`
	Liga          string             `bson:"liga,omitempty" json:"liga,omitempty" example:"La liga"`
	Tahun_Berdiri int                `bson:"tahun_berdiri,omitempty" json:"tahun_berdiri,omitempty" example:"1985"`
	Stadion       string             `bson:"stadion,omitempty" json:"stadion,omitempty" example:"Si Jalak"`
	Manajer       string             `bson:"manajer,omitempty" json:"manajer,omitempty" example:"Jajang"`
	Jumlah_Pemain int                `bson:"jumlah_pemain,omitempty" json:"jumlah_pemain,omitempty" example:"11"`
	Logo          string             `bson:"logo,omitempty" json:"logo,omitempty" example:"https://upload.wikimedia.org/wikipedia/id/4/43/Al-Nassr_fc.png"`
}

type Pemain struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" example:"12345"`
	NamaPemain    string             `bson:"nama_pemain,omitempty" json:"nama_pemain,omitempty" example:"haaland"`
	Tim           ReqClub            `bson:"tim,omitempty" json:"tim,omitempty"`
	Posisi        string             `bson:"posisi,omitempty" json:"posisi,omitempty" example:"Striker"`
	Tinggi        float64            `bson:"tinggi,omitempty" json:"tinggi,omitempty" example:"192.0"`
	Berat         float64            `bson:"berat,omitempty" json:"berat,omitempty" example:"80.0"`
	TanggalLahir  primitive.DateTime `bson:"tanggal_lahir,omitempty" json:"tanggal_lahir,omitempty" swaggertype:"string" example:"2024-09-01T00:00:00Z" format:"date-time"`
	Negara        string             `bson:"negara,omitempty" json:"negara,omitempty" example:"Denmark"`
	NoPunggung    int                `bson:"no_punggung,omitempty" json:"no_punggung,omitempty" example:"9"`
}


type Admin struct{
	ID   		primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Username	string 		       `bson:"username,omitempty" json:"username,omitempty"`
	Password 	string			   `bson:"password,omitempty" json:"password,omitempty"`
}

type ReqPemain struct {
	NamaPemain    string             `bson:"nama_pemain,omitempty" json:"nama_pemain,omitempty" example:"haaland"`
	Tim           ReqClub            `bson:"tim,omitempty" json:"tim,omitempty"`
	Posisi        string             `bson:"posisi,omitempty" json:"posisi,omitempty" example:"Striker"`
	Tinggi        float64            `bson:"tinggi,omitempty" json:"tinggi,omitempty" example:"192.0"`
	Berat         float64            `bson:"berat,omitempty" json:"berat,omitempty" example:"80.0"`
	TanggalLahir  primitive.DateTime `bson:"tanggal_lahir,omitempty" json:"tanggal_lahir,omitempty" swaggertype:"string" example:"2024-09-01T00:00:00Z" format:"date-time"`
	Negara        string             `bson:"negara,omitempty" json:"negara,omitempty" example:"Denmark"`
	NoPunggung    int                `bson:"no_punggung,omitempty" json:"no_punggung,omitempty" example:"9"`
}

type ReqClub struct {
	Nama_Club     string             `bson:"nama_club,omitempty" json:"nama_club,omitempty" example:"Real Madrid"`
	Liga          string             `bson:"liga,omitempty" json:"liga,omitempty" example:"La liga"`
	Tahun_Berdiri int                `bson:"tahun_berdiri,omitempty" json:"tahun_berdiri,omitempty" example:"1985"`
	Stadion       string             `bson:"stadion,omitempty" json:"stadion,omitempty" example:"Si Jalak"`
	Manajer       string             `bson:"manajer,omitempty" json:"manajer,omitempty" example:"Jajang"`
	Jumlah_Pemain int                `bson:"jumlah_pemain,omitempty" json:"jumlah_pemain,omitempty" example:"11"`
	Logo          string             `bson:"logo,omitempty" json:"logo,omitempty" example:"https://upload.wikimedia.org/wikipedia/id/4/43/Al-Nassr_fc.png"`
}