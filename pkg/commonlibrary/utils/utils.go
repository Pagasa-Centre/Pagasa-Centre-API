package utils

import "strings"

// todo: might need to move into user service
func NormalizePhoneNumber(phone string) string {
	if strings.HasPrefix(phone, "07") {
		// Convert UK local (07...) to E.164 (+44...)
		return "+44" + phone[1:]
	}

	return phone
}
