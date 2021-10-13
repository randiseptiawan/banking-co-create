package mysql

func InitMigrationReg() (client *client, err error) {
	init_db := Init_db()
	client = NewMysqlClient(*init_db)
	if err != nil {
		return nil, err
	}

	client.dbConnection.AutoMigrate(User{})
	return client, err
}

func Register(user *User) (err error) {
	db, err := InitMigrationReg()
	if err != nil {
		return err
	}
	db.dbConnection.Create(user)
	return nil
}
