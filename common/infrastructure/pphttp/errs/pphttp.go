package pphttp_errs

import "common/utils/pperrs"

type CustomHTTPError struct {
	*pperrs.CustomError
}
