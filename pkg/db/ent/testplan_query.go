// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/testplan"
)

// TestPlanQuery is the builder for querying TestPlan entities.
type TestPlanQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.TestPlan
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TestPlanQuery builder.
func (tpq *TestPlanQuery) Where(ps ...predicate.TestPlan) *TestPlanQuery {
	tpq.predicates = append(tpq.predicates, ps...)
	return tpq
}

// Limit adds a limit step to the query.
func (tpq *TestPlanQuery) Limit(limit int) *TestPlanQuery {
	tpq.limit = &limit
	return tpq
}

// Offset adds an offset step to the query.
func (tpq *TestPlanQuery) Offset(offset int) *TestPlanQuery {
	tpq.offset = &offset
	return tpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tpq *TestPlanQuery) Unique(unique bool) *TestPlanQuery {
	tpq.unique = &unique
	return tpq
}

// Order adds an order step to the query.
func (tpq *TestPlanQuery) Order(o ...OrderFunc) *TestPlanQuery {
	tpq.order = append(tpq.order, o...)
	return tpq
}

// First returns the first TestPlan entity from the query.
// Returns a *NotFoundError when no TestPlan was found.
func (tpq *TestPlanQuery) First(ctx context.Context) (*TestPlan, error) {
	nodes, err := tpq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{testplan.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tpq *TestPlanQuery) FirstX(ctx context.Context) *TestPlan {
	node, err := tpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TestPlan ID from the query.
// Returns a *NotFoundError when no TestPlan ID was found.
func (tpq *TestPlanQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = tpq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{testplan.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tpq *TestPlanQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := tpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TestPlan entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TestPlan entity is found.
// Returns a *NotFoundError when no TestPlan entities are found.
func (tpq *TestPlanQuery) Only(ctx context.Context) (*TestPlan, error) {
	nodes, err := tpq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{testplan.Label}
	default:
		return nil, &NotSingularError{testplan.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tpq *TestPlanQuery) OnlyX(ctx context.Context) *TestPlan {
	node, err := tpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TestPlan ID in the query.
// Returns a *NotSingularError when more than one TestPlan ID is found.
// Returns a *NotFoundError when no entities are found.
func (tpq *TestPlanQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = tpq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{testplan.Label}
	default:
		err = &NotSingularError{testplan.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tpq *TestPlanQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := tpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TestPlans.
func (tpq *TestPlanQuery) All(ctx context.Context) ([]*TestPlan, error) {
	if err := tpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return tpq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (tpq *TestPlanQuery) AllX(ctx context.Context) []*TestPlan {
	nodes, err := tpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TestPlan IDs.
func (tpq *TestPlanQuery) IDs(ctx context.Context) ([]uint32, error) {
	var ids []uint32
	if err := tpq.Select(testplan.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tpq *TestPlanQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := tpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tpq *TestPlanQuery) Count(ctx context.Context) (int, error) {
	if err := tpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return tpq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (tpq *TestPlanQuery) CountX(ctx context.Context) int {
	count, err := tpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tpq *TestPlanQuery) Exist(ctx context.Context) (bool, error) {
	if err := tpq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return tpq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (tpq *TestPlanQuery) ExistX(ctx context.Context) bool {
	exist, err := tpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TestPlanQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tpq *TestPlanQuery) Clone() *TestPlanQuery {
	if tpq == nil {
		return nil
	}
	return &TestPlanQuery{
		config:     tpq.config,
		limit:      tpq.limit,
		offset:     tpq.offset,
		order:      append([]OrderFunc{}, tpq.order...),
		predicates: append([]predicate.TestPlan{}, tpq.predicates...),
		// clone intermediate query.
		sql:    tpq.sql.Clone(),
		path:   tpq.path,
		unique: tpq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TestPlan.Query().
//		GroupBy(testplan.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (tpq *TestPlanQuery) GroupBy(field string, fields ...string) *TestPlanGroupBy {
	grbuild := &TestPlanGroupBy{config: tpq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := tpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return tpq.sqlQuery(ctx), nil
	}
	grbuild.label = testplan.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//	}
//
//	client.TestPlan.Query().
//		Select(testplan.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (tpq *TestPlanQuery) Select(fields ...string) *TestPlanSelect {
	tpq.fields = append(tpq.fields, fields...)
	selbuild := &TestPlanSelect{TestPlanQuery: tpq}
	selbuild.label = testplan.Label
	selbuild.flds, selbuild.scan = &tpq.fields, selbuild.Scan
	return selbuild
}

func (tpq *TestPlanQuery) prepareQuery(ctx context.Context) error {
	for _, f := range tpq.fields {
		if !testplan.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tpq.path != nil {
		prev, err := tpq.path(ctx)
		if err != nil {
			return err
		}
		tpq.sql = prev
	}
	if testplan.Policy == nil {
		return errors.New("ent: uninitialized testplan.Policy (forgotten import ent/runtime?)")
	}
	if err := testplan.Policy.EvalQuery(ctx, tpq); err != nil {
		return err
	}
	return nil
}

func (tpq *TestPlanQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TestPlan, error) {
	var (
		nodes = []*TestPlan{}
		_spec = tpq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*TestPlan).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &TestPlan{config: tpq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(tpq.modifiers) > 0 {
		_spec.Modifiers = tpq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (tpq *TestPlanQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tpq.querySpec()
	if len(tpq.modifiers) > 0 {
		_spec.Modifiers = tpq.modifiers
	}
	_spec.Node.Columns = tpq.fields
	if len(tpq.fields) > 0 {
		_spec.Unique = tpq.unique != nil && *tpq.unique
	}
	return sqlgraph.CountNodes(ctx, tpq.driver, _spec)
}

func (tpq *TestPlanQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := tpq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (tpq *TestPlanQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   testplan.Table,
			Columns: testplan.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: testplan.FieldID,
			},
		},
		From:   tpq.sql,
		Unique: true,
	}
	if unique := tpq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := tpq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, testplan.FieldID)
		for i := range fields {
			if fields[i] != testplan.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tpq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tpq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tpq *TestPlanQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tpq.driver.Dialect())
	t1 := builder.Table(testplan.Table)
	columns := tpq.fields
	if len(columns) == 0 {
		columns = testplan.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tpq.sql != nil {
		selector = tpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tpq.unique != nil && *tpq.unique {
		selector.Distinct()
	}
	for _, m := range tpq.modifiers {
		m(selector)
	}
	for _, p := range tpq.predicates {
		p(selector)
	}
	for _, p := range tpq.order {
		p(selector)
	}
	if offset := tpq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tpq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (tpq *TestPlanQuery) ForUpdate(opts ...sql.LockOption) *TestPlanQuery {
	if tpq.driver.Dialect() == dialect.Postgres {
		tpq.Unique(false)
	}
	tpq.modifiers = append(tpq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return tpq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (tpq *TestPlanQuery) ForShare(opts ...sql.LockOption) *TestPlanQuery {
	if tpq.driver.Dialect() == dialect.Postgres {
		tpq.Unique(false)
	}
	tpq.modifiers = append(tpq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return tpq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (tpq *TestPlanQuery) Modify(modifiers ...func(s *sql.Selector)) *TestPlanSelect {
	tpq.modifiers = append(tpq.modifiers, modifiers...)
	return tpq.Select()
}

// TestPlanGroupBy is the group-by builder for TestPlan entities.
type TestPlanGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tpgb *TestPlanGroupBy) Aggregate(fns ...AggregateFunc) *TestPlanGroupBy {
	tpgb.fns = append(tpgb.fns, fns...)
	return tpgb
}

// Scan applies the group-by query and scans the result into the given value.
func (tpgb *TestPlanGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := tpgb.path(ctx)
	if err != nil {
		return err
	}
	tpgb.sql = query
	return tpgb.sqlScan(ctx, v)
}

func (tpgb *TestPlanGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range tpgb.fields {
		if !testplan.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := tpgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tpgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tpgb *TestPlanGroupBy) sqlQuery() *sql.Selector {
	selector := tpgb.sql.Select()
	aggregation := make([]string, 0, len(tpgb.fns))
	for _, fn := range tpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(tpgb.fields)+len(tpgb.fns))
		for _, f := range tpgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(tpgb.fields...)...)
}

// TestPlanSelect is the builder for selecting fields of TestPlan entities.
type TestPlanSelect struct {
	*TestPlanQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (tps *TestPlanSelect) Scan(ctx context.Context, v interface{}) error {
	if err := tps.prepareQuery(ctx); err != nil {
		return err
	}
	tps.sql = tps.TestPlanQuery.sqlQuery(ctx)
	return tps.sqlScan(ctx, v)
}

func (tps *TestPlanSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := tps.sql.Query()
	if err := tps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (tps *TestPlanSelect) Modify(modifiers ...func(s *sql.Selector)) *TestPlanSelect {
	tps.modifiers = append(tps.modifiers, modifiers...)
	return tps
}
