package flow

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/jexia/semaphore/pkg/codec/json"
	"github.com/jexia/semaphore/pkg/core/instance"
	"github.com/jexia/semaphore/pkg/core/logger"
	"github.com/jexia/semaphore/pkg/refs"
	"github.com/jexia/semaphore/pkg/specs"
	"github.com/jexia/semaphore/pkg/specs/labels"
	"github.com/jexia/semaphore/pkg/specs/types"
)

func NewMockNode(name string, caller Call, rollback Call) *Node {
	ctx := instance.NewContext()
	logger := ctx.Logger(logger.Flow)

	return &Node{
		ctx:        ctx,
		logger:     logger,
		Name:       name,
		Call:       caller,
		Revert:     rollback,
		OnError:    NewMockOnError(),
		DependsOn:  map[string]*specs.Node{},
		References: map[string]*specs.PropertyReference{},
	}
}

func NewMockOnError() *specs.OnError {
	return &specs.OnError{
		Response: &specs.ParameterMap{
			Property: &specs.Property{
				Type:  types.Message,
				Label: labels.Optional,
				Nested: map[string]*specs.Property{
					"status": {
						Type:  types.Int64,
						Label: labels.Optional,
						Reference: &specs.PropertyReference{
							Resource: "error",
							Path:     "status",
						},
					},
					"message": {
						Type:  types.String,
						Label: labels.Optional,
						Reference: &specs.PropertyReference{
							Resource: "error",
							Path:     "message",
						},
					},
				},
			},
		},
		Status: &specs.Property{
			Type:    types.Int64,
			Label:   labels.Optional,
			Default: 500,
		},
		Message: &specs.Property{
			Type:    types.String,
			Label:   labels.Optional,
			Default: "mock error message",
		},
	}
}

func BenchmarkSingleNodeCallingJSONCodecParallel(b *testing.B) {
	ctx := instance.NewContext()
	constructor := json.NewConstructor()

	req, err := constructor.New("first.request", &specs.ParameterMap{
		Property: &specs.Property{
			Type:  types.Message,
			Label: labels.Optional,
			Nested: map[string]*specs.Property{
				"key": {
					Name:    "key",
					Path:    "key",
					Type:    types.String,
					Label:   labels.Optional,
					Default: "message",
				},
			},
		},
	})

	if err != nil {
		b.Fatal(err)
	}

	res, err := constructor.New("first.response", &specs.ParameterMap{
		Property: &specs.Property{
			Type:  types.Message,
			Label: labels.Optional,
			Nested: map[string]*specs.Property{
				"key": {
					Name:    "key",
					Path:    "key",
					Type:    types.String,
					Label:   labels.Optional,
					Default: "message",
				},
			},
		},
	})

	if err != nil {
		b.Fatal(err)
	}

	options := &CallOptions{
		Request:  NewRequest(nil, req, nil),
		Response: NewRequest(nil, res, nil),
	}

	call := NewCall(ctx, nil, options)
	node := NewMockNode("first", call, nil)

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ctx := context.Background()
			tracker := NewTracker("", 1)
			processes := NewProcesses(1)
			refs := refs.NewReferenceStore(0)

			node.Do(ctx, tracker, processes, refs)
		}
	})
}

func BenchmarkSingleNodeCallingJSONCodecSerial(b *testing.B) {
	ctx := instance.NewContext()
	constructor := json.NewConstructor()

	req, err := constructor.New("first.request", &specs.ParameterMap{
		Property: &specs.Property{
			Type:  types.Message,
			Label: labels.Optional,
			Nested: map[string]*specs.Property{
				"key": {
					Name:    "key",
					Path:    "key",
					Type:    types.String,
					Label:   labels.Optional,
					Default: "message",
				},
			},
		},
	})

	if err != nil {
		b.Fatal(err)
	}

	res, err := constructor.New("first.response", &specs.ParameterMap{
		Property: &specs.Property{
			Type:  types.Message,
			Label: labels.Optional,
			Nested: map[string]*specs.Property{
				"key": {
					Name:    "key",
					Path:    "key",
					Type:    types.String,
					Label:   labels.Optional,
					Default: "message",
				},
			},
		},
	})

	if err != nil {
		b.Fatal(err)
	}

	options := &CallOptions{
		Request:  NewRequest(nil, req, nil),
		Response: NewRequest(nil, res, nil),
	}

	call := NewCall(ctx, nil, options)
	node := NewMockNode("first", call, nil)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		tracker := NewTracker("", 1)
		processes := NewProcesses(1)
		refs := refs.NewReferenceStore(0)

		node.Do(ctx, tracker, processes, refs)
	}
}

func BenchmarkSingleNodeCallingParallel(b *testing.B) {
	caller := &caller{}
	node := NewMockNode("first", caller, nil)

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ctx := context.Background()
			tracker := NewTracker("", 1)
			processes := NewProcesses(1)
			refs := refs.NewReferenceStore(0)

			node.Do(ctx, tracker, processes, refs)
		}
	})
}

func BenchmarkSingleNodeCallingSerial(b *testing.B) {
	caller := &caller{}
	node := NewMockNode("first", caller, nil)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		tracker := NewTracker("", 1)
		processes := NewProcesses(1)
		refs := refs.NewReferenceStore(0)

		node.Do(ctx, tracker, processes, refs)
	}
}

func BenchmarkBranchedNodeCallingParallel(b *testing.B) {
	caller := &caller{}
	nodes := []*Node{
		NewMockNode("first", caller, nil),
		NewMockNode("second", caller, nil),
		NewMockNode("third", caller, nil),
	}

	nodes[0].Next = []*Node{nodes[1]}
	nodes[1].Previous = []*Node{nodes[0]}
	nodes[1].Next = []*Node{nodes[2]}
	nodes[2].Previous = []*Node{nodes[1]}

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ctx := context.Background()
			tracker := NewTracker("", len(nodes))
			processes := NewProcesses(1)
			refs := refs.NewReferenceStore(0)

			nodes[0].Do(ctx, tracker, processes, refs)
		}
	})
}

func BenchmarkBranchedNodeCallingSerial(b *testing.B) {
	caller := &caller{}
	nodes := []*Node{
		NewMockNode("first", caller, nil),
		NewMockNode("second", caller, nil),
		NewMockNode("third", caller, nil),
	}

	nodes[0].Next = []*Node{nodes[1]}
	nodes[1].Previous = []*Node{nodes[0]}
	nodes[1].Next = []*Node{nodes[2]}
	nodes[2].Previous = []*Node{nodes[1]}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		tracker := NewTracker("", len(nodes))
		processes := NewProcesses(1)
		refs := refs.NewReferenceStore(0)

		nodes[0].Do(ctx, tracker, processes, refs)
	}
}

func TestConstructingNode(t *testing.T) {
	type test struct {
		Node     *specs.Node
		Call     Call
		Rollback Call
		Expected int
	}

	tests := map[string]*test{
		"node call": {
			Expected: 2,
			Node: &specs.Node{
				Call: &specs.Call{
					Request: &specs.ParameterMap{
						Property: &specs.Property{
							Nested: map[string]*specs.Property{
								"first": {
									Reference: &specs.PropertyReference{
										Resource: "input",
										Path:     "first",
									},
								},
								"second": {
									Reference: &specs.PropertyReference{
										Resource: "input",
										Path:     "second",
									},
								},
							},
						},
					},
				},
			},
		},
		"combination": {
			Expected: 1,
			Node: &specs.Node{
				Call: &specs.Call{
					Request: &specs.ParameterMap{
						Property: &specs.Property{
							Nested: map[string]*specs.Property{
								"first": {
									Reference: &specs.PropertyReference{
										Resource: "input",
										Path:     "first",
									},
								},
							},
						},
					},
				},
			},
			Call: &caller{
				references: []*specs.Property{
					{
						Reference: &specs.PropertyReference{
							Resource: "input",
							Path:     "first",
						},
					},
				},
			},
			Rollback: &caller{
				references: []*specs.Property{
					{
						Reference: &specs.PropertyReference{
							Resource: "input",
							Path:     "first",
						},
					},
				},
			},
		},
		"call references": {
			Expected: 2,
			Node:     &specs.Node{},
			Call: &caller{
				references: []*specs.Property{
					{
						Reference: &specs.PropertyReference{
							Resource: "input",
							Path:     "first",
						},
					},
					{
						Reference: &specs.PropertyReference{
							Resource: "input",
							Path:     "second",
						},
					},
				},
			},
		},
		"rollback references": {
			Expected: 2,
			Node:     &specs.Node{},
			Rollback: &caller{
				references: []*specs.Property{
					{
						Reference: &specs.PropertyReference{
							Resource: "input",
							Path:     "first",
						},
					},
					{
						Reference: &specs.PropertyReference{
							Resource: "input",
							Path:     "second",
						},
					},
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ctx := instance.NewContext()
			result := NewNode(ctx, test.Node, nil, test.Call, test.Rollback, nil)

			if len(result.References) != test.Expected {
				t.Fatalf("unexpected amount of references %d, expected %d", len(result.References), test.Expected)
			}
		})
	}
}

func TestConstructingNodeReferences(t *testing.T) {
	ctx := instance.NewContext()
	call := &caller{}
	rollback := &caller{}

	node := &specs.Node{
		Name: "mock",
	}

	result := NewNode(ctx, node, nil, call, rollback, nil)
	if result == nil {
		t.Fatal("nil node returned")
	}
}

func TestNodeHas(t *testing.T) {
	nodes := make(Nodes, 2)

	nodes[0] = &Node{Name: "first"}
	nodes[1] = &Node{Name: "second"}

	if !nodes.Has("first") {
		t.Fatal("unexpected result, expected 'first' to be available")
	}

	if nodes.Has("unexpected") {
		t.Fatal("unexpected result, expected 'unexpected' to be unavailable")
	}
}

func TestNodeCalling(t *testing.T) {
	caller := &caller{}

	nodes := []*Node{
		NewMockNode("first", caller, nil),
		NewMockNode("second", caller, nil),
		NewMockNode("third", caller, nil),
	}

	nodes[0].Next = []*Node{nodes[1]}
	nodes[1].Previous = []*Node{nodes[0]}
	nodes[1].Next = []*Node{nodes[2]}
	nodes[2].Previous = []*Node{nodes[1]}

	tracker := NewTracker("", len(nodes))
	processes := NewProcesses(1)
	refs := refs.NewReferenceStore(0)

	nodes[0].Do(context.Background(), tracker, processes, refs)
	processes.Wait()

	if processes.Err() != nil {
		t.Error(processes.Err())
	}

	if caller.Counter != len(nodes) {
		t.Errorf("unexpected counter total %d, expected %d", caller.Counter, len(nodes))
	}
}

func TestSlowNodeAbortingOnErr(t *testing.T) {
	slow := &caller{name: "slow"}
	failed := &caller{name: "failed", Err: errors.New("unexpected")}
	caller := &caller{}

	nodes := []*Node{
		NewMockNode("first", caller, nil),
		NewMockNode("second", slow, nil),
		NewMockNode("third", failed, nil),
		NewMockNode("fourth", caller, nil),
	}

	nodes[0].Next = []*Node{nodes[1], nodes[2]}

	nodes[1].Previous = []*Node{nodes[0]}
	nodes[1].Next = []*Node{nodes[3]}

	nodes[2].Previous = []*Node{nodes[0]}
	nodes[2].Next = []*Node{nodes[3]}

	nodes[3].Previous = []*Node{nodes[1], nodes[2]}

	tracker := NewTracker("", len(nodes))
	processes := NewProcesses(1)
	refs := refs.NewReferenceStore(0)

	slow.mutex.Lock()
	failed.mutex.Lock()

	go func() {
		failed.mutex.Unlock()
		time.Sleep(100 * time.Millisecond)
		slow.mutex.Unlock()
	}()

	nodes[0].Do(context.Background(), tracker, processes, refs)

	processes.Wait()

	counter := (caller.Counter + slow.Counter + failed.Counter)
	if counter != 3 {
		t.Fatalf("unexpected counter total %d, expected %d", counter, 3)
	}
}

func TestNodeRevert(t *testing.T) {
	rollback := &caller{}

	nodes := []*Node{
		NewMockNode("first", nil, rollback),
		NewMockNode("second", nil, rollback),
		NewMockNode("third", nil, rollback),
	}

	nodes[0].Next = []*Node{nodes[1]}
	nodes[1].Previous = []*Node{nodes[0]}
	nodes[1].Next = []*Node{nodes[2]}
	nodes[2].Previous = []*Node{nodes[1]}

	tracker := NewTracker("", len(nodes))
	processes := NewProcesses(1)
	refs := refs.NewReferenceStore(0)

	nodes[len(nodes)-1].Rollback(context.Background(), tracker, processes, refs)
	processes.Wait()

	if processes.Err() != nil {
		t.Error(processes.Err())
	}

	if rollback.Counter != len(nodes) {
		t.Errorf("unexpected counter total %d, expected %d", rollback.Counter, len(nodes))
	}
}

func TestNodeBranchesCalling(t *testing.T) {
	caller := &caller{}

	nodes := []*Node{
		NewMockNode("first", caller, nil),
		NewMockNode("second", caller, nil),
		NewMockNode("third", caller, nil),
		NewMockNode("fourth", caller, nil),
	}

	nodes[0].Next = []*Node{nodes[1], nodes[2]}

	nodes[1].Previous = []*Node{nodes[0]}
	nodes[1].Next = []*Node{nodes[3]}
	nodes[2].Previous = []*Node{nodes[0]}
	nodes[2].Next = []*Node{nodes[3]}

	nodes[3].Previous = []*Node{nodes[1], nodes[2]}

	tracker := NewTracker("", len(nodes))
	processes := NewProcesses(1)
	refs := refs.NewReferenceStore(0)

	nodes[0].Do(context.Background(), tracker, processes, refs)
	processes.Wait()

	if processes.Err() != nil {
		t.Error(processes.Err())
	}

	if caller.Counter != len(nodes) {
		t.Errorf("unexpected counter total %d, expected %d", caller.Counter, len(nodes))
	}
}

func TestBeforeDoNode(t *testing.T) {
	counter := 0
	call := &caller{}
	node := NewMockNode("mock", call, nil)

	node.BeforeDo = func(ctx context.Context, node *Node, tracker *Tracker, processes *Processes, store refs.Store) (context.Context, error) {
		counter++
		return ctx, nil
	}

	processes := NewProcesses(1)
	node.Do(context.Background(), NewTracker("", 1), processes, nil)
	if processes.Err() != nil {
		t.Error(processes.Err())
	}

	if counter != 1 {
		t.Fatalf("unexpected counter %d, expected after rollback function to be called", counter)
	}
}

func TestBeforeDoNodeErr(t *testing.T) {
	expected := errors.New("unexpected err")
	counter := 0
	call := &caller{}
	node := NewMockNode("mock", call, nil)

	node.BeforeDo = func(ctx context.Context, node *Node, tracker *Tracker, processes *Processes, store refs.Store) (context.Context, error) {
		counter++
		return ctx, expected
	}

	processes := NewProcesses(1)
	node.Do(context.Background(), NewTracker("", 1), processes, nil)
	if !errors.Is(processes.Err(), expected) {
		t.Errorf("unexpected err '%s', expected '%s' to be thrown", processes.Err(), expected)
	}

	if counter != 1 {
		t.Fatalf("unexpected counter %d, expected after rollback function to be called", counter)
	}
}

func TestAfterDoNode(t *testing.T) {
	counter := 0
	call := &caller{}
	node := NewMockNode("mock", call, nil)

	node.AfterDo = func(ctx context.Context, node *Node, tracker *Tracker, processes *Processes, store refs.Store) (context.Context, error) {
		counter++
		return ctx, nil
	}

	processes := NewProcesses(1)
	node.Do(context.Background(), NewTracker("", 1), processes, nil)
	if processes.Err() != nil {
		t.Error(processes.Err())
	}

	if counter != 1 {
		t.Fatalf("unexpected counter %d, expected after rollback function to be called", counter)
	}
}

func TestAfterDoNodeErr(t *testing.T) {
	expected := errors.New("unexpected err")
	counter := 0
	call := &caller{}
	node := NewMockNode("mock", call, nil)

	node.AfterDo = func(ctx context.Context, node *Node, tracker *Tracker, processes *Processes, store refs.Store) (context.Context, error) {
		counter++
		return ctx, expected
	}

	processes := NewProcesses(1)
	node.Do(context.Background(), NewTracker("", 1), processes, nil)
	if !errors.Is(processes.Err(), expected) {
		t.Errorf("unexpected err '%s', expected '%s' to be thrown", processes.Err(), expected)
	}

	if counter != 1 {
		t.Fatalf("unexpected counter %d, expected after rollback function to be called", counter)
	}
}

func TestBeforeRevertNode(t *testing.T) {
	counter := 0
	call := &caller{}
	node := NewMockNode("mock", call, nil)

	node.BeforeRevert = func(ctx context.Context, node *Node, tracker *Tracker, processes *Processes, store refs.Store) (context.Context, error) {
		counter++
		return ctx, nil
	}

	processes := NewProcesses(1)
	node.Rollback(context.Background(), NewTracker("", 1), processes, nil)
	if processes.Err() != nil {
		t.Error(processes.Err())
	}

	if counter != 1 {
		t.Fatalf("unexpected counter %d, expected after revert function to be called", counter)
	}
}

func TestBeforeRevertNodeErr(t *testing.T) {
	expected := errors.New("unexpected err")
	counter := 0
	call := &caller{}
	node := NewMockNode("mock", call, nil)

	node.BeforeRevert = func(ctx context.Context, node *Node, tracker *Tracker, processes *Processes, store refs.Store) (context.Context, error) {
		counter++
		return ctx, expected
	}

	processes := NewProcesses(1)
	node.Rollback(context.Background(), NewTracker("", 1), processes, nil)
	if !errors.Is(processes.Err(), expected) {
		t.Errorf("unexpected err '%s', expected '%s' to be thrown", processes.Err(), expected)
	}

	if counter != 1 {
		t.Fatalf("unexpected counter %d, expected after revert function to be called", counter)
	}
}

func TestAfterRevertNode(t *testing.T) {
	counter := 0
	call := &caller{}
	node := NewMockNode("mock", call, nil)

	node.AfterRevert = func(ctx context.Context, node *Node, tracker *Tracker, processes *Processes, store refs.Store) (context.Context, error) {
		counter++
		return ctx, nil
	}

	processes := NewProcesses(1)
	node.Rollback(context.Background(), NewTracker("", 1), processes, nil)
	if processes.Err() != nil {
		t.Error(processes.Err())
	}

	if counter != 1 {
		t.Fatalf("unexpected counter %d, expected after revert function to be called", counter)
	}
}

func TestAfterRevertNodeErr(t *testing.T) {
	expected := errors.New("unexpected err")
	counter := 0
	call := &caller{}
	node := NewMockNode("mock", call, nil)

	node.AfterRevert = func(ctx context.Context, node *Node, tracker *Tracker, processes *Processes, store refs.Store) (context.Context, error) {
		counter++
		return ctx, expected
	}

	processes := NewProcesses(1)
	node.Rollback(context.Background(), NewTracker("", 1), processes, nil)
	if !errors.Is(processes.Err(), expected) {
		t.Errorf("unexpected err '%s', expected '%s' to be thrown", processes.Err(), expected)
	}

	if counter != 1 {
		t.Fatalf("unexpected counter %d, expected after revert function to be called", counter)
	}
}
