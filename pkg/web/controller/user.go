package controller

import (
	view "gin-demo/pkg/common/model"
	"gin-demo/pkg/config"
	"gin-demo/pkg/service"
	"gin-demo/pkg/service/bo"
	dao_model "gin-demo/pkg/service/dao/model"
	"gin-demo/pkg/util"
	"gin-demo/pkg/web/model"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Logger defines log in controller
var Logger = config.GetLogger()

// ListUsers returns list of users
func ListUsers(c *gin.Context) {
	name := util.GetQueryStr(c, "name")
	phoneNo := util.GetQueryStr(c, "phoneNo")
	pageNo := util.GetQueryInt(c, "pageNo")
	pageSize := util.GetQueryInt(c, "pageSize")

	query := &bo.UserQueryBO{
		Name:     name,
		PhoneNo:  phoneNo,
		PageNo:   pageNo,
		PageSize: pageSize,
	}

	pager := service.UserServiceInstance.GetByCondition(query)
	res := view.Success(0, util.SuccessInfo, pager)
	c.JSON(http.StatusOK, res)
}

// AddUser creates model User in db
func AddUser(c *gin.Context) {
	var userVO model.UserVO
	var res *view.JSONResponse
	if err := c.ShouldBindJSON(&userVO); err != nil {
		res = view.Fail(-1, util.FailInfo, err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}

	status := util.StatusNormal
	user := &dao_model.User{
		Name:        userVO.Name,
		Email:       userVO.Email,
		Phone:       userVO.Phone,
		Description: userVO.Description,
		Status:      &status,
	}
	service.UserServiceInstance.Add(user)
	// Logger.Infof("Log: %s", user.Name)
	res = view.Success(0, util.SuccessInfo, true)
	c.JSON(http.StatusOK, res)
}

func UpdateUser(c *gin.Context) {
	var userVO model.UserVO
	var res *view.JSONResponse

	userID := c.Param("token")
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		res = view.Fail(-1, util.FailInfo, nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := c.ShouldBindJSON(&userVO); err != nil {
		res = view.Fail(-1, util.FailInfo, err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}

	user := service.UserServiceInstance.GetByID(&id)
	if user == nil {
		res = view.Fail(-1, "user not found", nil)
		c.JSON(http.StatusNotFound, res)
		return
	}
	user.Name = userVO.Name
	user.Email = userVO.Email
	user.Description = userVO.Description
	user.Phone = userVO.Phone

	service.UserServiceInstance.Update(user)
	res = view.Success(0, util.SuccessInfo, true)
	c.JSON(http.StatusOK, res)
}

// GetUserByID returns User by id
func GetUserByID(c *gin.Context) {
	var res *view.JSONResponse
	userID := c.Param("token")
	Logger.Infof("Request By id: %s", userID)
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		res = view.Fail(-1, util.FailInfo, nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	user := service.UserServiceInstance.GetByID(&id)
	if user == nil {
		res = view.Fail(-1, "user not found", nil)
		c.JSON(http.StatusNotFound, res)
		return
	}

	res = view.Success(0, util.SuccessInfo, user)
	c.JSON(http.StatusOK, res)
}
