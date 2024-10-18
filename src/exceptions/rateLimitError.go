package exceptions

type RateLimitError struct {
	Message    string
	RetryAfter string
}

func NewRateLimitError(retryAfter string) RateLimitError {
	return RateLimitError{
		Message:    "Rate limit exceeded. Please try again later.",
		RetryAfter: retryAfter,
	}
}

func (e RateLimitError) Error() string {
	return e.Message
}
