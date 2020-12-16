package compare

import (
	"fmt"
	"log"

	"github.com/jexia/semaphore/pkg/broker"
	"github.com/jexia/semaphore/pkg/broker/logger"
	"github.com/jexia/semaphore/pkg/specs"
	"github.com/jexia/semaphore/pkg/specs/types"
	"go.uber.org/zap"
)

// Types compares the types defined insde the schema definitions against the configured specification
func Types(ctx *broker.Context, services specs.ServiceList, objects specs.Schemas, flows specs.FlowListInterface) (err error) {
	logger.Info(ctx, "Comparing manifest types")

	for _, flow := range flows {
		log.Println("FLOW INPUT:", flow.GetInput())
		log.Println("FLOW OUTPUT:", flow.GetOutput())

		err := FlowTypes(ctx, services, objects, flow)
		if err != nil {
			return err
		}
	}

	return nil
}

// FlowTypes compares the flow types against the configured schema types
func FlowTypes(ctx *broker.Context, services specs.ServiceList, objects specs.Schemas, flow specs.FlowInterface) (err error) {
	logger.Info(ctx, "Comparing flow types", zap.String("flow", flow.GetName()))

	// omit typecheck if schema is not defined
	if input := flow.GetInput(); input != nil && input.Property != nil {
		err = CheckParameterMapTypes(ctx, input, objects, flow)
		if err != nil {
			return err
		}
	}

	if object := flow.GetOnError(); object != nil {
		err = CheckParameterMapTypes(ctx, object.Response, objects, flow)
		if err != nil {
			return err
		}
	}

	for _, node := range flow.GetNodes() {
		log.Println("___NODE:", node.Intermediate)

		err = CallTypes(ctx, services, objects, node, node.Call, flow)
		if err != nil {
			log.Println("CALL TYPES ERROR")

			return err
		}

		err = CallTypes(ctx, services, objects, node, node.Rollback, flow)
		if err != nil {
			return err
		}
	}

	if output := flow.GetOutput(); output != nil && output.Property != nil {
		message := objects.Get(output.Property.Schema)
		if message == nil {
			return ErrUndefinedObject{
				Flow:   flow.GetName(),
				Schema: output.Property.Schema,
			}
		}

		err = CheckParameterMapTypes(ctx, flow.GetOutput(), objects, flow)
		if err != nil {
			log.Println("PARSE OUTPUT ERROR")

			return err
		}
	}

	if flow.GetForward() != nil {
		if flow.GetForward().Request != nil && flow.GetForward().Request.Header != nil {
			err = CheckHeader(flow.GetForward().Request.Header, flow)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// CallTypes compares the given call types against the configured schema types
func CallTypes(ctx *broker.Context, services specs.ServiceList, objects specs.Schemas, node *specs.Node, call *specs.Call, flow specs.FlowInterface) (err error) {
	if call == nil {
		return nil
	}

	if call.Method == "" {
		return nil
	}

	logger.Info(ctx, "Comparing call types", zap.String("call", node.ID), zap.String("method", call.Method), zap.String("service", call.Service))

	service := services.Get(call.Service)
	if service == nil {
		return ErrUndefinedService{
			Flow:    flow.GetName(),
			Service: call.Service,
		}
	}

	method := service.GetMethod(call.Method)
	if method == nil {
		return ErrUndefinedMethod{
			Flow:   flow.GetName(),
			Method: call.Method,
		}
	}

	err = CheckParameterMapTypes(ctx, call.Request, objects, flow)
	if err != nil {
		return err
	}

	err = CheckParameterMapTypes(ctx, call.Response, objects, flow)
	if err != nil {
		return err
	}

	if node.OnError != nil {
		err = CheckParameterMapTypes(ctx, node.OnError.Response, objects, flow)
		if err != nil {
			return err
		}
	}

	return nil
}

// CheckParameterMapTypes checks the given parameter map against the configured schema property
func CheckParameterMapTypes(ctx *broker.Context, parameters *specs.ParameterMap, objects specs.Schemas, flow specs.FlowInterface) error {
	if parameters == nil {
		return nil
	}

	if parameters.Header != nil {
		err := CheckHeader(parameters.Header, flow)
		if err != nil {
			return err
		}
	}

	if parameters.Property != nil /*&& parameters.Property.Schema != "" */ {
		// WACHT: schema is empty here!!!

		schema := objects.Get(parameters.Property.Schema)

		err := parameters.Property.Compare(schema)
		if err != nil {
			return fmt.Errorf("flow '%s' mismatch: %w", flow.GetName(), err)
		}

		parameters.Property.Define(schema)
	}

	return nil
}

// CheckHeader compares the given header types
func CheckHeader(header specs.Header, flow specs.FlowInterface) error {
	for _, header := range header {
		if header.Type() != types.String {
			return ErrHeaderTypeMismatch{
				Type:     header.Type(),
				Path:     header.Path,
				Flow:     flow.GetName(),
				Expected: types.String,
			}
		}
	}

	return nil
}
