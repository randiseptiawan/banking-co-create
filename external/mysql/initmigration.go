package mysql

func InitMigrationUser() (client *client, err error) {
	init_db := Init_db()
	client = NewMysqlClient(*init_db)
	if err != nil {
		return nil, err
	}

	client.dbConnection.AutoMigrate(User{})
	return client, err
}

func InitMigrationPro() (client *client, err error) {
	init_db := Init_db()
	client = NewMysqlClient(*init_db)
	if err != nil {
		return nil, err
	}

	client.dbConnection.AutoMigrate(Project{})
	return client, err
}

func InitMigrationArt() (client *client, err error) {
	init_db := Init_db()
	client = NewMysqlClient(*init_db)
	if err != nil {
		return nil, err
	}

	client.dbConnection.AutoMigrate(Artikel{})
	return client, err
}

func InitMigrationInv() (client *client, err error) {
	init_db := Init_db()
	client = NewMysqlClient(*init_db)
	if err != nil {
		return nil, err
	}

	client.dbConnection.AutoMigrate(Invited{})
	return client, err
}

func InitMigrationCol() (client *client, err error) {
	init_db := Init_db()
	client = NewMysqlClient(*init_db)
	if err != nil {
		return nil, err
	}

	client.dbConnection.AutoMigrate(Collaborator{})
	return client, err
}
