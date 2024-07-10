package controller

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Club struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_Club     string             `bson:"nama_club,omitempty" json:"nama_club,omitempty"`
	Liga          string             `bson:"liga,omitempty" json:"liga,omitempty"`
	Tahun_Berdiri int                `bson:"tahun_berdiri,omitempty" json:"tahun_berdiri,omitempty"`
	Stadion       string             `bson:"stadion,omitempty" json:"stadion,omitempty"`
	Manajer       string             `bson:"manajer,omitempty" json:"manajer,omitempty"`
	Jumlah_Pemain int                `bson:"jumlah_pemain,omitempty" json:"jumlah_pemain,omitempty"`
	Logo          string             `bson:"logo,omitempty" json:"logo,omitempty"`
}

type Pemain struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_Pemain   string             `bson:"nama_pemain,omitempty" json:"nama_pemain,omitempty"`
	Tim           Club               `bson:"tim,omitempty" json:"tim,omitempty"` 
	Posisi        string             `bson:"posisi,omitempty" json:"posisi,omitempty"`
	Tinggi        float64            `bson:"tinggi,omitempty" json:"tinggi,omitempty"`
	Berat         float64            `bson:"berat,omitempty" json:"berat,omitempty"`
	Tanggal_Lahir primitive.DateTime `bson:"tanggal_lahir,omitempty" json:"tanggal_lahir,omitempty"`
	Negara        string             `bson:"negara,omitempty" json:"negara,omitempty"`
	No_Punggung   int                `bson:"no_punggung,omitempty" json:"no_punggung,omitempty"`
}


type Admin struct{
	ID   		primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Username	string 		       `bson:"username,omitempty" json:"username,omitempty"`
	Password 	string			   `bson:"password,omitempty" json:"password,omitempty"`
}

type ReqPemain struct {
	NamaPemain    string             `bson:"nama_pemain,omitempty" json:"nama_pemain,omitempty" example:"haaland"`
	Tim           string             `bson:"tim,omitempty" json:"tim,omitempty" example:"Manchester City"`
	Posisi        string             `bson:"posisi,omitempty" json:"posisi,omitempty" example:"Striker"`
	Tinggi        float64            `bson:"tinggi,omitempty" json:"tinggi,omitempty" example:"192.0"`
	Berat         float64            `bson:"berat,omitempty" json:"berat,omitempty" example:"80.0"`
	TanggalLahir  primitive.DateTime `bson:"tanggal_lahir,omitempty" json:"tanggal_lahir,omitempty" swaggertype:"string" example:"2024-09-01T00:00:00Z" format:"date-time"`
	Negara        string             `bson:"negara,omitempty" json:"negara,omitempty" example:"Denmark"`
	NoPunggung    int                `bson:"no_punggung,omitempty" json:"no_punggung,omitempty" example:"9"`
}

type ReqClub struct {
	Nama_Club     string             `bson:"nama_club,omitempty" json:"nama_club,omitempty"`
	Liga          string             `bson:"liga,omitempty" json:"liga,omitempty"`
	Tahun_Berdiri int                `bson:"tahun_berdiri,omitempty" json:"tahun_berdiri,omitempty"`
	Stadion       string             `bson:"stadion,omitempty" json:"stadion,omitempty"`
	Manajer       string             `bson:"manajer,omitempty" json:"manajer,omitempty"`
	Jumlah_Pemain int                `bson:"jumlah_pemain,omitempty" json:"jumlah_pemain,omitempty"`
	Logo          string             `bson:"logo,omitempty" json:"logo,omitempty"`
}