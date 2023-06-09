// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CondsColumns holds the columns for the "conds" table.
	CondsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "cond_type", Type: field.TypeString, Nullable: true, Default: "DefaultCondType"},
		{Name: "test_case_id", Type: field.TypeUUID, Nullable: true},
		{Name: "cond_test_case_id", Type: field.TypeUUID, Nullable: true},
		{Name: "argument_map", Type: field.TypeString, Nullable: true, Size: 2147483647, Default: ""},
		{Name: "index", Type: field.TypeUint32, Nullable: true, Default: 0},
	}
	// CondsTable holds the schema information for the "conds" table.
	CondsTable = &schema.Table{
		Name:       "conds",
		Columns:    CondsColumns,
		PrimaryKey: []*schema.Column{CondsColumns[0]},
	}
	// ModulesColumns holds the columns for the "modules" table.
	ModulesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "name", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "description", Type: field.TypeString, Nullable: true, Default: ""},
	}
	// ModulesTable holds the schema information for the "modules" table.
	ModulesTable = &schema.Table{
		Name:       "modules",
		Columns:    ModulesColumns,
		PrimaryKey: []*schema.Column{ModulesColumns[0]},
	}
	// PlanTestCasesColumns holds the columns for the "plan_test_cases" table.
	PlanTestCasesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "test_plan_id", Type: field.TypeUUID, Nullable: true},
		{Name: "test_case_id", Type: field.TypeUUID, Nullable: true},
		{Name: "input", Type: field.TypeString, Nullable: true, Size: 2147483647, Default: ""},
		{Name: "output", Type: field.TypeString, Nullable: true, Size: 2147483647, Default: ""},
		{Name: "description", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "test_user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "run_duration", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "result", Type: field.TypeString, Nullable: true, Default: "DefaultTestCaseResult"},
		{Name: "index", Type: field.TypeUint32, Nullable: true, Default: 0},
	}
	// PlanTestCasesTable holds the schema information for the "plan_test_cases" table.
	PlanTestCasesTable = &schema.Table{
		Name:       "plan_test_cases",
		Columns:    PlanTestCasesColumns,
		PrimaryKey: []*schema.Column{PlanTestCasesColumns[0]},
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
		{Name: "input", Type: field.TypeString, Nullable: true, Size: 2147483647, Default: ""},
		{Name: "input_desc", Type: field.TypeString, Nullable: true, Size: 2147483647, Default: ""},
		{Name: "expectation", Type: field.TypeString, Nullable: true, Size: 2147483647, Default: ""},
		{Name: "output_desc", Type: field.TypeString, Nullable: true, Size: 2147483647, Default: ""},
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
		{Name: "state", Type: field.TypeString, Nullable: true, Default: "WaitStart"},
		{Name: "created_by", Type: field.TypeUUID, Nullable: true},
		{Name: "executor", Type: field.TypeUUID, Nullable: true},
		{Name: "fails", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "passes", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "skips", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "run_duration", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "deadline", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "result", Type: field.TypeString, Nullable: true, Default: "DefaultTestResultState"},
	}
	// TestPlansTable holds the schema information for the "test_plans" table.
	TestPlansTable = &schema.Table{
		Name:       "test_plans",
		Columns:    TestPlansColumns,
		PrimaryKey: []*schema.Column{TestPlansColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CondsTable,
		ModulesTable,
		PlanTestCasesTable,
		TestCasesTable,
		TestPlansTable,
	}
)

func init() {
}
