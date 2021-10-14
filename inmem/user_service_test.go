package inmem

import (
	"context"
	"qaplagql/graph/model"
	"testing"
)

func TestCreateUser(t *testing.T) {
	// Setup
	userMap := make(map[string]*model.User)
	projectMap := make(map[string]*model.Project)
	projectUserMap := make(map[string][]string)
	userService := &UserServiceInmem{
		users:       userMap,
		projects:    projectMap,
		projectUser: projectUserMap,
	}
	ctx := context.Background()
	input := &model.NewUser{
		Email:    "test@eloquia.io",
		Password: "SimplePasswordForTesting",
	}

	// Test
	result, err := userService.CreateUser(ctx, input)

	// Assert
	if err != nil {
		t.Log("Error should be nil", err)
		t.Fail()
	}

	if result == nil {
		t.Log("userService should return created User with ID")
		t.Fail()
	}

	if result.ID == "" {
		t.Log("Expected created User to have a non-empty ID")
		t.Fail()
	}
}

func TestGetAllUsers(t *testing.T) {
	// Setup
	userMap := make(map[string]*model.User)
	projectMap := make(map[string]*model.Project)
	projectUserMap := make(map[string][]string)
	userService := &UserServiceInmem{
		users:       userMap,
		projects:    projectMap,
		projectUser: projectUserMap,
	}

	// Test
	_, err := userService.GetAll()

	// Assert
	if err != nil {
		t.Log("Expected no errors", err)
		t.Fail()
	}

}
