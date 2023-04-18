package cond

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	condmgrpb "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testcase/cond"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/cond"
	"github.com/google/uuid"
)

type Req struct {
	ID             *uuid.UUID
	CondType       *condmgrpb.CondType
	TestCaseID     *uuid.UUID
	CondTestCaseID *uuid.UUID
	Index          *uint32
	ArgumentMap    *string
}

func CreateSet(c *ent.CondCreate, req *Req) *ent.CondCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.CondType != nil {
		c.SetCondType(req.CondType.String())
	}
	if req.TestCaseID != nil {
		c.SetTestCaseID(*req.TestCaseID)
	}
	if req.CondTestCaseID != nil {
		c.SetCondTestCaseID(*req.CondTestCaseID)
	}
	if req.Index != nil {
		c.SetIndex(*req.Index)
	}
	if req.ArgumentMap != nil {
		c.SetArgumentMap(*req.ArgumentMap)
	}
	return c
}

func UpdateSet(ctx context.Context, u *ent.CondUpdateOne, req *Req) (*ent.CondUpdateOne, error) {
	if req.CondType != nil {
		u.SetCondType(req.CondType.String())
	}
	if req.Index != nil {
		u.SetIndex(*req.Index)
	}
	if req.ArgumentMap != nil {
		u.SetArgumentMap(*req.ArgumentMap)
	}
	return u, nil
}

type Conds struct {
	ID             *cruder.Cond
	TestCaseID     *cruder.Cond
	CondTestCaseID *cruder.Cond
	CondType       *cruder.Cond
}

//nolint:nolintlint,gocyclo
func SetQueryConds(q *ent.CondQuery, conds *Conds) (*ent.CondQuery, error) {
	if conds == nil {
		return q, nil
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(cond.ID(id))
		default:
			return nil, fmt.Errorf("invalid id field")
		}
	}
	if conds.TestCaseID != nil {
		testCaseID, ok := conds.TestCaseID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.TestCaseID.Op {
		case cruder.EQ:
			q.Where(cond.TestCaseID(testCaseID))
		default:
			return nil, fmt.Errorf("invalid testcase id field")
		}
	}
	if conds.CondTestCaseID != nil {
		condTestCaseID, ok := conds.CondTestCaseID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid cond testcase id")
		}
		switch conds.CondTestCaseID.Op {
		case cruder.EQ:
			q.Where(cond.CondTestCaseID(condTestCaseID))
		default:
			return nil, fmt.Errorf("invalid cond testcase id field")
		}
	}
	if conds.CondType != nil {
		condType, ok := conds.CondType.Val.(condmgrpb.CondType)
		if !ok {
			return nil, fmt.Errorf("invalid cond type")
		}
		switch conds.CondType.Op {
		case cruder.EQ:
			q.Where(cond.CondType(condType.String()))
		}
	}
	return q, nil
}