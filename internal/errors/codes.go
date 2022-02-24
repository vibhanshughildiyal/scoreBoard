package serviceerrors

type ErrorDetail struct {
	Description       string
	DetailDescription string
}

var (
	ErrCodes = map[Code]ErrorDetail{
		/*
			Parse Error
		*/
		"1.0": {"parsing error", "parsing error"},
	}
)
