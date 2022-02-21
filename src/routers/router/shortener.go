package router

import (
	"net/http"
	"shortener/src/controllers"
)

var routersShortener = []Router{

	{
		URI:                    "/shorteners",
		Method:                 http.MethodGet,
		Function:               controllers.GetAllShortener,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/shorteners/user",
		Method:                 http.MethodGet,
		Function:               controllers.FindByUserID,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/shortener/{shortenerID}",
		Method:                 http.MethodGet,
		Function:               controllers.GetByIDShortener,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/shortener",
		Method:                 http.MethodPost,
		Function:               controllers.CreateShortener,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/shortener/{shortenerID}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeleteByID,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/{shortener}",
		Method:                 http.MethodGet,
		Function:               controllers.RedirectURL,
		RequiresAuthentication: false,
	},
}
