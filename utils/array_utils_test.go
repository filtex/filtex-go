package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsInAny_ShouldReturnFalse_WhenSourcesDoesNotContainItem(t *testing.T) {
	// Act
	// Assert
	assert.False(t, IsInAny("test", []string{"test1"}, []string{"test2"}))
	assert.False(t, IsInAny(100, []int{90}, []int{110}))
	assert.False(t, IsInAny(true, []bool{false}))
}

func TestIsInAny_ShouldReturnTrue_WhenSourcesContainsItem(t *testing.T) {
	// Act
	// Assert
	assert.True(t, IsInAny("test", []string{"test"}, []string{"test2"}, []string{"test3"}))
	assert.True(t, IsInAny(100, []int{90}, []int{100, 110}))
	assert.True(t, IsInAny(true, []bool{false}, []bool{true, false}))
}
