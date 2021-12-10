package mongodb

import "github.com/dgkg/gomasterclasses/api/model"

func (db *MongoDB) CreateUser(u *model.User) (*model.User, error) {
	return nil, nil
}

func (db *MongoDB) Delete(uuid string) error {
	return nil
}
func (db *MongoDB) Update(uuid string, data map[string]interface{}) (*model.User, error) {
	return nil, nil
}
func (db *MongoDB) Get(uuid string) (*model.User, error) {
	return nil, nil
}
func (db *MongoDB) GetAll() (map[string]*model.User, error) {
	return nil, nil
}
func (db *MongoDB) GetByEmail(email string) (*model.User, error) {
	return nil, nil
}
