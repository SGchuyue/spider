package v1

import (
	"github.com/gin-gonic/gin"
)
type Title struct{}
func  NewTitle() Title {
	return Title{}
}

func (t Title) Get(c *gin.Context) {
	
}

func (t Title) Create(c *gin.Context) {}

func (t Title) Update(c *gin.Context) {}

func (t Title) Delete(c *gin.Context) {}

