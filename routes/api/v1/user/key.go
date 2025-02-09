// Copyright 2015 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package user

import (
	api "github.com/gogs/go-gogs-client"
	"net/http"

	"github.com/nameisalreadytakenexception/gogs/models"
	"github.com/nameisalreadytakenexception/gogs/models/errors"
	"github.com/nameisalreadytakenexception/gogs/pkg/context"
	"github.com/nameisalreadytakenexception/gogs/pkg/setting"
	"github.com/nameisalreadytakenexception/gogs/routes/api/v1/convert"
	"github.com/nameisalreadytakenexception/gogs/routes/api/v1/repo"
)

func GetUserByParamsName(c *context.APIContext, name string) *models.User {
	user, err := models.GetUserByName(c.Params(name))
	if err != nil {
		c.NotFoundOrServerError("GetUserByName", errors.IsUserNotExist, err)
		return nil
	}
	return user
}

// GetUserByParams returns user whose name is presented in URL paramenter.
func GetUserByParams(c *context.APIContext) *models.User {
	return GetUserByParamsName(c, ":username")
}

func composePublicKeysAPILink() string {
	return setting.AppURL + "api/v1/user/keys/"
}

func listPublicKeys(c *context.APIContext, uid int64) {
	keys, err := models.ListPublicKeys(uid)
	if err != nil {
		c.ServerError("ListPublicKeys", err)
		return
	}

	apiLink := composePublicKeysAPILink()
	apiKeys := make([]*api.PublicKey, len(keys))
	for i := range keys {
		apiKeys[i] = convert.ToPublicKey(apiLink, keys[i])
	}

	c.JSONSuccess(&apiKeys)
}

func ListMyPublicKeys(c *context.APIContext) {
	listPublicKeys(c, c.User.ID)
}

func ListPublicKeys(c *context.APIContext) {
	user := GetUserByParams(c)
	if c.Written() {
		return
	}
	listPublicKeys(c, user.ID)
}

func GetPublicKey(c *context.APIContext) {
	key, err := models.GetPublicKeyByID(c.ParamsInt64(":id"))
	if err != nil {
		c.NotFoundOrServerError("GetPublicKeyByID", models.IsErrKeyNotExist, err)
		return
	}

	apiLink := composePublicKeysAPILink()
	c.JSONSuccess(convert.ToPublicKey(apiLink, key))
}

// CreateUserPublicKey creates new public key to given user by ID.
func CreateUserPublicKey(c *context.APIContext, form api.CreateKeyOption, uid int64) {
	content, err := models.CheckPublicKeyString(form.Key)
	if err != nil {
		repo.HandleCheckKeyStringError(c, err)
		return
	}

	key, err := models.AddPublicKey(uid, form.Title, content)
	if err != nil {
		repo.HandleAddKeyError(c, err)
		return
	}
	apiLink := composePublicKeysAPILink()
	c.JSON(http.StatusCreated, convert.ToPublicKey(apiLink, key))
}

func CreatePublicKey(c *context.APIContext, form api.CreateKeyOption) {
	CreateUserPublicKey(c, form, c.User.ID)
}

func DeletePublicKey(c *context.APIContext) {
	if err := models.DeletePublicKey(c.User, c.ParamsInt64(":id")); err != nil {
		if models.IsErrKeyAccessDenied(err) {
			c.Error(http.StatusForbidden, "", "you do not have access to this key")
		} else {
			c.Error(http.StatusInternalServerError, "DeletePublicKey", err)
		}
		return
	}

	c.NoContent()
}
