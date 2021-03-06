// Code generated by entc, DO NOT EDIT.

package title

const (
	// Label holds the string label denoting the title type in the database.
	Label = "title"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitleType holds the string denoting the title_type field in the database.
	FieldTitleType = "title_type"

	// EdgeTitles holds the string denoting the titles edge name in mutations.
	EdgeTitles = "titles"

	// Table holds the table name of the title in the database.
	Table = "titles"
	// TitlesTable is the table the holds the titles relation/edge.
	TitlesTable = "patients"
	// TitlesInverseTable is the table name for the Patient entity.
	// It exists in this package in order to avoid circular dependency with the "patient" package.
	TitlesInverseTable = "patients"
	// TitlesColumn is the table column denoting the titles relation/edge.
	TitlesColumn = "title_titles"
)

// Columns holds all SQL columns for title fields.
var Columns = []string{
	FieldID,
	FieldTitleType,
}

var (
	// TitleTypeValidator is a validator for the "Title_type" field. It is called by the builders before save.
	TitleTypeValidator func(string) error
)
