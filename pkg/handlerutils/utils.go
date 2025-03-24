package handlerutils

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
)

// Function which attempts to unmarshal both as a struct and as a list of structs.
// Useful for getting data from a post request and figuring out if it's a multiple type or not in the process
func AttemptUnmarshal[T any](body io.ReadCloser) (singular bool, dto *T, dtos []T, err error) {

	data, err := io.ReadAll(body)
	if err != nil {
		return false, nil, nil, err
	}

	item, singularErr := unmarshal[T](data)
	list, multipleErr := unmarshal[[]T](data)

	if singularErr == nil && multipleErr == nil {
		return false, nil, nil, fmt.Errorf("this shouldn't be possible")
	} else if singularErr != nil && multipleErr != nil {
		slog.Error("unable to unmarshal body", "singularErr", singularErr, "multipleErr", multipleErr)
		return false, nil, nil, singularErr
	}

	singular = list == nil

	return singular, item, *list, nil
}

func unmarshal[T any](data []byte) (*T, error) {
	var unmarshalled T
	if err := json.Unmarshal(data, &unmarshalled); err != nil {
		return nil, err
	}

	fmt.Println(unmarshalled)
	return &unmarshalled, nil
}
