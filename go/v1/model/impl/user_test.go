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

func deleteUserByName(db *sql.DB, name string) {
	s, _ := db.Prepare("DELETE FROM user WHERE login_name=?")
	defer s.Close()
	s.Exec(name)
}

func insertUser(db *sql.DB, id, name, password string) {
	s, err := db.Prepare("INSERT INTO user" +
		"(id,login_name,display_name,role,tel,password,created_time,modified_time,deleted)" +
		"VALUES(?,?,'demo','Read','',?,0,0,0)")
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

func TestUserDAO_0100_GetList(t *testing.T) {
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
	list, err := dao.GetList()
	if err != nil {
		t.Errorf("Failed to get user by name / password : %s", err)
		return
	}
	l := len(list)

	insertUser(db, id, name, password)
	list, err = dao.GetList()
	if err != nil {
		t.Errorf("Failed to get user by name / password : %s", err)
		return
	}
	if len(list) != l+1 {
		t.Errorf("Wrong length : %d", len(list))
	}
}

func TestUserDAO_0200_Create(t *testing.T) {
	db, err := connect()
	if err != nil {
		t.Errorf("Failed to connect")
		return
	}
	defer db.Close()
	// prepare
	loginName := "test0100"
	displayName := "Demo user"
	tel := "08011112222"
	password := "pass1234"
	deleteUserByName(db, loginName)

	dao := createUserDAO(db)
	user, err := dao.Create(loginName, displayName, "Read", tel, password)
	if err != nil {
		t.Errorf("Failed to create user : %s", err)
		return
	}
	assertUser(t, user, user.Id, loginName, displayName, "Read", tel)
	// login

	user2, err := dao.GetByNamePassword(loginName, password)
	if err != nil {
		t.Errorf("Failed to get user : %s", err)
		return
	}
	assertUser(t, user2, user.Id, loginName, displayName, "Read", tel)
	if !user2.Role.CanRead() {
		t.Errorf("This user must be able to Read")
		return
	}
}

func TestUserDAO_0300_Update(t *testing.T) {
	db, err := connect()
	if err != nil {
		t.Errorf("Failed to connect")
		return
	}
	defer db.Close()
	// prepare
	loginName := "test0100"
	displayName := "Demo user"
	tel := "08011112222"
	password := "pass1234"
	deleteUserByName(db, loginName)

	dao := createUserDAO(db)
	user, err := dao.Create(loginName, displayName, "Read", tel, password)
	if err != nil {
		t.Errorf("Failed to create user : %s", err)
		return
	}
	assertUser(t, user, user.Id, loginName, displayName, "Read", tel)

	// update
	user2, err := dao.Update(user.Id, loginName, "NewName", "Read", "09022223333", "")
	if err != nil {
		t.Errorf("Failed to update user : %s", err)
		return
	}
	assertUser(t, user2, user.Id, loginName, "NewName", "Read", "09022223333")

	// login
	user3, err := dao.GetByNamePassword(loginName, password)
	if err != nil {
		t.Errorf("Failed to get user : %s", err)
		return
	}
	assertUser(t, user3, user.Id, loginName, "NewName", "Read", "09022223333")

	// update with new password
	user4, err := dao.Update(user.Id, loginName, "NewName2", "Read", "09033334444", "NewPassword")
	if err != nil {
		t.Errorf("Failed to update user : %s", err)
		return
	}
	assertUser(t, user4, user.Id, loginName, "NewName2", "Read", "09033334444")

	// login
	user5, err := dao.GetByNamePassword(loginName, "NewPassword")
	if err != nil {
		t.Errorf("Failed to get user : %s", err)
		return
	}
	assertUser(t, user5, user.Id, loginName, "NewName2", "Read", "09033334444")

}
