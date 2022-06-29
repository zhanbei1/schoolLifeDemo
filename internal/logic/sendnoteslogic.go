/*
 * @Author: Desmond.zhan
 * @Date: 2022-05-07 18:15:33
 * @Description:
 */
package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"schoolLifeDemo/internal/config"
	"schoolLifeDemo/internal/model/sql/model"
	"schoolLifeDemo/internal/svc"
	"schoolLifeDemo/internal/types"
	"schoolLifeDemo/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type SendNotesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendNotesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendNotesLogic {
	return &SendNotesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendNotesLogic) SendNotes(req *types.SchoolNotesRequest) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line
	var res types.CommonResponse

	// 从JWT中获取用户账号，然后判断用户是否活跃
	schoolNo := fmt.Sprintf("%v", l.ctx.Value("userId"))

	if schoolNo == "" {
		res.Code = config.LoginError.Code()
		res.Message = config.LoginError.String()
		return &res, nil
	}

	studentModel := model.NewStudentBaseInfoModel(sqlx.NewSqlConn("mysql", l.svcCtx.Config.Mysql.DataSource), l.svcCtx.Config.CacheRedis)
	studentInfo, err := studentModel.FindOneByStudentNo(l.ctx, schoolNo)

	if studentInfo == nil || studentInfo.IsDeleted != 0 {
		res.Code = config.NoneUser.Code()
		res.Message = config.NoneUser.String()
		return &res, nil
	}
	esClient, err := utils.GetEsClient(l.svcCtx)
	if err != nil {
		res.Code = config.ErrorCode.Code()
		res.Message = err.Error()
		return &res, nil
	}

	notesInfo := types.SchoolWallNotes{
		StudentNo:    schoolNo,
		UserPetName:  studentInfo.PetName.String,
		NoteType:     int64(req.NoteType),
		NoteTitle:    req.NoteTitle,
		NoteTag:      req.NoteTag,
		NotesContent: req.NoteContent,
		SchoolName:   studentInfo.SchoolName,
		Watermark: func(mark bool) int64 {
			if mark {
				return 1
			}
			return 0
		}(req.Watermark),
		CreateTime: time.Now().Format(time.RFC3339),
		UpdateTime: time.Now().Format(time.RFC3339),
	}
	notesStr, err := json.Marshal(notesInfo)

	esRes, err := esClient.Index(
		config.SchoolNotesESIndex,           // Index name
		strings.NewReader(string(notesStr)), // Document body
		esClient.Index.WithRefresh("true"),  // Refresh
	)
	if err != nil {
		res.Code = config.ErrorCode.Code()
		res.Message = config.ErrorCode.String()
		res.Description = err.Error()
		return &res, nil
	}
	defer esRes.Body.Close()

	res.Code = config.SuccessCode.Code()
	res.Message = config.SuccessCode.String()
	return &res, nil
}
