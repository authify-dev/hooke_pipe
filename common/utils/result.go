package utils

import (
	"common/utils/pperrs"
)

// --------------------------------
// UTILS
// --------------------------------
// Utils
//--------------------------------

type Result[R any] struct {
	Data R
	Err  pperrs.CustomErrorInterface
}
