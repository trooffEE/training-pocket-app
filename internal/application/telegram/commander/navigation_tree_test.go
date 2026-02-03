package commander

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const filePath = "./test__mocks/client_navigation_tree_test.yaml"

var tree = LoadNavigationTree(filePath)

func TestLoadNavigationTreeSmokeTest(t *testing.T) {
	tree := LoadNavigationTree(filePath)
	assert.NotNil(t, tree)
}

func TestGetSectionByPathSmokeTest(t *testing.T) {
	val := tree.GetSectionByPath("")
	assert.Nil(t, val, "Key is NOT found!")

	val = tree.GetSectionByPath("start_training")
	assert.NotNil(t, val, "Key is found!")

	val = tree.GetSectionByPath("bla_bla_bla")
	assert.Nil(t, val, "Key is NOT found!")

	val = tree.GetSectionByPath("settings.schedule")
	assert.NotNil(t, val, "Nested key is found!")

	val = tree.GetSectionByPath("settings.schedule123")
	assert.Nil(t, val, "Nested key is NOT found!")

	val = tree.GetSectionByPath("settings.schedule.bla.bla")
	assert.Nil(t, val, "Tree depth is respected! (here is 2, not 4)")
}
