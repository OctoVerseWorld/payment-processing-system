package responses

import (
	"github.com/go-chi/render"
	"net/http"
)

func RenderResponse(w http.ResponseWriter, r *http.Request, res interface{}, status int) {
	resp := Response{
		ResultResponse: res,
	}
	render.Status(r, status)
	render.JSON(w, r, &resp)
}
