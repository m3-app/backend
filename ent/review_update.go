// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/m3-app/backend/ent/location"
	"github.com/m3-app/backend/ent/predicate"
	"github.com/m3-app/backend/ent/review"
)

// ReviewUpdate is the builder for updating Review entities.
type ReviewUpdate struct {
	config
	hooks    []Hook
	mutation *ReviewMutation
}

// Where appends a list predicates to the ReviewUpdate builder.
func (ru *ReviewUpdate) Where(ps ...predicate.Review) *ReviewUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetReviewID sets the "review_id" field.
func (ru *ReviewUpdate) SetReviewID(i int64) *ReviewUpdate {
	ru.mutation.ResetReviewID()
	ru.mutation.SetReviewID(i)
	return ru
}

// AddReviewID adds i to the "review_id" field.
func (ru *ReviewUpdate) AddReviewID(i int64) *ReviewUpdate {
	ru.mutation.AddReviewID(i)
	return ru
}

// SetRating sets the "rating" field.
func (ru *ReviewUpdate) SetRating(f float64) *ReviewUpdate {
	ru.mutation.ResetRating()
	ru.mutation.SetRating(f)
	return ru
}

// AddRating adds f to the "rating" field.
func (ru *ReviewUpdate) AddRating(f float64) *ReviewUpdate {
	ru.mutation.AddRating(f)
	return ru
}

// SetMessage sets the "message" field.
func (ru *ReviewUpdate) SetMessage(s string) *ReviewUpdate {
	ru.mutation.SetMessage(s)
	return ru
}

// AddLocationIDs adds the "location" edge to the Location entity by IDs.
func (ru *ReviewUpdate) AddLocationIDs(ids ...int) *ReviewUpdate {
	ru.mutation.AddLocationIDs(ids...)
	return ru
}

// AddLocation adds the "location" edges to the Location entity.
func (ru *ReviewUpdate) AddLocation(l ...*Location) *ReviewUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ru.AddLocationIDs(ids...)
}

// Mutation returns the ReviewMutation object of the builder.
func (ru *ReviewUpdate) Mutation() *ReviewMutation {
	return ru.mutation
}

// ClearLocation clears all "location" edges to the Location entity.
func (ru *ReviewUpdate) ClearLocation() *ReviewUpdate {
	ru.mutation.ClearLocation()
	return ru
}

// RemoveLocationIDs removes the "location" edge to Location entities by IDs.
func (ru *ReviewUpdate) RemoveLocationIDs(ids ...int) *ReviewUpdate {
	ru.mutation.RemoveLocationIDs(ids...)
	return ru
}

// RemoveLocation removes "location" edges to Location entities.
func (ru *ReviewUpdate) RemoveLocation(l ...*Location) *ReviewUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ru.RemoveLocationIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *ReviewUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, ReviewMutation](ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *ReviewUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *ReviewUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *ReviewUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ru *ReviewUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(review.Table, review.Columns, sqlgraph.NewFieldSpec(review.FieldID, field.TypeInt))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.ReviewID(); ok {
		_spec.SetField(review.FieldReviewID, field.TypeInt64, value)
	}
	if value, ok := ru.mutation.AddedReviewID(); ok {
		_spec.AddField(review.FieldReviewID, field.TypeInt64, value)
	}
	if value, ok := ru.mutation.Rating(); ok {
		_spec.SetField(review.FieldRating, field.TypeFloat64, value)
	}
	if value, ok := ru.mutation.AddedRating(); ok {
		_spec.AddField(review.FieldRating, field.TypeFloat64, value)
	}
	if value, ok := ru.mutation.Message(); ok {
		_spec.SetField(review.FieldMessage, field.TypeString, value)
	}
	if ru.mutation.LocationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   review.LocationTable,
			Columns: review.LocationPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: location.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedLocationIDs(); len(nodes) > 0 && !ru.mutation.LocationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   review.LocationTable,
			Columns: review.LocationPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: location.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.LocationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   review.LocationTable,
			Columns: review.LocationPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: location.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{review.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// ReviewUpdateOne is the builder for updating a single Review entity.
type ReviewUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ReviewMutation
}

// SetReviewID sets the "review_id" field.
func (ruo *ReviewUpdateOne) SetReviewID(i int64) *ReviewUpdateOne {
	ruo.mutation.ResetReviewID()
	ruo.mutation.SetReviewID(i)
	return ruo
}

// AddReviewID adds i to the "review_id" field.
func (ruo *ReviewUpdateOne) AddReviewID(i int64) *ReviewUpdateOne {
	ruo.mutation.AddReviewID(i)
	return ruo
}

// SetRating sets the "rating" field.
func (ruo *ReviewUpdateOne) SetRating(f float64) *ReviewUpdateOne {
	ruo.mutation.ResetRating()
	ruo.mutation.SetRating(f)
	return ruo
}

// AddRating adds f to the "rating" field.
func (ruo *ReviewUpdateOne) AddRating(f float64) *ReviewUpdateOne {
	ruo.mutation.AddRating(f)
	return ruo
}

// SetMessage sets the "message" field.
func (ruo *ReviewUpdateOne) SetMessage(s string) *ReviewUpdateOne {
	ruo.mutation.SetMessage(s)
	return ruo
}

// AddLocationIDs adds the "location" edge to the Location entity by IDs.
func (ruo *ReviewUpdateOne) AddLocationIDs(ids ...int) *ReviewUpdateOne {
	ruo.mutation.AddLocationIDs(ids...)
	return ruo
}

// AddLocation adds the "location" edges to the Location entity.
func (ruo *ReviewUpdateOne) AddLocation(l ...*Location) *ReviewUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ruo.AddLocationIDs(ids...)
}

// Mutation returns the ReviewMutation object of the builder.
func (ruo *ReviewUpdateOne) Mutation() *ReviewMutation {
	return ruo.mutation
}

// ClearLocation clears all "location" edges to the Location entity.
func (ruo *ReviewUpdateOne) ClearLocation() *ReviewUpdateOne {
	ruo.mutation.ClearLocation()
	return ruo
}

// RemoveLocationIDs removes the "location" edge to Location entities by IDs.
func (ruo *ReviewUpdateOne) RemoveLocationIDs(ids ...int) *ReviewUpdateOne {
	ruo.mutation.RemoveLocationIDs(ids...)
	return ruo
}

// RemoveLocation removes "location" edges to Location entities.
func (ruo *ReviewUpdateOne) RemoveLocation(l ...*Location) *ReviewUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ruo.RemoveLocationIDs(ids...)
}

// Where appends a list predicates to the ReviewUpdate builder.
func (ruo *ReviewUpdateOne) Where(ps ...predicate.Review) *ReviewUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *ReviewUpdateOne) Select(field string, fields ...string) *ReviewUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Review entity.
func (ruo *ReviewUpdateOne) Save(ctx context.Context) (*Review, error) {
	return withHooks[*Review, ReviewMutation](ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *ReviewUpdateOne) SaveX(ctx context.Context) *Review {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *ReviewUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *ReviewUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ruo *ReviewUpdateOne) sqlSave(ctx context.Context) (_node *Review, err error) {
	_spec := sqlgraph.NewUpdateSpec(review.Table, review.Columns, sqlgraph.NewFieldSpec(review.FieldID, field.TypeInt))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Review.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, review.FieldID)
		for _, f := range fields {
			if !review.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != review.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.ReviewID(); ok {
		_spec.SetField(review.FieldReviewID, field.TypeInt64, value)
	}
	if value, ok := ruo.mutation.AddedReviewID(); ok {
		_spec.AddField(review.FieldReviewID, field.TypeInt64, value)
	}
	if value, ok := ruo.mutation.Rating(); ok {
		_spec.SetField(review.FieldRating, field.TypeFloat64, value)
	}
	if value, ok := ruo.mutation.AddedRating(); ok {
		_spec.AddField(review.FieldRating, field.TypeFloat64, value)
	}
	if value, ok := ruo.mutation.Message(); ok {
		_spec.SetField(review.FieldMessage, field.TypeString, value)
	}
	if ruo.mutation.LocationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   review.LocationTable,
			Columns: review.LocationPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: location.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedLocationIDs(); len(nodes) > 0 && !ruo.mutation.LocationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   review.LocationTable,
			Columns: review.LocationPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: location.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.LocationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   review.LocationTable,
			Columns: review.LocationPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: location.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Review{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{review.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}
