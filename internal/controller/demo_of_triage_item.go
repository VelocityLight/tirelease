package controller

import (
	"tirelease/internal/entity"
	"tirelease/internal/repository"
	"tirelease/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type TriageOption struct {
	Owner string `json:"owner"`
	Repo  string `json:"repo"`
}

type TriageOprate struct {
	Owner  string   `json:"owner"`
	Repo   string   `json:"repo"`
	Number int      `json:"number"`
	Lables []string `json:"labels"`
}

// Rest-API controller
func InsertTriageItems(c *gin.Context) {
	// Params
	triageOption := TriageOption{}
	c.ShouldBindWith(&triageOption, binding.JSON)

	// Action
	triageItems, err := service.CollectTriageItemByRepo(triageOption.Owner, triageOption.Repo)
	if nil != err {
		c.JSON(500, err.Error())
		return
	}

	err2 := service.SavaTriageItems(triageItems)
	if nil != err2 {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, gin.H{"status": "ok"})
}

func SelectTriageItems(c *gin.Context) {
	// Params
	option := entity.TriageItemOption{}
	c.ShouldBindWith(&option, binding.JSON)

	// Action
	triageItems, err := repository.TriageItemSelect(&option)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, gin.H{"data": triageItems})
}

func AddLabelsToIssue(c *gin.Context) {
	// Params
	operate := TriageOprate{}
	c.ShouldBindWith(&operate, binding.JSON)

	// Action
	err := service.AddLabelOfAccept(operate.Owner, operate.Repo, operate.Number, operate.Lables)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, gin.H{"status": "ok"})
}
