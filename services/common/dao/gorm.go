package dao

import (
	"errors"
	"time"

	"github.com/0x726f6f6b6965/mono-service/services/common/consts"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDB(dbType consts.DBType, dsn string) (*gorm.DB, error) {
	var dialector gorm.Dialector
	switch dbType {
	case consts.PSQL:
		dialector = postgres.Open(dsn)
	default:
		return nil, errors.New("not support the db type")
	}
	db, err := gorm.Open(dialector, &gorm.Config{
		NowFunc: func() time.Time {
			ti, _ := time.LoadLocation("Asia/Taipei")
			return time.Now().In(ti)
		},
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}
