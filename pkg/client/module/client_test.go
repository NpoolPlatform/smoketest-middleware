package module

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"bou.ke/monkey"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/module"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/testinit"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var (
	ret = npool.Module{
		Name:        uuid.NewString(),
		Description: uuid.NewString(),
	}
)

func createModule(t *testing.T) {
	var (
		req = &npool.ModuleReq{
			Name:        &ret.Name,
			Description: &ret.Description,
		}
	)

	info, err := CreateModule(context.Background(), req)
	if assert.Nil(t, err) {
		ret.ID = info.ID
		ret.EntID = info.EntID
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func updateModule(t *testing.T) {
	ret.Name = uuid.NewString()
	ret.Description = uuid.NewString()
	var (
		req = &npool.ModuleReq{
			ID:          &ret.ID,
			Name:        &ret.Name,
			Description: &ret.Description,
		}
	)

	info, err := UpdateModule(context.Background(), req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getModule(t *testing.T) {
	info, err := GetModule(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getModules(t *testing.T) {
	infos, _, err := GetModules(context.Background(), &npool.Conds{}, 0, 1)
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func existModule(t *testing.T) {
	exist, err := ExistModule(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.True(t, exist)
	}
}

func deleteModule(t *testing.T) {
	info, err := DeleteModule(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = GetModule(context.Background(), ret.EntID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestMainOrder(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})
	monkey.Patch(grpc2.GetGRPCConnV1, func(service string, recvMsgBytes int, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createModule", createModule)
	t.Run("updateModule", updateModule)
	t.Run("getModule", getModule)
	t.Run("existModule", existModule)
	t.Run("getModules", getModules)
	t.Run("deleteModule", deleteModule)
}
