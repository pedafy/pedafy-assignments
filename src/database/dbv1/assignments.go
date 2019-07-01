package dbv1

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/pedafy/pedafy-assignments/src/database"
)

// GetAllAssignments return all the assignments found in the database
func (d *DBV1) GetAllAssignments() ([]database.Assignments, error) {
	sql := "SELECT * FROM `assignments`"

	resp, err := d.dbc.Query(sql)
	if err != nil {
		return nil, err
	}
	defer resp.Close()

	assignment := make([]database.Assignments, 0)

	for resp.Next() {
		var curr database.Assignments
		err = resp.Scan(&curr.ID, &curr.CreatorID, &curr.AssignedID, &curr.StatusID, &curr.TaskID, &curr.CreatedAt, &curr.LastEdit, &curr.DueDate, &curr.CompletionDate, &curr.Title, &curr.Description)
		if err != nil {
			return nil, err
		}
		assignment = append(assignment, curr)
	}
	return assignment, nil
}

// GetAllOrderAssignments get all assignments by order
func (d *DBV1) GetAllOrderAssignments(order string) ([]database.Assignments, error) {
	var sql string

	switch order {
	case "status":
		sql = "SELECT * FROM `assignments` ORDER BY `status_id` DESC"
	case "due":
		sql = "SELECT * FROM `assignments` ORDER BY `due_date` DESC"
	case "new":
		sql = "SELECT * FROM `assignments` ORDER BY `id` DESC"
	default:
		return d.GetAllAssignments()
	}

	resp, err := d.dbc.Query(sql)
	if err != nil {
		return nil, err
	}
	defer resp.Close()

	assignment := make([]database.Assignments, 0)

	for resp.Next() {
		var curr database.Assignments
		err = resp.Scan(&curr.ID, &curr.CreatorID, &curr.AssignedID, &curr.StatusID, &curr.TaskID, &curr.CreatedAt, &curr.LastEdit, &curr.DueDate, &curr.CompletionDate, &curr.Title, &curr.Description)
		if err != nil {
			return nil, err
		}
		assignment = append(assignment, curr)
	}
	return assignment, nil
}

// GetAllByFilterAssignments allows to get every assignments with a filter,
// when the filter is "status_id" and the value is "2" the function
// will query and return only assignments with a status_id equal to 2
func (d *DBV1) GetAllByFilterAssignments(filter, value string) ([]database.Assignments, error) {
	sql := fmt.Sprintf("SELECT * FROM `assignments` WHERE `%s` = \"%s\"", filter, value)
	log.Println(sql)
	resp, err := d.dbc.Query(sql)
	if err != nil {
		return nil, err
	}
	defer resp.Close()

	assignment := make([]database.Assignments, 0)

	for resp.Next() {
		var curr database.Assignments
		err = resp.Scan(&curr.ID, &curr.CreatorID, &curr.AssignedID, &curr.StatusID, &curr.TaskID, &curr.CreatedAt, &curr.LastEdit, &curr.DueDate, &curr.CompletionDate, &curr.Title, &curr.Description)
		if err != nil {
			return nil, err
		}
		assignment = append(assignment, curr)
	}
	return assignment, nil
}

// NewAssignment will add the given assignment to the database and return it
func (d *DBV1) NewAssignment(assignment database.Assignments) (database.Assignments, error) {
	query, err := d.dbc.Prepare("INSERT INTO `assignments` (creator_id,assigned_id,status_id,task_id,due_date,last_edit,title,description) VALUES(?,?,?,?,?,?,?,?)")
	if err != nil {
		return database.Assignments{}, err
	}
	defer query.Close()

	result, err := query.Exec(assignment.CreatorID, assignment.AssignedID, assignment.StatusID, assignment.TaskID, assignment.DueDate, time.Now(), assignment.Title, assignment.Description)
	if err != nil {
		return database.Assignments{}, err
	}

	// Get the last inserted ID in order to retreive the new todo and return it.
	newID, err := result.LastInsertId()
	if err != nil {
		return database.Assignments{}, err
	}

	newAssignment, err := d.GetAllByFilterAssignments("id", strconv.Itoa(int(newID)))
	if err != nil {
		return database.Assignments{}, err
	}
	return newAssignment[0], nil
}

// ModifyAssignment the whole assignment with the given ID and data
func (d *DBV1) ModifyAssignment(assignment database.Assignments, ID int) (database.Assignments, error) {
	if d.isInDatabase(ID) == false {
		return database.Assignments{}, nil
	}

	sql := "UPDATE `assignments` SET creator_id=?, assigned_id=?, status_id=?, task_id=?, last_edit=?, due_date=?, completion_date=?, title=?, description=? WHERE id=?"
	query, err := d.dbc.Prepare(sql)
	if err != nil {
		return database.Assignments{}, err
	}
	defer query.Close()

	_, err = query.Exec(assignment.CreatorID, assignment.AssignedID, assignment.StatusID, assignment.TaskID, time.Now(), assignment.DueDate, assignment.CompletionDate, assignment.Title, assignment.Description, ID)
	if err != nil {
		return database.Assignments{}, err
	}
	assignments, err := d.GetAllByFilterAssignments("id", strconv.Itoa(ID))
	if err != nil {
		return database.Assignments{}, err
	}
	return assignments[0], nil
}

// ArchiveAssignment will directy archive the assignment
func (d *DBV1) ArchiveAssignment(ID int) (database.Assignments, error) {
	if d.isInDatabase(ID) == false {
		return database.Assignments{}, nil
	}

	archiveID, err := d.GetStatusByName("archived")
	if err != nil {
		return database.Assignments{}, err
	}
	sql := "UPDATE `assignments` SET status_id=?, last_edit=? WHERE id=?"
	query, err := d.dbc.Prepare(sql)
	if err != nil {
		return database.Assignments{}, err
	}
	defer query.Close()

	_, err = query.Exec(strconv.Itoa(archiveID.ID), time.Now(), strconv.Itoa(ID))
	if err != nil {
		return database.Assignments{}, err
	}
	assignments, err := d.GetAllByFilterAssignments("id", strconv.Itoa(ID))
	if err != nil {
		return database.Assignments{}, err
	}
	return assignments[0], nil
}

func (d *DBV1) isInDatabase(ID int) bool {
	assignments, _ := d.GetAllByFilterAssignments("id", strconv.Itoa(ID))
	if len(assignments) == 0 || assignments[0].ID == 0 {
		return false
	}
	return true
}
