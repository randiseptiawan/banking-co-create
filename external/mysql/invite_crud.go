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
