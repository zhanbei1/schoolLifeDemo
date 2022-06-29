package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"schoolLifeDemo/internal/logic"
	"schoolLifeDemo/internal/svc"
	"schoolLifeDemo/internal/types"
)

func UserOrdersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserOrdersRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserOrdersLogic(r.Context(), svcCtx)
		resp, err := l.UserOrders(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
