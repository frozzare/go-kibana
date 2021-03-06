package kibana

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SavedObjectsGetByType(t *testing.T) {
	client := NewClient(NewDefaultConfig())

	result, err := client.SavedObjects().GetByType(
		NewSavedObjectRequestBuilder().
			WithFields("title").
			WithType("index-pattern").
			WithPerPage(15).
			Build())

	assert.Nil(t, err)

	assert.NotNil(t, result)
	assert.Equal(t, 1, result.Page)
	assert.Equal(t, 15, result.PerPage)
	assert.Equal(t, 1, result.Total)

	assert.Len(t, result.SavedObjects, 1)
	assert.Len(t, result.SavedObjects[0].Id, 36)
	assert.Equal(t, "index-pattern", result.SavedObjects[0].Type)
	assert.NotZero(t, result.SavedObjects[0].Version)
	assert.NotNil(t, result.SavedObjects[0].Attributes)
	assert.Equal(t, "logstash-*", result.SavedObjects[0].Attributes["title"])
}
