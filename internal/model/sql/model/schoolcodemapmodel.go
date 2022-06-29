/*
 * @Author: Desmond.zhan
 * @Date: 2022-05-15 11:49:27
 * @Description:
 */
package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SchoolCodeMapModel = (*customSchoolCodeMapModel)(nil)

type (
	// SchoolCodeMapModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSchoolCodeMapModel.
	SchoolCodeMapModel interface {
		schoolCodeMapModel
	}

	customSchoolCodeMapModel struct {
		*defaultSchoolCodeMapModel
	}
)

// NewSchoolCodeMapModel returns a model for the database table.
func NewSchoolCodeMapModel(conn sqlx.SqlConn, c cache.CacheConf) SchoolCodeMapModel {
	return &customSchoolCodeMapModel{
		defaultSchoolCodeMapModel: newSchoolCodeMapModel(conn, c),
	}
}

func (m *defaultSchoolCodeMapModel) FindAll(ctx context.Context) ([]*SchoolCodeMap, error) {
	schoolCodeMapIdKey := fmt.Sprintf("%s%v", cacheSchoolCodeMapIdPrefix, "allCode")
	codes := make([]*SchoolCodeMap, 0)
	query := fmt.Sprintf("select %s from %s where status = 0", schoolCodeMapRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, codes, schoolCodeMapIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		return conn.QueryRowsCtx(ctx, v, query)
	})
	switch err {
	case nil:
		return codes, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
