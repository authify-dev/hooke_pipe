package pphttp_errs

import "common/utils/pperrs"

type RequestOptions struct {
	Method string
	URL    string
}

type MakingRequestError struct {
	CustomHTTPError
}

func NewMakingRequestError(options RequestOptions) *MakingRequestError {
	return &MakingRequestError{
		CustomHTTPError: CustomHTTPError{
			CustomError: &pperrs.CustomError{
				Code:    500,
				Message: "error making the request to " + options.URL + "." + options.Method,
				Scope:   "konectus.knhttp.MakingRequestError",
			},
		},
	}
}

type DoingRequestError struct {
	CustomHTTPError
}

func NewDoingRequestError(options RequestOptions) *DoingRequestError {
	return &DoingRequestError{
		CustomHTTPError: CustomHTTPError{
			CustomError: &pperrs.CustomError{
				Code:    500,
				Message: "error doing the request to " + options.URL + "." + options.Method,
				Scope:   "konectus.knhttp.DoingRequestError",
			},
		},
	}
}
