package say

import (
	"net/http"
	"strings"
)

func Say(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")


	switch message {
	case "sayhello":
		{
			w.Write([]byte("Hello World!!!"))

		}
	case "saygoodbye":
		{
			w.Write([]byte("Goodbye!!!"))
		}
	default:
		w.Write([]byte("Default message"))
	}

}

