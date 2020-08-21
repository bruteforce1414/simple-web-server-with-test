package say

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"strings"
)

type Note struct {
	ID string `json:"key"`
	Text string `json:"value"`
}

var storage = map[string]Note{
	"833745ca-dc0f-4087-a9d0-3643a3b51cbc": {ID:"833745ca-dc0f-4087-a9d0-3643a3b51cbc", Text: "dog1"},
	"d0aa78dc-d67f-4fa1-bdfa-86e1ff8ad7c7": {ID:"d0aa78dc-d67f-4fa1-bdfa-86e1ff8ad7c7", Text:"dog2"},
	"2751cc6d-6e39-4ae0-af78-98a55edcaed9": {ID:"2751cc6d-6e39-4ae0-af78-98a55edcaed9",Text: "dog3"},
}


type All struct{
	Notes [] Note
}

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

	all:=All{
		Notes: make([]Note,0),
	}
	for _, v :=range storage{
		all.Notes=append(all.Notes, Note{
			ID:   v.ID,
			Text: v.Text,
		})
	}

	fmt.Println("all.Notes", all.Notes)

	return c.JSON(http.StatusOK, all.Notes)
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

