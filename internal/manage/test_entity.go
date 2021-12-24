package manage

import (
	"fmt"
	"time"

	"tirelease/commons/database"
	"tirelease/internal/entity"

	"github.com/pkg/errors"
)

// Implement
func TestEntityInsert(testEntity *entity.TestEntity) error {
	testEntity.CreateTime = time.Now()
	testEntity.UpdateTime = time.Now()

	if err := database.DBConn.DB.Create(&testEntity).Error; err != nil {
		return errors.Wrap(err, fmt.Sprintf("create test entity: %+v failed", testEntity))
	}
	return nil
}

func TestEntitySelect(option *entity.ListOption) (*[]entity.TestEntity, error) {
	var testEntities []entity.TestEntity

	if err := database.DBConn.DB.Find(&testEntities).Where(option).Error; err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("find test entity: %+v failed", option))
	}
	return &testEntities, nil
}
