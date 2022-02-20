package router

import (
	"net/http"
	"shortener/src/controllers"
)

var routersShortener = []Router{
	{
		URI:                    "/shortener",
		Method:                 http.MethodGet,
		Function:               controllers.GetShortener,
		RequiresAuthentication: false,
	},
}
