// Code generated by mockery v2.1.0. DO NOT EDIT.

package processor

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	processor "github.com/goharbor/harbor/src/controller/artifact/processor"
	artifact "github.com/goharbor/harbor/src/pkg/artifact"
)

// Processor is an autogenerated mock type for the Processor type
type Processor struct {
	mock.Mock
}

// AbstractAddition provides a mock function with given fields: ctx, _a1, additionType
func (_m *Processor) AbstractAddition(ctx context.Context, _a1 *artifact.Artifact, additionType string) (*processor.Addition, error) {
	ret := _m.Called(ctx, _a1, additionType)

	var r0 *processor.Addition
	if rf, ok := ret.Get(0).(func(context.Context, *artifact.Artifact, string) *processor.Addition); ok {
		r0 = rf(ctx, _a1, additionType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*processor.Addition)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *artifact.Artifact, string) error); ok {
		r1 = rf(ctx, _a1, additionType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AbstractMetadata provides a mock function with given fields: ctx, _a1, manifest
func (_m *Processor) AbstractMetadata(ctx context.Context, _a1 *artifact.Artifact, manifest []byte) error {
	ret := _m.Called(ctx, _a1, manifest)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *artifact.Artifact, []byte) error); ok {
		r0 = rf(ctx, _a1, manifest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetArtifactType provides a mock function with given fields: ctx, _a1
func (_m *Processor) GetArtifactType(ctx context.Context, _a1 *artifact.Artifact) string {
	ret := _m.Called(ctx, _a1)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *artifact.Artifact) string); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ListAdditionTypes provides a mock function with given fields: ctx, _a1
func (_m *Processor) ListAdditionTypes(ctx context.Context, _a1 *artifact.Artifact) []string {
	ret := _m.Called(ctx, _a1)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context, *artifact.Artifact) []string); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}
