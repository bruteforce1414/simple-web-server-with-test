package say

import (
	"github.com/labstack/echo"
	"net/http"
	"strings"
)


func TasksGet(c echo.Context) error {
	//p := new(request.TaskSample)
	var errOffset, errLimit error

	//p.Offset, errOffset = strconv.ParseUint(c.QueryParam("offset"), 0, 64)
	//p.Limit, errLimit = strconv.ParseUint(c.QueryParam("limit"), 0, 64)
	//ctx := c.Request().Context()
	if (errOffset != nil) || (errLimit != nil) {
		return c.JSON(http.StatusNotAcceptable, "Offset or Limit aren't numbers")
	}
	//r, e := h.Ts.TasksGet(ctx, p.Offset, p.Limit)
	/*if e != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}*/

	return c.JSON(http.StatusOK, "test"/*r*/)
}

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

