// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/detail"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/module"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/planrelatedtestcase"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/relatedtestcase"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/testcase"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/testplan"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 6)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   detail.Table,
			Columns: detail.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: detail.FieldID,
			},
		},
		Type: "Detail",
		Fields: map[string]*sqlgraph.FieldSpec{
			detail.FieldCreatedAt:       {Type: field.TypeUint32, Column: detail.FieldCreatedAt},
			detail.FieldUpdatedAt:       {Type: field.TypeUint32, Column: detail.FieldUpdatedAt},
			detail.FieldDeletedAt:       {Type: field.TypeUint32, Column: detail.FieldDeletedAt},
			detail.FieldAppID:           {Type: field.TypeUUID, Column: detail.FieldAppID},
			detail.FieldUserID:          {Type: field.TypeUUID, Column: detail.FieldUserID},
			detail.FieldCoinTypeID:      {Type: field.TypeUUID, Column: detail.FieldCoinTypeID},
			detail.FieldIoType:          {Type: field.TypeString, Column: detail.FieldIoType},
			detail.FieldIoSubType:       {Type: field.TypeString, Column: detail.FieldIoSubType},
			detail.FieldAmount:          {Type: field.TypeOther, Column: detail.FieldAmount},
			detail.FieldFromCoinTypeID:  {Type: field.TypeUUID, Column: detail.FieldFromCoinTypeID},
			detail.FieldCoinUsdCurrency: {Type: field.TypeOther, Column: detail.FieldCoinUsdCurrency},
			detail.FieldIoExtra:         {Type: field.TypeString, Column: detail.FieldIoExtra},
			detail.FieldFromOldID:       {Type: field.TypeUUID, Column: detail.FieldFromOldID},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   module.Table,
			Columns: module.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: module.FieldID,
			},
		},
		Type: "Module",
		Fields: map[string]*sqlgraph.FieldSpec{
			module.FieldCreatedAt:   {Type: field.TypeUint32, Column: module.FieldCreatedAt},
			module.FieldUpdatedAt:   {Type: field.TypeUint32, Column: module.FieldUpdatedAt},
			module.FieldDeletedAt:   {Type: field.TypeUint32, Column: module.FieldDeletedAt},
			module.FieldName:        {Type: field.TypeString, Column: module.FieldName},
			module.FieldDescription: {Type: field.TypeString, Column: module.FieldDescription},
		},
	}
	graph.Nodes[2] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   planrelatedtestcase.Table,
			Columns: planrelatedtestcase.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: planrelatedtestcase.FieldID,
			},
		},
		Type: "PlanRelatedTestCase",
		Fields: map[string]*sqlgraph.FieldSpec{
			planrelatedtestcase.FieldCreatedAt:      {Type: field.TypeUint32, Column: planrelatedtestcase.FieldCreatedAt},
			planrelatedtestcase.FieldUpdatedAt:      {Type: field.TypeUint32, Column: planrelatedtestcase.FieldUpdatedAt},
			planrelatedtestcase.FieldDeletedAt:      {Type: field.TypeUint32, Column: planrelatedtestcase.FieldDeletedAt},
			planrelatedtestcase.FieldTestPlanID:     {Type: field.TypeUUID, Column: planrelatedtestcase.FieldTestPlanID},
			planrelatedtestcase.FieldTestCaseID:     {Type: field.TypeUUID, Column: planrelatedtestcase.FieldTestCaseID},
			planrelatedtestcase.FieldTestCaseOutput: {Type: field.TypeString, Column: planrelatedtestcase.FieldTestCaseOutput},
			planrelatedtestcase.FieldDescription:    {Type: field.TypeString, Column: planrelatedtestcase.FieldDescription},
			planrelatedtestcase.FieldTestUserID:     {Type: field.TypeUUID, Column: planrelatedtestcase.FieldTestUserID},
			planrelatedtestcase.FieldRunDuration:    {Type: field.TypeUint32, Column: planrelatedtestcase.FieldRunDuration},
			planrelatedtestcase.FieldTestCaseResult: {Type: field.TypeString, Column: planrelatedtestcase.FieldTestCaseResult},
		},
	}
	graph.Nodes[3] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   relatedtestcase.Table,
			Columns: relatedtestcase.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: relatedtestcase.FieldID,
			},
		},
		Type: "RelatedTestCase",
		Fields: map[string]*sqlgraph.FieldSpec{
			relatedtestcase.FieldCreatedAt:         {Type: field.TypeUint32, Column: relatedtestcase.FieldCreatedAt},
			relatedtestcase.FieldUpdatedAt:         {Type: field.TypeUint32, Column: relatedtestcase.FieldUpdatedAt},
			relatedtestcase.FieldDeletedAt:         {Type: field.TypeUint32, Column: relatedtestcase.FieldDeletedAt},
			relatedtestcase.FieldCondType:          {Type: field.TypeString, Column: relatedtestcase.FieldCondType},
			relatedtestcase.FieldTestCaseID:        {Type: field.TypeUUID, Column: relatedtestcase.FieldTestCaseID},
			relatedtestcase.FieldRelatedTestCaseID: {Type: field.TypeUUID, Column: relatedtestcase.FieldRelatedTestCaseID},
			relatedtestcase.FieldArgumentsTransfer: {Type: field.TypeString, Column: relatedtestcase.FieldArgumentsTransfer},
			relatedtestcase.FieldIndex:             {Type: field.TypeUint32, Column: relatedtestcase.FieldIndex},
		},
	}
	graph.Nodes[4] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   testcase.Table,
			Columns: testcase.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: testcase.FieldID,
			},
		},
		Type: "TestCase",
		Fields: map[string]*sqlgraph.FieldSpec{
			testcase.FieldCreatedAt:         {Type: field.TypeUint32, Column: testcase.FieldCreatedAt},
			testcase.FieldUpdatedAt:         {Type: field.TypeUint32, Column: testcase.FieldUpdatedAt},
			testcase.FieldDeletedAt:         {Type: field.TypeUint32, Column: testcase.FieldDeletedAt},
			testcase.FieldName:              {Type: field.TypeString, Column: testcase.FieldName},
			testcase.FieldDescription:       {Type: field.TypeString, Column: testcase.FieldDescription},
			testcase.FieldModuleID:          {Type: field.TypeUUID, Column: testcase.FieldModuleID},
			testcase.FieldAPIID:             {Type: field.TypeUUID, Column: testcase.FieldAPIID},
			testcase.FieldArguments:         {Type: field.TypeString, Column: testcase.FieldArguments},
			testcase.FieldExpectationResult: {Type: field.TypeString, Column: testcase.FieldExpectationResult},
			testcase.FieldTestCaseType:      {Type: field.TypeString, Column: testcase.FieldTestCaseType},
			testcase.FieldDeprecated:        {Type: field.TypeBool, Column: testcase.FieldDeprecated},
		},
	}
	graph.Nodes[5] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   testplan.Table,
			Columns: testplan.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: testplan.FieldID,
			},
		},
		Type: "TestPlan",
		Fields: map[string]*sqlgraph.FieldSpec{
			testplan.FieldCreatedAt:            {Type: field.TypeUint32, Column: testplan.FieldCreatedAt},
			testplan.FieldUpdatedAt:            {Type: field.TypeUint32, Column: testplan.FieldUpdatedAt},
			testplan.FieldDeletedAt:            {Type: field.TypeUint32, Column: testplan.FieldDeletedAt},
			testplan.FieldName:                 {Type: field.TypeString, Column: testplan.FieldName},
			testplan.FieldState:                {Type: field.TypeString, Column: testplan.FieldState},
			testplan.FieldOwnerID:              {Type: field.TypeUUID, Column: testplan.FieldOwnerID},
			testplan.FieldResponsibleUserID:    {Type: field.TypeUUID, Column: testplan.FieldResponsibleUserID},
			testplan.FieldFailedTestCaseCount:  {Type: field.TypeUint32, Column: testplan.FieldFailedTestCaseCount},
			testplan.FieldPassedTestCaseCount:  {Type: field.TypeUint32, Column: testplan.FieldPassedTestCaseCount},
			testplan.FieldSkippedTestCaseCount: {Type: field.TypeUint32, Column: testplan.FieldSkippedTestCaseCount},
			testplan.FieldRunDuration:          {Type: field.TypeUint32, Column: testplan.FieldRunDuration},
			testplan.FieldDeadline:             {Type: field.TypeUint32, Column: testplan.FieldDeadline},
			testplan.FieldTestResult:           {Type: field.TypeString, Column: testplan.FieldTestResult},
		},
	}
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (dq *DetailQuery) addPredicate(pred func(s *sql.Selector)) {
	dq.predicates = append(dq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the DetailQuery builder.
func (dq *DetailQuery) Filter() *DetailFilter {
	return &DetailFilter{config: dq.config, predicateAdder: dq}
}

// addPredicate implements the predicateAdder interface.
func (m *DetailMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the DetailMutation builder.
func (m *DetailMutation) Filter() *DetailFilter {
	return &DetailFilter{config: m.config, predicateAdder: m}
}

// DetailFilter provides a generic filtering capability at runtime for DetailQuery.
type DetailFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *DetailFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *DetailFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(detail.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *DetailFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(detail.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *DetailFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(detail.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *DetailFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(detail.FieldDeletedAt))
}

// WhereAppID applies the entql [16]byte predicate on the app_id field.
func (f *DetailFilter) WhereAppID(p entql.ValueP) {
	f.Where(p.Field(detail.FieldAppID))
}

// WhereUserID applies the entql [16]byte predicate on the user_id field.
func (f *DetailFilter) WhereUserID(p entql.ValueP) {
	f.Where(p.Field(detail.FieldUserID))
}

// WhereCoinTypeID applies the entql [16]byte predicate on the coin_type_id field.
func (f *DetailFilter) WhereCoinTypeID(p entql.ValueP) {
	f.Where(p.Field(detail.FieldCoinTypeID))
}

// WhereIoType applies the entql string predicate on the io_type field.
func (f *DetailFilter) WhereIoType(p entql.StringP) {
	f.Where(p.Field(detail.FieldIoType))
}

// WhereIoSubType applies the entql string predicate on the io_sub_type field.
func (f *DetailFilter) WhereIoSubType(p entql.StringP) {
	f.Where(p.Field(detail.FieldIoSubType))
}

// WhereAmount applies the entql other predicate on the amount field.
func (f *DetailFilter) WhereAmount(p entql.OtherP) {
	f.Where(p.Field(detail.FieldAmount))
}

// WhereFromCoinTypeID applies the entql [16]byte predicate on the from_coin_type_id field.
func (f *DetailFilter) WhereFromCoinTypeID(p entql.ValueP) {
	f.Where(p.Field(detail.FieldFromCoinTypeID))
}

// WhereCoinUsdCurrency applies the entql other predicate on the coin_usd_currency field.
func (f *DetailFilter) WhereCoinUsdCurrency(p entql.OtherP) {
	f.Where(p.Field(detail.FieldCoinUsdCurrency))
}

// WhereIoExtra applies the entql string predicate on the io_extra field.
func (f *DetailFilter) WhereIoExtra(p entql.StringP) {
	f.Where(p.Field(detail.FieldIoExtra))
}

// WhereFromOldID applies the entql [16]byte predicate on the from_old_id field.
func (f *DetailFilter) WhereFromOldID(p entql.ValueP) {
	f.Where(p.Field(detail.FieldFromOldID))
}

// addPredicate implements the predicateAdder interface.
func (mq *ModuleQuery) addPredicate(pred func(s *sql.Selector)) {
	mq.predicates = append(mq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the ModuleQuery builder.
func (mq *ModuleQuery) Filter() *ModuleFilter {
	return &ModuleFilter{config: mq.config, predicateAdder: mq}
}

// addPredicate implements the predicateAdder interface.
func (m *ModuleMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the ModuleMutation builder.
func (m *ModuleMutation) Filter() *ModuleFilter {
	return &ModuleFilter{config: m.config, predicateAdder: m}
}

// ModuleFilter provides a generic filtering capability at runtime for ModuleQuery.
type ModuleFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *ModuleFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *ModuleFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(module.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *ModuleFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(module.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *ModuleFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(module.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *ModuleFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(module.FieldDeletedAt))
}

// WhereName applies the entql string predicate on the name field.
func (f *ModuleFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(module.FieldName))
}

// WhereDescription applies the entql string predicate on the description field.
func (f *ModuleFilter) WhereDescription(p entql.StringP) {
	f.Where(p.Field(module.FieldDescription))
}

// addPredicate implements the predicateAdder interface.
func (prtcq *PlanRelatedTestCaseQuery) addPredicate(pred func(s *sql.Selector)) {
	prtcq.predicates = append(prtcq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the PlanRelatedTestCaseQuery builder.
func (prtcq *PlanRelatedTestCaseQuery) Filter() *PlanRelatedTestCaseFilter {
	return &PlanRelatedTestCaseFilter{config: prtcq.config, predicateAdder: prtcq}
}

// addPredicate implements the predicateAdder interface.
func (m *PlanRelatedTestCaseMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the PlanRelatedTestCaseMutation builder.
func (m *PlanRelatedTestCaseMutation) Filter() *PlanRelatedTestCaseFilter {
	return &PlanRelatedTestCaseFilter{config: m.config, predicateAdder: m}
}

// PlanRelatedTestCaseFilter provides a generic filtering capability at runtime for PlanRelatedTestCaseQuery.
type PlanRelatedTestCaseFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *PlanRelatedTestCaseFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[2].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *PlanRelatedTestCaseFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(planrelatedtestcase.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *PlanRelatedTestCaseFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(planrelatedtestcase.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *PlanRelatedTestCaseFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(planrelatedtestcase.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *PlanRelatedTestCaseFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(planrelatedtestcase.FieldDeletedAt))
}

// WhereTestPlanID applies the entql [16]byte predicate on the test_plan_id field.
func (f *PlanRelatedTestCaseFilter) WhereTestPlanID(p entql.ValueP) {
	f.Where(p.Field(planrelatedtestcase.FieldTestPlanID))
}

// WhereTestCaseID applies the entql [16]byte predicate on the test_case_id field.
func (f *PlanRelatedTestCaseFilter) WhereTestCaseID(p entql.ValueP) {
	f.Where(p.Field(planrelatedtestcase.FieldTestCaseID))
}

// WhereTestCaseOutput applies the entql string predicate on the test_case_output field.
func (f *PlanRelatedTestCaseFilter) WhereTestCaseOutput(p entql.StringP) {
	f.Where(p.Field(planrelatedtestcase.FieldTestCaseOutput))
}

// WhereDescription applies the entql string predicate on the description field.
func (f *PlanRelatedTestCaseFilter) WhereDescription(p entql.StringP) {
	f.Where(p.Field(planrelatedtestcase.FieldDescription))
}

// WhereTestUserID applies the entql [16]byte predicate on the test_user_id field.
func (f *PlanRelatedTestCaseFilter) WhereTestUserID(p entql.ValueP) {
	f.Where(p.Field(planrelatedtestcase.FieldTestUserID))
}

// WhereRunDuration applies the entql uint32 predicate on the run_duration field.
func (f *PlanRelatedTestCaseFilter) WhereRunDuration(p entql.Uint32P) {
	f.Where(p.Field(planrelatedtestcase.FieldRunDuration))
}

// WhereTestCaseResult applies the entql string predicate on the test_case_result field.
func (f *PlanRelatedTestCaseFilter) WhereTestCaseResult(p entql.StringP) {
	f.Where(p.Field(planrelatedtestcase.FieldTestCaseResult))
}

// addPredicate implements the predicateAdder interface.
func (rtcq *RelatedTestCaseQuery) addPredicate(pred func(s *sql.Selector)) {
	rtcq.predicates = append(rtcq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the RelatedTestCaseQuery builder.
func (rtcq *RelatedTestCaseQuery) Filter() *RelatedTestCaseFilter {
	return &RelatedTestCaseFilter{config: rtcq.config, predicateAdder: rtcq}
}

// addPredicate implements the predicateAdder interface.
func (m *RelatedTestCaseMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the RelatedTestCaseMutation builder.
func (m *RelatedTestCaseMutation) Filter() *RelatedTestCaseFilter {
	return &RelatedTestCaseFilter{config: m.config, predicateAdder: m}
}

// RelatedTestCaseFilter provides a generic filtering capability at runtime for RelatedTestCaseQuery.
type RelatedTestCaseFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *RelatedTestCaseFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[3].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *RelatedTestCaseFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(relatedtestcase.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *RelatedTestCaseFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(relatedtestcase.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *RelatedTestCaseFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(relatedtestcase.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *RelatedTestCaseFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(relatedtestcase.FieldDeletedAt))
}

// WhereCondType applies the entql string predicate on the cond_type field.
func (f *RelatedTestCaseFilter) WhereCondType(p entql.StringP) {
	f.Where(p.Field(relatedtestcase.FieldCondType))
}

// WhereTestCaseID applies the entql [16]byte predicate on the test_case_id field.
func (f *RelatedTestCaseFilter) WhereTestCaseID(p entql.ValueP) {
	f.Where(p.Field(relatedtestcase.FieldTestCaseID))
}

// WhereRelatedTestCaseID applies the entql [16]byte predicate on the related_test_case_id field.
func (f *RelatedTestCaseFilter) WhereRelatedTestCaseID(p entql.ValueP) {
	f.Where(p.Field(relatedtestcase.FieldRelatedTestCaseID))
}

// WhereArgumentsTransfer applies the entql string predicate on the arguments_transfer field.
func (f *RelatedTestCaseFilter) WhereArgumentsTransfer(p entql.StringP) {
	f.Where(p.Field(relatedtestcase.FieldArgumentsTransfer))
}

// WhereIndex applies the entql uint32 predicate on the index field.
func (f *RelatedTestCaseFilter) WhereIndex(p entql.Uint32P) {
	f.Where(p.Field(relatedtestcase.FieldIndex))
}

// addPredicate implements the predicateAdder interface.
func (tcq *TestCaseQuery) addPredicate(pred func(s *sql.Selector)) {
	tcq.predicates = append(tcq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the TestCaseQuery builder.
func (tcq *TestCaseQuery) Filter() *TestCaseFilter {
	return &TestCaseFilter{config: tcq.config, predicateAdder: tcq}
}

// addPredicate implements the predicateAdder interface.
func (m *TestCaseMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the TestCaseMutation builder.
func (m *TestCaseMutation) Filter() *TestCaseFilter {
	return &TestCaseFilter{config: m.config, predicateAdder: m}
}

// TestCaseFilter provides a generic filtering capability at runtime for TestCaseQuery.
type TestCaseFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *TestCaseFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[4].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *TestCaseFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(testcase.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *TestCaseFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(testcase.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *TestCaseFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(testcase.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *TestCaseFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(testcase.FieldDeletedAt))
}

// WhereName applies the entql string predicate on the name field.
func (f *TestCaseFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(testcase.FieldName))
}

// WhereDescription applies the entql string predicate on the description field.
func (f *TestCaseFilter) WhereDescription(p entql.StringP) {
	f.Where(p.Field(testcase.FieldDescription))
}

// WhereModuleID applies the entql [16]byte predicate on the module_id field.
func (f *TestCaseFilter) WhereModuleID(p entql.ValueP) {
	f.Where(p.Field(testcase.FieldModuleID))
}

// WhereAPIID applies the entql [16]byte predicate on the api_id field.
func (f *TestCaseFilter) WhereAPIID(p entql.ValueP) {
	f.Where(p.Field(testcase.FieldAPIID))
}

// WhereArguments applies the entql string predicate on the arguments field.
func (f *TestCaseFilter) WhereArguments(p entql.StringP) {
	f.Where(p.Field(testcase.FieldArguments))
}

// WhereExpectationResult applies the entql string predicate on the expectation_result field.
func (f *TestCaseFilter) WhereExpectationResult(p entql.StringP) {
	f.Where(p.Field(testcase.FieldExpectationResult))
}

// WhereTestCaseType applies the entql string predicate on the test_case_type field.
func (f *TestCaseFilter) WhereTestCaseType(p entql.StringP) {
	f.Where(p.Field(testcase.FieldTestCaseType))
}

// WhereDeprecated applies the entql bool predicate on the deprecated field.
func (f *TestCaseFilter) WhereDeprecated(p entql.BoolP) {
	f.Where(p.Field(testcase.FieldDeprecated))
}

// addPredicate implements the predicateAdder interface.
func (tpq *TestPlanQuery) addPredicate(pred func(s *sql.Selector)) {
	tpq.predicates = append(tpq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the TestPlanQuery builder.
func (tpq *TestPlanQuery) Filter() *TestPlanFilter {
	return &TestPlanFilter{config: tpq.config, predicateAdder: tpq}
}

// addPredicate implements the predicateAdder interface.
func (m *TestPlanMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the TestPlanMutation builder.
func (m *TestPlanMutation) Filter() *TestPlanFilter {
	return &TestPlanFilter{config: m.config, predicateAdder: m}
}

// TestPlanFilter provides a generic filtering capability at runtime for TestPlanQuery.
type TestPlanFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *TestPlanFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[5].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *TestPlanFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(testplan.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *TestPlanFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(testplan.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *TestPlanFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(testplan.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *TestPlanFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(testplan.FieldDeletedAt))
}

// WhereName applies the entql string predicate on the name field.
func (f *TestPlanFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(testplan.FieldName))
}

// WhereState applies the entql string predicate on the state field.
func (f *TestPlanFilter) WhereState(p entql.StringP) {
	f.Where(p.Field(testplan.FieldState))
}

// WhereOwnerID applies the entql [16]byte predicate on the owner_id field.
func (f *TestPlanFilter) WhereOwnerID(p entql.ValueP) {
	f.Where(p.Field(testplan.FieldOwnerID))
}

// WhereResponsibleUserID applies the entql [16]byte predicate on the responsible_user_id field.
func (f *TestPlanFilter) WhereResponsibleUserID(p entql.ValueP) {
	f.Where(p.Field(testplan.FieldResponsibleUserID))
}

// WhereFailedTestCaseCount applies the entql uint32 predicate on the failed_test_case_count field.
func (f *TestPlanFilter) WhereFailedTestCaseCount(p entql.Uint32P) {
	f.Where(p.Field(testplan.FieldFailedTestCaseCount))
}

// WherePassedTestCaseCount applies the entql uint32 predicate on the passed_test_case_count field.
func (f *TestPlanFilter) WherePassedTestCaseCount(p entql.Uint32P) {
	f.Where(p.Field(testplan.FieldPassedTestCaseCount))
}

// WhereSkippedTestCaseCount applies the entql uint32 predicate on the skipped_test_case_count field.
func (f *TestPlanFilter) WhereSkippedTestCaseCount(p entql.Uint32P) {
	f.Where(p.Field(testplan.FieldSkippedTestCaseCount))
}

// WhereRunDuration applies the entql uint32 predicate on the run_duration field.
func (f *TestPlanFilter) WhereRunDuration(p entql.Uint32P) {
	f.Where(p.Field(testplan.FieldRunDuration))
}

// WhereDeadline applies the entql uint32 predicate on the deadline field.
func (f *TestPlanFilter) WhereDeadline(p entql.Uint32P) {
	f.Where(p.Field(testplan.FieldDeadline))
}

// WhereTestResult applies the entql string predicate on the test_result field.
func (f *TestPlanFilter) WhereTestResult(p entql.StringP) {
	f.Where(p.Field(testplan.FieldTestResult))
}
