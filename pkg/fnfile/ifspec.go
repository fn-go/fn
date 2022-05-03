package fnfile

import (
	"context"
)

type IfSpec struct {
}

// ConditionVisitor is part of the Decoupled Visitor Pattern
// https://making.pusher.com/alternatives-to-sum-types-in-go/
type ConditionVisitor interface {
	VisitFileCondition(ctx context.Context, fileCond *FileCondition) error
}

type Condition interface {
	Visit(ctx context.Context, v ConditionVisitor) error
}

type FileCondition struct {
	IfSpec

	// Sources matched files have their contents hashed.
	// If the content hashes change between runs, this fn will be marked as "out-of-date".
	Sources FileGlobs `json:"from,omitempty"`

	// Generates matched files are checked to exist.
	// If these files do not exist, this fn will be marked as "out-of-date".
	// When providing a glob, only 1 match is required to keep this fn "up-to-date".
	Generates FileGlobs `json:"makes,omitempty"`
}

func (f *FileCondition) Visit(ctx context.Context, v ConditionVisitor) error {
	return nil
}

type FnCondition struct {
	IfSpec

	Fn Fn `json:"fn"`
}

func (f *FnCondition) Visit(ctx context.Context, v ConditionVisitor) error {
	return nil
}

type StepOutcomeCondition struct {
	IfSpec
}
