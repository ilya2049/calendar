package errors

import "errors"

func ExtractCode(err error) Code {
	for err != nil {
		if errorWithCode, ok := err.(interface{ Code() Code }); ok {
			return errorWithCode.Code()
		}

		err = errors.Unwrap(err)
	}

	return 0
}

func ExtractDetails(err error) map[string]string {
	if err == nil {
		return nil
	}

	details := make(map[string]string)

	for err != nil {
		if errorWithDetails, ok := err.(interface{ Details() map[string]string }); ok {
			for key, detail := range errorWithDetails.Details() {
				details[key] = detail
			}
		}

		err = errors.Unwrap(err)
	}

	return details
}
