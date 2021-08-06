package svc

import (
	config2 "go_zero_mall/admin/internal/config"
)

type ServiceContext struct {
	Config config2.Config
}

func NewServiceContext(c config2.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
