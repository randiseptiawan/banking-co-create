package mysql

func GetUser(email string) (user User, err error) {
	db, err := InitMigrationUser()
	if err != nil {
		return User{}, err
	}
	res := db.dbConnection.First(&user, "email=?", email)
	if res.Error != nil {
		return User{}, res.Error
	}
	return user, nil
}

func GetUserById(id uint) (user User, err error) {
	db, err := InitMigrationUser()
	if err != nil {
		return User{}, err
	}
	res := db.dbConnection.First(&user, "id=?", id)
	if res.Error != nil {
		return User{}, res.Error
	}
	return user, nil
}

func GetAllUser() (user []User, err error) {
	db, _ := InitMigrationUser()
	db.dbConnection.Find(&user)
	return user, nil
}

func UpdateUser(id uint, user User) (err error) {
	db, err := InitMigrationUser()
	if err != nil {
		return err
	}
	db.dbConnection.Model(&user).Where("id=?", id).Updates(&user)
	return nil
}

func CreateUser(user *User) (err error) {
	db, err := InitMigrationUser()
	if err != nil {
		return err
	}
	db.dbConnection.Create(user)
	return nil
}
