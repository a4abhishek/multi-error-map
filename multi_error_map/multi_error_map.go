package multi_error_map

import "fmt"

//
type MultiErrorMap struct {
	errors map[string]error
}

func NewMultiErrorMap() *MultiErrorMap {
	return &MultiErrorMap{errors: map[string]error{}}
}

func (c *MultiErrorMap) AddError(key string, e error) { c.errors[key] = e }

func (c *MultiErrorMap) Error() (err string) {
	for k, e := range c.errors {
		err += fmt.Sprintf("\t%s: %s\n", k, e.Error())
	}

	return err
}
