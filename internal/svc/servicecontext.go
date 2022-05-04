/*
 * @Author: Desmond.zhan
 * @Date: 2022-05-04 15:07:53
 * @Description:
 */
package svc

import (
	"schoolLifeDemo/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
