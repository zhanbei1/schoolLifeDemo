/*
 * @Author: Desmond.zhan
 * @Date: 2022-05-04 15:07:53
 * @Description:
 */
package logic

import (
	"context"

	"schoolLifeDemo/internal/svc"
	"schoolLifeDemo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SchoolLifeDemoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSchoolLifeDemoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SchoolLifeDemoLogic {
	return &SchoolLifeDemoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SchoolLifeDemoLogic) SchoolLifeDemo(req *types.Request) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
