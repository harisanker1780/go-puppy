package puppy

import (
	"fmt"
	"log"
)

type User struct {
	ID    int
	First string
}

type MockDatastore struct {
	Users map[int]User
}

type Datastore interface {
	GetUser(id int) (User, error)
	SaveUser(u User) error
}

type Service struct {
	ds Datastore
}

func (md MockDatastore) GetUser(id int) (User, error) {
	user, ok := md.Users[id]
	if !ok {
		return User{}, fmt.Errorf("User %v not found", id)
	}
	return user, nil
}

func (md MockDatastore) SaveUser(u User) error {
	_, ok := md.Users[u.ID]
	if ok {
		return fmt.Errorf("User %v already exist", u.ID)
	}
	md.Users[u.ID] = u
	return nil
}

func (s Service) GetUser(id int) (User, error) {
	return s.ds.GetUser(id)
}

func (s Service) SaveUser(u User) error {
	return s.ds.SaveUser(u)
}

func Mock() {
	db := MockDatastore{
		Users: make(map[int]User),
	}

	svc := Service{
		ds: db,
	}

	u1 := User{
		ID:    1,
		First: "Test",
	}

	err := svc.SaveUser(u1)
	if err != nil {
		log.Fatalf("error %s", err)
	}

	u1Returned, err := svc.GetUser(u1.ID)
	if err != nil {
		log.Fatalf("error %s", err)
	}

	fmt.Println(u1)
	fmt.Println(u1Returned)
}
