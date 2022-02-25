package crud

import (
	"database/sql"
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetDetailsById(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("An error %s occured when opening a database connection", err)
		return
	}

	testCases := []struct {
		desc        string
		id          int
		emp         *Employee
		expectedErr error
		mockQuery   interface{}
	}{
		{
			desc:        "success",
			id:          1,
			emp:         &Employee{1, "John", "jhn@yahoo.com", "intern"},
			expectedErr: nil,
			mockQuery:   mock.ExpectQuery("select * from employee where id=?").WithArgs(1).WillReturnRows(mock.NewRows([]string{"id", "name", "email", "role"}).AddRow(1, "John", "jhn@yahoo.com", "intern")),
		},
		{
			desc:        "success",
			id:          2,
			emp:         &Employee{2, "Jane", "jane@yahoo.com", "intern"},
			expectedErr: nil,
			mockQuery:   mock.ExpectQuery("select * from employee where id=?").WithArgs(2).WillReturnRows(mock.NewRows([]string{"id", "name", "email", "role"}).AddRow(2, "Jane", "jane@yahoo.com", "intern")),
		},

		{
			desc:        "faliure",
			id:          2,
			emp:         nil,
			expectedErr: sql.ErrNoRows,
			mockQuery:   mock.ExpectQuery("select * from employee where id=?").WithArgs(2).WillReturnError(sql.ErrNoRows),
		},
	}

	for i, testCase := range testCases {

		t.Run(testCase.desc, func(t *testing.T) {
			emp, err := GetDetailsById(db, testCase.id)
			if err != nil && err.Error() != testCase.expectedErr.Error() {
				t.Errorf("expected error:%v, got:%v", testCase.expectedErr, err)
				return
			}
			if !reflect.DeepEqual(emp, testCase.emp) {
				t.Errorf("expected userL%v, got:%v", testCases[i], emp)
			}
		})

	}

}

func TestGetDeleteById(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("An error %s occured when opening a database connection", err)
		return
	}

	testCases := []struct {
		desc        string
		id          int
		expectedErr error
		mockQuery   interface{}
	}{
		{
			desc:        "success",
			id:          1,
			expectedErr: sql.ErrNoRows,
			mockQuery:   mock.ExpectExec("delete from employee where id=?").WithArgs(1).WillReturnError(sql.ErrNoRows),
		},

		{
			desc:        "faliure",
			id:          2,
			expectedErr: nil,
			mockQuery:   mock.ExpectExec("delete from employee where id=?").WithArgs(2).WillReturnResult(sqlmock.NewResult(1, 0)),
		},
	}

	for _, testCase := range testCases {
		err := DeleteById(db, testCase.id)
		if err != nil && err.Error() != testCase.expectedErr.Error() {
			t.Errorf("expected error:%v, got:%v", testCase.expectedErr, err)

		}

	}
}

func TestUpdateById(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("An error %s occured when opening a database connection", err)
		return
	}
	testCases := []struct {
		id          int
		name        string
		email       string
		role        string
		mockQuery   interface{}
		expectError error
	}{
		{
			id:    1,
			name:  "John",
			email: "jhn@yahoo.com",
			role:  "intern",

			mockQuery:   mock.ExpectPrepare("UPDATE employee SET name = ?, email = ?, role = ? WHERE id = ?").ExpectExec().WithArgs("John", "jhn@yahoo.com", "intern", 1).WillReturnResult(sqlmock.NewResult(0, 1)),
			expectError: nil,
		},
		// Error
		{
			id:    1,
			name:  "John",
			email: "jhn@yahoo.com",
			role:  "intern",

			mockQuery:   mock.ExpectPrepare("UPDATE employee SET name = ?, email = ?, role = ? WHERE id = ?").WillReturnError(errors.New("query doesn't prepare")),
			expectError: errors.New("query doesn't prepare"),
		},
		// Failure
		{
			id:          2,
			name:        "John",
			email:       "jhn@yahoo.com",
			role:        "intern",
			mockQuery:   mock.ExpectPrepare("UPDATE employee SET name = ?, email = ?, role = ? WHERE id = ?").ExpectExec().WithArgs("John", "jhn@yahoo.com", "intern", 2).WillReturnError(sql.ErrNoRows),
			expectError: sql.ErrNoRows,
		},
	}
	for i, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			err := UpdateById(db, testCases[i].id, testCases[i].name, testCases[i].email, testCases[i].role)
			if err != nil && err.Error() != testCase.expectError.Error() {
				t.Errorf("expected error:%v, got:%v", testCase.expectError, err)
			}
		})
	}
}

func TestInsertData(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("An error %s occured when opening a database connection", err)
		return
	}
	testCases := []struct {
		id          int
		name        string
		email       string
		role        string
		mockQuery   interface{}
		expectError error
	}{
		{
			id:          1,
			name:        "John",
			email:       "jhn@yahoo.com",
			role:        "intern",
			mockQuery:   mock.ExpectExec("INSERT INTO employee (Name,Email,role) values(?,?,?)").WithArgs("John", "jhn@yahoo.com", "intern").WillReturnResult(sqlmock.NewResult(1, 1)),
			expectError: nil,
		},

		{
			id:          2,
			name:        "Johni",
			email:       "jhni@yahoo.com",
			role:        "sde",
			mockQuery:   mock.ExpectExec("INSERT INTO employee (Name,Email,role) values(?,?,?)").WithArgs("Johni", "jhni@yahoo.com", "sde").WillReturnResult(sqlmock.NewResult(1, 1)),
			expectError: nil,
		},
		// Error
		{
			id:          2,
			name:        "John",
			email:       "jhn@yahoo.com",
			role:        "intern",
			mockQuery:   mock.ExpectExec("INSERT INTO employee (Name,Email,role) values(?,?,?)").WithArgs("John", "jhn@yahoo.com", "intern").WillReturnError(errors.New("not inserted")),
			expectError: errors.New("not inserted"),
		},
	}

	for i, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			err := InsertData(db, testCases[i].name, testCases[i].email, testCases[i].role)
			if err != nil && err.Error() != testCase.expectError.Error() {
				t.Errorf("expected error:%v, got:%v", testCase.expectError, err)
			}
		})
	}

}
