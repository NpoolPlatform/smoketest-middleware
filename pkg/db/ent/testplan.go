// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/testplan"
	"github.com/google/uuid"
)

// TestPlan is the model entity for the TestPlan schema.
type TestPlan struct {
	config `json:"-"`
	// ID of the ent.
	ID uint32 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// EntID holds the value of the "ent_id" field.
	EntID uuid.UUID `json:"ent_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// State holds the value of the "state" field.
	State string `json:"state,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy uuid.UUID `json:"created_by,omitempty"`
	// Executor holds the value of the "executor" field.
	Executor uuid.UUID `json:"executor,omitempty"`
	// Fails holds the value of the "fails" field.
	Fails uint32 `json:"fails,omitempty"`
	// Passes holds the value of the "passes" field.
	Passes uint32 `json:"passes,omitempty"`
	// Skips holds the value of the "skips" field.
	Skips uint32 `json:"skips,omitempty"`
	// RunDuration holds the value of the "run_duration" field.
	RunDuration uint32 `json:"run_duration,omitempty"`
	// Deadline holds the value of the "deadline" field.
	Deadline uint32 `json:"deadline,omitempty"`
	// Result holds the value of the "result" field.
	Result string `json:"result,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TestPlan) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case testplan.FieldID, testplan.FieldCreatedAt, testplan.FieldUpdatedAt, testplan.FieldDeletedAt, testplan.FieldFails, testplan.FieldPasses, testplan.FieldSkips, testplan.FieldRunDuration, testplan.FieldDeadline:
			values[i] = new(sql.NullInt64)
		case testplan.FieldName, testplan.FieldState, testplan.FieldResult:
			values[i] = new(sql.NullString)
		case testplan.FieldEntID, testplan.FieldCreatedBy, testplan.FieldExecutor:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type TestPlan", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TestPlan fields.
func (tp *TestPlan) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case testplan.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			tp.ID = uint32(value.Int64)
		case testplan.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				tp.CreatedAt = uint32(value.Int64)
			}
		case testplan.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				tp.UpdatedAt = uint32(value.Int64)
			}
		case testplan.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				tp.DeletedAt = uint32(value.Int64)
			}
		case testplan.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				tp.EntID = *value
			}
		case testplan.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				tp.Name = value.String
			}
		case testplan.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				tp.State = value.String
			}
		case testplan.FieldCreatedBy:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value != nil {
				tp.CreatedBy = *value
			}
		case testplan.FieldExecutor:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field executor", values[i])
			} else if value != nil {
				tp.Executor = *value
			}
		case testplan.FieldFails:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field fails", values[i])
			} else if value.Valid {
				tp.Fails = uint32(value.Int64)
			}
		case testplan.FieldPasses:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field passes", values[i])
			} else if value.Valid {
				tp.Passes = uint32(value.Int64)
			}
		case testplan.FieldSkips:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field skips", values[i])
			} else if value.Valid {
				tp.Skips = uint32(value.Int64)
			}
		case testplan.FieldRunDuration:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field run_duration", values[i])
			} else if value.Valid {
				tp.RunDuration = uint32(value.Int64)
			}
		case testplan.FieldDeadline:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deadline", values[i])
			} else if value.Valid {
				tp.Deadline = uint32(value.Int64)
			}
		case testplan.FieldResult:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field result", values[i])
			} else if value.Valid {
				tp.Result = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this TestPlan.
// Note that you need to call TestPlan.Unwrap() before calling this method if this TestPlan
// was returned from a transaction, and the transaction was committed or rolled back.
func (tp *TestPlan) Update() *TestPlanUpdateOne {
	return (&TestPlanClient{config: tp.config}).UpdateOne(tp)
}

// Unwrap unwraps the TestPlan entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tp *TestPlan) Unwrap() *TestPlan {
	_tx, ok := tp.config.driver.(*txDriver)
	if !ok {
		panic("ent: TestPlan is not a transactional entity")
	}
	tp.config.driver = _tx.drv
	return tp
}

// String implements the fmt.Stringer.
func (tp *TestPlan) String() string {
	var builder strings.Builder
	builder.WriteString("TestPlan(")
	builder.WriteString(fmt.Sprintf("id=%v, ", tp.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", tp.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", tp.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", tp.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", tp.EntID))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(tp.Name)
	builder.WriteString(", ")
	builder.WriteString("state=")
	builder.WriteString(tp.State)
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", tp.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("executor=")
	builder.WriteString(fmt.Sprintf("%v", tp.Executor))
	builder.WriteString(", ")
	builder.WriteString("fails=")
	builder.WriteString(fmt.Sprintf("%v", tp.Fails))
	builder.WriteString(", ")
	builder.WriteString("passes=")
	builder.WriteString(fmt.Sprintf("%v", tp.Passes))
	builder.WriteString(", ")
	builder.WriteString("skips=")
	builder.WriteString(fmt.Sprintf("%v", tp.Skips))
	builder.WriteString(", ")
	builder.WriteString("run_duration=")
	builder.WriteString(fmt.Sprintf("%v", tp.RunDuration))
	builder.WriteString(", ")
	builder.WriteString("deadline=")
	builder.WriteString(fmt.Sprintf("%v", tp.Deadline))
	builder.WriteString(", ")
	builder.WriteString("result=")
	builder.WriteString(tp.Result)
	builder.WriteByte(')')
	return builder.String()
}

// TestPlans is a parsable slice of TestPlan.
type TestPlans []*TestPlan

func (tp TestPlans) config(cfg config) {
	for _i := range tp {
		tp[_i].config = cfg
	}
}
