// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: dsl.proto

package def

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on PopulationOption with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *PopulationOption) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Related

	// no validation rules for WithCounts

	// no validation rules for WithResources

	if v, ok := interface{}(m.GetHasRelationsWith()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PopulationOptionValidationError{
				field:  "HasRelationsWith",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// PopulationOptionValidationError is the validation error returned by
// PopulationOption.Validate if the designated constraints aren't met.
type PopulationOptionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PopulationOptionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PopulationOptionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PopulationOptionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PopulationOptionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PopulationOptionValidationError) ErrorName() string { return "PopulationOptionValidationError" }

// Error satisfies the builtin error interface
func (e PopulationOptionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPopulationOption.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PopulationOptionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PopulationOptionValidationError{}

// Validate checks the field values on EntityQuery with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *EntityQuery) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Qt

	// no validation rules for Type

	// no validation rules for Id

	if v, ok := interface{}(m.GetPopOp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EntityQueryValidationError{
				field:  "PopOp",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EntityQueryValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Action

	for idx, item := range m.GetQueries() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EntityQueryValidationError{
					field:  fmt.Sprintf("Queries[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Sorts

	// no validation rules for FromID

	// no validation rules for WithLimit

	return nil
}

// EntityQueryValidationError is the validation error returned by
// EntityQuery.Validate if the designated constraints aren't met.
type EntityQueryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EntityQueryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EntityQueryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EntityQueryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EntityQueryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EntityQueryValidationError) ErrorName() string { return "EntityQueryValidationError" }

// Error satisfies the builtin error interface
func (e EntityQueryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEntityQuery.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EntityQueryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EntityQueryValidationError{}

// Validate checks the field values on EntityCreate with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *EntityCreate) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Type

	if v, ok := interface{}(m.GetE_()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EntityCreateValidationError{
				field:  "E_",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetRelated_() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EntityCreateValidationError{
					field:  fmt.Sprintf("Related_[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetResources_() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EntityCreateValidationError{
					field:  fmt.Sprintf("Resources_[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EntityCreateValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Action

	return nil
}

// EntityCreateValidationError is the validation error returned by
// EntityCreate.Validate if the designated constraints aren't met.
type EntityCreateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EntityCreateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EntityCreateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EntityCreateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EntityCreateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EntityCreateValidationError) ErrorName() string { return "EntityCreateValidationError" }

// Error satisfies the builtin error interface
func (e EntityCreateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEntityCreate.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EntityCreateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EntityCreateValidationError{}

// Validate checks the field values on RelationQuery with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *RelationQuery) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Type

	if v, ok := interface{}(m.GetOfID()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RelationQueryValidationError{
				field:  "OfID",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Relation

	if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RelationQueryValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetPopOp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RelationQueryValidationError{
				field:  "PopOp",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for FromID

	// no validation rules for WithLimit

	return nil
}

// RelationQueryValidationError is the validation error returned by
// RelationQuery.Validate if the designated constraints aren't met.
type RelationQueryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RelationQueryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RelationQueryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RelationQueryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RelationQueryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RelationQueryValidationError) ErrorName() string { return "RelationQueryValidationError" }

// Error satisfies the builtin error interface
func (e RelationQueryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRelationQuery.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RelationQueryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RelationQueryValidationError{}

// Validate checks the field values on RelationCreate with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *RelationCreate) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetFrom_()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RelationCreateValidationError{
				field:  "From_",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetTo_()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RelationCreateValidationError{
				field:  "To_",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Verb_

	if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RelationCreateValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for IsAdd_

	return nil
}

// RelationCreateValidationError is the validation error returned by
// RelationCreate.Validate if the designated constraints aren't met.
type RelationCreateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RelationCreateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RelationCreateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RelationCreateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RelationCreateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RelationCreateValidationError) ErrorName() string { return "RelationCreateValidationError" }

// Error satisfies the builtin error interface
func (e RelationCreateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRelationCreate.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RelationCreateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RelationCreateValidationError{}
