package database

// DatabaseHandler interfaces all the version of the database code
type DatabaseHandler interface {
	ConnectionDatabase(user, pass, dbname, url string) error
	IsNewDB() bool
	IsOldDB() bool

	// Status
	GetAllStatus() ([]Status, error)
	GetStatusByID(ID int) (Status, error)
	GetStatusByName(name string) (Status, error)

	// Assignments
	GetAllAssignments() ([]Assignments, error)
	GetAllOrderAssignments(order string) ([]Assignments, error)
	GetAllByFilterAssignments(filter, value string) ([]Assignments, error)
	NewAssignment(assignment Assignments) (Assignments, error)
	ModifyAssignment(assignment Assignments, ID int) (Assignments, error)
	ArchiveAssignment(ID int) (Assignments, error)
}
