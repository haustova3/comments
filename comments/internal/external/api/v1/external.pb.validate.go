// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: external.proto

package external

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on CheckUserIDRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CheckUserIDRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CheckUserIDRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CheckUserIDRequestMultiError, or nil if none found.
func (m *CheckUserIDRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CheckUserIDRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetUserID() <= 0 {
		err := CheckUserIDRequestValidationError{
			field:  "UserID",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CheckUserIDRequestMultiError(errors)
	}

	return nil
}

// CheckUserIDRequestMultiError is an error wrapping multiple validation errors
// returned by CheckUserIDRequest.ValidateAll() if the designated constraints
// aren't met.
type CheckUserIDRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CheckUserIDRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CheckUserIDRequestMultiError) AllErrors() []error { return m }

// CheckUserIDRequestValidationError is the validation error returned by
// CheckUserIDRequest.Validate if the designated constraints aren't met.
type CheckUserIDRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckUserIDRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckUserIDRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckUserIDRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckUserIDRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckUserIDRequestValidationError) ErrorName() string {
	return "CheckUserIDRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CheckUserIDRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckUserIDRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckUserIDRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckUserIDRequestValidationError{}

// Validate checks the field values on CheckUserIDResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CheckUserIDResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CheckUserIDResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CheckUserIDResponseMultiError, or nil if none found.
func (m *CheckUserIDResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CheckUserIDResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for IsCorrect

	if len(errors) > 0 {
		return CheckUserIDResponseMultiError(errors)
	}

	return nil
}

// CheckUserIDResponseMultiError is an error wrapping multiple validation
// errors returned by CheckUserIDResponse.ValidateAll() if the designated
// constraints aren't met.
type CheckUserIDResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CheckUserIDResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CheckUserIDResponseMultiError) AllErrors() []error { return m }

// CheckUserIDResponseValidationError is the validation error returned by
// CheckUserIDResponse.Validate if the designated constraints aren't met.
type CheckUserIDResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckUserIDResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckUserIDResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckUserIDResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckUserIDResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckUserIDResponseValidationError) ErrorName() string {
	return "CheckUserIDResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CheckUserIDResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckUserIDResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckUserIDResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckUserIDResponseValidationError{}

// Validate checks the field values on GetOwnerRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetOwnerRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetOwnerRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetOwnerRequestMultiError, or nil if none found.
func (m *GetOwnerRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetOwnerRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetProductID() <= 0 {
		err := GetOwnerRequestValidationError{
			field:  "ProductID",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetOwnerRequestMultiError(errors)
	}

	return nil
}

// GetOwnerRequestMultiError is an error wrapping multiple validation errors
// returned by GetOwnerRequest.ValidateAll() if the designated constraints
// aren't met.
type GetOwnerRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetOwnerRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetOwnerRequestMultiError) AllErrors() []error { return m }

// GetOwnerRequestValidationError is the validation error returned by
// GetOwnerRequest.Validate if the designated constraints aren't met.
type GetOwnerRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetOwnerRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetOwnerRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetOwnerRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetOwnerRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetOwnerRequestValidationError) ErrorName() string { return "GetOwnerRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetOwnerRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetOwnerRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetOwnerRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetOwnerRequestValidationError{}

// Validate checks the field values on GetOwnerResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetOwnerResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetOwnerResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetOwnerResponseMultiError, or nil if none found.
func (m *GetOwnerResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetOwnerResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for OwnerID

	if len(errors) > 0 {
		return GetOwnerResponseMultiError(errors)
	}

	return nil
}

// GetOwnerResponseMultiError is an error wrapping multiple validation errors
// returned by GetOwnerResponse.ValidateAll() if the designated constraints
// aren't met.
type GetOwnerResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetOwnerResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetOwnerResponseMultiError) AllErrors() []error { return m }

// GetOwnerResponseValidationError is the validation error returned by
// GetOwnerResponse.Validate if the designated constraints aren't met.
type GetOwnerResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetOwnerResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetOwnerResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetOwnerResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetOwnerResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetOwnerResponseValidationError) ErrorName() string { return "GetOwnerResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetOwnerResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetOwnerResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetOwnerResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetOwnerResponseValidationError{}
