package psql

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/0x726f6f6b6965/mono-service/services/common/consts"
	"github.com/0x726f6f6b6965/mono-service/services/common/dao"
	"gorm.io/gorm"
)

var (
	UserDB *gorm.DB
)

func InitUserDB() {
	cmd := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Taipei",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PORT"))
	db, err := dao.OpenDB(consts.PSQL, cmd)
	if err != nil {
		log.Printf("database connect failed:%+v", err)
		for i := 1; i < 4; i++ {
			time.Sleep(time.Second * 5)
			db, err = dao.OpenDB(consts.PSQL, cmd)
			if err == nil {
				break
			}
			log.Printf("database re-connect failed. %d", i)
		}
	}
	UserDB = db
}
