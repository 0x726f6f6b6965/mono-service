package user

import (
	"github.com/0x726f6f6b6965/mono-service/services/common/dao/models"
	"github.com/0x726f6f6b6965/mono-service/services/user-service/internal/dao/psql"
)

type _user struct {
	Table *models.User
}

var (
	UserObj = new(_user)
)

func (m *_user) Add(data *models.User) (uint64, error) {
	err := psql.UserDB.Table(m.Table.TableName()).Create(data).Error
	return data.ID, err
}

func (m *_user) GetUser(username string) (models.User, error) {
	var (
		result models.User
		err    error
		db     = psql.UserDB.Table(m.Table.TableName()).
			Select("id, username, password, role_type").
			Where("username = ?", username)
	)

	err = db.First(&result).Error
	return result, err
}

func (m *_user) IsUsernameExist(username string) (bool, error) {
	var (
		result models.User
		err    error
		db     = psql.UserDB.Table(m.Table.TableName()).
			Select("id, username, password").
			Where("username = ?", username)
	)

	err = db.First(&result).Error
	return result.ID > 0, err
}

func (m *_user) SetUserRole(username string, role int) error {
	var (
		err = psql.UserDB.Table(m.Table.TableName()).
			Where("username = ?", username).UpdateColumn("role_type", role).Error
	)

	return err
}
