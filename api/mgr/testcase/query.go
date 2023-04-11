package testcase

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase"
	testcase1 "github.com/NpoolPlatform/smoketest-middleware/pkg/mw/testcase"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetTestCases(ctx context.Context, in *npool.GetTestCasesRequest) (*npool.GetTestCasesResponse, error) {
	handler, err := testcase1.NewHandler(
		ctx,
		testcase1.WithConds(
			in.GetConds(),
			in.GetOffset(),
			in.GetLimit(),
		),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetUsers",
			"In", in,
			"error", err,
		)
		return &npool.GetTestCasesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	infos, total, err := handler.GetTestCases(ctx)
	if err != nil {
		return &npool.GetTestCasesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetTestCasesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
