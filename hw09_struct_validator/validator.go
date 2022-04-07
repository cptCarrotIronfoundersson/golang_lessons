package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type ValidationError struct {
	Field string
	Err   error
}

var (
	ErrInvalidLen    = errors.New("invalid Length")
	ErrInvalidIn     = errors.New("invalid In")
	ErrInvalidMax    = errors.New("invalid Max")
	ErrInvalidMin    = errors.New("invalid Min")
	ErrInvalidRegexp = errors.New("invalid Regexp")
)

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	var sb strings.Builder
	for _, err := range v {
		fmt.Fprintf(&sb, "[f: %s, e: %v] ", err.Field, err.Err)
	}
	return sb.String()
}

type StructValidator interface {
	Validate(interface{}) []error
}

type Validator struct {
	// Тут можно забабахать всякие доп настройки, но впереди столько домашек, а времени нет
}

type intValidation struct {
	min int64
	max int64
	in  []int64
}

type stringValidation struct {
	len    int64
	regexp string
	in     []string
}

func (i Validator) PrepareIntValidation(tag string) (*intValidation, []error) {
	terms := strings.Split(tag, "|")
	var valTerms intValidation
	var validationErrors []error
	for _, term := range terms {
		splitedTag := strings.Split(term, ":")

		tagExp := splitedTag[0]
		tagValue := splitedTag[1]

		switch {
		case tagExp == "min":
			val, err := strconv.Atoi(tagValue)
			if err != nil {
				validationErrors = append(validationErrors, err)
			}
			valTerms.min = int64(val)

		case tagExp == "max":
			val, err := strconv.Atoi(tagValue)
			if err != nil {
				validationErrors = append(validationErrors, err)
			}
			valTerms.max = int64(val)

		case tagExp == "in":
			var inValues []int64
			for _, val := range strings.Split(tagValue, ",") {
				intVal, err := strconv.Atoi(val)
				if err != nil {
					validationErrors = append(validationErrors)
					continue
				}
				inValues = append(inValues, int64(intVal))
			}
			valTerms.in = inValues
		}
	}
	return &valTerms, validationErrors
}

func (i Validator) PrepareStringValidation(tag string) (*stringValidation, []error) {
	terms := strings.Split(tag, "|")
	var valTerms stringValidation
	var validationErrors []error
	for _, term := range terms {
		splitedTag := strings.Split(term, ":")

		tagExp := splitedTag[0]
		tagValue := splitedTag[1]

		switch {
		case tagExp == "len":
			val, err := strconv.Atoi(tagValue)
			if err != nil {
				validationErrors = append(validationErrors, err)
			}
			valTerms.len = int64(val)

		case tagExp == "regexp":
			valTerms.regexp = tagValue

		case tagExp == "in":
			var inValues []string
			for _, val := range strings.Split(tagValue, ",") {
				inValues = append(inValues, val)
			}
			valTerms.in = inValues
		}
	}
	return &valTerms, validationErrors
}

func (i Validator) Validate(StructToValidate interface{}) error {
	var vErr ValidationErrors
	Value := reflect.ValueOf(StructToValidate)
	valueType := Value.Type()
	for d := 0; d < valueType.NumField(); d++ {

		if valueType.Field(d).Tag.Get("validate") == "" {
			continue
		} else if Value.Field(d).Kind() == reflect.Int {
			valTerms, PrepareErrors := i.PrepareIntValidation(valueType.Field(d).Tag.Get("validate")) // use validation errors
			for _, err := range PrepareErrors {
				vErr = append(vErr, ValidationError{
					Field: valueType.Field(d).Name,
					Err:   err,
				})
			}
			err := valTerms.validateMin(Value.Field(d).Int())
			if err != nil {
				vErr = append(vErr, ValidationError{
					Field: valueType.Field(d).Name,
					Err:   err,
				})
			}
			err = valTerms.validateMax(Value.Field(d).Int())
			if err != nil {
				vErr = append(vErr, ValidationError{
					Field: valueType.Field(d).Name,
					Err:   err,
				})
			}
			err = valTerms.validateIn(Value.Field(d).Int())
			if err != nil {
				vErr = append(vErr, ValidationError{
					Field: valueType.Field(d).Name,
					Err:   err,
				})
			}
		} else if Value.Field(d).Kind() == reflect.String {
			valTerms, PrepareErrors := i.PrepareStringValidation(valueType.Field(d).Tag.Get("validate"))
			for _, err := range PrepareErrors {
				vErr = append(vErr, ValidationError{
					Field: valueType.Field(d).Name,
					Err:   err,
				})
			}
			err := valTerms.validateLen(Value.Field(d).String())
			if err != nil {
				vErr = append(vErr, ValidationError{
					Field: valueType.Field(d).Name,
					Err:   err,
				})
			}
			err = valTerms.validateRegexp(Value.Field(d).String())
			if err != nil {
				vErr = append(vErr, ValidationError{
					Field: valueType.Field(d).Name,
					Err:   err,
				})
			}
			err = valTerms.validateIn(Value.Field(d).String())
			if err != nil {
				vErr = append(vErr, ValidationError{
					Field: valueType.Field(d).Name,
					Err:   err,
				})
			}
		} else if Value.Field(d).Kind() == reflect.Slice {
			s := reflect.ValueOf(Value.Field(d))
			fmt.Println(Value.Field(d).Kind(), valueType.Field(d).Name, reflect.TypeOf(s), &s, s)
			for sl := 0; sl < s.Len(); sl++ {
				fmt.Println(s.Index(d))
			}
		}
	}
	if len(vErr) != 0 {
		return vErr
	}
	return nil
}

func (i intValidation) validateMin(val int64) error {
	if val < i.min && i.min > 0 {
		return ErrInvalidMin
	}
	return nil
}

func (i intValidation) validateMax(val int64) error {
	if val > i.max && i.max > 0 {
		return ErrInvalidMax
	}
	return nil
}

func (i intValidation) validateIn(val int64) error {
	var ValueIn bool
	if len(i.in) < 1 {
		return nil
	}
	for _, value := range i.in {
		if value == val {
			ValueIn = true
			break
		}
	}
	if ValueIn == false {
		return ErrInvalidIn
	}
	return nil
}

func (i stringValidation) validateLen(val string) error {
	if int64(len(val)) != i.len && i.len > 0 {
		return ErrInvalidLen
	}
	return nil
}

func (i stringValidation) validateRegexp(val string) error {
	matched, err := regexp.Match(i.regexp, []byte(val))
	if err != nil {
		return err
	}
	if !matched {
		return ErrInvalidRegexp
	}
	return nil
}

func (i stringValidation) validateIn(val string) error {
	var ValueIn bool
	if len(i.in) < 1 {
		return nil
	}
	for _, validVal := range i.in {
		if validVal == val {
			ValueIn = true
			break
		}
	}
	if ValueIn == false {
		return ErrInvalidIn
	}
	return nil
}
