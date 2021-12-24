package manage

import (
	"fmt"

	"tirelease/commons/database"
	"tirelease/internal/entity"

	"github.com/pkg/errors"
)

// Implement
func TestEntityInsert(testEntity *entity.TestEntity) error {
	if err := database.DBConn.DB.Create(&testEntity).Error; err != nil {
		return errors.Wrap(err, fmt.Sprintf("create test case: %+v failed", testEntity))
	}
	return nil
}
