package form

import "fmt"

func combineValidators(validators ...func([]string) string) func([]string) string {
	return func(paragraphs []string) string {
		for _, validator := range validators {
			result := validator(paragraphs)
			if result != "" {
				return result
			}
		}

		return ""
	}
}

func requiredValidator(name string) func([]string) string {
	return func(paragraphs []string) string {
		if len(paragraphs[0]) == 0 {
			return fmt.Sprintf("%s is required", name)
		}

		return ""
	}
}

func maxLengthValidator(name string, maxLength int) func([]string) string {
	return func(paragraphs []string) string {
		curLength := len(paragraphs[0])
		if curLength > maxLength {
			return fmt.Sprintf("%s length must be less than or equal to %d characters. Current length is %d characters.", name, maxLength, curLength)
		}

		return ""
	}
}
