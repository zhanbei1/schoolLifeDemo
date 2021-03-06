package handler

import (
	"net/http"

	"schoolLifeDemo/internal/logic"
	"schoolLifeDemo/internal/svc"
	"schoolLifeDemo/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterInfo
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewRegisterLogic(r.Context(), svcCtx)
		res := l.Register(&req)
		httpx.OkJson(w, res)
		// if err != nil {
		// 	httpx.OkJson()(w, err)
		// } else {
		// 	httpx.Ok(w)
		// }
	}
}
