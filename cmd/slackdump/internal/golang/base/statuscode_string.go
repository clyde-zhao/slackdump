// Code generated by "stringer -type StatusCode -trimprefix S"; DO NOT EDIT.

package base

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SNoError-0]
	_ = x[SGenericError-1]
	_ = x[SHelpRequested-2]
	_ = x[SInvalidParameters-3]
	_ = x[SAuthError-4]
	_ = x[SInitializationError-5]
	_ = x[SApplicationError-6]
	_ = x[SWorkspaceError-7]
	_ = x[SCacheError-8]
	_ = x[SUserError-9]
}

const _StatusCode_name = "NoErrorGenericErrorHelpRequestedInvalidParametersAuthErrorInitializationErrorApplicationErrorWorkspaceErrorCacheErrorUserError"

var _StatusCode_index = [...]uint8{0, 7, 19, 32, 49, 58, 77, 93, 107, 117, 126}

func (i StatusCode) String() string {
	if i >= StatusCode(len(_StatusCode_index)-1) {
		return "StatusCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _StatusCode_name[_StatusCode_index[i]:_StatusCode_index[i+1]]
}
