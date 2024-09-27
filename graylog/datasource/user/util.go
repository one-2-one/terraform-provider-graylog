package user

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	keyID       = "id"
	keyUsername = "username"
	keyEmail    = "email"
)

func setDataToResourceData(d *schema.ResourceData, data map[string]interface{}) error {
	if id, ok := data[keyID]; ok {
		d.SetId(id.(string))
	} else {
		return fmt.Errorf("error: id not found in user data")
	}

	if username, ok := data[keyUsername]; ok {
		if err := d.Set(keyUsername, username); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("error: username not found in user data")
	}

	if email, ok := data[keyEmail]; ok {
		if err := d.Set(keyEmail, email); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("error: email not found in user data")
	}

	return nil
}
