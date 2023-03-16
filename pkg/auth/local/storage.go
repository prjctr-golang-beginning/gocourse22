package local

import (
	"gopkg.in/yaml.v3"
	"os"
)

type permissionGroup struct {
	Name        string   `yaml:"name"`
	Permissions []string `yaml:"permissions"`
}

type storageUser struct {
	Name             string   `yaml:"name"`
	PermissionGroups []string `yaml:"permission-groups,omitempty"`
	Permissions      []string `yaml:"permissions,omitempty"`
}

type localStorage struct {
	PermissionGroups []permissionGroup `yaml:"permission-groups"`
	Users            []storageUser     `yaml:"users"`
}

func (d *localStorage) GetUser(name string) *storageUser {
	for _, user := range d.Users {
		if user.Name == name {
			return &user
		}
	}

	return nil
}

func parse(fileName string) (*localStorage, error) {
	bytes, readErr := os.ReadFile(fileName)
	if readErr != nil {
		return nil, readErr
	}

	var s localStorage
	if parseErr := yaml.Unmarshal(bytes, &s); parseErr != nil {
		return nil, parseErr
	}

	// fill permissions for each user
	for i := range s.Users {
		for _, group := range s.Users[i].PermissionGroups {
			for _, permission := range s.PermissionGroups {
				if permission.Name == group {
					s.Users[i].Permissions = append(s.Users[i].Permissions, permission.Permissions...)
				}
			}
		}
	}

	return &s, nil
}
