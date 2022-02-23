package router

import (
	"net/http"
	"shortener/src/controllers"
)

var routersHome = []Router{
	{
		URI:                    "/",
		Method:                 http.MethodGet,
		Function:               controllers.HomeRouter,
		RequiresAuthentication: false,
	},
}
