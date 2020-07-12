// generated by jsonenums -type=Kind; DO NOT EDIT

package main

import (
	"encoding/json"
	"fmt"
)

var (
	_KindNameToValue = map[string]Kind{
		"UI_LAUNCHER_DISPLAY":     UI_LAUNCHER_DISPLAY,
		"UI_LAUNCHER_HIDE":        UI_LAUNCHER_HIDE,
		"THEME_SWITCH":            THEME_SWITCH,
		"APPLICATION_GET":         APPLICATION_GET,
		"APPLICATION_GET_SUCCESS": APPLICATION_GET_SUCCESS,
		"MODULE_GET":              MODULE_GET,
		"MODULE_GET_SUCCESS":      MODULE_GET_SUCCESS,
		"ERROR":                   ERROR,
	}

	_KindValueToName = map[Kind]string{
		UI_LAUNCHER_DISPLAY:     "UI_LAUNCHER_DISPLAY",
		UI_LAUNCHER_HIDE:        "UI_LAUNCHER_HIDE",
		THEME_SWITCH:            "THEME_SWITCH",
		APPLICATION_GET:         "APPLICATION_GET",
		APPLICATION_GET_SUCCESS: "APPLICATION_GET_SUCCESS",
		MODULE_GET:              "MODULE_GET",
		MODULE_GET_SUCCESS:      "MODULE_GET_SUCCESS",
		ERROR:                   "ERROR",
	}
)

func init() {
	var v Kind
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_KindNameToValue = map[string]Kind{
			interface{}(UI_LAUNCHER_DISPLAY).(fmt.Stringer).String():     UI_LAUNCHER_DISPLAY,
			interface{}(UI_LAUNCHER_HIDE).(fmt.Stringer).String():        UI_LAUNCHER_HIDE,
			interface{}(THEME_SWITCH).(fmt.Stringer).String():            THEME_SWITCH,
			interface{}(APPLICATION_GET).(fmt.Stringer).String():         APPLICATION_GET,
			interface{}(APPLICATION_GET_SUCCESS).(fmt.Stringer).String(): APPLICATION_GET_SUCCESS,
			interface{}(MODULE_GET).(fmt.Stringer).String():              MODULE_GET,
			interface{}(MODULE_GET_SUCCESS).(fmt.Stringer).String():      MODULE_GET_SUCCESS,
			interface{}(ERROR).(fmt.Stringer).String():                   ERROR,
		}
	}
}

// MarshalJSON is generated so Kind satisfies json.Marshaler.
func (r Kind) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _KindValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid Kind: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so Kind satisfies json.Unmarshaler.
func (r *Kind) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Kind should be a string, got %s", data)
	}
	v, ok := _KindNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid Kind %q", s)
	}
	*r = v
	return nil
}
