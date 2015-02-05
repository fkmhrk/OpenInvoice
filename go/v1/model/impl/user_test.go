package impl

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func createUserDAO(db *sql.DB) *userDAO {
	c := NewConnection(db)
	return NewUserDAO(c)
}

func deleteByID(db *sql.DB, id string) {
	s, _ := db.Prepare("DELETE FROM user WHERE id=?")
	defer s.Close()
	s.Exec(id)
}

func insertUser(db *sql.DB, id, name, password string) {
	s, err := db.Prepare("INSERT INTO user" +
		"(id,login_name,password,deleted)" +
		"VALUES(?,?,?,0)")
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	defer s.Close()
	s.Exec(id, name, hashPassword(password))
}

func TestUserDAO_0000_GetByNamePassword(t *testing.T) {
	db, err := connect()
	if err != nil {
		t.Errorf("Failed to connect")
		return
	}
	defer db.Close()
	// prepare
	id := "testID0000"
	name := "test0000"
	password := "pass1234"
	deleteByID(db, id)
	insertUser(db, id, name, password)

	dao := createUserDAO(db)
	user, err := dao.GetByNamePassword(name, password)
	if err != nil {
		t.Errorf("Failed to get user by name / password : %s", err)
		return
	}
	if user.Id != id {
		t.Errorf("Invalid ID %s", user.Id)
		return
	}
}

func TestUserDAO_0001_GetByNamePassword_wrongPass(t *testing.T) {
	db, err := connect()
	if err != nil {
		t.Errorf("Failed to connect")
		return
	}
	defer db.Close()
	// prepare
	id := "testID0000"
	name := "test0000"
	password := "pass1234"
	deleteByID(db, id)
	insertUser(db, id, name, password)

	dao := createUserDAO(db)
	user, err := dao.GetByNamePassword(name, "dummyPass")
	if err != nil {
		t.Errorf("Unexpected error : %s", err)
		return
	}
	if user != nil {
		t.Errorf("User must be nil : %s", user.Id)
		return
	}
}

func TestUserDAO_0002_GetByNamePassword_notFound(t *testing.T) {
	db, err := connect()
	if err != nil {
		t.Errorf("Failed to connect")
		return
	}
	defer db.Close()
	// prepare
	id := "testID0000"
	name := "test0000"
	password := "pass1234"
	deleteByID(db, id)

	dao := createUserDAO(db)
	user, err := dao.GetByNamePassword(name, password)
	if err != nil {
		t.Errorf("Unexpected error : %s", err)
		return
	}
	if user != nil {
		t.Errorf("User must be nil : id=", user.Id)
		return
	}
}
