package repository

//
//import (
//	"errors"
//	"fmt"
//	"github.com/Sokol111/educational-go/internal/model"
//	"github.com/google/uuid"
//	"sync"
//	"time"
//)
//
//type UserMemoryRepository struct {
//	users sync.Map
//}
//
//func NewUserMemoryRepository() *UserMemoryRepository {
//	return &UserMemoryRepository{}
//}
//
//func (r *UserMemoryRepository) GetById(id string) (model.User, error) {
//	if user, ok := r.users.Load(id); ok {
//		return user.(model.User), nil
//	}
//
//	return model.User{}, fmt.Errorf("failed to find user by id [%v]", id)
//}
//
//func (r *UserMemoryRepository) GetByName(name string) (found model.User, err error) {
//	r.users.Range(func(_ any, u any) bool {
//		user := u.(model.User)
//		if user.Name == name {
//			found = user
//			return false
//		}
//		return true
//	})
//	if found.ID == "" {
//		err = fmt.Errorf("failed to find user by name [%v]", name)
//	}
//	return found, err
//}
//
//func (r *UserMemoryRepository) Create(user model.User) (model.User, error) {
//	user.ID = uuid.NewString()
//	user.Version = 1
//	user.CreatedDate = time.Now().UTC()
//	user.LastModifiedDate = user.CreatedDate
//	r.users.Store(user.ID, user)
//	return user, nil
//}
//
//func (r *UserMemoryRepository) Update(user model.User) (found model.User, err error) {
//	if u, ok := r.users.Load(user.ID); ok {
//		found = u.(model.User)
//		if found.Version != user.Version {
//			return model.User{}, errors.New("failed to update user because of different versions")
//		}
//		found.Version++
//		found.Name = user.Name
//		found.LastModifiedDate = time.Now().UTC()
//		r.users.Store(found.ID, found)
//	} else {
//		err = fmt.Errorf("failed to update user because couldn't find it by id [%v]", user.ID)
//	}
//	return found, err
//}
//
//func (r *UserMemoryRepository) GetUsers() (s []model.User, err error) {
//	r.users.Range(func(_ any, u any) bool {
//		s = append(s, u.(model.User))
//		return true
//	})
//
//	return s, err
//}
