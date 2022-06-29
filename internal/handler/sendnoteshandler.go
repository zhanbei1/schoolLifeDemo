package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"schoolLifeDemo/internal/logic"
	"schoolLifeDemo/internal/svc"
	"schoolLifeDemo/internal/types"
)

func SendNotesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SchoolNotesRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewSendNotesLogic(r.Context(), svcCtx)
		resp, err := l.SendNotes(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
