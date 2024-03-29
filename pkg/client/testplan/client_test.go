package testplan

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan"
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
	ret = npool.TestPlan{
		Name:      uuid.NewString(),
		CreatedBy: uuid.NewString(),
		Executor:  uuid.NewString(),
		State:     npool.TestPlanState_WaitStart,
		StateStr:  npool.TestPlanState_WaitStart.String(),
		Deadline:  uint32(time.Now().Unix() + 60*60*24),
	}
)

func createTestPlan(t *testing.T) {
	var (
		req = &npool.TestPlanReq{
			Name:      &ret.Name,
			CreatedBy: &ret.CreatedBy,
			Executor:  &ret.Executor,
			Deadline:  &ret.Deadline,
		}
	)

	info, err := CreateTestPlan(context.Background(), req)
	if assert.Nil(t, err) {
		ret.ID = info.ID
		ret.EntID = info.EntID
		ret.Result = info.Result
		ret.ResultStr = info.ResultStr
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func updateTestPlan(t *testing.T) {
	ret.State = npool.TestPlanState_InProgress
	ret.StateStr = npool.TestPlanState_InProgress.String()
	ret.Deadline = uint32(time.Now().Unix() + 60*60*24)
	ret.RunDuration = 100
	ret.Result = npool.TestResultState_Passed
	ret.ResultStr = npool.TestResultState_Passed.String()
	ret.Fails = 1
	ret.Passes = 10
	ret.Skips = 1

	var (
		req = &npool.TestPlanReq{
			ID:          &ret.ID,
			State:       &ret.State,
			Deadline:    &ret.Deadline,
			RunDuration: &ret.RunDuration,
			Result:      &ret.Result,
			Fails:       &ret.Fails,
			Passes:      &ret.Passes,
			Skips:       &ret.Skips,
		}
	)

	info, err := UpdateTestPlan(context.Background(), req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getTestPlan(t *testing.T) {
	info, err := GetTestPlan(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getTestPlans(t *testing.T) {
	infos, _, err := GetTestPlans(context.Background(), &npool.Conds{}, 0, 1)
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteTestPlan(t *testing.T) {
	info, err := DeleteTestPlan(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = GetTestPlan(context.Background(), ret.EntID)
	assert.NotNil(t, err)
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

	t.Run("createTestPlan", createTestPlan)
	t.Run("updateTestPlan", updateTestPlan)
	t.Run("getTestPlan", getTestPlan)
	t.Run("getTestPlans", getTestPlans)
	t.Run("deleteTestPlan", deleteTestPlan)
}
