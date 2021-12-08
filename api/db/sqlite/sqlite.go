package sqlite

import (
	"github.com/dgkg/gomasterclasses/api/db"
	"github.com/dgkg/gomasterclasses/api/model"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var _ db.Storage = &SQLite{}

type SQLite struct {
	conn *gorm.DB
}

func New(filename string) *SQLite {
	db, err := gorm.Open(sqlite.Open(filename), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// create tables.
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}

	return &SQLite{
		conn: db,
	}
}

func (db *SQLite) CreateUser(u *model.User) (*model.User, error) {
	u.UUID = uuid.NewString()
	return u, db.conn.Model(model.User{}).Create(u).Error
}

func (db *SQLite) Delete(uuid string) error {
	return db.conn.Where("uuid = ?", uuid).Delete(model.User{}).Error
}

func (db *SQLite) Update(uuid string, data map[string]interface{}) (*model.User, error) {
	var u model.User
	return &u, db.conn.Model(model.User{}).Where("uuid = ?", uuid).Updates(data).First(&u).Error
}

func (db *SQLite) Get(uuid string) (*model.User, error) {
	var u model.User
	return &u, db.conn.Where("uuid = ?", uuid).First(&u).Error
}

func (db *SQLite) GetAll() (map[string]*model.User, error) {
	var us []model.User
	err := db.conn.Find(&us).Error
	if err != nil {
		return nil, err
	}
	res := make(map[string]*model.User)
	for k := range us {
		res[us[k].UUID] = &us[k]
	}
	return res, nil
}

func (db *SQLite) GetByEmail(email string) (*model.User, error) {
	var u model.User
	return &u, db.conn.Where("email = ?", email).First(&u).Error
}
