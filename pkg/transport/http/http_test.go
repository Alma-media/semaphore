package http

import (
	"context"
	"encoding/json"
	"io"
	"net"
	"reflect"
	"testing"

	"github.com/jexia/maestro/internal/flow"
	"github.com/jexia/maestro/pkg/metadata"
	"github.com/jexia/maestro/pkg/refs"
	"github.com/jexia/maestro/pkg/specs"
	"github.com/jexia/maestro/pkg/specs/labels"
	"github.com/jexia/maestro/pkg/specs/types"
)

// JSONEqual compares the JSON from two Readers.
func JSONEqual(a, b io.Reader) (bool, interface{}, interface{}, error) {
	var j, j2 interface{}
	d := json.NewDecoder(a)
	if err := d.Decode(&j); err != nil {
		return false, j, j2, err
	}
	d = json.NewDecoder(b)
	if err := d.Decode(&j2); err != nil {
		return false, j, j2, err
	}
	return reflect.DeepEqual(j2, j), j, j2, nil
}

type caller struct {
	fn func(context.Context, refs.Store) error
}

func (caller *caller) Do(ctx context.Context, store refs.Store) error {
	return caller.fn(ctx, store)
}

func (caller *caller) References() []*specs.Property {
	return nil
}

func NewCallerFunc(fn func(context.Context, refs.Store) error) flow.Call {
	return &caller{fn: fn}
}

func NewSimpleMockSpecs() *specs.ParameterMap {
	return &specs.ParameterMap{
		Property: &specs.Property{
			Type:  types.Message,
			Label: labels.Optional,
			Nested: map[string]*specs.Property{
				"message": {
					Name: "message",
					Path: "message",
					Type: types.String,
				},
			},
		},
	}
}

func NewMockService(host string, method string, endpoint string) *specs.Service {
	return &specs.Service{
		Host: host,
		Methods: []*specs.Method{
			{
				Name: "mock",
				Options: specs.Options{
					MethodOption:   method,
					EndpointOption: endpoint,
				},
			},
		},
	}
}

type MockResponseWriter struct {
	header  metadata.MD
	writer  io.WriteCloser
	status  int
	message string
}

func (rw *MockResponseWriter) Header() metadata.MD {
	return rw.header
}

func (rw *MockResponseWriter) Write(bb []byte) (int, error) {
	return rw.writer.Write(bb)
}

func (rw *MockResponseWriter) HeaderStatus(status int) {
	rw.status = status
}

func (rw *MockResponseWriter) HeaderMessage(message string) {
	rw.message = message
}

func (rw *MockResponseWriter) Close() error {
	return rw.writer.Close()
}

func AvailablePort(t *testing.T) int {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatal(err)
	}

	defer listener.Close()
	return listener.Addr().(*net.TCPAddr).Port
}
