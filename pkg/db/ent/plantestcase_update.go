// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/plantestcase"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// PlanTestCaseUpdate is the builder for updating PlanTestCase entities.
type PlanTestCaseUpdate struct {
	config
	hooks     []Hook
	mutation  *PlanTestCaseMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the PlanTestCaseUpdate builder.
func (ptcu *PlanTestCaseUpdate) Where(ps ...predicate.PlanTestCase) *PlanTestCaseUpdate {
	ptcu.mutation.Where(ps...)
	return ptcu
}

// SetCreatedAt sets the "created_at" field.
func (ptcu *PlanTestCaseUpdate) SetCreatedAt(u uint32) *PlanTestCaseUpdate {
	ptcu.mutation.ResetCreatedAt()
	ptcu.mutation.SetCreatedAt(u)
	return ptcu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ptcu *PlanTestCaseUpdate) SetNillableCreatedAt(u *uint32) *PlanTestCaseUpdate {
	if u != nil {
		ptcu.SetCreatedAt(*u)
	}
	return ptcu
}

// AddCreatedAt adds u to the "created_at" field.
func (ptcu *PlanTestCaseUpdate) AddCreatedAt(u int32) *PlanTestCaseUpdate {
	ptcu.mutation.AddCreatedAt(u)
	return ptcu
}

// SetUpdatedAt sets the "updated_at" field.
func (ptcu *PlanTestCaseUpdate) SetUpdatedAt(u uint32) *PlanTestCaseUpdate {
	ptcu.mutation.ResetUpdatedAt()
	ptcu.mutation.SetUpdatedAt(u)
	return ptcu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (ptcu *PlanTestCaseUpdate) AddUpdatedAt(u int32) *PlanTestCaseUpdate {
	ptcu.mutation.AddUpdatedAt(u)
	return ptcu
}

// SetDeletedAt sets the "deleted_at" field.
func (ptcu *PlanTestCaseUpdate) SetDeletedAt(u uint32) *PlanTestCaseUpdate {
	ptcu.mutation.ResetDeletedAt()
	ptcu.mutation.SetDeletedAt(u)
	return ptcu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ptcu *PlanTestCaseUpdate) SetNillableDeletedAt(u *uint32) *PlanTestCaseUpdate {
	if u != nil {
		ptcu.SetDeletedAt(*u)
	}
	return ptcu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (ptcu *PlanTestCaseUpdate) AddDeletedAt(u int32) *PlanTestCaseUpdate {
	ptcu.mutation.AddDeletedAt(u)
	return ptcu
}

// SetEntID sets the "ent_id" field.
func (ptcu *PlanTestCaseUpdate) SetEntID(u uuid.UUID) *PlanTestCaseUpdate {
	ptcu.mutation.SetEntID(u)
	return ptcu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (ptcu *PlanTestCaseUpdate) SetNillableEntID(u *uuid.UUID) *PlanTestCaseUpdate {
	if u != nil {
		ptcu.SetEntID(*u)
	}
	return ptcu
}

// SetTestPlanID sets the "test_plan_id" field.
func (ptcu *PlanTestCaseUpdate) SetTestPlanID(u uuid.UUID) *PlanTestCaseUpdate {
	ptcu.mutation.SetTestPlanID(u)
	return ptcu
}

// SetNillableTestPlanID sets the "test_plan_id" field if the given value is not nil.
func (ptcu *PlanTestCaseUpdate) SetNillableTestPlanID(u *uuid.UUID) *PlanTestCaseUpdate {
	if u != nil {
		ptcu.SetTestPlanID(*u)
	}
	return ptcu
}

// ClearTestPlanID clears the value of the "test_plan_id" field.
func (ptcu *PlanTestCaseUpdate) ClearTestPlanID() *PlanTestCaseUpdate {
	ptcu.mutation.ClearTestPlanID()
	return ptcu
}

// SetTestCaseID sets the "test_case_id" field.
func (ptcu *PlanTestCaseUpdate) SetTestCaseID(u uuid.UUID) *PlanTestCaseUpdate {
	ptcu.mutation.SetTestCaseID(u)
	return ptcu
}

// SetNillableTestCaseID sets the "test_case_id" field if the given value is not nil.
func (ptcu *PlanTestCaseUpdate) SetNillableTestCaseID(u *uuid.UUID) *PlanTestCaseUpdate {
	if u != nil {
		ptcu.SetTestCaseID(*u)
	}
	return ptcu
}

// ClearTestCaseID clears the value of the "test_case_id" field.
func (ptcu *PlanTestCaseUpdate) ClearTestCaseID() *PlanTestCaseUpdate {
	ptcu.mutation.ClearTestCaseID()
	return ptcu
}

// SetInput sets the "input" field.
func (ptcu *PlanTestCaseUpdate) SetInput(s string) *PlanTestCaseUpdate {
	ptcu.mutation.SetInput(s)
	return ptcu
}

// SetNillableInput sets the "input" field if the given value is not nil.
func (ptcu *PlanTestCaseUpdate) SetNillableInput(s *string) *PlanTestCaseUpdate {
	if s != nil {
		ptcu.SetInput(*s)
	}
	return ptcu
}

// ClearInput clears the value of the "input" field.
func (ptcu *PlanTestCaseUpdate) ClearInput() *PlanTestCaseUpdate {
	ptcu.mutation.ClearInput()
	return ptcu
}

// SetOutput sets the "output" field.
func (ptcu *PlanTestCaseUpdate) SetOutput(s string) *PlanTestCaseUpdate {
	ptcu.mutation.SetOutput(s)
	return ptcu
}

// SetNillableOutput sets the "output" field if the given value is not nil.
func (ptcu *PlanTestCaseUpdate) SetNillableOutput(s *string) *PlanTestCaseUpdate {
	if s != nil {
		ptcu.SetOutput(*s)
	}
	return ptcu
}

// ClearOutput clears the value of the "output" field.
func (ptcu *PlanTestCaseUpdate) ClearOutput() *PlanTestCaseUpdate {
	ptcu.mutation.ClearOutput()
	return ptcu
}

// SetDescription sets the "description" field.
func (ptcu *PlanTestCaseUpdate) SetDescription(s string) *PlanTestCaseUpdate {
	ptcu.mutation.SetDescription(s)
	return ptcu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ptcu *PlanTestCaseUpdate) SetNillableDescription(s *string) *PlanTestCaseUpdate {
	if s != nil {
		ptcu.SetDescription(*s)
	}
	return ptcu
}

// ClearDescription clears the value of the "description" field.
func (ptcu *PlanTestCaseUpdate) ClearDescription() *PlanTestCaseUpdate {
	ptcu.mutation.ClearDescription()
	return ptcu
}

// SetTestUserID sets the "test_user_id" field.
func (ptcu *PlanTestCaseUpdate) SetTestUserID(u uuid.UUID) *PlanTestCaseUpdate {
	ptcu.mutation.SetTestUserID(u)
	return ptcu
}

// SetNillableTestUserID sets the "test_user_id" field if the given value is not nil.
func (ptcu *PlanTestCaseUpdate) SetNillableTestUserID(u *uuid.UUID) *PlanTestCaseUpdate {
	if u != nil {
		ptcu.SetTestUserID(*u)
	}
	return ptcu
}

// ClearTestUserID clears the value of the "test_user_id" field.
func (ptcu *PlanTestCaseUpdate) ClearTestUserID() *PlanTestCaseUpdate {
	ptcu.mutation.ClearTestUserID()
	return ptcu
}

// SetRunDuration sets the "run_duration" field.
func (ptcu *PlanTestCaseUpdate) SetRunDuration(u uint32) *PlanTestCaseUpdate {
	ptcu.mutation.ResetRunDuration()
	ptcu.mutation.SetRunDuration(u)
	return ptcu
}

// SetNillableRunDuration sets the "run_duration" field if the given value is not nil.
func (ptcu *PlanTestCaseUpdate) SetNillableRunDuration(u *uint32) *PlanTestCaseUpdate {
	if u != nil {
		ptcu.SetRunDuration(*u)
	}
	return ptcu
}

// AddRunDuration adds u to the "run_duration" field.
func (ptcu *PlanTestCaseUpdate) AddRunDuration(u int32) *PlanTestCaseUpdate {
	ptcu.mutation.AddRunDuration(u)
	return ptcu
}

// ClearRunDuration clears the value of the "run_duration" field.
func (ptcu *PlanTestCaseUpdate) ClearRunDuration() *PlanTestCaseUpdate {
	ptcu.mutation.ClearRunDuration()
	return ptcu
}

// SetResult sets the "result" field.
func (ptcu *PlanTestCaseUpdate) SetResult(s string) *PlanTestCaseUpdate {
	ptcu.mutation.SetResult(s)
	return ptcu
}

// SetNillableResult sets the "result" field if the given value is not nil.
func (ptcu *PlanTestCaseUpdate) SetNillableResult(s *string) *PlanTestCaseUpdate {
	if s != nil {
		ptcu.SetResult(*s)
	}
	return ptcu
}

// ClearResult clears the value of the "result" field.
func (ptcu *PlanTestCaseUpdate) ClearResult() *PlanTestCaseUpdate {
	ptcu.mutation.ClearResult()
	return ptcu
}

// SetIndex sets the "index" field.
func (ptcu *PlanTestCaseUpdate) SetIndex(u uint32) *PlanTestCaseUpdate {
	ptcu.mutation.ResetIndex()
	ptcu.mutation.SetIndex(u)
	return ptcu
}

// SetNillableIndex sets the "index" field if the given value is not nil.
func (ptcu *PlanTestCaseUpdate) SetNillableIndex(u *uint32) *PlanTestCaseUpdate {
	if u != nil {
		ptcu.SetIndex(*u)
	}
	return ptcu
}

// AddIndex adds u to the "index" field.
func (ptcu *PlanTestCaseUpdate) AddIndex(u int32) *PlanTestCaseUpdate {
	ptcu.mutation.AddIndex(u)
	return ptcu
}

// ClearIndex clears the value of the "index" field.
func (ptcu *PlanTestCaseUpdate) ClearIndex() *PlanTestCaseUpdate {
	ptcu.mutation.ClearIndex()
	return ptcu
}

// Mutation returns the PlanTestCaseMutation object of the builder.
func (ptcu *PlanTestCaseUpdate) Mutation() *PlanTestCaseMutation {
	return ptcu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ptcu *PlanTestCaseUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := ptcu.defaults(); err != nil {
		return 0, err
	}
	if len(ptcu.hooks) == 0 {
		affected, err = ptcu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlanTestCaseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ptcu.mutation = mutation
			affected, err = ptcu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ptcu.hooks) - 1; i >= 0; i-- {
			if ptcu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ptcu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ptcu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ptcu *PlanTestCaseUpdate) SaveX(ctx context.Context) int {
	affected, err := ptcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ptcu *PlanTestCaseUpdate) Exec(ctx context.Context) error {
	_, err := ptcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptcu *PlanTestCaseUpdate) ExecX(ctx context.Context) {
	if err := ptcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ptcu *PlanTestCaseUpdate) defaults() error {
	if _, ok := ptcu.mutation.UpdatedAt(); !ok {
		if plantestcase.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized plantestcase.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := plantestcase.UpdateDefaultUpdatedAt()
		ptcu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ptcu *PlanTestCaseUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PlanTestCaseUpdate {
	ptcu.modifiers = append(ptcu.modifiers, modifiers...)
	return ptcu
}

func (ptcu *PlanTestCaseUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   plantestcase.Table,
			Columns: plantestcase.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: plantestcase.FieldID,
			},
		},
	}
	if ps := ptcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ptcu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: plantestcase.FieldCreatedAt,
		})
	}
	if value, ok := ptcu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: plantestcase.FieldCreatedAt,
		})
	}
	if value, ok := ptcu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: plantestcase.FieldUpdatedAt,
		})
	}
	if value, ok := ptcu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: plantestcase.FieldUpdatedAt,
		})
	}
	if value, ok := ptcu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: plantestcase.FieldDeletedAt,
		})
	}
	if value, ok := ptcu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: plantestcase.FieldDeletedAt,
		})
	}
	if value, ok := ptcu.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: plantestcase.FieldEntID,
		})
	}
	if value, ok := ptcu.mutation.TestPlanID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: plantestcase.FieldTestPlanID,
		})
	}
	if ptcu.mutation.TestPlanIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: plantestcase.FieldTestPlanID,
		})
	}
	if value, ok := ptcu.mutation.TestCaseID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: plantestcase.FieldTestCaseID,
		})
	}
	if ptcu.mutation.TestCaseIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: plantestcase.FieldTestCaseID,
		})
	}
	if value, ok := ptcu.mutation.Input(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: plantestcase.FieldInput,
		})
	}
	if ptcu.mutation.InputCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: plantestcase.FieldInput,
		})
	}
	if value, ok := ptcu.mutation.Output(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: plantestcase.FieldOutput,
		})
	}
	if ptcu.mutation.OutputCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: plantestcase.FieldOutput,
		})
	}
	if value, ok := ptcu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: plantestcase.FieldDescription,
		})
	}
	if ptcu.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: plantestcase.FieldDescription,
		})
	}
	if value, ok := ptcu.mutation.TestUserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: plantestcase.FieldTestUserID,
		})
	}
	if ptcu.mutation.TestUserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: plantestcase.FieldTestUserID,
		})
	}
	if value, ok := ptcu.mutation.RunDuration(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: plantestcase.FieldRunDuration,
		})
	}
	if value, ok := ptcu.mutation.AddedRunDuration(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: plantestcase.FieldRunDuration,
		})
	}
	if ptcu.mutation.RunDurationCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: plantestcase.FieldRunDuration,
		})
	}
	if value, ok := ptcu.mutation.Result(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: plantestcase.FieldResult,
		})
	}
	if ptcu.mutation.ResultCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: plantestcase.FieldResult,
		})
	}
	if value, ok := ptcu.mutation.Index(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: plantestcase.FieldIndex,
		})
	}
	if value, ok := ptcu.mutation.AddedIndex(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: plantestcase.FieldIndex,
		})
	}
	if ptcu.mutation.IndexCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: plantestcase.FieldIndex,
		})
	}
	_spec.Modifiers = ptcu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, ptcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{plantestcase.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// PlanTestCaseUpdateOne is the builder for updating a single PlanTestCase entity.
type PlanTestCaseUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *PlanTestCaseMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (ptcuo *PlanTestCaseUpdateOne) SetCreatedAt(u uint32) *PlanTestCaseUpdateOne {
	ptcuo.mutation.ResetCreatedAt()
	ptcuo.mutation.SetCreatedAt(u)
	return ptcuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ptcuo *PlanTestCaseUpdateOne) SetNillableCreatedAt(u *uint32) *PlanTestCaseUpdateOne {
	if u != nil {
		ptcuo.SetCreatedAt(*u)
	}
	return ptcuo
}

// AddCreatedAt adds u to the "created_at" field.
func (ptcuo *PlanTestCaseUpdateOne) AddCreatedAt(u int32) *PlanTestCaseUpdateOne {
	ptcuo.mutation.AddCreatedAt(u)
	return ptcuo
}

// SetUpdatedAt sets the "updated_at" field.
func (ptcuo *PlanTestCaseUpdateOne) SetUpdatedAt(u uint32) *PlanTestCaseUpdateOne {
	ptcuo.mutation.ResetUpdatedAt()
	ptcuo.mutation.SetUpdatedAt(u)
	return ptcuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (ptcuo *PlanTestCaseUpdateOne) AddUpdatedAt(u int32) *PlanTestCaseUpdateOne {
	ptcuo.mutation.AddUpdatedAt(u)
	return ptcuo
}

// SetDeletedAt sets the "deleted_at" field.
func (ptcuo *PlanTestCaseUpdateOne) SetDeletedAt(u uint32) *PlanTestCaseUpdateOne {
	ptcuo.mutation.ResetDeletedAt()
	ptcuo.mutation.SetDeletedAt(u)
	return ptcuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ptcuo *PlanTestCaseUpdateOne) SetNillableDeletedAt(u *uint32) *PlanTestCaseUpdateOne {
	if u != nil {
		ptcuo.SetDeletedAt(*u)
	}
	return ptcuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (ptcuo *PlanTestCaseUpdateOne) AddDeletedAt(u int32) *PlanTestCaseUpdateOne {
	ptcuo.mutation.AddDeletedAt(u)
	return ptcuo
}

// SetEntID sets the "ent_id" field.
func (ptcuo *PlanTestCaseUpdateOne) SetEntID(u uuid.UUID) *PlanTestCaseUpdateOne {
	ptcuo.mutation.SetEntID(u)
	return ptcuo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (ptcuo *PlanTestCaseUpdateOne) SetNillableEntID(u *uuid.UUID) *PlanTestCaseUpdateOne {
	if u != nil {
		ptcuo.SetEntID(*u)
	}
	return ptcuo
}

// SetTestPlanID sets the "test_plan_id" field.
func (ptcuo *PlanTestCaseUpdateOne) SetTestPlanID(u uuid.UUID) *PlanTestCaseUpdateOne {
	ptcuo.mutation.SetTestPlanID(u)
	return ptcuo
}

// SetNillableTestPlanID sets the "test_plan_id" field if the given value is not nil.
func (ptcuo *PlanTestCaseUpdateOne) SetNillableTestPlanID(u *uuid.UUID) *PlanTestCaseUpdateOne {
	if u != nil {
		ptcuo.SetTestPlanID(*u)
	}
	return ptcuo
}

// ClearTestPlanID clears the value of the "test_plan_id" field.
func (ptcuo *PlanTestCaseUpdateOne) ClearTestPlanID() *PlanTestCaseUpdateOne {
	ptcuo.mutation.ClearTestPlanID()
	return ptcuo
}

// SetTestCaseID sets the "test_case_id" field.
func (ptcuo *PlanTestCaseUpdateOne) SetTestCaseID(u uuid.UUID) *PlanTestCaseUpdateOne {
	ptcuo.mutation.SetTestCaseID(u)
	return ptcuo
}

// SetNillableTestCaseID sets the "test_case_id" field if the given value is not nil.
func (ptcuo *PlanTestCaseUpdateOne) SetNillableTestCaseID(u *uuid.UUID) *PlanTestCaseUpdateOne {
	if u != nil {
		ptcuo.SetTestCaseID(*u)
	}
	return ptcuo
}

// ClearTestCaseID clears the value of the "test_case_id" field.
func (ptcuo *PlanTestCaseUpdateOne) ClearTestCaseID() *PlanTestCaseUpdateOne {
	ptcuo.mutation.ClearTestCaseID()
	return ptcuo
}

// SetInput sets the "input" field.
func (ptcuo *PlanTestCaseUpdateOne) SetInput(s string) *PlanTestCaseUpdateOne {
	ptcuo.mutation.SetInput(s)
	return ptcuo
}

// SetNillableInput sets the "input" field if the given value is not nil.
func (ptcuo *PlanTestCaseUpdateOne) SetNillableInput(s *string) *PlanTestCaseUpdateOne {
	if s != nil {
		ptcuo.SetInput(*s)
	}
	return ptcuo
}

// ClearInput clears the value of the "input" field.
func (ptcuo *PlanTestCaseUpdateOne) ClearInput() *PlanTestCaseUpdateOne {
	ptcuo.mutation.ClearInput()
	return ptcuo
}

// SetOutput sets the "output" field.
func (ptcuo *PlanTestCaseUpdateOne) SetOutput(s string) *PlanTestCaseUpdateOne {
	ptcuo.mutation.SetOutput(s)
	return ptcuo
}

// SetNillableOutput sets the "output" field if the given value is not nil.
func (ptcuo *PlanTestCaseUpdateOne) SetNillableOutput(s *string) *PlanTestCaseUpdateOne {
	if s != nil {
		ptcuo.SetOutput(*s)
	}
	return ptcuo
}

// ClearOutput clears the value of the "output" field.
func (ptcuo *PlanTestCaseUpdateOne) ClearOutput() *PlanTestCaseUpdateOne {
	ptcuo.mutation.ClearOutput()
	return ptcuo
}

// SetDescription sets the "description" field.
func (ptcuo *PlanTestCaseUpdateOne) SetDescription(s string) *PlanTestCaseUpdateOne {
	ptcuo.mutation.SetDescription(s)
	return ptcuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ptcuo *PlanTestCaseUpdateOne) SetNillableDescription(s *string) *PlanTestCaseUpdateOne {
	if s != nil {
		ptcuo.SetDescription(*s)
	}
	return ptcuo
}

// ClearDescription clears the value of the "description" field.
func (ptcuo *PlanTestCaseUpdateOne) ClearDescription() *PlanTestCaseUpdateOne {
	ptcuo.mutation.ClearDescription()
	return ptcuo
}

// SetTestUserID sets the "test_user_id" field.
func (ptcuo *PlanTestCaseUpdateOne) SetTestUserID(u uuid.UUID) *PlanTestCaseUpdateOne {
	ptcuo.mutation.SetTestUserID(u)
	return ptcuo
}

// SetNillableTestUserID sets the "test_user_id" field if the given value is not nil.
func (ptcuo *PlanTestCaseUpdateOne) SetNillableTestUserID(u *uuid.UUID) *PlanTestCaseUpdateOne {
	if u != nil {
		ptcuo.SetTestUserID(*u)
	}
	return ptcuo
}

// ClearTestUserID clears the value of the "test_user_id" field.
func (ptcuo *PlanTestCaseUpdateOne) ClearTestUserID() *PlanTestCaseUpdateOne {
	ptcuo.mutation.ClearTestUserID()
	return ptcuo
}

// SetRunDuration sets the "run_duration" field.
func (ptcuo *PlanTestCaseUpdateOne) SetRunDuration(u uint32) *PlanTestCaseUpdateOne {
	ptcuo.mutation.ResetRunDuration()
	ptcuo.mutation.SetRunDuration(u)
	return ptcuo
}

// SetNillableRunDuration sets the "run_duration" field if the given value is not nil.
func (ptcuo *PlanTestCaseUpdateOne) SetNillableRunDuration(u *uint32) *PlanTestCaseUpdateOne {
	if u != nil {
		ptcuo.SetRunDuration(*u)
	}
	return ptcuo
}

// AddRunDuration adds u to the "run_duration" field.
func (ptcuo *PlanTestCaseUpdateOne) AddRunDuration(u int32) *PlanTestCaseUpdateOne {
	ptcuo.mutation.AddRunDuration(u)
	return ptcuo
}

// ClearRunDuration clears the value of the "run_duration" field.
func (ptcuo *PlanTestCaseUpdateOne) ClearRunDuration() *PlanTestCaseUpdateOne {
	ptcuo.mutation.ClearRunDuration()
	return ptcuo
}

// SetResult sets the "result" field.
func (ptcuo *PlanTestCaseUpdateOne) SetResult(s string) *PlanTestCaseUpdateOne {
	ptcuo.mutation.SetResult(s)
	return ptcuo
}

// SetNillableResult sets the "result" field if the given value is not nil.
func (ptcuo *PlanTestCaseUpdateOne) SetNillableResult(s *string) *PlanTestCaseUpdateOne {
	if s != nil {
		ptcuo.SetResult(*s)
	}
	return ptcuo
}

// ClearResult clears the value of the "result" field.
func (ptcuo *PlanTestCaseUpdateOne) ClearResult() *PlanTestCaseUpdateOne {
	ptcuo.mutation.ClearResult()
	return ptcuo
}

// SetIndex sets the "index" field.
func (ptcuo *PlanTestCaseUpdateOne) SetIndex(u uint32) *PlanTestCaseUpdateOne {
	ptcuo.mutation.ResetIndex()
	ptcuo.mutation.SetIndex(u)
	return ptcuo
}

// SetNillableIndex sets the "index" field if the given value is not nil.
func (ptcuo *PlanTestCaseUpdateOne) SetNillableIndex(u *uint32) *PlanTestCaseUpdateOne {
	if u != nil {
		ptcuo.SetIndex(*u)
	}
	return ptcuo
}

// AddIndex adds u to the "index" field.
func (ptcuo *PlanTestCaseUpdateOne) AddIndex(u int32) *PlanTestCaseUpdateOne {
	ptcuo.mutation.AddIndex(u)
	return ptcuo
}

// ClearIndex clears the value of the "index" field.
func (ptcuo *PlanTestCaseUpdateOne) ClearIndex() *PlanTestCaseUpdateOne {
	ptcuo.mutation.ClearIndex()
	return ptcuo
}

// Mutation returns the PlanTestCaseMutation object of the builder.
func (ptcuo *PlanTestCaseUpdateOne) Mutation() *PlanTestCaseMutation {
	return ptcuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ptcuo *PlanTestCaseUpdateOne) Select(field string, fields ...string) *PlanTestCaseUpdateOne {
	ptcuo.fields = append([]string{field}, fields...)
	return ptcuo
}

// Save executes the query and returns the updated PlanTestCase entity.
func (ptcuo *PlanTestCaseUpdateOne) Save(ctx context.Context) (*PlanTestCase, error) {
	var (
		err  error
		node *PlanTestCase
	)
	if err := ptcuo.defaults(); err != nil {
		return nil, err
	}
	if len(ptcuo.hooks) == 0 {
		node, err = ptcuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlanTestCaseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ptcuo.mutation = mutation
			node, err = ptcuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ptcuo.hooks) - 1; i >= 0; i-- {
			if ptcuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ptcuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ptcuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*PlanTestCase)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from PlanTestCaseMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ptcuo *PlanTestCaseUpdateOne) SaveX(ctx context.Context) *PlanTestCase {
	node, err := ptcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ptcuo *PlanTestCaseUpdateOne) Exec(ctx context.Context) error {
	_, err := ptcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptcuo *PlanTestCaseUpdateOne) ExecX(ctx context.Context) {
	if err := ptcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ptcuo *PlanTestCaseUpdateOne) defaults() error {
	if _, ok := ptcuo.mutation.UpdatedAt(); !ok {
		if plantestcase.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized plantestcase.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := plantestcase.UpdateDefaultUpdatedAt()
		ptcuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ptcuo *PlanTestCaseUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PlanTestCaseUpdateOne {
	ptcuo.modifiers = append(ptcuo.modifiers, modifiers...)
	return ptcuo
}

func (ptcuo *PlanTestCaseUpdateOne) sqlSave(ctx context.Context) (_node *PlanTestCase, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   plantestcase.Table,
			Columns: plantestcase.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: plantestcase.FieldID,
			},
		},
	}
	id, ok := ptcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PlanTestCase.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ptcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, plantestcase.FieldID)
		for _, f := range fields {
			if !plantestcase.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != plantestcase.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ptcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ptcuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: plantestcase.FieldCreatedAt,
		})
	}
	if value, ok := ptcuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: plantestcase.FieldCreatedAt,
		})
	}
	if value, ok := ptcuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: plantestcase.FieldUpdatedAt,
		})
	}
	if value, ok := ptcuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: plantestcase.FieldUpdatedAt,
		})
	}
	if value, ok := ptcuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: plantestcase.FieldDeletedAt,
		})
	}
	if value, ok := ptcuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: plantestcase.FieldDeletedAt,
		})
	}
	if value, ok := ptcuo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: plantestcase.FieldEntID,
		})
	}
	if value, ok := ptcuo.mutation.TestPlanID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: plantestcase.FieldTestPlanID,
		})
	}
	if ptcuo.mutation.TestPlanIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: plantestcase.FieldTestPlanID,
		})
	}
	if value, ok := ptcuo.mutation.TestCaseID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: plantestcase.FieldTestCaseID,
		})
	}
	if ptcuo.mutation.TestCaseIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: plantestcase.FieldTestCaseID,
		})
	}
	if value, ok := ptcuo.mutation.Input(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: plantestcase.FieldInput,
		})
	}
	if ptcuo.mutation.InputCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: plantestcase.FieldInput,
		})
	}
	if value, ok := ptcuo.mutation.Output(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: plantestcase.FieldOutput,
		})
	}
	if ptcuo.mutation.OutputCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: plantestcase.FieldOutput,
		})
	}
	if value, ok := ptcuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: plantestcase.FieldDescription,
		})
	}
	if ptcuo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: plantestcase.FieldDescription,
		})
	}
	if value, ok := ptcuo.mutation.TestUserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: plantestcase.FieldTestUserID,
		})
	}
	if ptcuo.mutation.TestUserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: plantestcase.FieldTestUserID,
		})
	}
	if value, ok := ptcuo.mutation.RunDuration(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: plantestcase.FieldRunDuration,
		})
	}
	if value, ok := ptcuo.mutation.AddedRunDuration(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: plantestcase.FieldRunDuration,
		})
	}
	if ptcuo.mutation.RunDurationCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: plantestcase.FieldRunDuration,
		})
	}
	if value, ok := ptcuo.mutation.Result(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: plantestcase.FieldResult,
		})
	}
	if ptcuo.mutation.ResultCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: plantestcase.FieldResult,
		})
	}
	if value, ok := ptcuo.mutation.Index(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: plantestcase.FieldIndex,
		})
	}
	if value, ok := ptcuo.mutation.AddedIndex(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: plantestcase.FieldIndex,
		})
	}
	if ptcuo.mutation.IndexCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: plantestcase.FieldIndex,
		})
	}
	_spec.Modifiers = ptcuo.modifiers
	_node = &PlanTestCase{config: ptcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ptcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{plantestcase.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
