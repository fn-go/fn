package fnfile

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MyMockedObject is a mocked object that implements an interface
// that describes an object that the code I am testing relies on.
type MockedStepVisitor struct {
	mock.Mock
}

func (m *MockedStepVisitor) VisitDo(ctx context.Context, do *Do) error {
	args := m.Called(ctx, do)
	return args.Error(0)
}
func (m *MockedStepVisitor) VisitParallel(ctx context.Context, parallel *Parallel) error {
	args := m.Called(ctx, parallel)
	return args.Error(0)
}
func (m *MockedStepVisitor) VisitTry(ctx context.Context, try *Try) error {
	args := m.Called(ctx, try)
	return args.Error(0)
}
func (m *MockedStepVisitor) VisitSh(ctx context.Context, sh *Sh) error {
	args := m.Called(ctx, sh)
	return args.Error(0)
}
func (m *MockedStepVisitor) VisitDefer(ctx context.Context, spec *DeferSpec) error {
	args := m.Called(ctx, spec)
	return args.Error(0)
}
func (m *MockedStepVisitor) VisitReturn(ctx context.Context, spec *ReturnSpec) error {
	args := m.Called(ctx, spec)
	return args.Error(0)
}
func (m *MockedStepVisitor) VisitMatrix(ctx context.Context, matrix *Matrix) error {
	args := m.Called(ctx, matrix)
	return args.Error(0)
}
func (m *MockedStepVisitor) VisitWait(ctx context.Context, wait *Wait) error {
	args := m.Called(ctx, wait)
	return args.Error(0)
}

func visitTest[T Step](t *testing.T, name string, step T) {
	mockVisitor := new(MockedStepVisitor)
	expectedCtx := context.Background()

	mockVisitor.On("Visit"+name, expectedCtx, step).Return(nil)
	err := step.Visit(expectedCtx, mockVisitor)
	assert.NoError(t, err)
	mockVisitor.AssertExpectations(t)
}
