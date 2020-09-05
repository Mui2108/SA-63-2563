// Code generated by entc, DO NOT EDIT.

package title

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/panupong/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Title {
	return predicate.Title(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Title {
	return predicate.Title(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Title {
	return predicate.Title(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Title {
	return predicate.Title(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Title {
	return predicate.Title(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Title {
	return predicate.Title(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Title {
	return predicate.Title(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Title {
	return predicate.Title(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Title {
	return predicate.Title(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// TitleType applies equality check predicate on the "Title_type" field. It's identical to TitleTypeEQ.
func TitleType(v string) predicate.Title {
	return predicate.Title(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitleType), v))
	})
}

// TitleTypeEQ applies the EQ predicate on the "Title_type" field.
func TitleTypeEQ(v string) predicate.Title {
	return predicate.Title(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitleType), v))
	})
}

// TitleTypeNEQ applies the NEQ predicate on the "Title_type" field.
func TitleTypeNEQ(v string) predicate.Title {
	return predicate.Title(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitleType), v))
	})
}

// TitleTypeIn applies the In predicate on the "Title_type" field.
func TitleTypeIn(vs ...string) predicate.Title {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Title(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTitleType), v...))
	})
}

// TitleTypeNotIn applies the NotIn predicate on the "Title_type" field.
func TitleTypeNotIn(vs ...string) predicate.Title {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Title(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTitleType), v...))
	})
}

// TitleTypeGT applies the GT predicate on the "Title_type" field.
func TitleTypeGT(v string) predicate.Title {
	return predicate.Title(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitleType), v))
	})
}

// TitleTypeGTE applies the GTE predicate on the "Title_type" field.
func TitleTypeGTE(v string) predicate.Title {
	return predicate.Title(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitleType), v))
	})
}

// TitleTypeLT applies the LT predicate on the "Title_type" field.
func TitleTypeLT(v string) predicate.Title {
	return predicate.Title(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitleType), v))
	})
}

// TitleTypeLTE applies the LTE predicate on the "Title_type" field.
func TitleTypeLTE(v string) predicate.Title {
	return predicate.Title(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitleType), v))
	})
}

// TitleTypeContains applies the Contains predicate on the "Title_type" field.
func TitleTypeContains(v string) predicate.Title {
	return predicate.Title(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitleType), v))
	})
}

// TitleTypeHasPrefix applies the HasPrefix predicate on the "Title_type" field.
func TitleTypeHasPrefix(v string) predicate.Title {
	return predicate.Title(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitleType), v))
	})
}

// TitleTypeHasSuffix applies the HasSuffix predicate on the "Title_type" field.
func TitleTypeHasSuffix(v string) predicate.Title {
	return predicate.Title(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitleType), v))
	})
}

// TitleTypeEqualFold applies the EqualFold predicate on the "Title_type" field.
func TitleTypeEqualFold(v string) predicate.Title {
	return predicate.Title(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitleType), v))
	})
}

// TitleTypeContainsFold applies the ContainsFold predicate on the "Title_type" field.
func TitleTypeContainsFold(v string) predicate.Title {
	return predicate.Title(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitleType), v))
	})
}

// HasTitle applies the HasEdge predicate on the "title" edge.
func HasTitle() predicate.Title {
	return predicate.Title(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TitleTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TitleTable, TitleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTitleWith applies the HasEdge predicate on the "title" edge with a given conditions (other predicates).
func HasTitleWith(preds ...predicate.Patient) predicate.Title {
	return predicate.Title(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TitleInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TitleTable, TitleColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Title) predicate.Title {
	return predicate.Title(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Title) predicate.Title {
	return predicate.Title(func(s *sql.Selector) {
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
func Not(p predicate.Title) predicate.Title {
	return predicate.Title(func(s *sql.Selector) {
		p(s.Not())
	})
}