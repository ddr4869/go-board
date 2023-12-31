// Code generated by ent, DO NOT EDIT.

package board

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/go-board/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint) predicate.Board {
	return predicate.Board(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint) predicate.Board {
	return predicate.Board(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint) predicate.Board {
	return predicate.Board(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint) predicate.Board {
	return predicate.Board(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint) predicate.Board {
	return predicate.Board(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint) predicate.Board {
	return predicate.Board(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint) predicate.Board {
	return predicate.Board(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint) predicate.Board {
	return predicate.Board(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint) predicate.Board {
	return predicate.Board(sql.FieldLTE(FieldID, id))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Board {
	return predicate.Board(sql.FieldEQ(FieldTitle, v))
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.Board {
	return predicate.Board(sql.FieldEQ(FieldContent, v))
}

// Writer applies equality check predicate on the "writer" field. It's identical to WriterEQ.
func Writer(v string) predicate.Board {
	return predicate.Board(sql.FieldEQ(FieldWriter, v))
}

// View applies equality check predicate on the "view" field. It's identical to ViewEQ.
func View(v uint) predicate.Board {
	return predicate.Board(sql.FieldEQ(FieldView, v))
}

// Recommend applies equality check predicate on the "recommend" field. It's identical to RecommendEQ.
func Recommend(v uint) predicate.Board {
	return predicate.Board(sql.FieldEQ(FieldRecommend, v))
}

// Hot applies equality check predicate on the "hot" field. It's identical to HotEQ.
func Hot(v bool) predicate.Board {
	return predicate.Board(sql.FieldEQ(FieldHot, v))
}

// CreatedDate applies equality check predicate on the "created_date" field. It's identical to CreatedDateEQ.
func CreatedDate(v time.Time) predicate.Board {
	return predicate.Board(sql.FieldEQ(FieldCreatedDate, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Board {
	return predicate.Board(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Board {
	return predicate.Board(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Board {
	return predicate.Board(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Board {
	return predicate.Board(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Board {
	return predicate.Board(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Board {
	return predicate.Board(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Board {
	return predicate.Board(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Board {
	return predicate.Board(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Board {
	return predicate.Board(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Board {
	return predicate.Board(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Board {
	return predicate.Board(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Board {
	return predicate.Board(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Board {
	return predicate.Board(sql.FieldContainsFold(FieldTitle, v))
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.Board {
	return predicate.Board(sql.FieldEQ(FieldContent, v))
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.Board {
	return predicate.Board(sql.FieldNEQ(FieldContent, v))
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.Board {
	return predicate.Board(sql.FieldIn(FieldContent, vs...))
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.Board {
	return predicate.Board(sql.FieldNotIn(FieldContent, vs...))
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.Board {
	return predicate.Board(sql.FieldGT(FieldContent, v))
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.Board {
	return predicate.Board(sql.FieldGTE(FieldContent, v))
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.Board {
	return predicate.Board(sql.FieldLT(FieldContent, v))
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.Board {
	return predicate.Board(sql.FieldLTE(FieldContent, v))
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.Board {
	return predicate.Board(sql.FieldContains(FieldContent, v))
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.Board {
	return predicate.Board(sql.FieldHasPrefix(FieldContent, v))
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.Board {
	return predicate.Board(sql.FieldHasSuffix(FieldContent, v))
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.Board {
	return predicate.Board(sql.FieldEqualFold(FieldContent, v))
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.Board {
	return predicate.Board(sql.FieldContainsFold(FieldContent, v))
}

// WriterEQ applies the EQ predicate on the "writer" field.
func WriterEQ(v string) predicate.Board {
	return predicate.Board(sql.FieldEQ(FieldWriter, v))
}

// WriterNEQ applies the NEQ predicate on the "writer" field.
func WriterNEQ(v string) predicate.Board {
	return predicate.Board(sql.FieldNEQ(FieldWriter, v))
}

// WriterIn applies the In predicate on the "writer" field.
func WriterIn(vs ...string) predicate.Board {
	return predicate.Board(sql.FieldIn(FieldWriter, vs...))
}

// WriterNotIn applies the NotIn predicate on the "writer" field.
func WriterNotIn(vs ...string) predicate.Board {
	return predicate.Board(sql.FieldNotIn(FieldWriter, vs...))
}

// WriterGT applies the GT predicate on the "writer" field.
func WriterGT(v string) predicate.Board {
	return predicate.Board(sql.FieldGT(FieldWriter, v))
}

// WriterGTE applies the GTE predicate on the "writer" field.
func WriterGTE(v string) predicate.Board {
	return predicate.Board(sql.FieldGTE(FieldWriter, v))
}

// WriterLT applies the LT predicate on the "writer" field.
func WriterLT(v string) predicate.Board {
	return predicate.Board(sql.FieldLT(FieldWriter, v))
}

// WriterLTE applies the LTE predicate on the "writer" field.
func WriterLTE(v string) predicate.Board {
	return predicate.Board(sql.FieldLTE(FieldWriter, v))
}

// WriterContains applies the Contains predicate on the "writer" field.
func WriterContains(v string) predicate.Board {
	return predicate.Board(sql.FieldContains(FieldWriter, v))
}

// WriterHasPrefix applies the HasPrefix predicate on the "writer" field.
func WriterHasPrefix(v string) predicate.Board {
	return predicate.Board(sql.FieldHasPrefix(FieldWriter, v))
}

// WriterHasSuffix applies the HasSuffix predicate on the "writer" field.
func WriterHasSuffix(v string) predicate.Board {
	return predicate.Board(sql.FieldHasSuffix(FieldWriter, v))
}

// WriterEqualFold applies the EqualFold predicate on the "writer" field.
func WriterEqualFold(v string) predicate.Board {
	return predicate.Board(sql.FieldEqualFold(FieldWriter, v))
}

// WriterContainsFold applies the ContainsFold predicate on the "writer" field.
func WriterContainsFold(v string) predicate.Board {
	return predicate.Board(sql.FieldContainsFold(FieldWriter, v))
}

// ViewEQ applies the EQ predicate on the "view" field.
func ViewEQ(v uint) predicate.Board {
	return predicate.Board(sql.FieldEQ(FieldView, v))
}

// ViewNEQ applies the NEQ predicate on the "view" field.
func ViewNEQ(v uint) predicate.Board {
	return predicate.Board(sql.FieldNEQ(FieldView, v))
}

// ViewIn applies the In predicate on the "view" field.
func ViewIn(vs ...uint) predicate.Board {
	return predicate.Board(sql.FieldIn(FieldView, vs...))
}

// ViewNotIn applies the NotIn predicate on the "view" field.
func ViewNotIn(vs ...uint) predicate.Board {
	return predicate.Board(sql.FieldNotIn(FieldView, vs...))
}

// ViewGT applies the GT predicate on the "view" field.
func ViewGT(v uint) predicate.Board {
	return predicate.Board(sql.FieldGT(FieldView, v))
}

// ViewGTE applies the GTE predicate on the "view" field.
func ViewGTE(v uint) predicate.Board {
	return predicate.Board(sql.FieldGTE(FieldView, v))
}

// ViewLT applies the LT predicate on the "view" field.
func ViewLT(v uint) predicate.Board {
	return predicate.Board(sql.FieldLT(FieldView, v))
}

// ViewLTE applies the LTE predicate on the "view" field.
func ViewLTE(v uint) predicate.Board {
	return predicate.Board(sql.FieldLTE(FieldView, v))
}

// RecommendEQ applies the EQ predicate on the "recommend" field.
func RecommendEQ(v uint) predicate.Board {
	return predicate.Board(sql.FieldEQ(FieldRecommend, v))
}

// RecommendNEQ applies the NEQ predicate on the "recommend" field.
func RecommendNEQ(v uint) predicate.Board {
	return predicate.Board(sql.FieldNEQ(FieldRecommend, v))
}

// RecommendIn applies the In predicate on the "recommend" field.
func RecommendIn(vs ...uint) predicate.Board {
	return predicate.Board(sql.FieldIn(FieldRecommend, vs...))
}

// RecommendNotIn applies the NotIn predicate on the "recommend" field.
func RecommendNotIn(vs ...uint) predicate.Board {
	return predicate.Board(sql.FieldNotIn(FieldRecommend, vs...))
}

// RecommendGT applies the GT predicate on the "recommend" field.
func RecommendGT(v uint) predicate.Board {
	return predicate.Board(sql.FieldGT(FieldRecommend, v))
}

// RecommendGTE applies the GTE predicate on the "recommend" field.
func RecommendGTE(v uint) predicate.Board {
	return predicate.Board(sql.FieldGTE(FieldRecommend, v))
}

// RecommendLT applies the LT predicate on the "recommend" field.
func RecommendLT(v uint) predicate.Board {
	return predicate.Board(sql.FieldLT(FieldRecommend, v))
}

// RecommendLTE applies the LTE predicate on the "recommend" field.
func RecommendLTE(v uint) predicate.Board {
	return predicate.Board(sql.FieldLTE(FieldRecommend, v))
}

// HotEQ applies the EQ predicate on the "hot" field.
func HotEQ(v bool) predicate.Board {
	return predicate.Board(sql.FieldEQ(FieldHot, v))
}

// HotNEQ applies the NEQ predicate on the "hot" field.
func HotNEQ(v bool) predicate.Board {
	return predicate.Board(sql.FieldNEQ(FieldHot, v))
}

// CreatedDateEQ applies the EQ predicate on the "created_date" field.
func CreatedDateEQ(v time.Time) predicate.Board {
	return predicate.Board(sql.FieldEQ(FieldCreatedDate, v))
}

// CreatedDateNEQ applies the NEQ predicate on the "created_date" field.
func CreatedDateNEQ(v time.Time) predicate.Board {
	return predicate.Board(sql.FieldNEQ(FieldCreatedDate, v))
}

// CreatedDateIn applies the In predicate on the "created_date" field.
func CreatedDateIn(vs ...time.Time) predicate.Board {
	return predicate.Board(sql.FieldIn(FieldCreatedDate, vs...))
}

// CreatedDateNotIn applies the NotIn predicate on the "created_date" field.
func CreatedDateNotIn(vs ...time.Time) predicate.Board {
	return predicate.Board(sql.FieldNotIn(FieldCreatedDate, vs...))
}

// CreatedDateGT applies the GT predicate on the "created_date" field.
func CreatedDateGT(v time.Time) predicate.Board {
	return predicate.Board(sql.FieldGT(FieldCreatedDate, v))
}

// CreatedDateGTE applies the GTE predicate on the "created_date" field.
func CreatedDateGTE(v time.Time) predicate.Board {
	return predicate.Board(sql.FieldGTE(FieldCreatedDate, v))
}

// CreatedDateLT applies the LT predicate on the "created_date" field.
func CreatedDateLT(v time.Time) predicate.Board {
	return predicate.Board(sql.FieldLT(FieldCreatedDate, v))
}

// CreatedDateLTE applies the LTE predicate on the "created_date" field.
func CreatedDateLTE(v time.Time) predicate.Board {
	return predicate.Board(sql.FieldLTE(FieldCreatedDate, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Board {
	return predicate.Board(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Board {
	return predicate.Board(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Board) predicate.Board {
	return predicate.Board(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Board) predicate.Board {
	return predicate.Board(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Board) predicate.Board {
	return predicate.Board(sql.NotPredicates(p))
}
