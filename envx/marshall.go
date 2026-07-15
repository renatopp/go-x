package envx

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// EnvUnmarshaler is the interface implemented by types that can unmarshal
// themselves from a raw environment variable string. If a field's type (or its
// pointer) implements this interface, FromEnv is called instead of the
// built-in scalar logic — the same pattern as encoding/json.Unmarshaler.
//
// Example:
//
//	func (e *Environment) FromEnv(raw string) error {
//	    switch raw {
//	    case "local", "":
//	        *e = EnvLocal
//	    case "production":
//	        *e = EnvProduction
//	    default:
//	        return fmt.Errorf("unknown environment %q", raw)
//	    }
//	    return nil
//	}
type EnvUnmarshaler interface {
	FromEnv(raw string) error
}

// Unmarshal populates the exported fields of a struct pointer using environment
// variables declared with the `env:"VAR_NAME[,option...]"` struct tag.
//
// Supported types: string, bool, int family, uint family, float32/float64,
// time.Duration, slices of any of those types (comma-separated values), and any
// type that implements EnvUnmarshaler.
//
// Tag options (comma-separated after the key):
//   - required         — returns an error when the variable is not set
//   - default=<value>  — used when the variable is not set
//
// Examples:
//
//	Port    uint32      `env:"PORT,default=8080"`
//	Name    string      `env:"NAME,required"`
//	Tags    []string    `env:"TAGS,default=a,b,c"`
//	Env     Environment `env:"ENVIRONMENT,default=local"`
//
// Exported struct fields without an `env` tag are traversed recursively,
// allowing nested configuration objects.
func Unmarshal(v any) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Pointer || rv.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("envx: Unmarshal requires a non-nil pointer to a struct")
	}
	return unmarshalStruct(rv.Elem())
}

func unmarshalStruct(rv reflect.Value) error {
	rt := rv.Type()

	for i := range rt.NumField() {
		field := rt.Field(i)
		fv := rv.Field(i)

		tag, ok := field.Tag.Lookup("env")
		if !ok || tag == "" {
			// Recurse into untagged nested structs (skip EnvUnmarshaler types
			// and named scalars like time.Duration).
			if fv.Kind() == reflect.Struct && fv.CanSet() {
				if _, isUnmarshaler := fv.Addr().Interface().(EnvUnmarshaler); !isUnmarshaler {
					if err := unmarshalStruct(fv); err != nil {
						return err
					}
				}
			}
			continue
		}

		opts := parseTag(tag)

		raw, set := os.LookupEnv(opts.key)
		if !set {
			if opts.required {
				return fmt.Errorf("envx: %q is required but not set", opts.key)
			}
			if !opts.hasDefault {
				continue
			}
			raw = opts.defaultVal
		}

		if !fv.CanSet() {
			continue
		}

		if err := setField(fv, raw, opts.key); err != nil {
			return err
		}
	}
	return nil
}

type tagOpts struct {
	key        string
	required   bool
	hasDefault bool
	defaultVal string
}

// parseTag parses `KEY[,required][,default=VALUE]`.
// The default value may itself contain commas (e.g. "a,b,c" for slice fields),
// so "default=" is always the last recognised option and consumes the remainder.
func parseTag(tag string) tagOpts {
	key, rest, _ := strings.Cut(tag, ",")
	opts := tagOpts{key: key}

	for rest != "" {
		var part string
		part, rest, _ = strings.Cut(rest, ",")

		if part == "required" {
			opts.required = true
			continue
		}

		if val, found := strings.CutPrefix(part, "default="); found {
			// Everything from here (including any remaining commas) is the default.
			if rest != "" {
				val = val + "," + rest
				rest = ""
			}
			opts.hasDefault = true
			opts.defaultVal = val
		}
	}
	return opts
}

func setField(fv reflect.Value, raw, key string) error {
	if fv.Kind() == reflect.Slice {
		return setSliceField(fv, raw, key)
	}
	return setScalarField(fv, raw, key)
}

func setSliceField(fv reflect.Value, raw, key string) error {
	parts := strings.Split(raw, ",")
	slice := reflect.MakeSlice(fv.Type(), len(parts), len(parts))
	for i, p := range parts {
		if err := setScalarField(slice.Index(i), strings.TrimSpace(p), key); err != nil {
			return err
		}
	}
	fv.Set(slice)
	return nil
}

func setScalarField(fv reflect.Value, raw, key string) error {
	// If the type (or a pointer to it) implements EnvUnmarshaler, delegate.
	if u, ok := fv.Addr().Interface().(EnvUnmarshaler); ok {
		if err := u.FromEnv(raw); err != nil {
			return fmt.Errorf("envx: %q: %w", key, err)
		}
		return nil
	}

	// time.Duration is a named int64 — handle it before the generic int path.
	if fv.Type() == reflect.TypeOf(time.Duration(0)) {
		d, err := time.ParseDuration(raw)
		if err != nil {
			return fmt.Errorf("envx: %q: cannot parse %q as duration: %w", key, raw, err)
		}
		fv.SetInt(int64(d))
		return nil
	}

	switch fv.Kind() {
	case reflect.String:
		fv.SetString(raw)

	case reflect.Bool:
		b, err := strconv.ParseBool(raw)
		if err != nil {
			return fmt.Errorf("envx: %q: cannot parse %q as bool: %w", key, raw, err)
		}
		fv.SetBool(b)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n, err := strconv.ParseInt(raw, 10, fv.Type().Bits())
		if err != nil {
			return fmt.Errorf("envx: %q: cannot parse %q as int: %w", key, raw, err)
		}
		fv.SetInt(n)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		n, err := strconv.ParseUint(raw, 10, fv.Type().Bits())
		if err != nil {
			return fmt.Errorf("envx: %q: cannot parse %q as uint: %w", key, raw, err)
		}
		fv.SetUint(n)

	case reflect.Float32, reflect.Float64:
		n, err := strconv.ParseFloat(raw, fv.Type().Bits())
		if err != nil {
			return fmt.Errorf("envx: %q: cannot parse %q as float: %w", key, raw, err)
		}
		fv.SetFloat(n)

	default:
		return fmt.Errorf("envx: %q: unsupported field type %s", key, fv.Type())
	}
	return nil
}
