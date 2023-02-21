package models

type Source struct {
	FoodSources *string `json:"foodsources"`
}

// func MigrateSource(db *gorm.DB) error {
// 	err := db.AutoMigrate(&Source{})
// 	return err
// }