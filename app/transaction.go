package app

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func BeginTransaction(r RouteContext, db *gorm.DB, code string) (bool, *gorm.DB) {
	tx := db.Begin()
	if tx.Error != nil {
		SetAndSaveInternalError(r, "Error creating tx", tx.Error, "ME00600")
		return false, nil
	}
	return true, tx
}
func CommitTransaction(r RouteContext, tx *gorm.DB, code string) bool {
	if tx.Commit().Error != nil {
		SetAndSaveInternalError(r, "commit tx", tx.Error, code)
		return false
	}
	return true
}

func getURL(c *fiber.Ctx) string {
	url := ""
	if c != nil {
		url = c.OriginalURL()
	}
	return url
}
