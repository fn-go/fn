package fnfile

import (
	"encoding/json"
)

type Do struct {
	StepCommon
	Steps []Step `json:"steps"`
}

type DoOptions struct {
	Steps             []Step
	StepCommonOptions []StepCommonOption
}

type DoOption func(options *DoOptions)

func NewDo(name string, options ...DoOption) Do {
	opts := &DoOptions{}
	for _, o := range options {
		o(opts)
	}

	return Do{
		StepCommon: NewStepCommon(name, opts.StepCommonOptions...),
		Steps:      opts.Steps,
	}
}

func (do *Do) UnmarshalJSON(data []byte) error {
	type DoAlias Do
	var tmpDo DoAlias

	err := json.Unmarshal(data, &tmpDo)
	if err != nil {
		return err
	}

	*do = Do(tmpDo)
	return nil
}

func (do Do) Exec(w ResponseWriter, c *CallInfo) {
	validateHandlerParams(w, c)
	ctx := c.Context()

	w2, d2 := withNewDeferrals(w)

	// for now, you cannot defer from a defer
	deferWriter, _ := withNewDeferrals(w)

	defer func() {
		// execute deferrals in reverse order
		for i := len(*d2) - 1; i >= 0; i-- {
			deferral := (*d2)[i]
			deferral.Exec(deferWriter, c)
			if w2.ErrorOrNil() != nil {
				return
			}
		}
	}()

	for _, s := range do.Steps {
		select {
		case <-ctx.Done():
			w.Error(NewCancelledStepError(do.Name, ctx.Err()))
			return
		default:
			s.Exec(w2, c)
			if w2.ErrorOrNil() != nil {
				return
			}
		}
	}
}
