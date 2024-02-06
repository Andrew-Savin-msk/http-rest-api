package model

import validation "github.com/go-ozzo/ozzo-validation"

// Custom validator createed because when we getting user out of table we dont have password field - which makes user non valid
// so we need to write our own validator
func requiredIf(cond bool) validation.RuleFunc {
	return func(value interface{}) error {
		if cond {
			return validation.Validate(value, validation.Required)
		}

		return nil
	}
}
