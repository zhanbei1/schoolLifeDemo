package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ StudentBaseInfoModel = (*customStudentBaseInfoModel)(nil)

type (
	// StudentBaseInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStudentBaseInfoModel.
	StudentBaseInfoModel interface {
		studentBaseInfoModel
	}

	customStudentBaseInfoModel struct {
		*defaultStudentBaseInfoModel
	}
)

// NewStudentBaseInfoModel returns a model for the database table.
func NewStudentBaseInfoModel(conn sqlx.SqlConn, c cache.CacheConf) StudentBaseInfoModel {
	return &customStudentBaseInfoModel{
		defaultStudentBaseInfoModel: newStudentBaseInfoModel(conn, c),
	}
}
