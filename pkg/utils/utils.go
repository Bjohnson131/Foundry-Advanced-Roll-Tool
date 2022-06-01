package utils

import (
	"strconv"
)

func Int(i int) *int {
	return &i
}

type DynamicInt string

// String returns the literal text of the number.
func (n DynamicInt) String() string { return string(n) }

// Int64 returns the number as an int64.
func (n DynamicInt) Int64() (int64, error) {
	return strconv.ParseInt(string(n), 10, 64)
}

func (w *DynamicInt) UnmarshalJSON(data []byte) (err error) {
	if len(data) == 0 {
		var v DynamicInt = "0"
		*w = v
		return nil
	}

	if len(data) == 2 && data[0] == '"' && data[1] == '"' {
		var v DynamicInt = "0"
		*w = v
		return nil
	}

	if len(data) > 1 && data[0] == '"' && data[len(data)-1] == '"' {
		data = data[1 : len(data)-1]
	}

	if _, err := strconv.Atoi(string(data)); err != nil {
		return err
	}
	*w = DynamicInt(string(data))
	return nil
}

type DynamicFloat string

// String returns the literal text of the number.
func (n DynamicFloat) String() string { return string(n) }

// Float64 returns the number as a float64.
func (n DynamicFloat) Float64() (float64, error) {
	return strconv.ParseFloat(string(n), 64)
}

func (w *DynamicFloat) UnmarshalJSON(data []byte) (err error) {
	if len(data) == 0 {
		var v DynamicFloat = "0.0"
		*w = v
		return nil
	}

	if len(data) == 2 && data[0] == '"' && data[1] == '"' {
		var v DynamicFloat = "0.0"
		*w = v
		return nil
	}

	if len(data) > 1 && data[0] == '"' && data[len(data)-1] == '"' {
		data = data[1 : len(data)-1]
	}

	if _, err := strconv.ParseFloat(string(data), 64); err != nil {
		return err
	}
	*w = DynamicFloat(string(data))
	return nil
}
