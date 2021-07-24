package logic

import (
	"context"

	"go_zero_mall/internal/svc"
	"go_zero_mall/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type Go_zero_mallLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGo_zero_mallLogic(ctx context.Context, svcCtx *svc.ServiceContext) Go_zero_mallLogic {
	return Go_zero_mallLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Go_zero_mallLogic) Go_zero_mall(req types.Request) (*types.Response, error) {
	// todo: add your logic here and delete this line

	return &types.Response{}, nil
}
