package mysql

import (
	"gorm.io/gorm"
)

// type Role struct {
// 	gorm.Model
// 	Status	 string `gorm:"size:10;not null"`
// }

type User struct {
	gorm.Model
	NamaLengkap      string `gorm:"size:50;not null" json:"namaLengkap"`
	Email            string `gorm:"size:50;not null;unique" json:"email"`
	Password         string `gorm:"size:100;not null" json:"password"`
	TopikDiminati    string `gorm:"size:50;not null" json:"topikDiminati"`
	EnrollmentStatus string `gorm:"size:20;not null" json:"enrollmentStatus"`
	RoleStatus       string `gorm:"size:10;not null" json:"roleStatus"`
}

type Project struct {
	gorm.Model
	KategoriProject  string `gorm:"size:20;not null" json:"kategoriProject"`
	NamaProject      string `gorm:"size:20;not null" json:"namaProject"`
	TanggalMulai     string `gorm:"not null" json:"tanggalMulai"`
	LinkTrello       string `gorm:"size:50;not null" json:"linkTrello"`
	DeskripsiProject string `gorm:"size:255;not null" json:"deskripsiProject"`
	// InvitedUsername      pq.StringArray `json:"invitedUsername"`
	// InvitedEmail         pq.StringArray `json:"invitedEmail"`
	// CollaboratorUsername pq.StringArray `json:"collaboratorUsername"`
	// CollaboratorEmail    pq.StringArray `json:"collaboratorEmail"`
	ProjectAdmin      string `gorm:"size:50;not null" json:"projectAdmin"`
	ProjectAdminEmail string `gorm:"size:50;not null" json:"projectAdminEmail"`
}

type Artikel struct {
	gorm.Model
	Kategori     string `gorm:"size:50;not null" json:"kategori"`
	Judul        string `gorm:"size:100;not null" json:"judul"`
	IsiArtikel   string `gorm:"not null" json:"isiArtikel"`
	NamaPenulis  string `gorm:"size:50;not null json:"namaPenulis;"`
	EmailPenulis string `gorm:"size:50;not null json:"emailPenulis;"`
}
