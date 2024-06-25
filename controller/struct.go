package controller

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Pemain struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" example:"123456789"`
	NamaPemain    string             `bson:"nama_pemain,omitempty" json:"nama_pemain,omitempty" example:"haaland"`
	Tim           string             `bson:"tim,omitempty" json:"tim,omitempty" example:"Manchester City"`
	Posisi        string             `bson:"posisi,omitempty" json:"posisi,omitempty" example:"Striker"`
	Tinggi        float64            `bson:"tinggi,omitempty" json:"tinggi,omitempty" example:"192.0"`
	Berat         float64            `bson:"berat,omitempty" json:"berat,omitempty" example:"80.0"`
	TanggalLahir  primitive.DateTime `bson:"tanggal_lahir,omitempty" json:"tanggal_lahir,omitempty" swaggertype:"string" example:"2024-09-01T00:00:00Z" format:"date-time"`
	Negara        string             `bson:"negara,omitempty" json:"negara,omitempty" example:"Denmark"`
	NoPunggung    int                `bson:"no_punggung,omitempty" json:"no_punggung,omitempty" example:"9"`
}


type Admin struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Username string             `bson:"username,omitempty" json:"username,omitempty"`
	Password string             `bson:"password,omitempty" json:"password,omitempty"`
}

type ReqPemain struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" example:"123456789"`
	NamaPemain    string             `bson:"nama_pemain,omitempty" json:"nama_pemain,omitempty" example:"haaland"`
	Tim           string             `bson:"tim,omitempty" json:"tim,omitempty" example:"Manchester City"`
	Posisi        string             `bson:"posisi,omitempty" json:"posisi,omitempty" example:"Striker"`
	Tinggi        float64            `bson:"tinggi,omitempty" json:"tinggi,omitempty" example:"192.0"`
	Berat         float64            `bson:"berat,omitempty" json:"berat,omitempty" example:"80.0"`
	TanggalLahir  primitive.DateTime `bson:"tanggal_lahir,omitempty" json:"tanggal_lahir,omitempty" swaggertype:"string" example:"2024-09-01T00:00:00Z" format:"date-time"`
	Negara        string             `bson:"negara,omitempty" json:"negara,omitempty" example:"Denmark"`
	NoPunggung    int                `bson:"no_punggung,omitempty" json:"no_punggung,omitempty" example:"9"`
}
