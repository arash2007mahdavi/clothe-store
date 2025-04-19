package migrations

import (
	"store/src/database"
	"store/src/database/models"
)

func Up_1() {
	database := database.GetDB()
	tables := []interface{}{}

	hat := models.Hat{}
	shoes := models.Shoes{}
	pant := models.Pant{}
	shirt := models.Shirt{}

	if !database.Migrator().HasTable(hat) {
		tables = append(tables, hat)
	}
	if !database.Migrator().HasTable(shoes) {
		tables = append(tables, shoes)
	}
	if !database.Migrator().HasTable(pant) {
		tables = append(tables, pant)
	}
	if !database.Migrator().HasTable(shirt) {
		tables = append(tables, shirt)
	}

	database.Migrator().CreateTable(tables...)
	samplehat := &models.Hat{Price: 950.0, Currency: "toman"}
	database.Create(samplehat)
}
