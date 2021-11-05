package function

import (
	"strings"

	"github.com/mattermost/mattermost-plugin-apps/apps"
)

type Command struct {
	Name        string
	Hint        string
	Description string

	BaseSubmit  apps.Call
	BaseForm    apps.Form
	BaseBinding apps.Binding

	Handler func(CallRequest) apps.CallResponse
}

func (c Command) Path() string {
	return "/" + c.Name
}

func (c Command) Submit(creq CallRequest) *apps.Call {
	s := *c.BaseSubmit.PartialCopy()
	if s.Path == "" {
		s.Path = c.Path()
	}
	return &s
}

func (c Command) Form(creq CallRequest) *apps.Form {
	f := *c.BaseForm.PartialCopy()
	if f.Icon == "" {
		f.Icon = IconPath
	}
	if f.Call == nil {
		f.Call = c.Submit(creq)
	} else if f.Call.Path == "" {
		f.Call.Path = c.Path()
	}

	for i, field := range f.Fields {
		if field.Label == "" {
			f.Fields[i].Label = strings.ReplaceAll(field.Name, "_", "-")
		}
		if field.ModalLabel == "" {
			f.Fields[i].ModalLabel = strings.ReplaceAll(field.Name, "_", " ")
		}
	}
	return &f
}

func (c Command) Binding(creq CallRequest) apps.Binding {
	b := apps.Binding{
		Location:    apps.Location(c.Name),
		Icon:        IconPath,
		Label:       strings.ReplaceAll(c.Name, "_", "-"),
		Hint:        c.Hint,
		Description: c.Description,
		Call:        c.Submit(creq),
		Form:        c.Form(creq),
	}

	return b
}
