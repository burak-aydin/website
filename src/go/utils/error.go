package utils

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"strings"
)

type ErrorJson struct {
	Error string `json:"error"`
}

func HandleError(err error) error {
	if err == nil {
		return nil
	}
	var validationErrs validator.ValidationErrors
	ok := errors.As(err, &validationErrs)
	if !ok {
		log.Error("error is not a validation error")
		return err
	}

	var errs []string
	for _, valErr := range validationErrs {
		errs = append(errs, fmt.Sprintf("%s failed on %s", strings.ToLower(valErr.Field()), valErr.Tag()))
	}

	return fmt.Errorf("validation err: %s", strings.Join(errs, ","))
}
