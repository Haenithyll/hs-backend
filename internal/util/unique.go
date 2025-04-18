package util

import (
	"github.com/google/uuid"
)

func AreUUIDsUnique(ids []uuid.UUID) bool {
	seen := make(map[uuid.UUID]struct{}, len(ids))

	for _, id := range ids {
		if _, exists := seen[id]; exists {
			return false
		}
		seen[id] = struct{}{}
	}
	return true
}

func AreUintsUnique(ids []uint8) bool {
	seen := make(map[uint8]struct{}, len(ids))

	for _, id := range ids {
		if _, exists := seen[id]; exists {
			return false
		}
		seen[id] = struct{}{}
	}
	return true
}
