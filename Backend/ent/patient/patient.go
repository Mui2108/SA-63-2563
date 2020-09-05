// Code generated by entc, DO NOT EDIT.

package patient

const (
	// Label holds the string label denoting the patient type in the database.
	Label = "patient"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCardID holds the string denoting the card_id field in the database.
	FieldCardID = "card_id"
	// FieldFistName holds the string denoting the fist_name field in the database.
	FieldFistName = "fist_name"
	// FieldLastName holds the string denoting the last_name field in the database.
	FieldLastName = "last_name"
	// FieldAllergic holds the string denoting the allergic field in the database.
	FieldAllergic = "allergic"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"

	// EdgeGenders holds the string denoting the genders edge name in mutations.
	EdgeGenders = "genders"
	// EdgeTitles holds the string denoting the titles edge name in mutations.
	EdgeTitles = "titles"
	// EdgeJobs holds the string denoting the jobs edge name in mutations.
	EdgeJobs = "jobs"

	// Table holds the table name of the patient in the database.
	Table = "patients"
	// GendersTable is the table the holds the genders relation/edge.
	GendersTable = "genders"
	// GendersInverseTable is the table name for the Gender entity.
	// It exists in this package in order to avoid circular dependency with the "gender" package.
	GendersInverseTable = "genders"
	// GendersColumn is the table column denoting the genders relation/edge.
	GendersColumn = "patient_genders"
	// TitlesTable is the table the holds the titles relation/edge.
	TitlesTable = "titles"
	// TitlesInverseTable is the table name for the Title entity.
	// It exists in this package in order to avoid circular dependency with the "title" package.
	TitlesInverseTable = "titles"
	// TitlesColumn is the table column denoting the titles relation/edge.
	TitlesColumn = "patient_titles"
	// JobsTable is the table the holds the jobs relation/edge.
	JobsTable = "jobs"
	// JobsInverseTable is the table name for the Job entity.
	// It exists in this package in order to avoid circular dependency with the "job" package.
	JobsInverseTable = "jobs"
	// JobsColumn is the table column denoting the jobs relation/edge.
	JobsColumn = "patient_jobs"
)

// Columns holds all SQL columns for patient fields.
var Columns = []string{
	FieldID,
	FieldCardID,
	FieldFistName,
	FieldLastName,
	FieldAllergic,
	FieldAge,
}

var (
	// CardIDValidator is a validator for the "Card_id" field. It is called by the builders before save.
	CardIDValidator func(string) error
	// FistNameValidator is a validator for the "Fist_name" field. It is called by the builders before save.
	FistNameValidator func(string) error
	// LastNameValidator is a validator for the "Last_name" field. It is called by the builders before save.
	LastNameValidator func(string) error
	// AllergicValidator is a validator for the "Allergic" field. It is called by the builders before save.
	AllergicValidator func(string) error
	// AgeValidator is a validator for the "age" field. It is called by the builders before save.
	AgeValidator func(int) error
)