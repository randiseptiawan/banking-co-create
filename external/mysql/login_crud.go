package mysql

func InitMigrationLog() (client *client, err error) {
	init_db := Init_db()
	client = NewMysqlClient(*init_db)
	if err != nil {
		return nil, err
	}

	client.dbConnection.AutoMigrate(User{})
	return client, err
}

func GetUser(email string) (user User, err error) {
	db, err := InitMigrationLog()
	if err != nil {
		return User{}, err
	}
	res := db.dbConnection.First(&user, "email=?", email)
	if res.Error != nil {
		return User{}, err
	}
	return user, nil
}
