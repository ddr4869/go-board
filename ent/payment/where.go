// Code generated by ent, DO NOT EDIT.

package payment

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/go-board/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Payment {
	return predicate.Payment(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Payment {
	return predicate.Payment(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Payment {
	return predicate.Payment(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Payment {
	return predicate.Payment(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Payment {
	return predicate.Payment(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Payment {
	return predicate.Payment(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Payment {
	return predicate.Payment(sql.FieldLTE(FieldID, id))
}

// UserName applies equality check predicate on the "user_name" field. It's identical to UserNameEQ.
func UserName(v string) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldUserName, v))
}

// Grade applies equality check predicate on the "grade" field. It's identical to GradeEQ.
func Grade(v string) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldGrade, v))
}

// MovieID applies equality check predicate on the "movie_id" field. It's identical to MovieIDEQ.
func MovieID(v string) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldMovieID, v))
}

// CreatedDate applies equality check predicate on the "created_date" field. It's identical to CreatedDateEQ.
func CreatedDate(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldCreatedDate, v))
}

// UserNameEQ applies the EQ predicate on the "user_name" field.
func UserNameEQ(v string) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldUserName, v))
}

// UserNameNEQ applies the NEQ predicate on the "user_name" field.
func UserNameNEQ(v string) predicate.Payment {
	return predicate.Payment(sql.FieldNEQ(FieldUserName, v))
}

// UserNameIn applies the In predicate on the "user_name" field.
func UserNameIn(vs ...string) predicate.Payment {
	return predicate.Payment(sql.FieldIn(FieldUserName, vs...))
}

// UserNameNotIn applies the NotIn predicate on the "user_name" field.
func UserNameNotIn(vs ...string) predicate.Payment {
	return predicate.Payment(sql.FieldNotIn(FieldUserName, vs...))
}

// UserNameGT applies the GT predicate on the "user_name" field.
func UserNameGT(v string) predicate.Payment {
	return predicate.Payment(sql.FieldGT(FieldUserName, v))
}

// UserNameGTE applies the GTE predicate on the "user_name" field.
func UserNameGTE(v string) predicate.Payment {
	return predicate.Payment(sql.FieldGTE(FieldUserName, v))
}

// UserNameLT applies the LT predicate on the "user_name" field.
func UserNameLT(v string) predicate.Payment {
	return predicate.Payment(sql.FieldLT(FieldUserName, v))
}

// UserNameLTE applies the LTE predicate on the "user_name" field.
func UserNameLTE(v string) predicate.Payment {
	return predicate.Payment(sql.FieldLTE(FieldUserName, v))
}

// UserNameContains applies the Contains predicate on the "user_name" field.
func UserNameContains(v string) predicate.Payment {
	return predicate.Payment(sql.FieldContains(FieldUserName, v))
}

// UserNameHasPrefix applies the HasPrefix predicate on the "user_name" field.
func UserNameHasPrefix(v string) predicate.Payment {
	return predicate.Payment(sql.FieldHasPrefix(FieldUserName, v))
}

// UserNameHasSuffix applies the HasSuffix predicate on the "user_name" field.
func UserNameHasSuffix(v string) predicate.Payment {
	return predicate.Payment(sql.FieldHasSuffix(FieldUserName, v))
}

// UserNameEqualFold applies the EqualFold predicate on the "user_name" field.
func UserNameEqualFold(v string) predicate.Payment {
	return predicate.Payment(sql.FieldEqualFold(FieldUserName, v))
}

// UserNameContainsFold applies the ContainsFold predicate on the "user_name" field.
func UserNameContainsFold(v string) predicate.Payment {
	return predicate.Payment(sql.FieldContainsFold(FieldUserName, v))
}

// GradeEQ applies the EQ predicate on the "grade" field.
func GradeEQ(v string) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldGrade, v))
}

// GradeNEQ applies the NEQ predicate on the "grade" field.
func GradeNEQ(v string) predicate.Payment {
	return predicate.Payment(sql.FieldNEQ(FieldGrade, v))
}

// GradeIn applies the In predicate on the "grade" field.
func GradeIn(vs ...string) predicate.Payment {
	return predicate.Payment(sql.FieldIn(FieldGrade, vs...))
}

// GradeNotIn applies the NotIn predicate on the "grade" field.
func GradeNotIn(vs ...string) predicate.Payment {
	return predicate.Payment(sql.FieldNotIn(FieldGrade, vs...))
}

// GradeGT applies the GT predicate on the "grade" field.
func GradeGT(v string) predicate.Payment {
	return predicate.Payment(sql.FieldGT(FieldGrade, v))
}

// GradeGTE applies the GTE predicate on the "grade" field.
func GradeGTE(v string) predicate.Payment {
	return predicate.Payment(sql.FieldGTE(FieldGrade, v))
}

// GradeLT applies the LT predicate on the "grade" field.
func GradeLT(v string) predicate.Payment {
	return predicate.Payment(sql.FieldLT(FieldGrade, v))
}

// GradeLTE applies the LTE predicate on the "grade" field.
func GradeLTE(v string) predicate.Payment {
	return predicate.Payment(sql.FieldLTE(FieldGrade, v))
}

// GradeContains applies the Contains predicate on the "grade" field.
func GradeContains(v string) predicate.Payment {
	return predicate.Payment(sql.FieldContains(FieldGrade, v))
}

// GradeHasPrefix applies the HasPrefix predicate on the "grade" field.
func GradeHasPrefix(v string) predicate.Payment {
	return predicate.Payment(sql.FieldHasPrefix(FieldGrade, v))
}

// GradeHasSuffix applies the HasSuffix predicate on the "grade" field.
func GradeHasSuffix(v string) predicate.Payment {
	return predicate.Payment(sql.FieldHasSuffix(FieldGrade, v))
}

// GradeEqualFold applies the EqualFold predicate on the "grade" field.
func GradeEqualFold(v string) predicate.Payment {
	return predicate.Payment(sql.FieldEqualFold(FieldGrade, v))
}

// GradeContainsFold applies the ContainsFold predicate on the "grade" field.
func GradeContainsFold(v string) predicate.Payment {
	return predicate.Payment(sql.FieldContainsFold(FieldGrade, v))
}

// MovieIDEQ applies the EQ predicate on the "movie_id" field.
func MovieIDEQ(v string) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldMovieID, v))
}

// MovieIDNEQ applies the NEQ predicate on the "movie_id" field.
func MovieIDNEQ(v string) predicate.Payment {
	return predicate.Payment(sql.FieldNEQ(FieldMovieID, v))
}

// MovieIDIn applies the In predicate on the "movie_id" field.
func MovieIDIn(vs ...string) predicate.Payment {
	return predicate.Payment(sql.FieldIn(FieldMovieID, vs...))
}

// MovieIDNotIn applies the NotIn predicate on the "movie_id" field.
func MovieIDNotIn(vs ...string) predicate.Payment {
	return predicate.Payment(sql.FieldNotIn(FieldMovieID, vs...))
}

// MovieIDGT applies the GT predicate on the "movie_id" field.
func MovieIDGT(v string) predicate.Payment {
	return predicate.Payment(sql.FieldGT(FieldMovieID, v))
}

// MovieIDGTE applies the GTE predicate on the "movie_id" field.
func MovieIDGTE(v string) predicate.Payment {
	return predicate.Payment(sql.FieldGTE(FieldMovieID, v))
}

// MovieIDLT applies the LT predicate on the "movie_id" field.
func MovieIDLT(v string) predicate.Payment {
	return predicate.Payment(sql.FieldLT(FieldMovieID, v))
}

// MovieIDLTE applies the LTE predicate on the "movie_id" field.
func MovieIDLTE(v string) predicate.Payment {
	return predicate.Payment(sql.FieldLTE(FieldMovieID, v))
}

// MovieIDContains applies the Contains predicate on the "movie_id" field.
func MovieIDContains(v string) predicate.Payment {
	return predicate.Payment(sql.FieldContains(FieldMovieID, v))
}

// MovieIDHasPrefix applies the HasPrefix predicate on the "movie_id" field.
func MovieIDHasPrefix(v string) predicate.Payment {
	return predicate.Payment(sql.FieldHasPrefix(FieldMovieID, v))
}

// MovieIDHasSuffix applies the HasSuffix predicate on the "movie_id" field.
func MovieIDHasSuffix(v string) predicate.Payment {
	return predicate.Payment(sql.FieldHasSuffix(FieldMovieID, v))
}

// MovieIDEqualFold applies the EqualFold predicate on the "movie_id" field.
func MovieIDEqualFold(v string) predicate.Payment {
	return predicate.Payment(sql.FieldEqualFold(FieldMovieID, v))
}

// MovieIDContainsFold applies the ContainsFold predicate on the "movie_id" field.
func MovieIDContainsFold(v string) predicate.Payment {
	return predicate.Payment(sql.FieldContainsFold(FieldMovieID, v))
}

// CreatedDateEQ applies the EQ predicate on the "created_date" field.
func CreatedDateEQ(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldEQ(FieldCreatedDate, v))
}

// CreatedDateNEQ applies the NEQ predicate on the "created_date" field.
func CreatedDateNEQ(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldNEQ(FieldCreatedDate, v))
}

// CreatedDateIn applies the In predicate on the "created_date" field.
func CreatedDateIn(vs ...time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldIn(FieldCreatedDate, vs...))
}

// CreatedDateNotIn applies the NotIn predicate on the "created_date" field.
func CreatedDateNotIn(vs ...time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldNotIn(FieldCreatedDate, vs...))
}

// CreatedDateGT applies the GT predicate on the "created_date" field.
func CreatedDateGT(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldGT(FieldCreatedDate, v))
}

// CreatedDateGTE applies the GTE predicate on the "created_date" field.
func CreatedDateGTE(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldGTE(FieldCreatedDate, v))
}

// CreatedDateLT applies the LT predicate on the "created_date" field.
func CreatedDateLT(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldLT(FieldCreatedDate, v))
}

// CreatedDateLTE applies the LTE predicate on the "created_date" field.
func CreatedDateLTE(v time.Time) predicate.Payment {
	return predicate.Payment(sql.FieldLTE(FieldCreatedDate, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Payment) predicate.Payment {
	return predicate.Payment(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Payment) predicate.Payment {
	return predicate.Payment(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Payment) predicate.Payment {
	return predicate.Payment(sql.NotPredicates(p))
}
