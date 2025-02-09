// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package user

import (
	"net/http"

	"github.com/Unknwon/com"

	api "github.com/gogs/go-gogs-client"

	"github.com/nameisalreadytakenexception/gogs/models"
	"github.com/nameisalreadytakenexception/gogs/models/errors"
	"github.com/nameisalreadytakenexception/gogs/pkg/context"
	"github.com/nameisalreadytakenexception/gogs/pkg/markup"
)

func Search(c *context.APIContext) {
	opts := &models.SearchUserOptions{
		Keyword:  c.Query("q"),
		Type:     models.USER_TYPE_INDIVIDUAL,
		PageSize: com.StrTo(c.Query("limit")).MustInt(),
	}
	if opts.PageSize == 0 {
		opts.PageSize = 10
	}

	users, _, err := models.SearchUserByName(opts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	results := make([]*api.User, len(users))
	for i := range users {
		results[i] = &api.User{
			ID:        users[i].ID,
			UserName:  users[i].Name,
			AvatarUrl: users[i].AvatarLink(),
			FullName:  markup.Sanitize(users[i].FullName),
		}
		if c.IsLogged {
			results[i].Email = users[i].Email
		}
	}

	c.JSONSuccess(map[string]interface{}{
		"ok":   true,
		"data": results,
	})
}

func GetInfo(c *context.APIContext) {
	u, err := models.GetUserByName(c.Params(":username"))
	if err != nil {
		c.NotFoundOrServerError("GetUserByName", errors.IsUserNotExist, err)
		return
	}

	// Hide user e-mail when API caller isn't signed in.
	if !c.IsLogged {
		u.Email = ""
	}
	c.JSONSuccess(u.APIFormat())
}

func GetAuthenticatedUser(c *context.APIContext) {
	c.JSONSuccess(c.User.APIFormat())
}
