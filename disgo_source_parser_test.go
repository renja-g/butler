package main

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_searchSource(t *testing.T) {
	err := loadPackages()
	assert.NoError(t, err)

	result := findInPackages("disgo", "New")
	log.Printf("%+v", result)
}
