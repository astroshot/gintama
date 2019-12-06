package dao

import (
	// "fmt"

	"astroshot/gin-demo/pkg/service/bo"
	"astroshot/gin-demo/pkg/service/dao/model"
)

// UserDAO defines funcs to interfact with table `user`
type UserDAO interface {
	GetByID(id *int64) *model.User
	Add(user *model.User) bool
	GetByCondition(condition *bo.UserQueryBO) *bo.Pager
}

// UserDAOImpl implements interface UserDAO
type UserDAOImpl struct {
}

// GetByID returns User model by id
func (dao *UserDAOImpl) GetByID(id *int64) *model.User {
	user := model.User{}
	db.First(&user, *id)
	// or get a structure fulfilled with nil
	if user.ID == nil {
		return nil
	}

	return &user
}

// Add create User
func (dao *UserDAOImpl) Add(user *model.User) bool {
	if user == nil {
		return false
	}

	if err := db.Create(&user).Error; err != nil {
		panic(err)
	}
	return true
}

// Update User
func (dao *UserDAOImpl) Update(user *model.User) bool {
	if user == nil {
		return false
	}

	db.Update(&user)
	return true
}

// GetByCondition returns Users
func (dao *UserDAOImpl) GetByCondition(condition *bo.UserQueryBO) *bo.Pager {
	var users []model.User
	var totalCount int
	var pageCount int
	query := db

	if condition.Name != nil {
		query = query.Where("name LIKE ?", "%"+*condition.Name+"%")
	}

	if condition.PhoneNo != nil {
		query = query.Where("phone = ?", condition.PageNo)
	}

	offset := (*condition.PageNo - 1) * *condition.PageSize
	query.Find(&users).Count(&totalCount)
	query = query.Limit(*condition.PageSize)
	query = query.Offset(offset)
	query.Find(&users)

	pageCount = (totalCount + *condition.PageSize - 1) / *condition.PageSize

	pager := &bo.Pager{
		PageNo:     condition.PageNo,
		PageSize:   condition.PageSize,
		PageCount:  &pageCount,
		TotalCount: &totalCount,
		Data:       users,
	}

	return pager
}
