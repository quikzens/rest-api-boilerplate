package users

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/quikzens/rest-api-boilerplate/config"
	"github.com/quikzens/rest-api-boilerplate/db"
	"github.com/quikzens/rest-api-boilerplate/helper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserRegister(c *gin.Context) {
	var req userRegisterReq
	err := c.Bind(&req)
	if err != nil {
		helper.SendBadRequest(c, err)
		return
	}

	var checkUsername db.User
	err = db.FindOne(db.UserColl, bson.M{"username": req.Username}, &checkUsername)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			helper.SendServerError(c, err)
			return
		}
	} else {
		helper.SendBadRequest(c, errors.New("username is already registered"))
		return
	}

	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		helper.SendServerError(c, err)
		return
	}

	newUser := db.User{
		Id:        uuid.NewString(),
		Username:  req.Username,
		Password:  hashedPassword,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: 0,
	}

	err = db.InsertOne(db.UserColl, newUser)
	if err != nil {
		helper.SendServerError(c, err)
		return
	}

	payload := helper.UserPayload{
		UserId:    newUser.Id,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(config.TokenDuration),
	}
	newToken, _, err := helper.CreateToken(&payload)
	if err != nil {
		helper.SendServerError(c, err)
		return
	}

	resp := userRegisterResp{
		CreatedUserId:  newUser.Id,
		Token:          newToken,
		TokenExpiredAt: payload.ExpiredAt,
	}
	helper.SendSuccess(c, resp)
}

func UserLogin(c *gin.Context) {
	var req userLoginReq
	err := c.Bind(&req)
	if err != nil {
		helper.SendBadRequest(c, err)
		return
	}

	var user db.User
	err = db.FindOne(db.UserColl, bson.M{"username": req.Username}, &user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			helper.SendBadRequest(c, errors.New("username or password is wrong"))
			return
		}
		helper.SendServerError(c, err)
		return
	}

	err = checkPassword(req.Password, user.Password)
	if err != nil {
		helper.SendBadRequest(c, errors.New("username or password is wrong"))
		return
	}

	payload := helper.UserPayload{
		UserId:    user.Id,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(config.TokenDuration),
	}
	newToken, _, err := helper.CreateToken(&payload)
	if err != nil {
		helper.SendServerError(c, err)
		return
	}

	resp := userLoginResp{
		LoggedInUserId: user.Id,
		Token:          newToken,
		TokenExpiredAt: payload.ExpiredAt,
	}
	helper.SendSuccess(c, resp)
}

func UserCheckAuth(c *gin.Context) {
	userPayload := c.MustGet(authorizationPayloadKey).(*helper.UserPayload)

	var user db.User
	err := db.FindOne(db.UserColl, bson.M{"_id": userPayload.UserId}, &user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			helper.SendBadRequest(c, errors.New("username or password is wrong"))
			return
		}
		helper.SendServerError(c, err)
		return
	}

	resp := userCheckAuthResp{
		Id:        user.Id,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	helper.SendSuccess(c, resp)
}
