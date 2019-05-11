package dbv1

import (
	"github.com/pedafy/pedafy-assignments/src/database"
)

func (d *DBV1) GetAllAssignments() ([]database.Assignments, error) {
	return nil, nil
}

func (d *DBV1) GetAllOrderAssignments(order string) ([]database.Assignments, error) {
	return nil, nil
}

func (d *DBV1) GetAllByFilterAssignments(filter, value string) ([]database.Assignments, error) {
	return nil, nil
}

func (d *DBV1) GetAssignmentsByID(ID int) (database.Assignments, error) {
	return database.Assignments{}, nil
}

func (d *DBV1) NewAssignment(assignment database.Assignments) (database.Assignments, error) {
	return database.Assignments{}, nil
}

func (d *DBV1) ModifyAssignment(assignment database.Assignments, ID int) (database.Assignments, error) {
	return database.Assignments{}, nil
}

func (d *DBV1) ArchiveAssignment(ID int) (database.Assignments, error) {
	return database.Assignments{}, nil
}
