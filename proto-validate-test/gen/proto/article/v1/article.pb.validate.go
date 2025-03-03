// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/article/v1/article.proto

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

// Validate checks the field values on Article with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Article) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Article with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ArticleMultiError, or nil if none found.
func (m *Article) ValidateAll() error {
	return m.validate(true)
}

func (m *Article) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Title

	// no validation rules for Author

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	if len(errors) > 0 {
		return ArticleMultiError(errors)
	}

	return nil
}

// ArticleMultiError is an error wrapping multiple validation errors returned
// by Article.ValidateAll() if the designated constraints aren't met.
type ArticleMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ArticleMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ArticleMultiError) AllErrors() []error { return m }

// ArticleValidationError is the validation error returned by Article.Validate
// if the designated constraints aren't met.
type ArticleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ArticleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ArticleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ArticleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ArticleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ArticleValidationError) ErrorName() string { return "ArticleValidationError" }

// Error satisfies the builtin error interface
func (e ArticleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sArticle.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ArticleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ArticleValidationError{}

// Validate checks the field values on GetArticleRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetArticleRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetArticleRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetArticleRequestMultiError, or nil if none found.
func (m *GetArticleRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetArticleRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetArticleRequestMultiError(errors)
	}

	return nil
}

// GetArticleRequestMultiError is an error wrapping multiple validation errors
// returned by GetArticleRequest.ValidateAll() if the designated constraints
// aren't met.
type GetArticleRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetArticleRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetArticleRequestMultiError) AllErrors() []error { return m }

// GetArticleRequestValidationError is the validation error returned by
// GetArticleRequest.Validate if the designated constraints aren't met.
type GetArticleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetArticleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetArticleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetArticleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetArticleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetArticleRequestValidationError) ErrorName() string {
	return "GetArticleRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetArticleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetArticleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetArticleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetArticleRequestValidationError{}

// Validate checks the field values on GetArticleResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetArticleResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetArticleResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetArticleResponseMultiError, or nil if none found.
func (m *GetArticleResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetArticleResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetArticle()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetArticleResponseValidationError{
					field:  "Article",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetArticleResponseValidationError{
					field:  "Article",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetArticle()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetArticleResponseValidationError{
				field:  "Article",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetArticleResponseMultiError(errors)
	}

	return nil
}

// GetArticleResponseMultiError is an error wrapping multiple validation errors
// returned by GetArticleResponse.ValidateAll() if the designated constraints
// aren't met.
type GetArticleResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetArticleResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetArticleResponseMultiError) AllErrors() []error { return m }

// GetArticleResponseValidationError is the validation error returned by
// GetArticleResponse.Validate if the designated constraints aren't met.
type GetArticleResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetArticleResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetArticleResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetArticleResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetArticleResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetArticleResponseValidationError) ErrorName() string {
	return "GetArticleResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetArticleResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetArticleResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetArticleResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetArticleResponseValidationError{}

// Validate checks the field values on CreateArticleInput with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateArticleInput) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateArticleInput with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateArticleInputMultiError, or nil if none found.
func (m *CreateArticleInput) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateArticleInput) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetId()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateArticleInputValidationError{
					field:  "Id",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateArticleInputValidationError{
					field:  "Id",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateArticleInputValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Title

	// no validation rules for Author

	if len(errors) > 0 {
		return CreateArticleInputMultiError(errors)
	}

	return nil
}

// CreateArticleInputMultiError is an error wrapping multiple validation errors
// returned by CreateArticleInput.ValidateAll() if the designated constraints
// aren't met.
type CreateArticleInputMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateArticleInputMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateArticleInputMultiError) AllErrors() []error { return m }

// CreateArticleInputValidationError is the validation error returned by
// CreateArticleInput.Validate if the designated constraints aren't met.
type CreateArticleInputValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateArticleInputValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateArticleInputValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateArticleInputValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateArticleInputValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateArticleInputValidationError) ErrorName() string {
	return "CreateArticleInputValidationError"
}

// Error satisfies the builtin error interface
func (e CreateArticleInputValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateArticleInput.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateArticleInputValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateArticleInputValidationError{}

// Validate checks the field values on CreateArticlesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateArticlesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateArticlesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateArticlesRequestMultiError, or nil if none found.
func (m *CreateArticlesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateArticlesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetInputs() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CreateArticlesRequestValidationError{
						field:  fmt.Sprintf("Inputs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CreateArticlesRequestValidationError{
						field:  fmt.Sprintf("Inputs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CreateArticlesRequestValidationError{
					field:  fmt.Sprintf("Inputs[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CreateArticlesRequestMultiError(errors)
	}

	return nil
}

// CreateArticlesRequestMultiError is an error wrapping multiple validation
// errors returned by CreateArticlesRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateArticlesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateArticlesRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateArticlesRequestMultiError) AllErrors() []error { return m }

// CreateArticlesRequestValidationError is the validation error returned by
// CreateArticlesRequest.Validate if the designated constraints aren't met.
type CreateArticlesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateArticlesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateArticlesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateArticlesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateArticlesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateArticlesRequestValidationError) ErrorName() string {
	return "CreateArticlesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateArticlesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateArticlesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateArticlesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateArticlesRequestValidationError{}

// Validate checks the field values on CreateArticlesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateArticlesResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateArticlesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateArticlesResponseMultiError, or nil if none found.
func (m *CreateArticlesResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateArticlesResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetArticles() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CreateArticlesResponseValidationError{
						field:  fmt.Sprintf("Articles[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CreateArticlesResponseValidationError{
						field:  fmt.Sprintf("Articles[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CreateArticlesResponseValidationError{
					field:  fmt.Sprintf("Articles[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CreateArticlesResponseMultiError(errors)
	}

	return nil
}

// CreateArticlesResponseMultiError is an error wrapping multiple validation
// errors returned by CreateArticlesResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateArticlesResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateArticlesResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateArticlesResponseMultiError) AllErrors() []error { return m }

// CreateArticlesResponseValidationError is the validation error returned by
// CreateArticlesResponse.Validate if the designated constraints aren't met.
type CreateArticlesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateArticlesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateArticlesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateArticlesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateArticlesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateArticlesResponseValidationError) ErrorName() string {
	return "CreateArticlesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateArticlesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateArticlesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateArticlesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateArticlesResponseValidationError{}
