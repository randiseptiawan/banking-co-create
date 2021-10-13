package mysql

func InitMigrationArt() (client *client, err error) {
	init_db := Init_db()
	client = NewMysqlClient(*init_db)
	if err != nil {
		return nil, err
	}

	client.dbConnection.AutoMigrate(Artikel{})
	return client, err
}

func CreateArtikel(artikel *Artikel) (err error) {
	db, err := InitMigrationArt()
	if err != nil {
		return err
	}
	db.dbConnection.Create(artikel)
	return nil
}

func ReadAllArtikel() (artikel []Artikel, err error) {
	db, _ := InitMigrationArt()

	db.dbConnection.Find(&artikel)
	return artikel, nil
}

func ReadArtikel(id uint64) (artikel Artikel, err error) {
	db, err := InitMigrationArt()
	if err != nil {
		return Artikel{}, err
	}
	res := db.dbConnection.First(&artikel, "id=?", id)
	if res.Error != nil {
		return Artikel{}, err
	}
	return artikel, nil
}

func DeleteArtikel(id uint64) (err error) {
	db, err := InitMigrationArt()
	if err != nil {
		return err
	}
	var artikel Artikel
	db.dbConnection.Model(&artikel).Where("id=?", id).Delete(&artikel)
	return nil
}
