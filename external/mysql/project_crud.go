package mysql

func InitMigrationPro() (client *client, err error) {
	init_db := Init_db()
	client = NewMysqlClient(*init_db)
	if err != nil {
		return nil, err
	}

	client.dbConnection.AutoMigrate(Project{})
	return client, err
}

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

func DeleteProject(id uint64) (err error) {
	db, err := InitMigrationPro()
	if err != nil {
		return err
	}
	var project Project
	db.dbConnection.Model(&project).Where("id=?", id).Delete(&project)
	return nil
}
