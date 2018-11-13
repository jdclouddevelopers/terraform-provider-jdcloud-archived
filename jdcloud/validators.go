package jdcloud

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"strings"
)

func validateStringInSlice(validSlice []string, ignoreCase bool) schema.SchemaValidateFunc {
	return func(i interface{}, k string) (s []string, err []error) {
		v, ok := i.(string)
		if !ok {
			err = append(err, fmt.Errorf("type of %s must be string", k))
			return
		}

		for _, str := range validSlice {
			if v == str || (ignoreCase && strings.ToLower(v) == strings.ToLower(str)) {
				return
			}
		}

		err = append(err, fmt.Errorf("expected %s to be one of %v, got %s", k, validSlice, v))
		return
	}
}

func validateStringNoEmpty(v interface{}, k string) (s []string, errs []error) {

	value := v.(string)
	if len(value) < 1 {
		errs = append(errs, fmt.Errorf("%v can not be Empty characters.", k))
	}

	return
}
