package module

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/module"
	module1 "github.com/NpoolPlatform/smoketest-middleware/pkg/mw/module"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistModule(ctx context.Context, in *npool.ExistModuleRequest) (*npool.ExistModuleResponse, error) {
	handler, err := module1.NewHandler(
		ctx,
		module1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistModule",
			"In", in,
			"Error", err,
		)
		return &npool.ExistModuleResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistModule(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistModule",
			"In", in,
			"Error", err,
		)
		return &npool.ExistModuleResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistModuleResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistModuleConds(ctx context.Context, in *npool.ExistModuleCondsRequest) (*npool.ExistModuleCondsResponse, error) {
	handler, err := module1.NewHandler(
		ctx,
		module1.WithConds(in.Conds),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistModule",
			"In", in,
			"Error", err,
		)
		return &npool.ExistModuleCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistModuleConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistModule",
			"In", in,
			"Error", err,
		)
		return &npool.ExistModuleCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistModuleCondsResponse{
		Info: exist,
	}, nil
}
