// Copyright 2015 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package admin

import (
	"net/http"

	log "gopkg.in/clog.v1"

	api "github.com/gogs/go-gogs-client"

	"github.com/nameisalreadytakenexception/gogs/models"
	"github.com/nameisalreadytakenexception/gogs/models/errors"
	"github.com/nameisalreadytakenexception/gogs/pkg/context"
	"github.com/nameisalreadytakenexception/gogs/pkg/mailer"
	"github.com/nameisalreadytakenexception/gogs/pkg/setting"
	"github.com/nameisalreadytakenexception/gogs/routes/api/v1/user"
)

func parseLoginSource(c *context.APIContext, u *models.User, sourceID int64, loginName string) {
	if sourceID == 0 {
		return
	}

	source, err := models.GetLoginSourceByID(sourceID)
	if err != nil {
		if errors.IsLoginSourceNotExist(err) {
			c.Error(http.StatusUnprocessableEntity, "", err)
		} else {
			c.ServerError("GetLoginSourceByID", err)
		}
		return
	}

	u.LoginType = source.Type
	u.LoginSource = source.ID
	u.LoginName = loginName
}

func CreateUser(c *context.APIContext, form api.CreateUserOption) {
	u := &models.User{
		Name:      form.Username,
		FullName:  form.FullName,
		Email:     form.Email,
		Passwd:    form.Password,
		IsActive:  true,
		LoginType: models.LOGIN_PLAIN,
	}

	parseLoginSource(c, u, form.SourceID, form.LoginName)
	if c.Written() {
		return
	}

	if err := models.CreateUser(u); err != nil {
		if models.IsErrUserAlreadyExist(err) ||
			models.IsErrEmailAlreadyUsed(err) ||
			models.IsErrNameReserved(err) ||
			models.IsErrNamePatternNotAllowed(err) {
			c.Error(http.StatusUnprocessableEntity, "", err)
		} else {
			c.ServerError("CreateUser", err)
		}
		return
	}
	log.Trace("Account created by admin %q: %s", c.User.Name, u.Name)

	// Send email notification.
	if form.SendNotify && setting.MailService != nil {
		mailer.SendRegisterNotifyMail(c.Context.Context, models.NewMailerUser(u))
	}

	c.JSON(http.StatusCreated, u.APIFormat())
}

func EditUser(c *context.APIContext, form api.EditUserOption) {
	u := user.GetUserByParams(c)
	if c.Written() {
		return
	}

	parseLoginSource(c, u, form.SourceID, form.LoginName)
	if c.Written() {
		return
	}

	if len(form.Password) > 0 {
		u.Passwd = form.Password
		var err error
		if u.Salt, err = models.GetUserSalt(); err != nil {
			c.ServerError("GetUserSalt", err)
			return
		}
		u.EncodePasswd()
	}

	u.LoginName = form.LoginName
	u.FullName = form.FullName
	u.Email = form.Email
	u.Website = form.Website
	u.Location = form.Location
	if form.Active != nil {
		u.IsActive = *form.Active
	}
	if form.Admin != nil {
		u.IsAdmin = *form.Admin
	}
	if form.AllowGitHook != nil {
		u.AllowGitHook = *form.AllowGitHook
	}
	if form.AllowImportLocal != nil {
		u.AllowImportLocal = *form.AllowImportLocal
	}
	if form.MaxRepoCreation != nil {
		u.MaxRepoCreation = *form.MaxRepoCreation
	}

	if err := models.UpdateUser(u); err != nil {
		if models.IsErrEmailAlreadyUsed(err) {
			c.Error(http.StatusUnprocessableEntity, "", err)
		} else {
			c.ServerError("UpdateUser", err)
		}
		return
	}
	log.Trace("Account profile updated by admin %q: %s", c.User.Name, u.Name)

	c.JSONSuccess(u.APIFormat())
}

func DeleteUser(c *context.APIContext) {
	u := user.GetUserByParams(c)
	if c.Written() {
		return
	}

	if err := models.DeleteUser(u); err != nil {
		if models.IsErrUserOwnRepos(err) ||
			models.IsErrUserHasOrgs(err) {
			c.Error(http.StatusUnprocessableEntity, "", err)
		} else {
			c.ServerError("DeleteUser", err)
		}
		return
	}
	log.Trace("Account deleted by admin(%s): %s", c.User.Name, u.Name)

	c.NoContent()
}

func CreatePublicKey(c *context.APIContext, form api.CreateKeyOption) {
	u := user.GetUserByParams(c)
	if c.Written() {
		return
	}
	user.CreateUserPublicKey(c, form, u.ID)
}
