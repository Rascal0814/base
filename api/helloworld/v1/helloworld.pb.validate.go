// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/helloworld/v1/helloworld.proto

package v1

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

// Validate checks the field values on CreateHelloworldRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateHelloworldRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateHelloworldRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateHelloworldRequestMultiError, or nil if none found.
func (m *CreateHelloworldRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateHelloworldRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CreateHelloworldRequestMultiError(errors)
	}

	return nil
}

// CreateHelloworldRequestMultiError is an error wrapping multiple validation
// errors returned by CreateHelloworldRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateHelloworldRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateHelloworldRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateHelloworldRequestMultiError) AllErrors() []error { return m }

// CreateHelloworldRequestValidationError is the validation error returned by
// CreateHelloworldRequest.Validate if the designated constraints aren't met.
type CreateHelloworldRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateHelloworldRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateHelloworldRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateHelloworldRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateHelloworldRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateHelloworldRequestValidationError) ErrorName() string {
	return "CreateHelloworldRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateHelloworldRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateHelloworldRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateHelloworldRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateHelloworldRequestValidationError{}

// Validate checks the field values on CreateHelloworldReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateHelloworldReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateHelloworldReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateHelloworldReplyMultiError, or nil if none found.
func (m *CreateHelloworldReply) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateHelloworldReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CreateHelloworldReplyMultiError(errors)
	}

	return nil
}

// CreateHelloworldReplyMultiError is an error wrapping multiple validation
// errors returned by CreateHelloworldReply.ValidateAll() if the designated
// constraints aren't met.
type CreateHelloworldReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateHelloworldReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateHelloworldReplyMultiError) AllErrors() []error { return m }

// CreateHelloworldReplyValidationError is the validation error returned by
// CreateHelloworldReply.Validate if the designated constraints aren't met.
type CreateHelloworldReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateHelloworldReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateHelloworldReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateHelloworldReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateHelloworldReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateHelloworldReplyValidationError) ErrorName() string {
	return "CreateHelloworldReplyValidationError"
}

// Error satisfies the builtin error interface
func (e CreateHelloworldReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateHelloworldReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateHelloworldReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateHelloworldReplyValidationError{}

// Validate checks the field values on UpdateHelloworldRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateHelloworldRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateHelloworldRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateHelloworldRequestMultiError, or nil if none found.
func (m *UpdateHelloworldRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateHelloworldRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UpdateHelloworldRequestMultiError(errors)
	}

	return nil
}

// UpdateHelloworldRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateHelloworldRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateHelloworldRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateHelloworldRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateHelloworldRequestMultiError) AllErrors() []error { return m }

// UpdateHelloworldRequestValidationError is the validation error returned by
// UpdateHelloworldRequest.Validate if the designated constraints aren't met.
type UpdateHelloworldRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateHelloworldRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateHelloworldRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateHelloworldRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateHelloworldRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateHelloworldRequestValidationError) ErrorName() string {
	return "UpdateHelloworldRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateHelloworldRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateHelloworldRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateHelloworldRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateHelloworldRequestValidationError{}

// Validate checks the field values on UpdateHelloworldReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateHelloworldReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateHelloworldReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateHelloworldReplyMultiError, or nil if none found.
func (m *UpdateHelloworldReply) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateHelloworldReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UpdateHelloworldReplyMultiError(errors)
	}

	return nil
}

// UpdateHelloworldReplyMultiError is an error wrapping multiple validation
// errors returned by UpdateHelloworldReply.ValidateAll() if the designated
// constraints aren't met.
type UpdateHelloworldReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateHelloworldReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateHelloworldReplyMultiError) AllErrors() []error { return m }

// UpdateHelloworldReplyValidationError is the validation error returned by
// UpdateHelloworldReply.Validate if the designated constraints aren't met.
type UpdateHelloworldReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateHelloworldReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateHelloworldReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateHelloworldReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateHelloworldReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateHelloworldReplyValidationError) ErrorName() string {
	return "UpdateHelloworldReplyValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateHelloworldReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateHelloworldReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateHelloworldReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateHelloworldReplyValidationError{}

// Validate checks the field values on DeleteHelloworldRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteHelloworldRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteHelloworldRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteHelloworldRequestMultiError, or nil if none found.
func (m *DeleteHelloworldRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteHelloworldRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteHelloworldRequestMultiError(errors)
	}

	return nil
}

// DeleteHelloworldRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteHelloworldRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteHelloworldRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteHelloworldRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteHelloworldRequestMultiError) AllErrors() []error { return m }

// DeleteHelloworldRequestValidationError is the validation error returned by
// DeleteHelloworldRequest.Validate if the designated constraints aren't met.
type DeleteHelloworldRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteHelloworldRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteHelloworldRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteHelloworldRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteHelloworldRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteHelloworldRequestValidationError) ErrorName() string {
	return "DeleteHelloworldRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteHelloworldRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteHelloworldRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteHelloworldRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteHelloworldRequestValidationError{}

// Validate checks the field values on DeleteHelloworldReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteHelloworldReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteHelloworldReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteHelloworldReplyMultiError, or nil if none found.
func (m *DeleteHelloworldReply) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteHelloworldReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteHelloworldReplyMultiError(errors)
	}

	return nil
}

// DeleteHelloworldReplyMultiError is an error wrapping multiple validation
// errors returned by DeleteHelloworldReply.ValidateAll() if the designated
// constraints aren't met.
type DeleteHelloworldReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteHelloworldReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteHelloworldReplyMultiError) AllErrors() []error { return m }

// DeleteHelloworldReplyValidationError is the validation error returned by
// DeleteHelloworldReply.Validate if the designated constraints aren't met.
type DeleteHelloworldReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteHelloworldReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteHelloworldReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteHelloworldReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteHelloworldReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteHelloworldReplyValidationError) ErrorName() string {
	return "DeleteHelloworldReplyValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteHelloworldReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteHelloworldReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteHelloworldReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteHelloworldReplyValidationError{}

// Validate checks the field values on GetHelloworldRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetHelloworldRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetHelloworldRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetHelloworldRequestMultiError, or nil if none found.
func (m *GetHelloworldRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetHelloworldRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetHelloworldRequestMultiError(errors)
	}

	return nil
}

// GetHelloworldRequestMultiError is an error wrapping multiple validation
// errors returned by GetHelloworldRequest.ValidateAll() if the designated
// constraints aren't met.
type GetHelloworldRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetHelloworldRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetHelloworldRequestMultiError) AllErrors() []error { return m }

// GetHelloworldRequestValidationError is the validation error returned by
// GetHelloworldRequest.Validate if the designated constraints aren't met.
type GetHelloworldRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetHelloworldRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetHelloworldRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetHelloworldRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetHelloworldRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetHelloworldRequestValidationError) ErrorName() string {
	return "GetHelloworldRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetHelloworldRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetHelloworldRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetHelloworldRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetHelloworldRequestValidationError{}

// Validate checks the field values on GetHelloworldReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetHelloworldReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetHelloworldReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetHelloworldReplyMultiError, or nil if none found.
func (m *GetHelloworldReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetHelloworldReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if len(errors) > 0 {
		return GetHelloworldReplyMultiError(errors)
	}

	return nil
}

// GetHelloworldReplyMultiError is an error wrapping multiple validation errors
// returned by GetHelloworldReply.ValidateAll() if the designated constraints
// aren't met.
type GetHelloworldReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetHelloworldReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetHelloworldReplyMultiError) AllErrors() []error { return m }

// GetHelloworldReplyValidationError is the validation error returned by
// GetHelloworldReply.Validate if the designated constraints aren't met.
type GetHelloworldReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetHelloworldReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetHelloworldReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetHelloworldReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetHelloworldReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetHelloworldReplyValidationError) ErrorName() string {
	return "GetHelloworldReplyValidationError"
}

// Error satisfies the builtin error interface
func (e GetHelloworldReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetHelloworldReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetHelloworldReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetHelloworldReplyValidationError{}

// Validate checks the field values on ListHelloworldRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListHelloworldRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListHelloworldRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListHelloworldRequestMultiError, or nil if none found.
func (m *ListHelloworldRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListHelloworldRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListHelloworldRequestMultiError(errors)
	}

	return nil
}

// ListHelloworldRequestMultiError is an error wrapping multiple validation
// errors returned by ListHelloworldRequest.ValidateAll() if the designated
// constraints aren't met.
type ListHelloworldRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListHelloworldRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListHelloworldRequestMultiError) AllErrors() []error { return m }

// ListHelloworldRequestValidationError is the validation error returned by
// ListHelloworldRequest.Validate if the designated constraints aren't met.
type ListHelloworldRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListHelloworldRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListHelloworldRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListHelloworldRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListHelloworldRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListHelloworldRequestValidationError) ErrorName() string {
	return "ListHelloworldRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListHelloworldRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListHelloworldRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListHelloworldRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListHelloworldRequestValidationError{}

// Validate checks the field values on ListHelloworldReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListHelloworldReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListHelloworldReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListHelloworldReplyMultiError, or nil if none found.
func (m *ListHelloworldReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ListHelloworldReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListHelloworldReplyMultiError(errors)
	}

	return nil
}

// ListHelloworldReplyMultiError is an error wrapping multiple validation
// errors returned by ListHelloworldReply.ValidateAll() if the designated
// constraints aren't met.
type ListHelloworldReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListHelloworldReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListHelloworldReplyMultiError) AllErrors() []error { return m }

// ListHelloworldReplyValidationError is the validation error returned by
// ListHelloworldReply.Validate if the designated constraints aren't met.
type ListHelloworldReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListHelloworldReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListHelloworldReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListHelloworldReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListHelloworldReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListHelloworldReplyValidationError) ErrorName() string {
	return "ListHelloworldReplyValidationError"
}

// Error satisfies the builtin error interface
func (e ListHelloworldReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListHelloworldReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListHelloworldReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListHelloworldReplyValidationError{}
