package fnfile

const (
	DoStepType StepType = "do"
)

type Do struct {
	StepMeta
	Steps []Step `json:"steps"`
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
