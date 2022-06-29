package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SchoolWallNotesModel = (*customSchoolWallNotesModel)(nil)

type (
	// SchoolWallNotesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSchoolWallNotesModel.
	SchoolWallNotesModel interface {
		schoolWallNotesModel
	}

	customSchoolWallNotesModel struct {
		*defaultSchoolWallNotesModel
	}
)

// NewSchoolWallNotesModel returns a model for the database table.
func NewSchoolWallNotesModel(conn sqlx.SqlConn, c cache.CacheConf) SchoolWallNotesModel {
	return &customSchoolWallNotesModel{
		defaultSchoolWallNotesModel: newSchoolWallNotesModel(conn, c),
	}
}
