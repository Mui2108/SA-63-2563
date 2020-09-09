// Code generated by entc, DO NOT EDIT.

package gender

const (
	// Label holds the string label denoting the gender type in the database.
	Label = "gender"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldGenderType holds the string denoting the gender_type field in the database.
	FieldGenderType = "gender_type"

	// EdgeGenders holds the string denoting the genders edge name in mutations.
	EdgeGenders = "genders"

	// Table holds the table name of the gender in the database.
	Table = "genders"
	// GendersTable is the table the holds the genders relation/edge.
	GendersTable = "patients"
	// GendersInverseTable is the table name for the Patient entity.
	// It exists in this package in order to avoid circular dependency with the "patient" package.
	GendersInverseTable = "patients"
	// GendersColumn is the table column denoting the genders relation/edge.
	GendersColumn = "gender_genders"
)

// Columns holds all SQL columns for gender fields.
var Columns = []string{
	FieldID,
	FieldGenderType,
}

var (
	// GenderTypeValidator is a validator for the "Gender_type" field. It is called by the builders before save.
	GenderTypeValidator func(string) error
)
