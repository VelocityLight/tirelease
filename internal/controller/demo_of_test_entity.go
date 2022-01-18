package controller

import (
	"tirelease/internal/entity"
	"tirelease/internal/repository"

	"github.com/gin-gonic/gin"
)

// Rest-API controller
func TestEntityInsert(c *gin.Context) {
	testEntity := entity.TestEntity{}
	c.BindJSON(&testEntity)
	if err := repository.TestEntityInsert(&testEntity); err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, gin.H{"status": "ok"})
}

func TestEntitySelect(c *gin.Context) {
	option := entity.TestEntityOption{}
	c.BindJSON(&option)

	testEntities, err := repository.TestEntitySelect(&option)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, gin.H{"data": testEntities})
}
