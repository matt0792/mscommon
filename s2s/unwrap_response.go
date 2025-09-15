package s2s

import (
	"fmt"

	"github.com/matt0792/mscommon/commonmodels"
)

func UnwrapResponse[T any](resp commonmodels.Response[T], err error) (T, error) {
	var zero T

	if err != nil {
		return zero, err
	}

	if resp.Status != "success" {
		return zero, fmt.Errorf("service error: %s", resp.Message)
	}

	return resp.Data, nil
}
