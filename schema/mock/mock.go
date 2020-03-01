package mock

import (
	"github.com/jexia/maestro/schema"
	"github.com/jexia/maestro/specs/types"
)

// NewCollection constructs a new schema collection from the given descriptors
func NewCollection(descriptor Collection) *Collection {
	return &descriptor
}

// Exception represents a exception thrown during runtime
type Exception struct {
	File    string `yaml:"file"`
	Line    int    `yaml:"line"`
	Message string `yaml:"message"`
}

// Collection represents a mock YAML file
type Collection struct {
	Exception Exception           `yaml:"exception"`
	Services  map[string]*Service `yaml:"services"`
}

// GetService attempts to find the given service
func (collection *Collection) GetService(name string) schema.Service {
	for key, service := range collection.Services {
		if key != name {
			continue
		}

		return NewService(name, service)
	}

	return nil
}

// NewService constructs a new service with the given descriptor
func NewService(name string, service *Service) *Service {
	service.Name = name
	return service
}

// Service represents a mocking service
type Service struct {
	Name    string
	Methods map[string]*Method `yaml:"methods"`
	Options schema.Options     `yaml:"options"`
}

// GetName returns the service name
func (service *Service) GetName() string {
	return service.Name
}

// GetMethod attempts to return the given service method
func (service *Service) GetMethod(name string) schema.Method {
	for key, method := range service.Methods {
		if key != name {
			continue
		}

		return NewMethod(key, method)
	}

	return nil
}

// GetOptions returns the service options
func (service *Service) GetOptions() schema.Options {
	return service.Options
}

// GetMethods attempts to return the given service methods
func (service *Service) GetMethods() []schema.Method {
	result := make([]schema.Method, len(service.Methods))

	index := 0
	for key, method := range service.Methods {
		result[index] = NewMethod(key, method)
		index++
	}

	return result
}

// Method represents a mock YAML service method
type Method struct {
	Name    string
	Input   *Object        `yaml:"input"`
	Output  *Object        `yaml:"output"`
	Options schema.Options `yaml:"options"`
}

// NewMethod constructs a new method with the given descriptor
func NewMethod(name string, method *Method) *Method {
	method.Name = name
	return method
}

// GetName returns the method name
func (method *Method) GetName() string {
	return method.Name
}

// GetInput returns the method input
func (method *Method) GetInput() schema.Object {
	return method.Input
}

// GetOutput returns the method output
func (method *Method) GetOutput() schema.Object {
	return method.Output
}

// GetOptions returns the method options
func (method *Method) GetOptions() schema.Options {
	return method.Options
}

// Object represents a proto message
type Object struct {
	Fields  map[string]*Field `yaml:"fields"`
	Options schema.Options    `yaml:"options"`
}

// GetField attempts to return a field matching the given name
func (object *Object) GetField(name string) schema.Field {
	field, has := object.Fields[name]
	if !has {
		return nil
	}

	return NewField(name, field)
}

// GetFields returns the available field inside the given object
func (object *Object) GetFields() []schema.Field {
	result := []schema.Field{}

	for key, field := range object.Fields {
		result = append(result, NewField(key, field))
	}

	return result
}

// GetOptions returns the object options
func (object *Object) GetOptions() schema.Options {
	return object.Options
}

// NewField constructs a new object field with the given descriptor
func NewField(name string, field *Field) *Field {
	field.Name = name
	return field
}

// Field represents a proto message field
type Field struct {
	Name    string         `yaml:"name"`
	Type    types.Type     `yaml:"type"`
	Label   types.Label    `yaml:"label"`
	Object  *Object        `yaml:"message"`
	Options schema.Options `yaml:"options"`
}

// GetName returns the field name
func (field *Field) GetName() string {
	return field.Name
}

// GetType returns tye field type
func (field *Field) GetType() types.Type {
	return field.Type
}

// GetLabel returns the field label
func (field *Field) GetLabel() types.Label {
	return field.Label
}

// GetObject returns the field object
func (field *Field) GetObject() schema.Object {
	return field.Object
}

// GetOptions returns the field options
func (field *Field) GetOptions() schema.Options {
	return field.Options
}
