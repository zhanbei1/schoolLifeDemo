package logic

import (
	"context"
	"time"

	"schoolLifeDemo/internal/config"
	"schoolLifeDemo/internal/model/sql/model"
	"schoolLifeDemo/internal/svc"
	"schoolLifeDemo/internal/types"
	"schoolLifeDemo/internal/utils"

	"github.com/form3tech-oss/jwt-go"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginInfo) (resp *types.LoginResponse, err error) {
	var res types.LoginResponse
	// todo: add your logic here and delete this line
	// 1、查询这个用户
	studentModel := model.NewStudentBaseInfoModel(sqlx.NewSqlConn("mysql", l.svcCtx.Config.Mysql.DataSource), l.svcCtx.Config.CacheRedis)
	// 2、密码对比
	result, err := studentModel.FindOneByStudentNo(l.ctx, req.SchoolNum)
	if err != nil {
		res.Code = config.ErrorCode.Code()
		res.Message = err.Error()
		return &res, nil
	}
	compareResult := utils.ComparePassword(req.PassWord, result.Password.String)

	var data types.LoginResponseInfo

	if compareResult {
		//生成JWT
		now := time.Now().Unix()
		accessExpire := l.svcCtx.Config.Auth.AccessExpire
		jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, req.SchoolNum)
		if err != nil {
			return nil, err
		}
		// ---end---
		data = types.LoginResponseInfo{
			SchoolNo:     result.StudentNo,
			PetName:      result.PetName.String,
			AccessToken:  jwtToken,
			AccessExpire: now + accessExpire,
			RefreshAfter: now + accessExpire/2,
		}
		res.Code = config.SuccessCode.Code()
		res.Message = config.SuccessCode.String()
		res.Data = data
	} else {
		res.Code = config.LoginError.Code()
		res.Message = config.LoginError.String()
	}

	return &res, nil
}

// JWT token加密
func (l *LoginLogic) getJwtToken(secretKey string, iat int64, seconds int64, userId string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	return token.SignedString([]byte(secretKey))
}
