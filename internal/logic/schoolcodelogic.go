/*
 * @Author: Desmond.zhan
 * @Date: 2022-05-15 11:34:01
 * @Description:
 */
package logic

import (
	"context"

	"schoolLifeDemo/internal/config"
	"schoolLifeDemo/internal/model/sql/model"
	"schoolLifeDemo/internal/svc"
	"schoolLifeDemo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type SchoolCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSchoolCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SchoolCodeLogic {
	return &SchoolCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SchoolCodeLogic) SchoolCode() (resp []*model.SchoolCodeMap, err error) {
	var response types.CommonResponse
	// todo: add your logic here and delete this line
	schoolCodeModel := model.NewSchoolCodeMapModel(sqlx.NewSqlConn("mysql", l.svcCtx.Config.Mysql.DataSource), l.svcCtx.Config.CacheRedis)
	schoolCodeInfo, err := schoolCodeModel.FindAll(l.ctx)

	if err != nil {
		response.Code = config.ErrorCode.Code()
		response.Message = err.Error()
	}

	return schoolCodeInfo, nil
}
