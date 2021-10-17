package mysql

func CreateProject(project *Project) (err error) {
	db, err := InitMigrationPro()
	if err != nil {
		return err
	}
	db.dbConnection.Create(project)
	return nil
}

func ReadAllProject() (project []Project, err error) {
	db, _ := InitMigrationPro()
	db.dbConnection.Find(&project)
	return project, nil
}

func ReadProject(id uint64) (project Project, err error) {
	db, err := InitMigrationPro()
	if err != nil {
		return Project{}, err
	}
	res := db.dbConnection.First(&project, "id=?", id)
	if res.Error != nil {
		return Project{}, err
	}
	return project, err
}

func UpdateProject(id uint, project Project) (err error) {
	db, err := InitMigrationPro()
	if err != nil {
		return err
	}
	db.dbConnection.Model(&project).Where("id=?", id).Updates(&project)
	return nil
}

func DeleteProject(id uint64) (err error) {
	db, err := InitMigrationPro()
	if err != nil {
		return err
	}
	var project Project
	db.dbConnection.Model(&project).Where("id=?", id).Delete(&project)
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
