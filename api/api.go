package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/swaggo/swag/example/basic/web"
)

// GetUserByID example
//
//	@Summary		get a user to the store
//	@Description	get user by ID
//	@ID				get-user-by-string
//	@Accept			json
//	@Produce		json
//	@Param			some_id	path		string					true	"Some ID"
//
// @Success		200		{object}	model.User		"ok"
// @Failure		400		{object}	model.APIError	"We need ID!!"
// @Failure		404		{object}	model.APIError	"Can not find ID"
// @Router			/api/v1/{id} [get]
func GetUserByID() echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		var pet web.Pet
		err := c.Bind(&pet)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		return c.JSON(http.StatusOK, pet)
	})
	// write your code
}

// GetStructArrayByString example
//
//	@Description	get struct array by ID
//	@ID				get-struct-array-by-string
//	@Accept			json
//	@Produce		json
//	@Param			some_id	path		string			true	"Some ID"
//	@Param			offset	query		int				true	"Offset"
//	@Param			limit	query		int				true	"Offset"
//	@Success		200		{string}	string			"ok"
//	@Failure		400		{object}	model.APIError	"We need ID!!"
//	@Failure		404		{object}	model.APIError	"Can not find ID"
//	@Router			/testapi/get-struct-array-by-string/{some_id} [get]
func GetStructArrayByString(w http.ResponseWriter, r *http.Request) {
	// write your code
}

// SaveUser example
//
//	@Summary		SaveUser
//	@Description	SaveUser
//	@Accept			json
//	@Produce		json
//
//
//	@Param			some body		model.User			true	"Some ID"
//
// @Success		200		{string}	model.User			"ok"
// @Failure		400		{object}	model.APIError	"We need ID!!"
// @Failure		404		{object}	model.APIError	"Can not find ID"
// @Router			/api/v1/users [post]
func SaveUser() echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		return nil
	})
	// write your code
}
