// Code generated by ent, DO NOT EDIT.

package review

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/m3-app/backend/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Review {
	return predicate.Review(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Review {
	return predicate.Review(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Review {
	return predicate.Review(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Review {
	return predicate.Review(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Review {
	return predicate.Review(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Review {
	return predicate.Review(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Review {
	return predicate.Review(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Review {
	return predicate.Review(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Review {
	return predicate.Review(sql.FieldLTE(FieldID, id))
}

// ReviewID applies equality check predicate on the "review_id" field. It's identical to ReviewIDEQ.
func ReviewID(v int64) predicate.Review {
	return predicate.Review(sql.FieldEQ(FieldReviewID, v))
}

// Rating applies equality check predicate on the "rating" field. It's identical to RatingEQ.
func Rating(v float64) predicate.Review {
	return predicate.Review(sql.FieldEQ(FieldRating, v))
}

// Message applies equality check predicate on the "message" field. It's identical to MessageEQ.
func Message(v string) predicate.Review {
	return predicate.Review(sql.FieldEQ(FieldMessage, v))
}

// ReviewIDEQ applies the EQ predicate on the "review_id" field.
func ReviewIDEQ(v int64) predicate.Review {
	return predicate.Review(sql.FieldEQ(FieldReviewID, v))
}

// ReviewIDNEQ applies the NEQ predicate on the "review_id" field.
func ReviewIDNEQ(v int64) predicate.Review {
	return predicate.Review(sql.FieldNEQ(FieldReviewID, v))
}

// ReviewIDIn applies the In predicate on the "review_id" field.
func ReviewIDIn(vs ...int64) predicate.Review {
	return predicate.Review(sql.FieldIn(FieldReviewID, vs...))
}

// ReviewIDNotIn applies the NotIn predicate on the "review_id" field.
func ReviewIDNotIn(vs ...int64) predicate.Review {
	return predicate.Review(sql.FieldNotIn(FieldReviewID, vs...))
}

// ReviewIDGT applies the GT predicate on the "review_id" field.
func ReviewIDGT(v int64) predicate.Review {
	return predicate.Review(sql.FieldGT(FieldReviewID, v))
}

// ReviewIDGTE applies the GTE predicate on the "review_id" field.
func ReviewIDGTE(v int64) predicate.Review {
	return predicate.Review(sql.FieldGTE(FieldReviewID, v))
}

// ReviewIDLT applies the LT predicate on the "review_id" field.
func ReviewIDLT(v int64) predicate.Review {
	return predicate.Review(sql.FieldLT(FieldReviewID, v))
}

// ReviewIDLTE applies the LTE predicate on the "review_id" field.
func ReviewIDLTE(v int64) predicate.Review {
	return predicate.Review(sql.FieldLTE(FieldReviewID, v))
}

// RatingEQ applies the EQ predicate on the "rating" field.
func RatingEQ(v float64) predicate.Review {
	return predicate.Review(sql.FieldEQ(FieldRating, v))
}

// RatingNEQ applies the NEQ predicate on the "rating" field.
func RatingNEQ(v float64) predicate.Review {
	return predicate.Review(sql.FieldNEQ(FieldRating, v))
}

// RatingIn applies the In predicate on the "rating" field.
func RatingIn(vs ...float64) predicate.Review {
	return predicate.Review(sql.FieldIn(FieldRating, vs...))
}

// RatingNotIn applies the NotIn predicate on the "rating" field.
func RatingNotIn(vs ...float64) predicate.Review {
	return predicate.Review(sql.FieldNotIn(FieldRating, vs...))
}

// RatingGT applies the GT predicate on the "rating" field.
func RatingGT(v float64) predicate.Review {
	return predicate.Review(sql.FieldGT(FieldRating, v))
}

// RatingGTE applies the GTE predicate on the "rating" field.
func RatingGTE(v float64) predicate.Review {
	return predicate.Review(sql.FieldGTE(FieldRating, v))
}

// RatingLT applies the LT predicate on the "rating" field.
func RatingLT(v float64) predicate.Review {
	return predicate.Review(sql.FieldLT(FieldRating, v))
}

// RatingLTE applies the LTE predicate on the "rating" field.
func RatingLTE(v float64) predicate.Review {
	return predicate.Review(sql.FieldLTE(FieldRating, v))
}

// MessageEQ applies the EQ predicate on the "message" field.
func MessageEQ(v string) predicate.Review {
	return predicate.Review(sql.FieldEQ(FieldMessage, v))
}

// MessageNEQ applies the NEQ predicate on the "message" field.
func MessageNEQ(v string) predicate.Review {
	return predicate.Review(sql.FieldNEQ(FieldMessage, v))
}

// MessageIn applies the In predicate on the "message" field.
func MessageIn(vs ...string) predicate.Review {
	return predicate.Review(sql.FieldIn(FieldMessage, vs...))
}

// MessageNotIn applies the NotIn predicate on the "message" field.
func MessageNotIn(vs ...string) predicate.Review {
	return predicate.Review(sql.FieldNotIn(FieldMessage, vs...))
}

// MessageGT applies the GT predicate on the "message" field.
func MessageGT(v string) predicate.Review {
	return predicate.Review(sql.FieldGT(FieldMessage, v))
}

// MessageGTE applies the GTE predicate on the "message" field.
func MessageGTE(v string) predicate.Review {
	return predicate.Review(sql.FieldGTE(FieldMessage, v))
}

// MessageLT applies the LT predicate on the "message" field.
func MessageLT(v string) predicate.Review {
	return predicate.Review(sql.FieldLT(FieldMessage, v))
}

// MessageLTE applies the LTE predicate on the "message" field.
func MessageLTE(v string) predicate.Review {
	return predicate.Review(sql.FieldLTE(FieldMessage, v))
}

// MessageContains applies the Contains predicate on the "message" field.
func MessageContains(v string) predicate.Review {
	return predicate.Review(sql.FieldContains(FieldMessage, v))
}

// MessageHasPrefix applies the HasPrefix predicate on the "message" field.
func MessageHasPrefix(v string) predicate.Review {
	return predicate.Review(sql.FieldHasPrefix(FieldMessage, v))
}

// MessageHasSuffix applies the HasSuffix predicate on the "message" field.
func MessageHasSuffix(v string) predicate.Review {
	return predicate.Review(sql.FieldHasSuffix(FieldMessage, v))
}

// MessageEqualFold applies the EqualFold predicate on the "message" field.
func MessageEqualFold(v string) predicate.Review {
	return predicate.Review(sql.FieldEqualFold(FieldMessage, v))
}

// MessageContainsFold applies the ContainsFold predicate on the "message" field.
func MessageContainsFold(v string) predicate.Review {
	return predicate.Review(sql.FieldContainsFold(FieldMessage, v))
}

// HasLocation applies the HasEdge predicate on the "location" edge.
func HasLocation() predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, LocationTable, LocationPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLocationWith applies the HasEdge predicate on the "location" edge with a given conditions (other predicates).
func HasLocationWith(preds ...predicate.Location) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(LocationInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, LocationTable, LocationPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Review) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Review) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
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
func Not(p predicate.Review) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		p(s.Not())
	})
}
