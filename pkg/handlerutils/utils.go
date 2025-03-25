package handlerutils

import (
	"encoding/json"
	"fmt"
	"log/slog"
)

// Function which attempts to unmarshal both as a struct and as a list of structs.
// Useful for getting data from a post request and figuring out if it's a multiple type or not in the process
func AttemptUnmarshal[T any](data []byte) (singular bool, dto *T, dtos []T, err error) {

	item, singularErr := Unmarshal[T](data)
	list, multipleErr := Unmarshal[[]T](data)

	if singularErr == nil && multipleErr == nil {
		return false, nil, nil, fmt.Errorf("this shouldn't be possible")
	} else if singularErr != nil && multipleErr != nil {
		slog.Error("unable to unmarshal body", "singularErr", singularErr, "multipleErr", multipleErr)
		return false, nil, nil, singularErr
	}

	singular = list == nil

	return singular, item, *list, nil
}

func Unmarshal[T any](data []byte) (*T, error) {
	var unmarshalled T
	if err := json.Unmarshal(data, &unmarshalled); err != nil {
		return nil, err
	}

	fmt.Println(unmarshalled)
	return &unmarshalled, nil
}
