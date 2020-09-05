package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Patient holds the schema definition for the Patient entity.
type Patient struct {
	ent.Schema
}

// Fields of the Patient.
func (Patient) Fields() []ent.Field {
	return []ent.Field{
		field.String("Card_id").NotEmpty(),
		field.String("Fist_name").NotEmpty(),
		field.String("Last_name").NotEmpty(),
		field.String("Allergic").NotEmpty(),
		field.Int("age").Positive(),
	}

}

// Edges of the Patient.
func (Patient) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("genders", Gender.Type),
		edge.To("titles", Title.Type),
		edge.To("jobs", Job.Type),
	}
}
