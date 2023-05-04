// Code generated by ent, DO NOT EDIT.

package plantestcase

import (
	"entgo.io/ent"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the plantestcase type in the database.
	Label = "plan_test_case"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldTestPlanID holds the string denoting the test_plan_id field in the database.
	FieldTestPlanID = "test_plan_id"
	// FieldTestCaseID holds the string denoting the test_case_id field in the database.
	FieldTestCaseID = "test_case_id"
	// FieldInput holds the string denoting the input field in the database.
	FieldInput = "input"
	// FieldOutput holds the string denoting the output field in the database.
	FieldOutput = "output"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldTestUserID holds the string denoting the test_user_id field in the database.
	FieldTestUserID = "test_user_id"
	// FieldRunDuration holds the string denoting the run_duration field in the database.
	FieldRunDuration = "run_duration"
	// FieldResult holds the string denoting the result field in the database.
	FieldResult = "result"
	// FieldIndex holds the string denoting the index field in the database.
	FieldIndex = "index"
	// Table holds the table name of the plantestcase in the database.
	Table = "plan_test_cases"
)

// Columns holds all SQL columns for plantestcase fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldTestPlanID,
	FieldTestCaseID,
	FieldInput,
	FieldOutput,
	FieldDescription,
	FieldTestUserID,
	FieldRunDuration,
	FieldResult,
	FieldIndex,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/runtime"
//
var (
	Hooks  [1]ent.Hook
	Policy ent.Policy
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() uint32
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() uint32
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() uint32
	// DefaultDeletedAt holds the default value on creation for the "deleted_at" field.
	DefaultDeletedAt func() uint32
	// DefaultTestPlanID holds the default value on creation for the "test_plan_id" field.
	DefaultTestPlanID func() uuid.UUID
	// DefaultTestCaseID holds the default value on creation for the "test_case_id" field.
	DefaultTestCaseID func() uuid.UUID
	// DefaultInput holds the default value on creation for the "input" field.
	DefaultInput string
	// DefaultOutput holds the default value on creation for the "output" field.
	DefaultOutput string
	// DefaultDescription holds the default value on creation for the "description" field.
	DefaultDescription string
	// DefaultTestUserID holds the default value on creation for the "test_user_id" field.
	DefaultTestUserID func() uuid.UUID
	// DefaultRunDuration holds the default value on creation for the "run_duration" field.
	DefaultRunDuration uint32
	// DefaultResult holds the default value on creation for the "result" field.
	DefaultResult string
	// DefaultIndex holds the default value on creation for the "index" field.
	DefaultIndex uint32
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
