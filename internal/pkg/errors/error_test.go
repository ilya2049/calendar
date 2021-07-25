package errors_test

import (
	"calendar/internal/pkg/errors"

	stderrors "errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_genericError_Error(t *testing.T) {
	newTestError := func() error {
		err := stderrors.New("message 1")

		err = errors.Wrapf("message 2: %w", err)

		err = errors.Wrapf("message 3: %w", err)

		return err
	}

	wantErrorMessage := "message 3: message 2: message 1"

	assert.Equal(t, wantErrorMessage, newTestError().Error())
}
