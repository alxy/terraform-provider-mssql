package mssql

import (
	"context"

	"github.com/betr-io/terraform-provider-mssql/mssql/model"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type RoleConnector interface {
  CreateRole(ctx context.Context, roleName string) error
  GetRole(ctx context.Context, roleName string) (*model.Role, error)
  UpdateRole(ctx context.Context, roleName string) error
  DeleteRole(ctx context.Context, roleName string) error
}

func resourceRole() *schema.Resource {
  return &schema.Resource{
    CreateContext: resourceRoleCreate,
    ReadContext:   resourceRoleRead,
    UpdateContext: resourceRoleUpdate,
    DeleteContext: resourceRoleDelete,
    Importer: &schema.ResourceImporter{
      StateContext: resourceRoleImport,
    },
    Schema: map[string]*schema.Schema{
      serverProp: {
        Type:         schema.TypeList,
        MaxItems:     1,
        Required:     true,
        Elem: &schema.Resource{
          Schema: getServerSchema(serverProp),
        },
      },
      roleNameProp: {
        Type:     schema.TypeString,
        Required: true,
        ForceNew: true,
      },
      principalIdProp: {
        Type:     schema.TypeInt,
        Computed: true,
      },
    },
    Timeouts: &schema.ResourceTimeout{
      Default: defaultTimeout,
    },
  }
}

func resourceRoleCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	connector, err := getRoleConnector(meta, data)
	if err != nil {
		return diag.FromErr(err)
	}

	connector.CreateRole(ctx, "asdasd")
	return nil
}

func resourceRoleRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceRoleUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceRoleDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceRoleImport(ctx context.Context, data *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	return nil, nil
}

func getRoleConnector(meta interface{}, data *schema.ResourceData) (RoleConnector, error) {
	provider := meta.(model.Provider)
	connector, err := provider.GetConnector(serverProp, data)
	if err != nil {
	  return nil, err
	}
	return connector.(RoleConnector), nil
  }