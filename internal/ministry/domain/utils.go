package domain

import "fmt"

func ToMinistryLeaderRole(ministryName string) (string, error) {
	if ministryName == "" {
		return "", fmt.Errorf("ministry name cannot be empty")
	}
	// todo: maybe add check to see if ministry name is valid minstry name

	return fmt.Sprintf("%s Leader", ministryName), nil
}
