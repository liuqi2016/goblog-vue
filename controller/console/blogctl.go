// Solo.go - A small and beautiful blogging platform written in golang.
// Copyright (C) 2017, b3log.org
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package console

import (
	"net/http"
	"strconv"

	"github.com/b3log/solo.go/service"
	"github.com/b3log/solo.go/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func BlogSwitchCtl(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	idArg := c.Param("id")
	blogID, err := strconv.Atoi(idArg)
	if nil != err {
		result.Code = -1

		return
	}

	session := sessions.Default(c)
	userID := session.Get("id").(uint)

	userBlogs := service.User.GetUserBlogs(userID)
	if 1 > len(userBlogs) {
		result.Code = -1
		result.Msg = "switch blog failed"

		return
	}

	role := -1
	for _, userBlog := range userBlogs {
		if userBlog.ID == uint(blogID) {
			role = userBlog.UserRole

			break
		}
	}

	if -1 == role {
		result.Code = -1
		result.Msg = "switch blog failed"

		return
	}

	result.Data = role
}