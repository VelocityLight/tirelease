package service

import (
	"tirelease/internal/entity"
	"tirelease/internal/manage"

	"github.com/gin-gonic/gin"
)

func TestEntityInsert(c *gin.Context) {
	testEntity := entity.TestEntity{}
	c.BindJSON(&testEntity)
	if err := manage.TestEntityInsert(&testEntity); err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, gin.H{"status": "ok"})
}

func TestEntitySelect(c *gin.Context) {
	option := entity.ListOption{}
	c.BindJSON(&option)

	testEntity, err := manage.TestEntitySelect(&option)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, gin.H{"result": testEntity})
}
