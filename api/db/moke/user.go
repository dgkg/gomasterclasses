package moke

import (
	"time"

	"github.com/dgkg/gomasterclasses/api/db"
	"github.com/dgkg/gomasterclasses/api/model"
	"github.com/google/uuid"
)

func (m *MokeDB) CreateUser(u *model.User) (*model.User, error) {
	if _, err := m.GetByEmail(u.Email); err == nil {
		return nil, db.ErrUserAlreadyExists
	}
	u.UUID = uuid.NewString()
	u.CreateAt = time.Now()

	m.listUser[u.UUID] = u
	return u, nil
}

func (m *MokeDB) Delete(uuid string) error {
	_, ok := m.listUser[uuid]
	if !ok {
		return db.ErrUserNotFound
	}
	delete(m.listUser, uuid)
	return nil
}

func (m *MokeDB) Update(uuid string, data map[string]interface{}) (*model.User, error) {
	u, ok := m.listUser[uuid]
	if !ok {
		return nil, db.ErrUserNotFound
	}

	for k := range data {
		switch k {
		case "first_name":
			u.FirstName = data[k].(string)
		case "last_name":
			u.LastName = data[k].(string)
		case "pass":
			u.Password = data[k].(model.Password)
		}
	}

	return u, nil
}

func (m *MokeDB) Get(uuid string) (*model.User, error) {
	u, ok := m.listUser[uuid]
	if !ok {
		return nil, db.ErrUserNotFound
	}
	return u, nil
}

func (m *MokeDB) GetAll() (map[string]*model.User, error) {
	return m.listUser, nil
}

func (m *MokeDB) GetByEmail(email string) (*model.User, error) {
	for k := range m.listUser {
		if m.listUser[k].Email == email {
			return m.listUser[k], nil
		}
	}
	return nil, db.ErrUserNotFound
}
