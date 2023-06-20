package amiderrors

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

const COMMON_TYPE string = "common"

const (
	INTERNAL     = "internal"
	EMPTY_VALUES = "empty_values"
	EMPTY_VALUE  = "empty_value"
	WRONG_LEN    = "wrong_len"
)

// concatenation of etype + "_" + code
func MakeCode(etype string, code string) string {
	return etype + "_" + code
}

// config error for error mapping customization
type ConfigError interface {
	ErrorResponse(c *Config) *ErrorResponse
}

// base error config, parse from yaml file
type Config struct {
	Common map[string]string            `yaml:"common"`
	Errors map[string]map[string]string `yaml:"errors"`
}

// default config
var _DEFAULT *Config

func NewConfig() *Config {
	return &Config{make(map[string]string), make(map[string]map[string]string)}
}

// parse config from yaml file by path
func (c *Config) Parse(configPath string) error {
	file, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(file, c)
}

/*
common wrong length error implementation

	common.empty_value should contain ${param} to replace it with <param> value
	common.wrong_length should contain ${min} and ${max} to replace it with <min> and <max> value
*/
func (c *Config) WrongLength(param string, min, max int) *ErrorResponse {
	msg := c.Common[WRONG_LEN]
	replacer := strings.NewReplacer("${param}", param, "${min}", fmt.Sprint(min), "${max}", fmt.Sprint(max))
	msg = replacer.Replace(msg)
	return NewErrorResponse(msg, MakeCode(COMMON_TYPE, WRONG_LEN))
}

/*
common empty value error implementation

	common.empty_value should contain ${param} to replace it with <param> value
*/
func (c *Config) EmptyValue(param string) *ErrorResponse {
	msg := c.Common[EMPTY_VALUE]
	replacer := strings.NewReplacer("${param}", param)
	msg = replacer.Replace(msg)
	return NewErrorResponse(msg, MakeCode(COMMON_TYPE, EMPTY_VALUE))
}

/*
common empty valuess error implementation

	common.empty_value should contain ${param} to replace it with <param> value
*/
func (c *Config) EmptyValues(params ...string) *ErrorResponse {
	msg := c.Common[EMPTY_VALUES]
	replacer := strings.NewReplacer("${param}", strings.Join(params, ","))
	msg = replacer.Replace(msg)
	return NewErrorResponse(msg, MakeCode(COMMON_TYPE, EMPTY_VALUES))
}

// internal error implementation, return them not found error in config
func (c *Config) Internal() *ErrorResponse {
	msg := c.Common[INTERNAL]
	return NewErrorResponse(msg, MakeCode(COMMON_TYPE, INTERNAL))
}

// basic err to response map, if err is ConfigError return err.ErrorResponse(c)
func (c *Config) ErrorToResponse(err error) *ErrorResponse {
	switch err := err.(type) {
	case *Exception:
		return c.ExceptionToResponse(err)
	case *ErrorResponse:
		return err
	case ConfigError:
		return err.ErrorResponse(c)
	default:
		return c.Internal()
	}
}

// exception to response map method, if err with etype and code not found return Internal()
func (c *Config) ExceptionToResponse(exc *Exception) *ErrorResponse {
	if exc.Type == COMMON_TYPE {
		err, ok := c.Common[exc.Code]
		if !ok {
			return c.Internal()
		}
		return &ErrorResponse{Err: err, Code: MakeCode(COMMON_TYPE, exc.Code)}
	}
	typeErrors, ok := c.Errors[exc.Type]
	if !ok {
		return c.Internal()
	}
	err, ok := typeErrors[exc.Code]
	if !ok {
		return c.Internal()
	}
	return &ErrorResponse{Err: err, Code: MakeCode(exc.Type, exc.Code)}

}

// init config, set _DEFAULT
func Init(configPath string) {
	_DEFAULT = NewConfig()
	err := _DEFAULT.Parse(configPath)
	if err != nil {
		log.Fatalf("Couldn't parse err config, err is %v", err)
	}
}

var EMPTY_CONFIG = NewErrorResponse("error config is not initalize", INTERNAL)

func EmptyConfigInFunction(fname string) *ErrorResponse {
	return NewErrorResponse(fmt.Sprintf("funname is %s, error config is not initalize", fname), INTERNAL)
}

func EmptyValues(params ...string) error {
	if _DEFAULT == nil {
		return EmptyConfigInFunction("empty values")
	}
	return _DEFAULT.EmptyValues(params...)
}

func EmptyValue(param string) error {
	if _DEFAULT == nil {
		return EmptyConfigInFunction("empty value")
	}
	return _DEFAULT.EmptyValue(param)
}

func WrongLength(param string, min, max int) error {
	if _DEFAULT == nil {
		return EmptyConfigInFunction("wrong length")
	}
	return _DEFAULT.WrongLength(param, min, max)
}

func Internal() *ErrorResponse {
	if _DEFAULT == nil {
		return EmptyConfigInFunction("internal")
	}
	return _DEFAULT.Internal()
}

func ErrorToResponse(err error) *ErrorResponse {
	if _DEFAULT == nil {
		return EmptyConfigInFunction("error to response")
	}
	return _DEFAULT.ErrorToResponse(err)
}

func ExceptionToResponse(exc *Exception) *ErrorResponse {
	if _DEFAULT == nil {
		return EmptyConfigInFunction("exception to response")
	}
	return _DEFAULT.ExceptionToResponse(exc)
}

// return _DEFAULT value
func Default() *Config {
	return _DEFAULT
}
