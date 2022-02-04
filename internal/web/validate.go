package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func StrNotEmpty(value string, desc string, c *gin.Context) error {
	if len(value) <= 0 {
		return fmt.Errorf("%s不可为空", desc)
	}
	return nil
}
