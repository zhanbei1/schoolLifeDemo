package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"schoolLifeDemo/internal/logic"
	"schoolLifeDemo/internal/svc"
)

func SchoolCodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewSchoolCodeLogic(r.Context(), svcCtx)
		resp, err := l.SchoolCode()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
