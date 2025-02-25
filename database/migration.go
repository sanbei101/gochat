package database

import "go-chat/model"

func Migrate() error {
	if err := PG.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error; err != nil {
		return err
	}
	if err := PG.Migrator().DropTable(&model.TextMessage{}, &model.User{}); err != nil {
		return err
	}
	if err := PG.AutoMigrate(&model.TextMessage{}, &model.User{}); err != nil {
		return err
	}
	return nil
}
