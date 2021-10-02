package mysql

func InitMigration() (client *client) {
	init_db := Init_db()
	client = NewMysqlClient(*init_db)
	// if err != nil {
	// 	return nil, err
	// }

	client.dbConnection.AutoMigrate(Login{})
	return client
}

func Register_member(login Login) {
	db := InitMigration()
	// if err != nil {
	// 	return err
	// }
	db.dbConnection.Create(&login)
}
