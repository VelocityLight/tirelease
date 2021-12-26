package controller

import (
	"tirelease/internal/entity"
	"tirelease/internal/repository"
	"tirelease/internal/service"

	"github.com/gin-gonic/gin"
)

type TriageOption struct {
	Owner string `json:"owner"`
	Repo string `json:"repo"`
}

// Rest-API controller
func InsertTriageItems(c *gin.Context) {
	// Params
	triageOption := &TriageOption{}
	c.BindJSON(triageOption)

	// Action
	triageItems, err := service.CollectTriageItemByRepo(triageOption.Owner, triageOption.Repo)
	if nil != err {
		c.JSON(500, err.Error())
		return
	}

	err2 := service.SavaTriageItems(triageItems)
	if (nil != err2) {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, gin.H{"status": "ok"})
}

func SelectTriageItems(c *gin.Context) {
	// Params
	option := entity.TriageItemOption{}
	c.BindJSON(&option)

	// Action
	triageItems, err := repository.TriageItemSelect(&option)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, gin.H{"data": triageItems})
}
