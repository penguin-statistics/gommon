package constant

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDropTypeMap(t *testing.T) {
	for k, v := range DropTypeMap {
		if k == "REGULAR_DROP" {
			assert.Equal(t, "NORMAL_DROP", DropTypeReversedMap[v])
		} else {
			assert.Equal(t, k, DropTypeReversedMap[v])
		}
	}

	for _, v := range DropTypeMapKeys {
		assert.Contains(t, DropTypeMap, v)
	}
}

func TestDropTypeReversedMap(t *testing.T) {
	for k, v := range DropTypeReversedMap {
		assert.Equal(t, k, DropTypeMap[v])
	}
}
