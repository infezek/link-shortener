package router

import (
	"net/http"
	"shortener/src/controllers"
)

var routersSign = []Router{
	{
		URI:                    "/signin",
		Method:                 http.MethodPost,
		Function:               controllers.SignIn,
		RequiresAuthentication: false,
	},
	{

		URI:                    "/signup",
		Method:                 http.MethodPost,
		Function:               controllers.SignUp,
		RequiresAuthentication: false,
	},
}
