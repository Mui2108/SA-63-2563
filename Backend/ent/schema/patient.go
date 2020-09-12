package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"time"
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
		field.Time("Birthday").Default(time.Now).Immutable(),
            
            
	}

}

// Edges of the Patient.
func (Patient) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("patients", Gender.Type).Ref("genders").Unique(),
		edge.From("patients", Title.Type).Ref("titles").Unique(),
		edge.From("patients", Job.Type).Ref("jobs").Unique(),
	}
}
	