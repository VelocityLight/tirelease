package controller

import (
	"fmt"

	"tirelease/internal/entity"
	"tirelease/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// Rest-API controller
func TestEntityInsert(c *gin.Context) {
	testEntity := entity.TestEntity{}
	if err := c.ShouldBindWith(&testEntity, binding.JSON); err != nil {
		c.JSON(500, err.Error())
		return
	}

	if err := repository.TestEntityInsert(&testEntity); err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, gin.H{"status": "ok"})
}

func TestEntitySelect(c *gin.Context) {
	option := entity.TestEntityOption{}
	if err := c.ShouldBindWith(&option, binding.JSON); err != nil {
		c.JSON(500, err.Error())
		return
	}

	testEntities, err := repository.TestEntitySelect(&option)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, gin.H{"data": testEntities})
}

type TestPingPongStruct struct {
	Name  string `json:"name" form:"name"`
	Other string `json:"other" form:"other"`
}

func TestPingPongGet(c *gin.Context) {
	param := TestPingPongStruct{}
	if err := c.ShouldBindWith(&param, binding.Form); err != nil {
		c.JSON(500, err.Error())
		return
	}
	fmt.Println(param)

	c.JSON(200, gin.H{"status": "OK", "data": param})
}

func TestPingPongPost(c *gin.Context) {
	param := TestPingPongStruct{}
	if err := c.ShouldBindWith(&param, binding.JSON); err != nil {
		c.JSON(500, err.Error())
		return
	}
	fmt.Println(param)

	c.JSON(200, gin.H{"status": "OK", "data": param})
}
