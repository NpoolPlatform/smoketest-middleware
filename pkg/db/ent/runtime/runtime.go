// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"context"

	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/cond"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/module"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/plantestcase"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/schema"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/testcase"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/testplan"
	"github.com/google/uuid"

	"entgo.io/ent"
	"entgo.io/ent/privacy"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	condMixin := schema.Cond{}.Mixin()
	cond.Policy = privacy.NewPolicies(condMixin[0], schema.Cond{})
	cond.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := cond.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	condMixinFields0 := condMixin[0].Fields()
	_ = condMixinFields0
	condMixinFields1 := condMixin[1].Fields()
	_ = condMixinFields1
	condFields := schema.Cond{}.Fields()
	_ = condFields
	// condDescCreatedAt is the schema descriptor for created_at field.
	condDescCreatedAt := condMixinFields0[0].Descriptor()
	// cond.DefaultCreatedAt holds the default value on creation for the created_at field.
	cond.DefaultCreatedAt = condDescCreatedAt.Default.(func() uint32)
	// condDescUpdatedAt is the schema descriptor for updated_at field.
	condDescUpdatedAt := condMixinFields0[1].Descriptor()
	// cond.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	cond.DefaultUpdatedAt = condDescUpdatedAt.Default.(func() uint32)
	// cond.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	cond.UpdateDefaultUpdatedAt = condDescUpdatedAt.UpdateDefault.(func() uint32)
	// condDescDeletedAt is the schema descriptor for deleted_at field.
	condDescDeletedAt := condMixinFields0[2].Descriptor()
	// cond.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	cond.DefaultDeletedAt = condDescDeletedAt.Default.(func() uint32)
	// condDescEntID is the schema descriptor for ent_id field.
	condDescEntID := condMixinFields1[1].Descriptor()
	// cond.DefaultEntID holds the default value on creation for the ent_id field.
	cond.DefaultEntID = condDescEntID.Default.(func() uuid.UUID)
	// condDescCondType is the schema descriptor for cond_type field.
	condDescCondType := condFields[0].Descriptor()
	// cond.DefaultCondType holds the default value on creation for the cond_type field.
	cond.DefaultCondType = condDescCondType.Default.(string)
	// condDescTestCaseID is the schema descriptor for test_case_id field.
	condDescTestCaseID := condFields[1].Descriptor()
	// cond.DefaultTestCaseID holds the default value on creation for the test_case_id field.
	cond.DefaultTestCaseID = condDescTestCaseID.Default.(func() uuid.UUID)
	// condDescCondTestCaseID is the schema descriptor for cond_test_case_id field.
	condDescCondTestCaseID := condFields[2].Descriptor()
	// cond.DefaultCondTestCaseID holds the default value on creation for the cond_test_case_id field.
	cond.DefaultCondTestCaseID = condDescCondTestCaseID.Default.(func() uuid.UUID)
	// condDescArgumentMap is the schema descriptor for argument_map field.
	condDescArgumentMap := condFields[3].Descriptor()
	// cond.DefaultArgumentMap holds the default value on creation for the argument_map field.
	cond.DefaultArgumentMap = condDescArgumentMap.Default.(string)
	// condDescIndex is the schema descriptor for index field.
	condDescIndex := condFields[4].Descriptor()
	// cond.DefaultIndex holds the default value on creation for the index field.
	cond.DefaultIndex = condDescIndex.Default.(uint32)
	moduleMixin := schema.Module{}.Mixin()
	module.Policy = privacy.NewPolicies(moduleMixin[0], schema.Module{})
	module.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := module.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	moduleMixinFields0 := moduleMixin[0].Fields()
	_ = moduleMixinFields0
	moduleMixinFields1 := moduleMixin[1].Fields()
	_ = moduleMixinFields1
	moduleFields := schema.Module{}.Fields()
	_ = moduleFields
	// moduleDescCreatedAt is the schema descriptor for created_at field.
	moduleDescCreatedAt := moduleMixinFields0[0].Descriptor()
	// module.DefaultCreatedAt holds the default value on creation for the created_at field.
	module.DefaultCreatedAt = moduleDescCreatedAt.Default.(func() uint32)
	// moduleDescUpdatedAt is the schema descriptor for updated_at field.
	moduleDescUpdatedAt := moduleMixinFields0[1].Descriptor()
	// module.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	module.DefaultUpdatedAt = moduleDescUpdatedAt.Default.(func() uint32)
	// module.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	module.UpdateDefaultUpdatedAt = moduleDescUpdatedAt.UpdateDefault.(func() uint32)
	// moduleDescDeletedAt is the schema descriptor for deleted_at field.
	moduleDescDeletedAt := moduleMixinFields0[2].Descriptor()
	// module.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	module.DefaultDeletedAt = moduleDescDeletedAt.Default.(func() uint32)
	// moduleDescEntID is the schema descriptor for ent_id field.
	moduleDescEntID := moduleMixinFields1[1].Descriptor()
	// module.DefaultEntID holds the default value on creation for the ent_id field.
	module.DefaultEntID = moduleDescEntID.Default.(func() uuid.UUID)
	// moduleDescName is the schema descriptor for name field.
	moduleDescName := moduleFields[0].Descriptor()
	// module.DefaultName holds the default value on creation for the name field.
	module.DefaultName = moduleDescName.Default.(string)
	// moduleDescDescription is the schema descriptor for description field.
	moduleDescDescription := moduleFields[1].Descriptor()
	// module.DefaultDescription holds the default value on creation for the description field.
	module.DefaultDescription = moduleDescDescription.Default.(string)
	plantestcaseMixin := schema.PlanTestCase{}.Mixin()
	plantestcase.Policy = privacy.NewPolicies(plantestcaseMixin[0], schema.PlanTestCase{})
	plantestcase.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := plantestcase.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	plantestcaseMixinFields0 := plantestcaseMixin[0].Fields()
	_ = plantestcaseMixinFields0
	plantestcaseMixinFields1 := plantestcaseMixin[1].Fields()
	_ = plantestcaseMixinFields1
	plantestcaseFields := schema.PlanTestCase{}.Fields()
	_ = plantestcaseFields
	// plantestcaseDescCreatedAt is the schema descriptor for created_at field.
	plantestcaseDescCreatedAt := plantestcaseMixinFields0[0].Descriptor()
	// plantestcase.DefaultCreatedAt holds the default value on creation for the created_at field.
	plantestcase.DefaultCreatedAt = plantestcaseDescCreatedAt.Default.(func() uint32)
	// plantestcaseDescUpdatedAt is the schema descriptor for updated_at field.
	plantestcaseDescUpdatedAt := plantestcaseMixinFields0[1].Descriptor()
	// plantestcase.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	plantestcase.DefaultUpdatedAt = plantestcaseDescUpdatedAt.Default.(func() uint32)
	// plantestcase.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	plantestcase.UpdateDefaultUpdatedAt = plantestcaseDescUpdatedAt.UpdateDefault.(func() uint32)
	// plantestcaseDescDeletedAt is the schema descriptor for deleted_at field.
	plantestcaseDescDeletedAt := plantestcaseMixinFields0[2].Descriptor()
	// plantestcase.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	plantestcase.DefaultDeletedAt = plantestcaseDescDeletedAt.Default.(func() uint32)
	// plantestcaseDescEntID is the schema descriptor for ent_id field.
	plantestcaseDescEntID := plantestcaseMixinFields1[1].Descriptor()
	// plantestcase.DefaultEntID holds the default value on creation for the ent_id field.
	plantestcase.DefaultEntID = plantestcaseDescEntID.Default.(func() uuid.UUID)
	// plantestcaseDescTestPlanID is the schema descriptor for test_plan_id field.
	plantestcaseDescTestPlanID := plantestcaseFields[0].Descriptor()
	// plantestcase.DefaultTestPlanID holds the default value on creation for the test_plan_id field.
	plantestcase.DefaultTestPlanID = plantestcaseDescTestPlanID.Default.(func() uuid.UUID)
	// plantestcaseDescTestCaseID is the schema descriptor for test_case_id field.
	plantestcaseDescTestCaseID := plantestcaseFields[1].Descriptor()
	// plantestcase.DefaultTestCaseID holds the default value on creation for the test_case_id field.
	plantestcase.DefaultTestCaseID = plantestcaseDescTestCaseID.Default.(func() uuid.UUID)
	// plantestcaseDescInput is the schema descriptor for input field.
	plantestcaseDescInput := plantestcaseFields[2].Descriptor()
	// plantestcase.DefaultInput holds the default value on creation for the input field.
	plantestcase.DefaultInput = plantestcaseDescInput.Default.(string)
	// plantestcaseDescOutput is the schema descriptor for output field.
	plantestcaseDescOutput := plantestcaseFields[3].Descriptor()
	// plantestcase.DefaultOutput holds the default value on creation for the output field.
	plantestcase.DefaultOutput = plantestcaseDescOutput.Default.(string)
	// plantestcaseDescDescription is the schema descriptor for description field.
	plantestcaseDescDescription := plantestcaseFields[4].Descriptor()
	// plantestcase.DefaultDescription holds the default value on creation for the description field.
	plantestcase.DefaultDescription = plantestcaseDescDescription.Default.(string)
	// plantestcaseDescTestUserID is the schema descriptor for test_user_id field.
	plantestcaseDescTestUserID := plantestcaseFields[5].Descriptor()
	// plantestcase.DefaultTestUserID holds the default value on creation for the test_user_id field.
	plantestcase.DefaultTestUserID = plantestcaseDescTestUserID.Default.(func() uuid.UUID)
	// plantestcaseDescRunDuration is the schema descriptor for run_duration field.
	plantestcaseDescRunDuration := plantestcaseFields[6].Descriptor()
	// plantestcase.DefaultRunDuration holds the default value on creation for the run_duration field.
	plantestcase.DefaultRunDuration = plantestcaseDescRunDuration.Default.(uint32)
	// plantestcaseDescResult is the schema descriptor for result field.
	plantestcaseDescResult := plantestcaseFields[7].Descriptor()
	// plantestcase.DefaultResult holds the default value on creation for the result field.
	plantestcase.DefaultResult = plantestcaseDescResult.Default.(string)
	// plantestcaseDescIndex is the schema descriptor for index field.
	plantestcaseDescIndex := plantestcaseFields[8].Descriptor()
	// plantestcase.DefaultIndex holds the default value on creation for the index field.
	plantestcase.DefaultIndex = plantestcaseDescIndex.Default.(uint32)
	testcaseMixin := schema.TestCase{}.Mixin()
	testcase.Policy = privacy.NewPolicies(testcaseMixin[0], schema.TestCase{})
	testcase.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := testcase.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	testcaseMixinFields0 := testcaseMixin[0].Fields()
	_ = testcaseMixinFields0
	testcaseMixinFields1 := testcaseMixin[1].Fields()
	_ = testcaseMixinFields1
	testcaseFields := schema.TestCase{}.Fields()
	_ = testcaseFields
	// testcaseDescCreatedAt is the schema descriptor for created_at field.
	testcaseDescCreatedAt := testcaseMixinFields0[0].Descriptor()
	// testcase.DefaultCreatedAt holds the default value on creation for the created_at field.
	testcase.DefaultCreatedAt = testcaseDescCreatedAt.Default.(func() uint32)
	// testcaseDescUpdatedAt is the schema descriptor for updated_at field.
	testcaseDescUpdatedAt := testcaseMixinFields0[1].Descriptor()
	// testcase.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	testcase.DefaultUpdatedAt = testcaseDescUpdatedAt.Default.(func() uint32)
	// testcase.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	testcase.UpdateDefaultUpdatedAt = testcaseDescUpdatedAt.UpdateDefault.(func() uint32)
	// testcaseDescDeletedAt is the schema descriptor for deleted_at field.
	testcaseDescDeletedAt := testcaseMixinFields0[2].Descriptor()
	// testcase.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	testcase.DefaultDeletedAt = testcaseDescDeletedAt.Default.(func() uint32)
	// testcaseDescEntID is the schema descriptor for ent_id field.
	testcaseDescEntID := testcaseMixinFields1[1].Descriptor()
	// testcase.DefaultEntID holds the default value on creation for the ent_id field.
	testcase.DefaultEntID = testcaseDescEntID.Default.(func() uuid.UUID)
	// testcaseDescName is the schema descriptor for name field.
	testcaseDescName := testcaseFields[0].Descriptor()
	// testcase.DefaultName holds the default value on creation for the name field.
	testcase.DefaultName = testcaseDescName.Default.(string)
	// testcaseDescDescription is the schema descriptor for description field.
	testcaseDescDescription := testcaseFields[1].Descriptor()
	// testcase.DefaultDescription holds the default value on creation for the description field.
	testcase.DefaultDescription = testcaseDescDescription.Default.(string)
	// testcaseDescModuleID is the schema descriptor for module_id field.
	testcaseDescModuleID := testcaseFields[2].Descriptor()
	// testcase.DefaultModuleID holds the default value on creation for the module_id field.
	testcase.DefaultModuleID = testcaseDescModuleID.Default.(func() uuid.UUID)
	// testcaseDescAPIID is the schema descriptor for api_id field.
	testcaseDescAPIID := testcaseFields[3].Descriptor()
	// testcase.DefaultAPIID holds the default value on creation for the api_id field.
	testcase.DefaultAPIID = testcaseDescAPIID.Default.(func() uuid.UUID)
	// testcaseDescInput is the schema descriptor for input field.
	testcaseDescInput := testcaseFields[4].Descriptor()
	// testcase.DefaultInput holds the default value on creation for the input field.
	testcase.DefaultInput = testcaseDescInput.Default.(string)
	// testcaseDescInputDesc is the schema descriptor for input_desc field.
	testcaseDescInputDesc := testcaseFields[5].Descriptor()
	// testcase.DefaultInputDesc holds the default value on creation for the input_desc field.
	testcase.DefaultInputDesc = testcaseDescInputDesc.Default.(string)
	// testcaseDescExpectation is the schema descriptor for expectation field.
	testcaseDescExpectation := testcaseFields[6].Descriptor()
	// testcase.DefaultExpectation holds the default value on creation for the expectation field.
	testcase.DefaultExpectation = testcaseDescExpectation.Default.(string)
	// testcaseDescOutputDesc is the schema descriptor for output_desc field.
	testcaseDescOutputDesc := testcaseFields[7].Descriptor()
	// testcase.DefaultOutputDesc holds the default value on creation for the output_desc field.
	testcase.DefaultOutputDesc = testcaseDescOutputDesc.Default.(string)
	// testcaseDescTestCaseType is the schema descriptor for test_case_type field.
	testcaseDescTestCaseType := testcaseFields[8].Descriptor()
	// testcase.DefaultTestCaseType holds the default value on creation for the test_case_type field.
	testcase.DefaultTestCaseType = testcaseDescTestCaseType.Default.(string)
	// testcaseDescTestCaseClass is the schema descriptor for test_case_class field.
	testcaseDescTestCaseClass := testcaseFields[9].Descriptor()
	// testcase.DefaultTestCaseClass holds the default value on creation for the test_case_class field.
	testcase.DefaultTestCaseClass = testcaseDescTestCaseClass.Default.(string)
	// testcaseDescDeprecated is the schema descriptor for deprecated field.
	testcaseDescDeprecated := testcaseFields[10].Descriptor()
	// testcase.DefaultDeprecated holds the default value on creation for the deprecated field.
	testcase.DefaultDeprecated = testcaseDescDeprecated.Default.(bool)
	testplanMixin := schema.TestPlan{}.Mixin()
	testplan.Policy = privacy.NewPolicies(testplanMixin[0], schema.TestPlan{})
	testplan.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := testplan.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	testplanMixinFields0 := testplanMixin[0].Fields()
	_ = testplanMixinFields0
	testplanMixinFields1 := testplanMixin[1].Fields()
	_ = testplanMixinFields1
	testplanFields := schema.TestPlan{}.Fields()
	_ = testplanFields
	// testplanDescCreatedAt is the schema descriptor for created_at field.
	testplanDescCreatedAt := testplanMixinFields0[0].Descriptor()
	// testplan.DefaultCreatedAt holds the default value on creation for the created_at field.
	testplan.DefaultCreatedAt = testplanDescCreatedAt.Default.(func() uint32)
	// testplanDescUpdatedAt is the schema descriptor for updated_at field.
	testplanDescUpdatedAt := testplanMixinFields0[1].Descriptor()
	// testplan.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	testplan.DefaultUpdatedAt = testplanDescUpdatedAt.Default.(func() uint32)
	// testplan.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	testplan.UpdateDefaultUpdatedAt = testplanDescUpdatedAt.UpdateDefault.(func() uint32)
	// testplanDescDeletedAt is the schema descriptor for deleted_at field.
	testplanDescDeletedAt := testplanMixinFields0[2].Descriptor()
	// testplan.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	testplan.DefaultDeletedAt = testplanDescDeletedAt.Default.(func() uint32)
	// testplanDescEntID is the schema descriptor for ent_id field.
	testplanDescEntID := testplanMixinFields1[1].Descriptor()
	// testplan.DefaultEntID holds the default value on creation for the ent_id field.
	testplan.DefaultEntID = testplanDescEntID.Default.(func() uuid.UUID)
	// testplanDescName is the schema descriptor for name field.
	testplanDescName := testplanFields[0].Descriptor()
	// testplan.DefaultName holds the default value on creation for the name field.
	testplan.DefaultName = testplanDescName.Default.(string)
	// testplanDescState is the schema descriptor for state field.
	testplanDescState := testplanFields[1].Descriptor()
	// testplan.DefaultState holds the default value on creation for the state field.
	testplan.DefaultState = testplanDescState.Default.(string)
	// testplanDescCreatedBy is the schema descriptor for created_by field.
	testplanDescCreatedBy := testplanFields[2].Descriptor()
	// testplan.DefaultCreatedBy holds the default value on creation for the created_by field.
	testplan.DefaultCreatedBy = testplanDescCreatedBy.Default.(func() uuid.UUID)
	// testplanDescExecutor is the schema descriptor for executor field.
	testplanDescExecutor := testplanFields[3].Descriptor()
	// testplan.DefaultExecutor holds the default value on creation for the executor field.
	testplan.DefaultExecutor = testplanDescExecutor.Default.(func() uuid.UUID)
	// testplanDescFails is the schema descriptor for fails field.
	testplanDescFails := testplanFields[4].Descriptor()
	// testplan.DefaultFails holds the default value on creation for the fails field.
	testplan.DefaultFails = testplanDescFails.Default.(uint32)
	// testplanDescPasses is the schema descriptor for passes field.
	testplanDescPasses := testplanFields[5].Descriptor()
	// testplan.DefaultPasses holds the default value on creation for the passes field.
	testplan.DefaultPasses = testplanDescPasses.Default.(uint32)
	// testplanDescSkips is the schema descriptor for skips field.
	testplanDescSkips := testplanFields[6].Descriptor()
	// testplan.DefaultSkips holds the default value on creation for the skips field.
	testplan.DefaultSkips = testplanDescSkips.Default.(uint32)
	// testplanDescRunDuration is the schema descriptor for run_duration field.
	testplanDescRunDuration := testplanFields[7].Descriptor()
	// testplan.DefaultRunDuration holds the default value on creation for the run_duration field.
	testplan.DefaultRunDuration = testplanDescRunDuration.Default.(uint32)
	// testplanDescDeadline is the schema descriptor for deadline field.
	testplanDescDeadline := testplanFields[8].Descriptor()
	// testplan.DefaultDeadline holds the default value on creation for the deadline field.
	testplan.DefaultDeadline = testplanDescDeadline.Default.(uint32)
	// testplanDescResult is the schema descriptor for result field.
	testplanDescResult := testplanFields[9].Descriptor()
	// testplan.DefaultResult holds the default value on creation for the result field.
	testplan.DefaultResult = testplanDescResult.Default.(string)
}

const (
	Version = "v0.11.2"                                         // Version of ent codegen.
	Sum     = "h1:UM2/BUhF2FfsxPHRxLjQbhqJNaDdVlOwNIAMLs2jyto=" // Sum of ent codegen.
)
