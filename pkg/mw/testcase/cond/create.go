package cond

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testcase/cond"
	crud "github.com/NpoolPlatform/smoketest-middleware/pkg/crud/testcase/cond"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) validate() error {
	return nil
}

func (h *Handler) CreateCond(ctx context.Context) (info *npool.Cond, err error) {
	handler := &createHandler{
		Handler: h,
	}

	if err := handler.validate(); err != nil {
		return nil, err
	}

	id := uuid.New()
	if handler.ID == nil {
		handler.ID = &id
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := crud.CreateSet(
			cli.Cond.Create(),
			&crud.Req{
				ID:             h.ID,
				TestCaseID:     h.TestCaseID,
				CondTestCaseID: h.CondTestCaseID,
				CondType:       h.CondType,
				ArgumentMap:    h.ArgumentMap,
				Index:          h.Index,
			},
		).Save(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return h.GetCond(ctx)
}