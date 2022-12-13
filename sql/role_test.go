package sql

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/betr-io/terraform-provider-mssql/mssql/model"
)

func TestConnector_GetRole(t *testing.T) {
	localConnector, _ := getLocalConnector()
	role, err := localConnector.GetRole(context.Background(), "master", "my_role")

	if err != nil {
		t.Fatalf("Role couldn't be read.")
	}

	log.Print(role)
}

func getLocalConnector() (Connector, error) {
	var timeout time.Duration = 1000000000

	connector := &Connector{
		Host:     "localhost",
		Port:     "1433",
		Database: "master",
		Timeout:  timeout,
	}
	connector.Login = &LoginUser{
		Username: "SA",
		Password: "!!up3R!!3cR37",
	}

	return *connector, nil
}

func TestConnector_UpdateRole(t *testing.T) {
	localConnector, _ := getLocalConnector()
	role := &model.Role{
		Name:        "updated_name",
		PrincipalID: 7,
	}

	err := localConnector.UpdateRole(context.Background(), "master", role)

	if err != nil {
		t.Fatalf("Role couldn't be updated.")
	}
}

func TestConnector_CreateRole(t *testing.T) {
	localConnector, _ := getLocalConnector()
	role := &model.Role{
		Name: "my_create_role",
		PrincipalID: 0,
	}

	err := localConnector.CreateRole(context.Background(), "master", role)

	if err != nil {
		t.Fatalf("Role couldn't be created.")
	}
}
