package Mancala_Go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBowl_NewBowlHasFourBeads(t *testing.T) {
	// Act
	bowl := CreateBowl()

	// Assert
	assert.Equal(t, uint(4), bowl.NumberOfBeads)
}

func TestBowl_PassOn_NumberOfBeadsIncreasedByOne(t *testing.T) {
	// Arrange
	bowl := CreateBowl()

	// Act
	bowl.PassOn(4)

	// Assert
	assert.Equal(t, uint(5), bowl.NumberOfBeads)
}

func TestBowl_PassOn_ZeroBeads(t *testing.T) {
	// Arrange
	bowl := CreateBowl()

	// Act
	bowl.PassOn(0)

	// Assert
	assert.Equal(t, uint(4), bowl.NumberOfBeads)
}