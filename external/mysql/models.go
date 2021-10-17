package mysql

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	NamaLengkap      string `gorm:"not null" json:"namaLengkap"`
	Email            string `gorm:"size:50;not null;unique" json:"email"`
	Password         string `gorm:"size:100;not null" json:"password"`
	TopikDiminati    string `gorm:"not null" json:"topikDiminati"`
	EnrollmentStatus string `gorm:"size:20;not null" json:"enrollmentStatus"`
	RoleStatus       string `gorm:"size:10;not null" json:"roleStatus"`
}

type Project struct {
	gorm.Model
	KategoriProject      string   `gorm:"not null" json:"kategoriProject"`
	NamaProject          string   `gorm:"not null" json:"namaProject"`
	TanggalMulai         string   `gorm:"not null" json:"tanggalMulai"`
	LinkTrello           string   `gorm:"not null" json:"linkTrello"`
	DeskripsiProject     string   `gorm:"not null" json:"deskripsiProject"`
	ProjectAdminId       uint     `gorm:"foreignkey" json:"-"`
	InvitedUserName      []string `gorm:"-" json:"invitedUserName"`
	InvitedUserEmail     []string `gorm:"-" json:"invitedUserEmail"`
	CollaboratorUserName []string `gorm:"-" json:"CollaboratorUserName"`
	CollaboratorEmail    []string `gorm:"-" json:"CollaboratorEmail"`
	ProjectAdminName     string   `gorm:"-" json:"projectAdminName"`
	ProjectAdminEmail    string   `gorm:"-" json:"projectAdminEmail"`
}

type Invited struct {
	gorm.Model
	InvitedUserId uint `gorm:"foreignkey" json:"invitedUsername"`
	ProjectId     uint `gorm:"foreignkey" json:"projectId"`
}

type Collaborator struct {
	gorm.Model
	CollaboratorUserId uint `gorm:"foreignkey" json:"collaboratorUsername"`
	ProjectId          uint `gorm:"foreignkey" json:"projectId"`
}

type Artikel struct {
	gorm.Model
	Kategori   string `gorm:"not null" json:"kategori"`
	Judul      string `gorm:"not null" json:"judul"`
	IsiArtikel string `gorm:"not null" json:"isiArtikel"`
	UserId     uint   `gorm:"foreignkey" json:"-"`
	UserName   string `gorm:"-" json:"UserName"`
	UserEmail  string `gorm:"-" json:"UserEmail"`
}
