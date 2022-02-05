package router

import (
	"fmt"
	"net/http"
)

var routersSign = []Router{
	{

		URI:    "/sigin",
		Method: http.MethodPost,
		Function: func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Sign In")
		},
		RequiresAuthentication: false,
	},
	{

		URI:    "/signup",
		Method: http.MethodPost,
		Function: func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Sign Up")
		},
		RequiresAuthentication: false,
	},
}
