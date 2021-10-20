package mysql

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	NamaLengkap      string `gorm:"not null" json:"NamaLengkap"`
	Email            string `gorm:"size:50;not null;unique" json:"Email"`
	Password         string `gorm:"size:100;not null" json:"Password"`
	TopikDiminati    string `gorm:"not null" json:"TopikDiminati"`
	EnrollmentStatus string `gorm:"size:20;not null" json:"EnrollmentStatus"`
	RoleStatus       string `gorm:"size:10;not null" json:"RoleStatus"`
}

type Project struct {
	gorm.Model
	KategoriProject      string   `gorm:"not null" json:"KategoriProject"`
	NamaProject          string   `gorm:"not null" json:"NamaProject"`
	TanggalMulai         string   `gorm:"not null" json:"TanggalMulai"`
	LinkTrello           string   `gorm:"not null" json:"LinkTrello"`
	DeskripsiProject     string   `gorm:"not null" json:"DeskripsiProject"`
	ProjectAdminId       uint     `gorm:"foreignkey" json:"-"`
	InvitedUserName      []string `gorm:"-" json:"InvitedUserName"`
	InvitedUserEmail     []string `gorm:"-" json:"InvitedUserEmail"`
	CollaboratorUserName []string `gorm:"-" json:"CollaboratorUserName"`
	CollaboratorEmail    []string `gorm:"-" json:"CollaboratorEmail"`
	ProjectAdminName     string   `gorm:"-" json:"ProjectAdminName"`
	ProjectAdminEmail    string   `gorm:"-" json:"ProjectAdminEmail"`
}

type Invited struct {
	gorm.Model
	InvitedUserId uint `gorm:"foreignkey" json:"InvitedUserId"`
	ProjectId     uint `gorm:"foreignkey" json:"ProjectId"`
}

type Collaborator struct {
	gorm.Model
	CollaboratorUserId uint `gorm:"foreignkey" json:"CollaboratorUserId"`
	ProjectId          uint `gorm:"foreignkey" json:"ProjectId"`
}

type Artikel struct {
	gorm.Model
	Kategori   string `gorm:"not null" json:"Kategori"`
	Judul      string `gorm:"not null" json:"Judul"`
	IsiArtikel string `gorm:"not null" json:"IsiArtikel"`
	UserId     uint   `gorm:"foreignkey" json:"-"`
	UserName   string `gorm:"-" json:"UserName"`
	UserEmail  string `gorm:"-" json:"UserEmail"`
}
