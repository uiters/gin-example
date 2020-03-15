package repository

import (
	"context"
	"net/http"
	"github.com/jinzhu/gorm"
	"time"
)

func initContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	return ctx, cancel
}

func getHTTPCode(err error) int {
	if err == gorm.ErrRecordNotFound {
		return http.StatusNotFound
	} else if err == gorm.ErrUnaddressable || err == gorm.ErrCantStartTransaction || err == gorm.ErrInvalidSQL || err == gorm.ErrInvalidTransaction {
		return http.StatusInternalServerError
	}
	return http.StatusBadRequest
}