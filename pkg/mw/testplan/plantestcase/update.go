package plantestcase

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testplan/plantestcase"
	plantestcasecrud "github.com/NpoolPlatform/smoketest-middleware/pkg/crud/testplan/plantestcase"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
)

func (h *Handler) UpdatePlanTestCase(ctx context.Context) (info *npool.PlanTestCase, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := plantestcasecrud.UpdateSet(
			cli.PlanTestCase.UpdateOneID(*h.ID),
			&plantestcasecrud.Req{
				ID:             h.ID,
				TestCaseOutput: h.TestCaseOutput,
				Result:         h.TestCaseResult,
				Index:          &h.Index,
				RunDuration:    h.RunDuration,
				Description:    h.Description,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	info, err = h.GetPlanTestCase(ctx)
	if err != nil {
		return nil, err
	}

	return info, nil
}