package mysql

func DeleteInvitedUser(UserId uint64, ProjectId uint64) (err error) {
	db, err := InitMigrationInv()
	if err != nil {
		return err
	}
	var invited Invited
	db.dbConnection.Model(&invited).Where("invited_user_id=? AND project_id=?", UserId, ProjectId).Delete(&invited)
	return nil
}

func CreateCollaborator(collaborator *Collaborator) (err error) {
	db, err := InitMigrationCol()
	if err != nil {
		return err
	}
	db.dbConnection.Create(collaborator)
	return nil
}

func InviteUser(invited *Invited) (err error) {
	db, err := InitMigrationInv()
	if err != nil {
		return err
	}
	db.dbConnection.Create(invited)
	return nil
}

func ProjectInvited(userId uint) (project Project, err error) {
	db, err := InitMigrationPro()
	if err != nil {
		return Project{}, err
	}
	db.dbConnection.Find(&project).Where("invited_user_id = ?", userId)
	return project, nil
}

func GetInvitedUser(id uint64) (invited Invited, err error) {
	db, err := InitMigrationInv()
	if err != nil {
		return Invited{}, err
	}
	res := db.dbConnection.First(&invited, "project_id=?", id)
	if res.Error != nil {
		return Invited{}, res.Error
	}
	return invited, err
}

func GetCollaboratorUser(id uint64) (collaborator Collaborator, err error) {
	db, err := InitMigrationCol()
	if err != nil {
		return Collaborator{}, err
	}
	res := db.dbConnection.First(&collaborator, "project_id=?", id)
	if res.Error != nil {
		return Collaborator{}, res.Error
	}
	return collaborator, err
}

func GetInvitedUsername(id uint) (username []Username, err error) {
	db, _ := InitMigrationInv()
	// if res.Error != nil {
	// 	return User{}, res.Error
	// }
	// var invited Invited
	db.dbConnection.Table("users").Select("users.nama_lengkap, users.email, inviteds.project_id, inviteds.deleted_at").Joins("right join inviteds on inviteds.invited_user_id = users.id").Where("project_id = ? AND inviteds.deleted_at is NULL", id).Scan(&username)

	return username, nil
}

func GetCollaboratedUsername(id uint) (username []Username, err error) {
	db, _ := InitMigrationUser()
	// if res.Error != nil {
	// 	return User{}, res.Error
	// }
	db.dbConnection.Table("users").Select("users.nama_lengkap, users.email, collaborators.project_id,collaborators.deleted_at").Joins("right join collaborators on collaborators.collaborator_user_id = users.id").Where("project_id = ? AND collaborators.deleted_at is NULL", id).Scan(&username)

	return username, nil
}

func GetProjectAdmin(id uint) (username Username, err error) {
	db, _ := InitMigrationUser()
	// if res.Error != nil {
	// 	return User{}, res.Error
	// }
	db.dbConnection.Table("users").Select("users.nama_lengkap, users.email, projects.id").Joins("right join projects on projects.project_admin_id = users.id").Where("projects.id = ?", id).Scan(&username)

	return username, nil
}
