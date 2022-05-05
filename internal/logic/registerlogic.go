package logic

import (
	"context"
	"database/sql"
	"errors"

	"schoolLifeDemo/internal/config"
	"schoolLifeDemo/internal/model/sql/model"
	"schoolLifeDemo/internal/svc"
	"schoolLifeDemo/internal/types"
	"schoolLifeDemo/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterInfo) types.CommonResponse {
	// todo: add your logic here and delete this line
	var res types.CommonResponse
	var err error

	role := req.Role
	if role == 0 {
		res = StudentRegister(*l, req)
	} else if role == 1 {
		res = ManagerRegister(l, req)
	} else {
		err = errors.New("Role is not access, pleace check the value")
	}

	if err != nil {
		l.Logger.Errorf("User Register error ,  error info : %v", err.Error())
		res.Message = err.Error()
		res.Code = config.ErrorCode.Code()
	}

	return res
}

func StudentRegister(l RegisterLogic, registerInfo *types.RegisterInfo) types.CommonResponse {
	var res types.CommonResponse
	// 1、 用户校验，是否注册过，年纪是否正确，个人昵称是否重复，手机号是否重复，性别是否重复
	studentModel := model.NewStudentBaseInfoModel(sqlx.NewSqlConn("mysql", l.svcCtx.Config.Mysql.DataSource), l.svcCtx.Config.CacheRedis)

	sameStudent, err := studentModel.FindOneByStudentNo(l.ctx, registerInfo.SchoolNum)
	if err == nil && sameStudent != nil {
		res.Code = config.StudentNoRepeat.Code()
		res.Message = config.StudentNoRepeat.String()
		return res
	}

	pwHash, err := utils.UserPassWordCryption(registerInfo.PassWord)
	if err != nil {
		res.Code = config.ErrorCode.Code()
		res.Message = err.Error()
		return res
	}

	student := model.StudentBaseInfo{
		StudentNo:  registerInfo.SchoolNum,
		SchoolName: registerInfo.SchoolName,
		Grade:      sql.NullInt64{int64(registerInfo.Grade), registerInfo.Gender != 0},
		PetName:    sql.NullString{registerInfo.PetName, registerInfo.PetName != ""},
		IsDeleted:  0,
		PhoneNo:    sql.NullString{registerInfo.PhoneNo, registerInfo.PhoneNo != ""},
		Birthday:   sql.NullString{registerInfo.Birthday, registerInfo.Birthday != ""},
		Gender:     sql.NullInt64{int64(registerInfo.Gender), registerInfo.Gender != 0},
		Password:   sql.NullString{pwHash, pwHash != ""},
	}
	result, err := studentModel.Insert(l.ctx, &student)

	if err != nil {
		res.Code = config.ErrorCode.Code()
		res.Message = err.Error()
		return res
	}

	if rows, err := result.RowsAffected(); err == nil && rows > 0 {
		res.Code = config.SuccessCode.Code()
		res.Message = config.SuccessCode.String()
		return res
	} else {
		res.Code = config.UserRegisterError.Code()
		res.Message = config.UserRegisterError.String()
		return res
	}
}

func ManagerRegister(l *RegisterLogic, registerInfo *types.RegisterInfo) types.CommonResponse {
	// TODO: 未实现
	var res types.CommonResponse
	return res
}
