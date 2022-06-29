package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"schoolLifeDemo/internal/logic"
	"schoolLifeDemo/internal/svc"
	"schoolLifeDemo/internal/types"
)

func NotesListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.NotesListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewNotesListLogic(r.Context(), svcCtx)
		resp, err := l.NotesList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
