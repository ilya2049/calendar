package errors_test

import (
	"calendar/internal/pkg/errors"

	stdErrors "errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractDetails(t *testing.T) {
	newTestError := func() error {
		err := errors.New(stdErrors.New("message 1")).
			SetDetail("detail 1 key", "detail 1")

		err = errors.Wrapf("message 2: %w", err).
			SetDetail("detail 2 key", "detail 2").
			SetDetailf("detail 3 key", "detail %d", 3).
			SetDetail("", "").
			SetDetailf("", "detail 4")

		return err
	}

	wantDetails := map[string]string{
		"detail 1 key": "detail 1",
		"detail 2 key": "detail 2",
		"detail 3 key": "detail 3",
	}

	assert.Equal(t, wantDetails, errors.ExtractDetails(newTestError()))
}
