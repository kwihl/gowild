package handlers

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestToDTOSingle(t *testing.T) {
	type dto struct {
		ID  string `json:"id"`
		Foo string `json:"foo"`
	}

	body := io.NopCloser(bytes.NewReader([]byte(`{
		"id":"123",
		"foo":"something" 
	}`)))

	var expected = dto{
		ID:  "123",
		Foo: "something",
	}

	res, err := toDTO[dto](body)
	require.NoError(t, err)

	assert.Equal(t, expected.ID, res.ID)
	assert.Equal(t, expected.Foo, res.Foo)
}

func TestToDTOList(t *testing.T) {
	type dto struct {
		ID  string `json:"id"`
		Foo string `json:"foo"`
	}

	body := io.NopCloser(bytes.NewReader([]byte(`[{
		"id":"123",
		"foo":"something" 
	},
	{
		"id":"456",
		"foo":"something else" 
	},
	{
		"id":"789",
		"foo":"something entirely different" 
	}
	]`)))

	var expected = []dto{
		{
			ID:  "123",
			Foo: "something",
		},
		{
			ID:  "456",
			Foo: "something else",
		},
		{
			ID:  "789",
			Foo: "something entirely different",
		},
	}

	res, err := toDTO[[]dto](body)
	require.NoError(t, err)

	assert.Len(t, *res, len(expected))
	assert.Equal(t, (*res)[0].ID, "123")
	assert.Equal(t, (*res)[1].ID, "456")
	assert.Equal(t, (*res)[2].ID, "789")
}

func TestToDTOShouldFail(t *testing.T) {
	type dto struct {
		ID  string `json:"id"`
		Foo string `json:"foo"`
	}

	body := io.NopCloser(bytes.NewReader([]byte(`[{
		"id":"123",
		"foo":"something" 
	},
	{
		"id":"456",
		"foo":"something else" 
	},
	{
		"id":"789",
		"foo":"something entirely different" 
	}
	]`)))

	_, err := toDTO[dto](body) // <-- Parametrized as single object but it contains a list
	assert.Error(t, err)

	body = io.NopCloser(bytes.NewReader([]byte(`{
		"id":"123",
		"foo":"something" 
	}`)))

	_, err = toDTO[[]dto](body) // <-- Parametrized as list but it contains a single object
	assert.Error(t, err)
}
