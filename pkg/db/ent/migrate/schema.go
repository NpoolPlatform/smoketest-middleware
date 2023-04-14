// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// DetailsColumns holds the columns for the "details" table.
	DetailsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "io_type", Type: field.TypeString, Nullable: true, Default: "DefaultType"},
		{Name: "io_sub_type", Type: field.TypeString, Nullable: true, Default: "DefaultSubType"},
		{Name: "amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "from_coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "coin_usd_currency", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "io_extra", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "from_old_id", Type: field.TypeUUID, Nullable: true},
	}
	// DetailsTable holds the schema information for the "details" table.
	DetailsTable = &schema.Table{
		Name:       "details",
		Columns:    DetailsColumns,
		PrimaryKey: []*schema.Column{DetailsColumns[0]},
	}
	// ModulesColumns holds the columns for the "modules" table.
	ModulesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString, Nullable: true, Default: ""},
	}
	// ModulesTable holds the schema information for the "modules" table.
	ModulesTable = &schema.Table{
		Name:       "modules",
		Columns:    ModulesColumns,
		PrimaryKey: []*schema.Column{ModulesColumns[0]},
	}
	// PlanRelatedTestCasesColumns holds the columns for the "plan_related_test_cases" table.
	PlanRelatedTestCasesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "test_plan_id", Type: field.TypeUUID, Nullable: true},
		{Name: "test_case_id", Type: field.TypeUUID, Nullable: true},
		{Name: "test_case_output", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "description", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "test_user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "run_duration", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "test_case_result", Type: field.TypeString, Nullable: true, Default: "DefaultTestCaseResult"},
	}
	// PlanRelatedTestCasesTable holds the schema information for the "plan_related_test_cases" table.
	PlanRelatedTestCasesTable = &schema.Table{
		Name:       "plan_related_test_cases",
		Columns:    PlanRelatedTestCasesColumns,
		PrimaryKey: []*schema.Column{PlanRelatedTestCasesColumns[0]},
	}
	// RelatedTestCasesColumns holds the columns for the "related_test_cases" table.
	RelatedTestCasesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "cond_type", Type: field.TypeString, Nullable: true, Default: "DefaultCondType"},
		{Name: "test_case_id", Type: field.TypeUUID, Nullable: true},
		{Name: "related_test_case_id", Type: field.TypeUUID, Nullable: true},
		{Name: "arguments_transfer", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "index", Type: field.TypeUint32, Nullable: true, Default: 0},
	}
	// RelatedTestCasesTable holds the schema information for the "related_test_cases" table.
	RelatedTestCasesTable = &schema.Table{
		Name:       "related_test_cases",
		Columns:    RelatedTestCasesColumns,
		PrimaryKey: []*schema.Column{RelatedTestCasesColumns[0]},
	}
	// TestCasesColumns holds the columns for the "test_cases" table.
	TestCasesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "name", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "description", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "module_id", Type: field.TypeUUID, Nullable: true},
		{Name: "api_id", Type: field.TypeUUID, Nullable: true},
		{Name: "arguments", Type: field.TypeString, Nullable: true, Size: 2147483647, Default: ""},
		{Name: "arg_type_description", Type: field.TypeString, Nullable: true, Size: 2147483647, Default: ""},
		{Name: "expectation_result", Type: field.TypeString, Nullable: true, Size: 2147483647, Default: ""},
		{Name: "test_case_type", Type: field.TypeString, Nullable: true, Default: "DefaultTestCaseType"},
		{Name: "deprecated", Type: field.TypeBool, Nullable: true, Default: false},
	}
	// TestCasesTable holds the schema information for the "test_cases" table.
	TestCasesTable = &schema.Table{
		Name:       "test_cases",
		Columns:    TestCasesColumns,
		PrimaryKey: []*schema.Column{TestCasesColumns[0]},
	}
	// TestPlansColumns holds the columns for the "test_plans" table.
	TestPlansColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "name", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "state", Type: field.TypeString, Nullable: true, Default: "DefaultTestPlanState"},
		{Name: "owner_id", Type: field.TypeUUID, Nullable: true},
		{Name: "responsible_user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "failed_test_case_count", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "passed_test_case_count", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "skipped_test_case_count", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "run_duration", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "deadline", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "test_result", Type: field.TypeString, Nullable: true, Default: "DefaultTestTestState"},
	}
	// TestPlansTable holds the schema information for the "test_plans" table.
	TestPlansTable = &schema.Table{
		Name:       "test_plans",
		Columns:    TestPlansColumns,
		PrimaryKey: []*schema.Column{TestPlansColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DetailsTable,
		ModulesTable,
		PlanRelatedTestCasesTable,
		RelatedTestCasesTable,
		TestCasesTable,
		TestPlansTable,
	}
)

func init() {
}
