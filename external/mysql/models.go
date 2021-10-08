package mysql

import (
	"gorm.io/gorm"
	"time"
)

type Login struct {
	gorm.Model
	Id_user	 string `gorm:"size:10;not null;primaryKey"`
	Username string `gorm:"size:20;not null"`
	Password string `gorm:"size:20;not null"`
	Login_as int    `gorm:"size:1;not null"`
}

type User struct {
	gorm.Model
	// Id_user 		  string `gorm:"size:10;not null"`
	Nama_Lengkap      string `gorm:"size:50;not null"`
	Username          string `gorm:"size:20;not null"`
	Password          string `gorm:"size:20;not null"`
	Topik_Diminati    string `gorm:"size:50;not null"`
	// Enrollment_Status int    `gorm:"size:1;not null"`
}

type Project struct {
	gorm.Model
	ID        			uint        `gorm:"primaryKey"`
	Kategori_project  	string 		`gorm:"size:20;not null"`
	Nama_project      	string 		`gorm:"size:20;not null"`
	Tanggal_mulai     	time.Time 
	Link_trello    	  	string 		`gorm:"size:50;not null"`
	Deskripsi_project 	string 		`gorm:"size:255;not null"`
	// Invited_user_id   []
	Project_admin 	  	string 		`gorm:"size:20;not null"`
}

type Artikel struct {
	gorm.Model
	// Id_artikel   	uint64   	`gorm:"primaryKey;autoIncrement"`
	// Posting_Date   	time.Time 
	Kategori  		string 		`gorm:"size:20;not null"`
	Judul     		string 		`gorm:"size:50;not null"`
	Isi_artikel 	string 		`gorm:"size:255;not null"`
	Penulis 	  	string 		`gorm:"size:20;not null;"`
}
