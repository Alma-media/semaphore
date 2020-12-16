package hcl

import (
	"net/http"

	"github.com/jexia/semaphore/pkg/broker"
	"github.com/jexia/semaphore/pkg/specs"
	"github.com/jexia/semaphore/pkg/specs/labels"
	"github.com/jexia/semaphore/pkg/specs/template"
	"github.com/jexia/semaphore/pkg/specs/types"
)

// ParseIntermediateNestedParameterMap parses the given intermediate parameter map to a spec parameter map
func ParseIntermediateNestedParameterMap(ctx *broker.Context, params NestedParameterMap, path string) (*specs.Property, error) {
	message, err := parseBaseParameterMap(ctx, params.BaseParameterMap, path)
	if err != nil {
		return nil, err
	}

	return &specs.Property{
		Name:  params.Name,
		Path:  path,
		Label: labels.Optional,
		Template: specs.Template{
			Message: message,
		},
	}, nil
}

// ParseIntermediateRepeatedParameterMap parses the given intermediate repeated parameter map to a spec repeated parameter map
func ParseIntermediateRepeatedParameterMap(ctx *broker.Context, params RepeatedParameterMap, path string) (*specs.Property, error) {
	message, err := parseBaseParameterMap(ctx, params.BaseParameterMap, path)
	if err != nil {
		return nil, err
	}

	return &specs.Property{
		Name:  params.Name,
		Path:  path,
		Label: labels.Optional,
		Template: specs.Template{
			Reference: template.ParsePropertyReference(params.Template),
			Repeated: specs.Repeated{
				specs.Template{
					Message: message,
				},
			},
		},
	}, nil
}

// ParseIntermediateOutput parses the given intermediate parameter map to a spec parameter map
func ParseIntermediateOutput(ctx *broker.Context, params *Output) (*specs.ParameterMap, error) {
	if params == nil {
		return nil, nil
	}

	result := specs.ParameterMap{
		Options: make(specs.Options),
		Status:  http.StatusOK,
	}

	if params.Header != nil {
		header, err := ParseIntermediateHeader(ctx, params.Header)
		if err != nil {
			return nil, err
		}

		result.Header = header
	}

	if params.OutputParameterMap != nil {
		message, err := parseBaseParameterMap(ctx, params.BaseParameterMap, "")
		if err != nil {
			return nil, err
		}

		result.Property = &specs.Property{
			Schema: params.OutputParameterMap.Schema,
			Label:  labels.Optional,
			Template: specs.Template{
				Message: message,
			},
		}
	}

	if params.Options != nil {
		result.Options = ParseIntermediateSpecOptions(params.Options.Body)
	}

	return &result, nil
}

// ParseIntermediateCallParameterMap parses the given intermediate parameter map to a spec parameter map
func ParseIntermediateCallParameterMap(ctx *broker.Context, params *Call) (*specs.ParameterMap, error) {
	if params == nil {
		return nil, nil
	}

	result := specs.ParameterMap{
		Options: make(specs.Options),
	}

	if params.Parameters != nil {
		params, err := ParseIntermediateParameters(ctx, params.Parameters.Body)
		if err != nil {
			return nil, err
		}

		result.Params = params
	}

	if params.Options != nil {
		result.Options = ParseIntermediateSpecOptions(params.Options.Body)
	}

	if params.Header != nil {
		header, err := ParseIntermediateHeader(ctx, params.Header)
		if err != nil {
			return nil, err
		}

		result.Header = header
	}

	if params.BaseParameterMap != nil {
		message, err := parseBaseParameterMap(ctx, *params.BaseParameterMap, "")
		if err != nil {
			return nil, err
		}

		result.Property = &specs.Property{
			//
			// Schema: params.BaseParameterMap.Schema,
			//
			Label: labels.Optional,
			Template: specs.Template{
				Message: message,
			},
		}
	}

	return &result, nil
}

// ParseIntermediateInput parses the given input parameter map
func ParseIntermediateInput(ctx *broker.Context, params *Input) (*specs.ParameterMap, error) {
	if params == nil {
		return nil, nil
	}

	result := &specs.ParameterMap{
		Options: make(specs.Options),
		Header:  make(specs.Header, len(params.Header)),
	}

	if params.InputParameterMap != nil {
		result.Property = &specs.Property{
			Schema: params.InputParameterMap.Schema,
		}
	}

	for _, key := range params.Header {
		result.Header[key] = &specs.Property{
			Path:  key,
			Name:  key,
			Label: labels.Optional,
			Template: specs.Template{
				Scalar: &specs.Scalar{
					Type: types.String,
				},
			},
		}
	}

	if params.Options != nil {
		result.Options = ParseIntermediateSpecOptions(params.Options.Body)
	}

	return result, nil
}

func parseBaseParameterMap(ctx *broker.Context, params BaseParameterMap, path string) (specs.Message, error) {
	var (
		message       = make(specs.Message)
		properties, _ = params.Properties.JustAttributes()
	)

	for _, attr := range properties {
		val, _ := attr.Expr.Value(nil)
		returns, err := ParseIntermediateProperty(ctx, path, attr, val)
		if err != nil {
			return nil, err
		}

		message[attr.Name] = returns
	}

	for _, nested := range params.Nested {
		returns, err := ParseIntermediateNestedParameterMap(ctx, nested, template.JoinPath(path, nested.Name))
		if err != nil {
			return nil, err
		}

		message[nested.Name] = returns
	}

	for _, repeated := range params.Repeated {
		returns, err := ParseIntermediateRepeatedParameterMap(ctx, repeated, template.JoinPath(path, repeated.Name))
		if err != nil {
			return nil, err
		}

		message[repeated.Name] = returns
	}

	return message, nil
}
