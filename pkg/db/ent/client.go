// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/migrate"
	"github.com/google/uuid"

	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/detail"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/module"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/relatedtestcase"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/testcase"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/testplan"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Detail is the client for interacting with the Detail builders.
	Detail *DetailClient
	// Module is the client for interacting with the Module builders.
	Module *ModuleClient
	// RelatedTestCase is the client for interacting with the RelatedTestCase builders.
	RelatedTestCase *RelatedTestCaseClient
	// TestCase is the client for interacting with the TestCase builders.
	TestCase *TestCaseClient
	// TestPlan is the client for interacting with the TestPlan builders.
	TestPlan *TestPlanClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Detail = NewDetailClient(c.config)
	c.Module = NewModuleClient(c.config)
	c.RelatedTestCase = NewRelatedTestCaseClient(c.config)
	c.TestCase = NewTestCaseClient(c.config)
	c.TestPlan = NewTestPlanClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:             ctx,
		config:          cfg,
		Detail:          NewDetailClient(cfg),
		Module:          NewModuleClient(cfg),
		RelatedTestCase: NewRelatedTestCaseClient(cfg),
		TestCase:        NewTestCaseClient(cfg),
		TestPlan:        NewTestPlanClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:             ctx,
		config:          cfg,
		Detail:          NewDetailClient(cfg),
		Module:          NewModuleClient(cfg),
		RelatedTestCase: NewRelatedTestCaseClient(cfg),
		TestCase:        NewTestCaseClient(cfg),
		TestPlan:        NewTestPlanClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Detail.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Detail.Use(hooks...)
	c.Module.Use(hooks...)
	c.RelatedTestCase.Use(hooks...)
	c.TestCase.Use(hooks...)
	c.TestPlan.Use(hooks...)
}

// DetailClient is a client for the Detail schema.
type DetailClient struct {
	config
}

// NewDetailClient returns a client for the Detail from the given config.
func NewDetailClient(c config) *DetailClient {
	return &DetailClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `detail.Hooks(f(g(h())))`.
func (c *DetailClient) Use(hooks ...Hook) {
	c.hooks.Detail = append(c.hooks.Detail, hooks...)
}

// Create returns a builder for creating a Detail entity.
func (c *DetailClient) Create() *DetailCreate {
	mutation := newDetailMutation(c.config, OpCreate)
	return &DetailCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Detail entities.
func (c *DetailClient) CreateBulk(builders ...*DetailCreate) *DetailCreateBulk {
	return &DetailCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Detail.
func (c *DetailClient) Update() *DetailUpdate {
	mutation := newDetailMutation(c.config, OpUpdate)
	return &DetailUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DetailClient) UpdateOne(d *Detail) *DetailUpdateOne {
	mutation := newDetailMutation(c.config, OpUpdateOne, withDetail(d))
	return &DetailUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DetailClient) UpdateOneID(id uuid.UUID) *DetailUpdateOne {
	mutation := newDetailMutation(c.config, OpUpdateOne, withDetailID(id))
	return &DetailUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Detail.
func (c *DetailClient) Delete() *DetailDelete {
	mutation := newDetailMutation(c.config, OpDelete)
	return &DetailDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *DetailClient) DeleteOne(d *Detail) *DetailDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *DetailClient) DeleteOneID(id uuid.UUID) *DetailDeleteOne {
	builder := c.Delete().Where(detail.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DetailDeleteOne{builder}
}

// Query returns a query builder for Detail.
func (c *DetailClient) Query() *DetailQuery {
	return &DetailQuery{
		config: c.config,
	}
}

// Get returns a Detail entity by its id.
func (c *DetailClient) Get(ctx context.Context, id uuid.UUID) (*Detail, error) {
	return c.Query().Where(detail.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DetailClient) GetX(ctx context.Context, id uuid.UUID) *Detail {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *DetailClient) Hooks() []Hook {
	hooks := c.hooks.Detail
	return append(hooks[:len(hooks):len(hooks)], detail.Hooks[:]...)
}

// ModuleClient is a client for the Module schema.
type ModuleClient struct {
	config
}

// NewModuleClient returns a client for the Module from the given config.
func NewModuleClient(c config) *ModuleClient {
	return &ModuleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `module.Hooks(f(g(h())))`.
func (c *ModuleClient) Use(hooks ...Hook) {
	c.hooks.Module = append(c.hooks.Module, hooks...)
}

// Create returns a builder for creating a Module entity.
func (c *ModuleClient) Create() *ModuleCreate {
	mutation := newModuleMutation(c.config, OpCreate)
	return &ModuleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Module entities.
func (c *ModuleClient) CreateBulk(builders ...*ModuleCreate) *ModuleCreateBulk {
	return &ModuleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Module.
func (c *ModuleClient) Update() *ModuleUpdate {
	mutation := newModuleMutation(c.config, OpUpdate)
	return &ModuleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ModuleClient) UpdateOne(m *Module) *ModuleUpdateOne {
	mutation := newModuleMutation(c.config, OpUpdateOne, withModule(m))
	return &ModuleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ModuleClient) UpdateOneID(id uuid.UUID) *ModuleUpdateOne {
	mutation := newModuleMutation(c.config, OpUpdateOne, withModuleID(id))
	return &ModuleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Module.
func (c *ModuleClient) Delete() *ModuleDelete {
	mutation := newModuleMutation(c.config, OpDelete)
	return &ModuleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ModuleClient) DeleteOne(m *Module) *ModuleDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *ModuleClient) DeleteOneID(id uuid.UUID) *ModuleDeleteOne {
	builder := c.Delete().Where(module.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ModuleDeleteOne{builder}
}

// Query returns a query builder for Module.
func (c *ModuleClient) Query() *ModuleQuery {
	return &ModuleQuery{
		config: c.config,
	}
}

// Get returns a Module entity by its id.
func (c *ModuleClient) Get(ctx context.Context, id uuid.UUID) (*Module, error) {
	return c.Query().Where(module.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ModuleClient) GetX(ctx context.Context, id uuid.UUID) *Module {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ModuleClient) Hooks() []Hook {
	hooks := c.hooks.Module
	return append(hooks[:len(hooks):len(hooks)], module.Hooks[:]...)
}

// RelatedTestCaseClient is a client for the RelatedTestCase schema.
type RelatedTestCaseClient struct {
	config
}

// NewRelatedTestCaseClient returns a client for the RelatedTestCase from the given config.
func NewRelatedTestCaseClient(c config) *RelatedTestCaseClient {
	return &RelatedTestCaseClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `relatedtestcase.Hooks(f(g(h())))`.
func (c *RelatedTestCaseClient) Use(hooks ...Hook) {
	c.hooks.RelatedTestCase = append(c.hooks.RelatedTestCase, hooks...)
}

// Create returns a builder for creating a RelatedTestCase entity.
func (c *RelatedTestCaseClient) Create() *RelatedTestCaseCreate {
	mutation := newRelatedTestCaseMutation(c.config, OpCreate)
	return &RelatedTestCaseCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of RelatedTestCase entities.
func (c *RelatedTestCaseClient) CreateBulk(builders ...*RelatedTestCaseCreate) *RelatedTestCaseCreateBulk {
	return &RelatedTestCaseCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for RelatedTestCase.
func (c *RelatedTestCaseClient) Update() *RelatedTestCaseUpdate {
	mutation := newRelatedTestCaseMutation(c.config, OpUpdate)
	return &RelatedTestCaseUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RelatedTestCaseClient) UpdateOne(rtc *RelatedTestCase) *RelatedTestCaseUpdateOne {
	mutation := newRelatedTestCaseMutation(c.config, OpUpdateOne, withRelatedTestCase(rtc))
	return &RelatedTestCaseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RelatedTestCaseClient) UpdateOneID(id uuid.UUID) *RelatedTestCaseUpdateOne {
	mutation := newRelatedTestCaseMutation(c.config, OpUpdateOne, withRelatedTestCaseID(id))
	return &RelatedTestCaseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for RelatedTestCase.
func (c *RelatedTestCaseClient) Delete() *RelatedTestCaseDelete {
	mutation := newRelatedTestCaseMutation(c.config, OpDelete)
	return &RelatedTestCaseDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *RelatedTestCaseClient) DeleteOne(rtc *RelatedTestCase) *RelatedTestCaseDeleteOne {
	return c.DeleteOneID(rtc.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *RelatedTestCaseClient) DeleteOneID(id uuid.UUID) *RelatedTestCaseDeleteOne {
	builder := c.Delete().Where(relatedtestcase.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RelatedTestCaseDeleteOne{builder}
}

// Query returns a query builder for RelatedTestCase.
func (c *RelatedTestCaseClient) Query() *RelatedTestCaseQuery {
	return &RelatedTestCaseQuery{
		config: c.config,
	}
}

// Get returns a RelatedTestCase entity by its id.
func (c *RelatedTestCaseClient) Get(ctx context.Context, id uuid.UUID) (*RelatedTestCase, error) {
	return c.Query().Where(relatedtestcase.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RelatedTestCaseClient) GetX(ctx context.Context, id uuid.UUID) *RelatedTestCase {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *RelatedTestCaseClient) Hooks() []Hook {
	hooks := c.hooks.RelatedTestCase
	return append(hooks[:len(hooks):len(hooks)], relatedtestcase.Hooks[:]...)
}

// TestCaseClient is a client for the TestCase schema.
type TestCaseClient struct {
	config
}

// NewTestCaseClient returns a client for the TestCase from the given config.
func NewTestCaseClient(c config) *TestCaseClient {
	return &TestCaseClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `testcase.Hooks(f(g(h())))`.
func (c *TestCaseClient) Use(hooks ...Hook) {
	c.hooks.TestCase = append(c.hooks.TestCase, hooks...)
}

// Create returns a builder for creating a TestCase entity.
func (c *TestCaseClient) Create() *TestCaseCreate {
	mutation := newTestCaseMutation(c.config, OpCreate)
	return &TestCaseCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of TestCase entities.
func (c *TestCaseClient) CreateBulk(builders ...*TestCaseCreate) *TestCaseCreateBulk {
	return &TestCaseCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for TestCase.
func (c *TestCaseClient) Update() *TestCaseUpdate {
	mutation := newTestCaseMutation(c.config, OpUpdate)
	return &TestCaseUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TestCaseClient) UpdateOne(tc *TestCase) *TestCaseUpdateOne {
	mutation := newTestCaseMutation(c.config, OpUpdateOne, withTestCase(tc))
	return &TestCaseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TestCaseClient) UpdateOneID(id uuid.UUID) *TestCaseUpdateOne {
	mutation := newTestCaseMutation(c.config, OpUpdateOne, withTestCaseID(id))
	return &TestCaseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for TestCase.
func (c *TestCaseClient) Delete() *TestCaseDelete {
	mutation := newTestCaseMutation(c.config, OpDelete)
	return &TestCaseDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TestCaseClient) DeleteOne(tc *TestCase) *TestCaseDeleteOne {
	return c.DeleteOneID(tc.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *TestCaseClient) DeleteOneID(id uuid.UUID) *TestCaseDeleteOne {
	builder := c.Delete().Where(testcase.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TestCaseDeleteOne{builder}
}

// Query returns a query builder for TestCase.
func (c *TestCaseClient) Query() *TestCaseQuery {
	return &TestCaseQuery{
		config: c.config,
	}
}

// Get returns a TestCase entity by its id.
func (c *TestCaseClient) Get(ctx context.Context, id uuid.UUID) (*TestCase, error) {
	return c.Query().Where(testcase.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TestCaseClient) GetX(ctx context.Context, id uuid.UUID) *TestCase {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *TestCaseClient) Hooks() []Hook {
	hooks := c.hooks.TestCase
	return append(hooks[:len(hooks):len(hooks)], testcase.Hooks[:]...)
}

// TestPlanClient is a client for the TestPlan schema.
type TestPlanClient struct {
	config
}

// NewTestPlanClient returns a client for the TestPlan from the given config.
func NewTestPlanClient(c config) *TestPlanClient {
	return &TestPlanClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `testplan.Hooks(f(g(h())))`.
func (c *TestPlanClient) Use(hooks ...Hook) {
	c.hooks.TestPlan = append(c.hooks.TestPlan, hooks...)
}

// Create returns a builder for creating a TestPlan entity.
func (c *TestPlanClient) Create() *TestPlanCreate {
	mutation := newTestPlanMutation(c.config, OpCreate)
	return &TestPlanCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of TestPlan entities.
func (c *TestPlanClient) CreateBulk(builders ...*TestPlanCreate) *TestPlanCreateBulk {
	return &TestPlanCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for TestPlan.
func (c *TestPlanClient) Update() *TestPlanUpdate {
	mutation := newTestPlanMutation(c.config, OpUpdate)
	return &TestPlanUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TestPlanClient) UpdateOne(tp *TestPlan) *TestPlanUpdateOne {
	mutation := newTestPlanMutation(c.config, OpUpdateOne, withTestPlan(tp))
	return &TestPlanUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TestPlanClient) UpdateOneID(id uuid.UUID) *TestPlanUpdateOne {
	mutation := newTestPlanMutation(c.config, OpUpdateOne, withTestPlanID(id))
	return &TestPlanUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for TestPlan.
func (c *TestPlanClient) Delete() *TestPlanDelete {
	mutation := newTestPlanMutation(c.config, OpDelete)
	return &TestPlanDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TestPlanClient) DeleteOne(tp *TestPlan) *TestPlanDeleteOne {
	return c.DeleteOneID(tp.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *TestPlanClient) DeleteOneID(id uuid.UUID) *TestPlanDeleteOne {
	builder := c.Delete().Where(testplan.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TestPlanDeleteOne{builder}
}

// Query returns a query builder for TestPlan.
func (c *TestPlanClient) Query() *TestPlanQuery {
	return &TestPlanQuery{
		config: c.config,
	}
}

// Get returns a TestPlan entity by its id.
func (c *TestPlanClient) Get(ctx context.Context, id uuid.UUID) (*TestPlan, error) {
	return c.Query().Where(testplan.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TestPlanClient) GetX(ctx context.Context, id uuid.UUID) *TestPlan {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *TestPlanClient) Hooks() []Hook {
	hooks := c.hooks.TestPlan
	return append(hooks[:len(hooks):len(hooks)], testplan.Hooks[:]...)
}
