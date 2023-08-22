package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRU(t *testing.T) {
	cache := NewLRUCache(3)
	cache.Put("name", "tangguh")

	result := cache.Get("name")
	assert.Equal(t, "tangguh", result)

	cache.Put("kelas", "asd")
	cache.Put("alamat", "bandung")
	cache.Put("name", "adalah")

	result2 := cache.Get("name")
	assert.Equal(t, "adalah", result2)
}
