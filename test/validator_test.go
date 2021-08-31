package main

import (
	"testing"

	"github.com/coretrix/hitrix/pkg/helper"
	"github.com/stretchr/testify/assert"
)

func TestBasicValidation(t *testing.T) {
	err := helper.NewValidator().Validate("no-email-string", "email")
	assert.NotNil(t, err)

	err = helper.NewValidator().Validate("awesome-dude@awesome-com.com", "email")
	assert.Nil(t, err)
}

func TestCountryCodeValidation(t *testing.T) {
	err := helper.NewValidator().Validate(1, "country_code")
	assert.NotNil(t, err)

	err = helper.NewValidator().Validate("SE", "country_code")
	assert.Nil(t, err)
}
