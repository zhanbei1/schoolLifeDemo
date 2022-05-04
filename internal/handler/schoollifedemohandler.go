/*
 * @Author: Desmond.zhan
 * @Date: 2022-05-04 15:07:53
 * @Description:
 */
package handler

import (
	"net/http"

	"schoolLifeDemo/internal/logic"
	"schoolLifeDemo/internal/svc"
	"schoolLifeDemo/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func SchoolLifeDemoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewSchoolLifeDemoLogic(r.Context(), svcCtx)
		resp, err := l.SchoolLifeDemo(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
