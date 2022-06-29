/*
 * @Author: Desmond.zhan
 * @Date: 2022-06-02 23:36:07
 * @Description:
 */
package logic

import (
	"context"
	"encoding/json"
	"strings"

	"schoolLifeDemo/internal/config"
	"schoolLifeDemo/internal/svc"
	"schoolLifeDemo/internal/types"
	"schoolLifeDemo/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type NotesListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNotesListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NotesListLogic {
	return &NotesListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NotesListLogic) NotesList(req *types.NotesListRequest) (resp *types.NotesListResponse, err error) {
	// todo: add your logic here and delete this line
	var res types.NotesListResponse
	var r map[string]interface{}
	// 根据条件，获取ES上的notes，并按照时间排序
	// TODO： 按照热度排序
	esClient, err := utils.GetEsClient(l.svcCtx)
	if err != nil {
		res.Code = config.ErrorCode.Code()
		res.Description = err.Error()
		return &res, nil
	}

	notesStr, err := json.Marshal(req)
	result, err := esClient.Search(
		esClient.Search.WithContext(l.ctx),
		esClient.Search.WithIndex(config.SchoolNotesESIndex),
		esClient.Search.WithBody(strings.NewReader(string(notesStr))),
	)
	defer result.Body.Close()
	// 解析返回内容
	if err := json.NewDecoder(result.Body).Decode(&r); err != nil {
		logx.Error("Error parsing the response body: %s", err)
	}

	if result.IsError() || err != nil {
		res.Code = config.ErrorCode.Code()
		res.Message = err.Error()
		return &res, nil
	}

	return &res, nil
}
