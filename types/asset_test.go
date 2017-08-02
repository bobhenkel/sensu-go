package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFixtureCreatesValidAsset(t *testing.T) {
	assert := assert.New(t)
	a := FixtureAsset("one")
	assert.Equal("one", a.Name)
	assert.NoError(a.Validate())
}

func TestValidator(t *testing.T) {
	assert := assert.New(t)
	asset := FixtureAsset("one")

	// Given valid asset it should pass
	assert.NoError(asset.Validate())

	// Given asset without a name it should not pass
	asset = FixtureAsset("")
	assert.Error(asset.Validate())

	// Given asset without a URL it should not pass
	asset = FixtureAsset("name")
	asset.URL = ""
	assert.Error(asset.Validate())

	// Given asset without an organization it should not pass
	asset = FixtureAsset("name")
	asset.Organization = ""
	assert.Error(asset.Validate())

	// Given asset with an invalid URL it should not pass
	asset = FixtureAsset("name")
	asset.URL = "asdfasdf"
	assert.Error(asset.Validate())

	// Given asset with a non-HTTP URL it should not pass
	asset = FixtureAsset("name")
	asset.URL = "file:///root/my_script.sh"
	assert.Error(asset.Validate())

	// Given asset with a filter that has a modifier token
	asset = FixtureAsset("name")
	asset.Filters = []string{"lol+1.0"}
	assert.Error(asset.Validate())

	// Given asset with a filter that has a modifier token
	asset = FixtureAsset("name")
	asset.Filters = []string{"!6!!6"}
	assert.Error(asset.Validate())

	// Given asset with valid filters
	asset = FixtureAsset("name")
	asset.Filters = []string{`entity.OS in ("macos", "linux")`}
	assert.NoError(asset.Validate())
}
