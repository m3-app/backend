// Code generated by ent, DO NOT EDIT.

package address

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/m3-app/backend/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldID, id))
}

// AddressID applies equality check predicate on the "address_id" field. It's identical to AddressIDEQ.
func AddressID(v int64) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldAddressID, v))
}

// Latitude applies equality check predicate on the "latitude" field. It's identical to LatitudeEQ.
func Latitude(v float64) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldLatitude, v))
}

// Longitude applies equality check predicate on the "longitude" field. It's identical to LongitudeEQ.
func Longitude(v float64) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldLongitude, v))
}

// MapsLink applies equality check predicate on the "maps_link" field. It's identical to MapsLinkEQ.
func MapsLink(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldMapsLink, v))
}

// FullAddress applies equality check predicate on the "full_address" field. It's identical to FullAddressEQ.
func FullAddress(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldFullAddress, v))
}

// AddressIDEQ applies the EQ predicate on the "address_id" field.
func AddressIDEQ(v int64) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldAddressID, v))
}

// AddressIDNEQ applies the NEQ predicate on the "address_id" field.
func AddressIDNEQ(v int64) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldAddressID, v))
}

// AddressIDIn applies the In predicate on the "address_id" field.
func AddressIDIn(vs ...int64) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldAddressID, vs...))
}

// AddressIDNotIn applies the NotIn predicate on the "address_id" field.
func AddressIDNotIn(vs ...int64) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldAddressID, vs...))
}

// AddressIDGT applies the GT predicate on the "address_id" field.
func AddressIDGT(v int64) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldAddressID, v))
}

// AddressIDGTE applies the GTE predicate on the "address_id" field.
func AddressIDGTE(v int64) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldAddressID, v))
}

// AddressIDLT applies the LT predicate on the "address_id" field.
func AddressIDLT(v int64) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldAddressID, v))
}

// AddressIDLTE applies the LTE predicate on the "address_id" field.
func AddressIDLTE(v int64) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldAddressID, v))
}

// LatitudeEQ applies the EQ predicate on the "latitude" field.
func LatitudeEQ(v float64) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldLatitude, v))
}

// LatitudeNEQ applies the NEQ predicate on the "latitude" field.
func LatitudeNEQ(v float64) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldLatitude, v))
}

// LatitudeIn applies the In predicate on the "latitude" field.
func LatitudeIn(vs ...float64) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldLatitude, vs...))
}

// LatitudeNotIn applies the NotIn predicate on the "latitude" field.
func LatitudeNotIn(vs ...float64) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldLatitude, vs...))
}

// LatitudeGT applies the GT predicate on the "latitude" field.
func LatitudeGT(v float64) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldLatitude, v))
}

// LatitudeGTE applies the GTE predicate on the "latitude" field.
func LatitudeGTE(v float64) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldLatitude, v))
}

// LatitudeLT applies the LT predicate on the "latitude" field.
func LatitudeLT(v float64) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldLatitude, v))
}

// LatitudeLTE applies the LTE predicate on the "latitude" field.
func LatitudeLTE(v float64) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldLatitude, v))
}

// LongitudeEQ applies the EQ predicate on the "longitude" field.
func LongitudeEQ(v float64) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldLongitude, v))
}

// LongitudeNEQ applies the NEQ predicate on the "longitude" field.
func LongitudeNEQ(v float64) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldLongitude, v))
}

// LongitudeIn applies the In predicate on the "longitude" field.
func LongitudeIn(vs ...float64) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldLongitude, vs...))
}

// LongitudeNotIn applies the NotIn predicate on the "longitude" field.
func LongitudeNotIn(vs ...float64) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldLongitude, vs...))
}

// LongitudeGT applies the GT predicate on the "longitude" field.
func LongitudeGT(v float64) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldLongitude, v))
}

// LongitudeGTE applies the GTE predicate on the "longitude" field.
func LongitudeGTE(v float64) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldLongitude, v))
}

// LongitudeLT applies the LT predicate on the "longitude" field.
func LongitudeLT(v float64) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldLongitude, v))
}

// LongitudeLTE applies the LTE predicate on the "longitude" field.
func LongitudeLTE(v float64) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldLongitude, v))
}

// MapsLinkEQ applies the EQ predicate on the "maps_link" field.
func MapsLinkEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldMapsLink, v))
}

// MapsLinkNEQ applies the NEQ predicate on the "maps_link" field.
func MapsLinkNEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldMapsLink, v))
}

// MapsLinkIn applies the In predicate on the "maps_link" field.
func MapsLinkIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldMapsLink, vs...))
}

// MapsLinkNotIn applies the NotIn predicate on the "maps_link" field.
func MapsLinkNotIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldMapsLink, vs...))
}

// MapsLinkGT applies the GT predicate on the "maps_link" field.
func MapsLinkGT(v string) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldMapsLink, v))
}

// MapsLinkGTE applies the GTE predicate on the "maps_link" field.
func MapsLinkGTE(v string) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldMapsLink, v))
}

// MapsLinkLT applies the LT predicate on the "maps_link" field.
func MapsLinkLT(v string) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldMapsLink, v))
}

// MapsLinkLTE applies the LTE predicate on the "maps_link" field.
func MapsLinkLTE(v string) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldMapsLink, v))
}

// MapsLinkContains applies the Contains predicate on the "maps_link" field.
func MapsLinkContains(v string) predicate.Address {
	return predicate.Address(sql.FieldContains(FieldMapsLink, v))
}

// MapsLinkHasPrefix applies the HasPrefix predicate on the "maps_link" field.
func MapsLinkHasPrefix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasPrefix(FieldMapsLink, v))
}

// MapsLinkHasSuffix applies the HasSuffix predicate on the "maps_link" field.
func MapsLinkHasSuffix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasSuffix(FieldMapsLink, v))
}

// MapsLinkEqualFold applies the EqualFold predicate on the "maps_link" field.
func MapsLinkEqualFold(v string) predicate.Address {
	return predicate.Address(sql.FieldEqualFold(FieldMapsLink, v))
}

// MapsLinkContainsFold applies the ContainsFold predicate on the "maps_link" field.
func MapsLinkContainsFold(v string) predicate.Address {
	return predicate.Address(sql.FieldContainsFold(FieldMapsLink, v))
}

// FullAddressEQ applies the EQ predicate on the "full_address" field.
func FullAddressEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldFullAddress, v))
}

// FullAddressNEQ applies the NEQ predicate on the "full_address" field.
func FullAddressNEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldFullAddress, v))
}

// FullAddressIn applies the In predicate on the "full_address" field.
func FullAddressIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldFullAddress, vs...))
}

// FullAddressNotIn applies the NotIn predicate on the "full_address" field.
func FullAddressNotIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldFullAddress, vs...))
}

// FullAddressGT applies the GT predicate on the "full_address" field.
func FullAddressGT(v string) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldFullAddress, v))
}

// FullAddressGTE applies the GTE predicate on the "full_address" field.
func FullAddressGTE(v string) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldFullAddress, v))
}

// FullAddressLT applies the LT predicate on the "full_address" field.
func FullAddressLT(v string) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldFullAddress, v))
}

// FullAddressLTE applies the LTE predicate on the "full_address" field.
func FullAddressLTE(v string) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldFullAddress, v))
}

// FullAddressContains applies the Contains predicate on the "full_address" field.
func FullAddressContains(v string) predicate.Address {
	return predicate.Address(sql.FieldContains(FieldFullAddress, v))
}

// FullAddressHasPrefix applies the HasPrefix predicate on the "full_address" field.
func FullAddressHasPrefix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasPrefix(FieldFullAddress, v))
}

// FullAddressHasSuffix applies the HasSuffix predicate on the "full_address" field.
func FullAddressHasSuffix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasSuffix(FieldFullAddress, v))
}

// FullAddressEqualFold applies the EqualFold predicate on the "full_address" field.
func FullAddressEqualFold(v string) predicate.Address {
	return predicate.Address(sql.FieldEqualFold(FieldFullAddress, v))
}

// FullAddressContainsFold applies the ContainsFold predicate on the "full_address" field.
func FullAddressContainsFold(v string) predicate.Address {
	return predicate.Address(sql.FieldContainsFold(FieldFullAddress, v))
}

// HasLocation applies the HasEdge predicate on the "location" edge.
func HasLocation() predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, LocationTable, LocationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLocationWith applies the HasEdge predicate on the "location" edge with a given conditions (other predicates).
func HasLocationWith(preds ...predicate.Location) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(LocationInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, LocationTable, LocationColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Address) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Address) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
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
func Not(p predicate.Address) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		p(s.Not())
	})
}