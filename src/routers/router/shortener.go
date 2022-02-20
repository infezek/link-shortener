package router

import (
	"net/http"
	"shortener/src/controllers"
)

var routersShortener = []Router{
	{
		URI:                    "/shortener",
		Method:                 http.MethodPost,
		Function:               controllers.GetShortener,
		RequiresAuthentication: true,
	},
}
