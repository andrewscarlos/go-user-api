package repository

import (
	"go-web/db"
	"go-web/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const UsersCollection = "users"

type UserRepositoryInterface interface {
	Find(id string) (user *models.User, err error)
	FindAll() (user []*models.User, err error)
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(id string) error
}

type usersRepository struct {
	c *mgo.Collection
}

func NewUserRepository(conn db.Connection) UserRepositoryInterface {
	return &usersRepository{conn.DB().C(UsersCollection)}
}

func (r *usersRepository) Create(user *models.User) error {
	return r.c.Insert(user)
}

func (r *usersRepository) Update(user *models.User) error {
	return r.c.UpdateId(user.Id, user)
}

func (r *usersRepository) Find(id string) (user *models.User, err error) {
	err = r.c.FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

func (r *usersRepository) FindAll() (users []*models.User, err error) {
	err = r.c.Find(bson.M{}).All(&users)
	return users, err
}
func (r *usersRepository) Delete(id string) error {
	return r.c.RemoveId(bson.ObjectIdHex(id))
}
