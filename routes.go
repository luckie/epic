package main

import (
  "net/http"
	"github.com/gorilla/mux"
)

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		IndexHandler,
	},
	Route{
		"ReadContent",
		"GET",
		"/content/{id}",
		ReadContentHandler,
	},
	Route{
		"UpdateContent",
		"PUT",
		"/content/{id}",
		UpdateContentHandler,
	},
	Route{
		"ListTags",
		"GET",
		"/tag",
		ListTagsHandler,
	},
	Route{
		"CreateTag",
		"POST",
		"/app/{app-uuid}/tag/{tag}",
		CreateTagHandler,
	},
	Route{
		"DeleteTag",
		"DELETE",
		"/tag/{id}",
		DeleteTagHandler,
	},
	Route{
		"ReadAllContentForTag",
		"GET",
		"/app/{app-uuid}/tag/{tag}",
		ReadAllContentForTagHandler,
	},
	Route{
		"AssignTagToContent",
		"POST",
		"/content/{content-uuid}/tag/{tag}",
		AssignTagToContentHandler,
	},
	Route{
		"Login",
		"POST",
		"/auth/login",
		LoginHandler,
	},
	Route{
		"Logout",
		"DELETE",
		"/auth/logout",
		LogoutHandler,
	},
	Route{
		"CreateUser",
		"POST",
		"/auth/user",
		CreateUserHandler,
	},
	Route{
		"AuthenticateToken",
		"POST",
		"/auth/token",
		AuthenticateTokenHandler,
	},
	Route{
		"NewUUID",
		"GET",
		"/auth/uuid",
		NewUUIDHandler,
	},
	Route{
		"AssetUploadURL",
		"GET", // or should this be POST?
		"/asset/url",
		AssetUploadURLHandler,
	},
	Route{
		"UserCryptoBootstrap",
		"POST",
		"/auth/crypto",
		UserCryptoBootstrapHandler,
	},
}

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	return router
}
