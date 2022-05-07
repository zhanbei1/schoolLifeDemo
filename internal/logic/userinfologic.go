package logic

import (
	"context"

	"schoolLifeDemo/internal/config"
	"schoolLifeDemo/internal/model/sql/model"
	"schoolLifeDemo/internal/svc"
	"schoolLifeDemo/internal/types"
	"schoolLifeDemo/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoRequest) (resp *types.UserInfoResponse, err error) {
	// todo: add your logic here and delete this line
	var response types.UserInfoResponse

	schoolNo := req.SchoolNum
	studentModel := model.NewStudentBaseInfoModel(sqlx.NewSqlConn("mysql", l.svcCtx.Config.Mysql.DataSource), l.svcCtx.Config.CacheRedis)

	result, err := studentModel.FindOneByStudentNo(l.ctx, schoolNo)
	if err != nil {
		response.Code = config.ErrorCode.Code()
		response.Message = err.Error()
		return &response, nil
	}

	studentInfoBody := types.UserInfoBody{
		SchoolNum:  result.StudentNo,
		SchoolName: result.SchoolName,
		IconUrl:    "Null",
		PetName:    result.PetName.String,
		PhoneNo:    utils.PhoneNoCryption(result.PhoneNo.String),
		Birthday:   result.Birthday.String,
		Gender:     int(result.Gender.Int64),
		Grade:      int(result.Grade.Int64),
		Role:       1,
	}

	response.Data = studentInfoBody
	response.Code = config.SuccessCode.Code()
	response.Message = config.SuccessCode.String()

	return &response, nil
}
