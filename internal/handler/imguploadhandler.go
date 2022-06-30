package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"schoolLifeDemo/internal/config"
	"schoolLifeDemo/internal/logic"
	"schoolLifeDemo/internal/svc"
	"schoolLifeDemo/internal/types"
)

func ImgUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewImgUploadLogic(r.Context(), svcCtx, r, w)
		//var req types.UploadFileRequest
		//if err := httpx.Parse(r, &req); err != nil {
		//	httpx.Error(w, err)
		//	return
		//}
		filePath, err := l.ImgUpload()
		var res = &types.UploadFileResponse{}
		if err != nil {
			res.Code = config.InvalidFileError.Code()
			res.Message = err.Error()
		} else {
			res.Code = config.SuccessCode.Code()
			res.Message = config.SuccessCode.String()
			res.FilePath = filePath
		}
		httpx.OkJson(w, res)
	}
}
