// Code generated by ent, DO NOT EDIT.

package plantestcase

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// TestPlanID applies equality check predicate on the "test_plan_id" field. It's identical to TestPlanIDEQ.
func TestPlanID(v uuid.UUID) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTestPlanID), v))
	})
}

// TestCaseID applies equality check predicate on the "test_case_id" field. It's identical to TestCaseIDEQ.
func TestCaseID(v uuid.UUID) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTestCaseID), v))
	})
}

// Input applies equality check predicate on the "input" field. It's identical to InputEQ.
func Input(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInput), v))
	})
}

// Output applies equality check predicate on the "output" field. It's identical to OutputEQ.
func Output(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutput), v))
	})
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// TestUserID applies equality check predicate on the "test_user_id" field. It's identical to TestUserIDEQ.
func TestUserID(v uuid.UUID) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTestUserID), v))
	})
}

// RunDuration applies equality check predicate on the "run_duration" field. It's identical to RunDurationEQ.
func RunDuration(v uint32) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRunDuration), v))
	})
}

// Result applies equality check predicate on the "result" field. It's identical to ResultEQ.
func Result(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldResult), v))
	})
}

// Index applies equality check predicate on the "index" field. It's identical to IndexEQ.
func Index(v uint32) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIndex), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.PlanTestCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.PlanTestCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.PlanTestCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.PlanTestCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.PlanTestCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.PlanTestCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// TestPlanIDEQ applies the EQ predicate on the "test_plan_id" field.
func TestPlanIDEQ(v uuid.UUID) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTestPlanID), v))
	})
}

// TestPlanIDNEQ applies the NEQ predicate on the "test_plan_id" field.
func TestPlanIDNEQ(v uuid.UUID) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTestPlanID), v))
	})
}

// TestPlanIDIn applies the In predicate on the "test_plan_id" field.
func TestPlanIDIn(vs ...uuid.UUID) predicate.PlanTestCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTestPlanID), v...))
	})
}

// TestPlanIDNotIn applies the NotIn predicate on the "test_plan_id" field.
func TestPlanIDNotIn(vs ...uuid.UUID) predicate.PlanTestCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTestPlanID), v...))
	})
}

// TestPlanIDGT applies the GT predicate on the "test_plan_id" field.
func TestPlanIDGT(v uuid.UUID) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTestPlanID), v))
	})
}

// TestPlanIDGTE applies the GTE predicate on the "test_plan_id" field.
func TestPlanIDGTE(v uuid.UUID) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTestPlanID), v))
	})
}

// TestPlanIDLT applies the LT predicate on the "test_plan_id" field.
func TestPlanIDLT(v uuid.UUID) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTestPlanID), v))
	})
}

// TestPlanIDLTE applies the LTE predicate on the "test_plan_id" field.
func TestPlanIDLTE(v uuid.UUID) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTestPlanID), v))
	})
}

// TestPlanIDIsNil applies the IsNil predicate on the "test_plan_id" field.
func TestPlanIDIsNil() predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTestPlanID)))
	})
}

// TestPlanIDNotNil applies the NotNil predicate on the "test_plan_id" field.
func TestPlanIDNotNil() predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTestPlanID)))
	})
}

// TestCaseIDEQ applies the EQ predicate on the "test_case_id" field.
func TestCaseIDEQ(v uuid.UUID) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTestCaseID), v))
	})
}

// TestCaseIDNEQ applies the NEQ predicate on the "test_case_id" field.
func TestCaseIDNEQ(v uuid.UUID) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTestCaseID), v))
	})
}

// TestCaseIDIn applies the In predicate on the "test_case_id" field.
func TestCaseIDIn(vs ...uuid.UUID) predicate.PlanTestCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTestCaseID), v...))
	})
}

// TestCaseIDNotIn applies the NotIn predicate on the "test_case_id" field.
func TestCaseIDNotIn(vs ...uuid.UUID) predicate.PlanTestCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTestCaseID), v...))
	})
}

// TestCaseIDGT applies the GT predicate on the "test_case_id" field.
func TestCaseIDGT(v uuid.UUID) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTestCaseID), v))
	})
}

// TestCaseIDGTE applies the GTE predicate on the "test_case_id" field.
func TestCaseIDGTE(v uuid.UUID) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTestCaseID), v))
	})
}

// TestCaseIDLT applies the LT predicate on the "test_case_id" field.
func TestCaseIDLT(v uuid.UUID) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTestCaseID), v))
	})
}

// TestCaseIDLTE applies the LTE predicate on the "test_case_id" field.
func TestCaseIDLTE(v uuid.UUID) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTestCaseID), v))
	})
}

// TestCaseIDIsNil applies the IsNil predicate on the "test_case_id" field.
func TestCaseIDIsNil() predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTestCaseID)))
	})
}

// TestCaseIDNotNil applies the NotNil predicate on the "test_case_id" field.
func TestCaseIDNotNil() predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTestCaseID)))
	})
}

// InputEQ applies the EQ predicate on the "input" field.
func InputEQ(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInput), v))
	})
}

// InputNEQ applies the NEQ predicate on the "input" field.
func InputNEQ(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInput), v))
	})
}

// InputIn applies the In predicate on the "input" field.
func InputIn(vs ...string) predicate.PlanTestCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldInput), v...))
	})
}

// InputNotIn applies the NotIn predicate on the "input" field.
func InputNotIn(vs ...string) predicate.PlanTestCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldInput), v...))
	})
}

// InputGT applies the GT predicate on the "input" field.
func InputGT(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldInput), v))
	})
}

// InputGTE applies the GTE predicate on the "input" field.
func InputGTE(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldInput), v))
	})
}

// InputLT applies the LT predicate on the "input" field.
func InputLT(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldInput), v))
	})
}

// InputLTE applies the LTE predicate on the "input" field.
func InputLTE(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldInput), v))
	})
}

// InputContains applies the Contains predicate on the "input" field.
func InputContains(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldInput), v))
	})
}

// InputHasPrefix applies the HasPrefix predicate on the "input" field.
func InputHasPrefix(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldInput), v))
	})
}

// InputHasSuffix applies the HasSuffix predicate on the "input" field.
func InputHasSuffix(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldInput), v))
	})
}

// InputIsNil applies the IsNil predicate on the "input" field.
func InputIsNil() predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldInput)))
	})
}

// InputNotNil applies the NotNil predicate on the "input" field.
func InputNotNil() predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldInput)))
	})
}

// InputEqualFold applies the EqualFold predicate on the "input" field.
func InputEqualFold(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldInput), v))
	})
}

// InputContainsFold applies the ContainsFold predicate on the "input" field.
func InputContainsFold(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldInput), v))
	})
}

// OutputEQ applies the EQ predicate on the "output" field.
func OutputEQ(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutput), v))
	})
}

// OutputNEQ applies the NEQ predicate on the "output" field.
func OutputNEQ(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutput), v))
	})
}

// OutputIn applies the In predicate on the "output" field.
func OutputIn(vs ...string) predicate.PlanTestCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldOutput), v...))
	})
}

// OutputNotIn applies the NotIn predicate on the "output" field.
func OutputNotIn(vs ...string) predicate.PlanTestCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldOutput), v...))
	})
}

// OutputGT applies the GT predicate on the "output" field.
func OutputGT(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutput), v))
	})
}

// OutputGTE applies the GTE predicate on the "output" field.
func OutputGTE(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutput), v))
	})
}

// OutputLT applies the LT predicate on the "output" field.
func OutputLT(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutput), v))
	})
}

// OutputLTE applies the LTE predicate on the "output" field.
func OutputLTE(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutput), v))
	})
}

// OutputContains applies the Contains predicate on the "output" field.
func OutputContains(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutput), v))
	})
}

// OutputHasPrefix applies the HasPrefix predicate on the "output" field.
func OutputHasPrefix(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutput), v))
	})
}

// OutputHasSuffix applies the HasSuffix predicate on the "output" field.
func OutputHasSuffix(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutput), v))
	})
}

// OutputIsNil applies the IsNil predicate on the "output" field.
func OutputIsNil() predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldOutput)))
	})
}

// OutputNotNil applies the NotNil predicate on the "output" field.
func OutputNotNil() predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldOutput)))
	})
}

// OutputEqualFold applies the EqualFold predicate on the "output" field.
func OutputEqualFold(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutput), v))
	})
}

// OutputContainsFold applies the ContainsFold predicate on the "output" field.
func OutputContainsFold(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutput), v))
	})
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescription), v))
	})
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.PlanTestCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDescription), v...))
	})
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.PlanTestCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDescription), v...))
	})
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescription), v))
	})
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescription), v))
	})
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescription), v))
	})
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescription), v))
	})
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescription), v))
	})
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescription), v))
	})
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescription), v))
	})
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDescription)))
	})
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDescription)))
	})
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescription), v))
	})
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescription), v))
	})
}

// TestUserIDEQ applies the EQ predicate on the "test_user_id" field.
func TestUserIDEQ(v uuid.UUID) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTestUserID), v))
	})
}

// TestUserIDNEQ applies the NEQ predicate on the "test_user_id" field.
func TestUserIDNEQ(v uuid.UUID) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTestUserID), v))
	})
}

// TestUserIDIn applies the In predicate on the "test_user_id" field.
func TestUserIDIn(vs ...uuid.UUID) predicate.PlanTestCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTestUserID), v...))
	})
}

// TestUserIDNotIn applies the NotIn predicate on the "test_user_id" field.
func TestUserIDNotIn(vs ...uuid.UUID) predicate.PlanTestCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTestUserID), v...))
	})
}

// TestUserIDGT applies the GT predicate on the "test_user_id" field.
func TestUserIDGT(v uuid.UUID) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTestUserID), v))
	})
}

// TestUserIDGTE applies the GTE predicate on the "test_user_id" field.
func TestUserIDGTE(v uuid.UUID) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTestUserID), v))
	})
}

// TestUserIDLT applies the LT predicate on the "test_user_id" field.
func TestUserIDLT(v uuid.UUID) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTestUserID), v))
	})
}

// TestUserIDLTE applies the LTE predicate on the "test_user_id" field.
func TestUserIDLTE(v uuid.UUID) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTestUserID), v))
	})
}

// TestUserIDIsNil applies the IsNil predicate on the "test_user_id" field.
func TestUserIDIsNil() predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTestUserID)))
	})
}

// TestUserIDNotNil applies the NotNil predicate on the "test_user_id" field.
func TestUserIDNotNil() predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTestUserID)))
	})
}

// RunDurationEQ applies the EQ predicate on the "run_duration" field.
func RunDurationEQ(v uint32) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRunDuration), v))
	})
}

// RunDurationNEQ applies the NEQ predicate on the "run_duration" field.
func RunDurationNEQ(v uint32) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRunDuration), v))
	})
}

// RunDurationIn applies the In predicate on the "run_duration" field.
func RunDurationIn(vs ...uint32) predicate.PlanTestCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRunDuration), v...))
	})
}

// RunDurationNotIn applies the NotIn predicate on the "run_duration" field.
func RunDurationNotIn(vs ...uint32) predicate.PlanTestCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRunDuration), v...))
	})
}

// RunDurationGT applies the GT predicate on the "run_duration" field.
func RunDurationGT(v uint32) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRunDuration), v))
	})
}

// RunDurationGTE applies the GTE predicate on the "run_duration" field.
func RunDurationGTE(v uint32) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRunDuration), v))
	})
}

// RunDurationLT applies the LT predicate on the "run_duration" field.
func RunDurationLT(v uint32) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRunDuration), v))
	})
}

// RunDurationLTE applies the LTE predicate on the "run_duration" field.
func RunDurationLTE(v uint32) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRunDuration), v))
	})
}

// RunDurationIsNil applies the IsNil predicate on the "run_duration" field.
func RunDurationIsNil() predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldRunDuration)))
	})
}

// RunDurationNotNil applies the NotNil predicate on the "run_duration" field.
func RunDurationNotNil() predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldRunDuration)))
	})
}

// ResultEQ applies the EQ predicate on the "result" field.
func ResultEQ(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldResult), v))
	})
}

// ResultNEQ applies the NEQ predicate on the "result" field.
func ResultNEQ(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldResult), v))
	})
}

// ResultIn applies the In predicate on the "result" field.
func ResultIn(vs ...string) predicate.PlanTestCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldResult), v...))
	})
}

// ResultNotIn applies the NotIn predicate on the "result" field.
func ResultNotIn(vs ...string) predicate.PlanTestCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldResult), v...))
	})
}

// ResultGT applies the GT predicate on the "result" field.
func ResultGT(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldResult), v))
	})
}

// ResultGTE applies the GTE predicate on the "result" field.
func ResultGTE(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldResult), v))
	})
}

// ResultLT applies the LT predicate on the "result" field.
func ResultLT(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldResult), v))
	})
}

// ResultLTE applies the LTE predicate on the "result" field.
func ResultLTE(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldResult), v))
	})
}

// ResultContains applies the Contains predicate on the "result" field.
func ResultContains(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldResult), v))
	})
}

// ResultHasPrefix applies the HasPrefix predicate on the "result" field.
func ResultHasPrefix(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldResult), v))
	})
}

// ResultHasSuffix applies the HasSuffix predicate on the "result" field.
func ResultHasSuffix(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldResult), v))
	})
}

// ResultIsNil applies the IsNil predicate on the "result" field.
func ResultIsNil() predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldResult)))
	})
}

// ResultNotNil applies the NotNil predicate on the "result" field.
func ResultNotNil() predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldResult)))
	})
}

// ResultEqualFold applies the EqualFold predicate on the "result" field.
func ResultEqualFold(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldResult), v))
	})
}

// ResultContainsFold applies the ContainsFold predicate on the "result" field.
func ResultContainsFold(v string) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldResult), v))
	})
}

// IndexEQ applies the EQ predicate on the "index" field.
func IndexEQ(v uint32) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIndex), v))
	})
}

// IndexNEQ applies the NEQ predicate on the "index" field.
func IndexNEQ(v uint32) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIndex), v))
	})
}

// IndexIn applies the In predicate on the "index" field.
func IndexIn(vs ...uint32) predicate.PlanTestCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldIndex), v...))
	})
}

// IndexNotIn applies the NotIn predicate on the "index" field.
func IndexNotIn(vs ...uint32) predicate.PlanTestCase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldIndex), v...))
	})
}

// IndexGT applies the GT predicate on the "index" field.
func IndexGT(v uint32) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIndex), v))
	})
}

// IndexGTE applies the GTE predicate on the "index" field.
func IndexGTE(v uint32) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIndex), v))
	})
}

// IndexLT applies the LT predicate on the "index" field.
func IndexLT(v uint32) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIndex), v))
	})
}

// IndexLTE applies the LTE predicate on the "index" field.
func IndexLTE(v uint32) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIndex), v))
	})
}

// IndexIsNil applies the IsNil predicate on the "index" field.
func IndexIsNil() predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldIndex)))
	})
}

// IndexNotNil applies the NotNil predicate on the "index" field.
func IndexNotNil() predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldIndex)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PlanTestCase) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PlanTestCase) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PlanTestCase) predicate.PlanTestCase {
	return predicate.PlanTestCase(func(s *sql.Selector) {
		p(s.Not())
	})
}
