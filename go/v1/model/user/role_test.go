package user

import (
	"testing"
)

func TestRole0000_Admin(t *testing.T) {
	r := Role("Admin")
	if !r.IsAdmin() {
		t.Errorf("r must be Admin : %s", string(r))
	}

	r = Role("Read,Admin")
	if !r.IsAdmin() {
		t.Errorf("r must be Admin : %s", string(r))
	}
}

func TestRole0001_Not_Admin(t *testing.T) {
	r := Role("Read")
	if r.IsAdmin() {
		t.Errorf("r must not be Admin : %s", string(r))
	}

	r = Role("Read,Write")
	if r.IsAdmin() {
		t.Errorf("r must not be Admin : %s", string(r))
	}
}

func TestRole0100_CanRead(t *testing.T) {
	r := Role("Read")
	if !r.CanRead() {
		t.Errorf("r must be able to read : %s", string(r))
	}

	r = Role("Read,Write")
	if !r.CanRead() {
		t.Errorf("r must be able to read : %s", string(r))
	}
}

func TestRole0101_NoCanRead(t *testing.T) {
	r := Role("Write")
	if r.CanRead() {
		t.Errorf("r must not be able to read : %s", string(r))
	}

	r = Role("Admin,Write")
	if r.CanRead() {
		t.Errorf("r must not be able to read : %s", string(r))
	}
}

func TestRole0200_CanWrite(t *testing.T) {
	r := Role("Write")
	if !r.CanWrite() {
		t.Errorf("r must be able to write : %s", string(r))
	}

	r = Role("Read,Write")
	if !r.CanWrite() {
		t.Errorf("r must be able to write : %s", string(r))
	}
}
