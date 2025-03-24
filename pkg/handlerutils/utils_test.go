package handlerutils

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnmarshalSingle(t *testing.T) {
	type testDTO struct {
		ID  string `json:"id"`
		Foo string `json:"foo"`
	}

	body := []byte(`{
		"id":"123",
		"foo":"something" 
	}`)

	var expected = testDTO{
		ID:  "123",
		Foo: "something",
	}

	res, err := unmarshal[testDTO](body)
	require.NoError(t, err)

	assert.Equal(t, expected.ID, res.ID)
	assert.Equal(t, expected.Foo, res.Foo)
}

func TestUnmarshalList(t *testing.T) {
	type testDTO struct {
		ID  string `json:"id"`
		Foo string `json:"foo"`
	}

	body := []byte(`[{
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
	]`)

	var expected = []testDTO{
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

	res, err := unmarshal[[]testDTO](body)
	require.NoError(t, err)

	assert.Len(t, *res, len(expected))
	assert.Equal(t, (*res)[0].ID, "123")
	assert.Equal(t, (*res)[1].ID, "456")
	assert.Equal(t, (*res)[2].ID, "789")
}

func TestUnmarshalShouldFail(t *testing.T) {
	type testDTO struct {
		ID  string `json:"id"`
		Foo string `json:"foo"`
	}

	body := []byte(`[{
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
	]`)

	_, err := unmarshal[testDTO](body) // <-- Parametrized as single object but it contains a list
	assert.Error(t, err)

	body = []byte(`{
		"id":"123",
		"foo":"something" 
	}`)

	_, err = unmarshal[[]testDTO](body) // <-- Parametrized as list but it contains a single object
	assert.Error(t, err)
}

func TestAttemptUnmarshal(t *testing.T) {
	type testDTO struct {
		ID  string `json:"id"`
		Foo string `json:"foo"`
	}

	bodyMultiple := io.NopCloser(bytes.NewReader([]byte(`[{
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

	singular, dto, dtos, err := AttemptUnmarshal[testDTO](bodyMultiple)

	require.NoError(t, err)

	assert.False(t, singular)
	assert.Nil(t, dto)
	assert.Len(t, dtos, 3)
}
