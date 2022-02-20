package router

import (
	"net/http"
	"shortener/src/controllers"
)

var routersShortener = []Router{
	{
		URI:                    "/shortener",
		Method:                 http.MethodGet,
		Function:               controllers.GetAllShortener,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/shortener/{shortenerID}",
		Method:                 http.MethodGet,
		Function:               controllers.GetByIDShortener,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/shortener",
		Method:                 http.MethodPost,
		Function:               controllers.CreateShortener,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/shortener/{shortenerID}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeleteByID,
		RequiresAuthentication: true,
	},
}
