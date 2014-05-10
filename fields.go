package wtforms

import (
	"fmt"
	"html/template"
	"strings"
)

type IField interface {
	RenderLabel(attrs ...string) template.HTML
	RenderInput(attrs ...string) template.HTML
	Validate() bool
	GetName() string
	GetValue() string
	SetValue(value string)
	IsName(name string) bool
	HasErrors() bool
	RenderErrors() template.HTML
	AddError(err string)
	Errors() []string
}

type BaseField struct {
	Name       string
	Label      string
	Value      string
	errors     []string
	Validators []IValidator
}

func (field *BaseField) RenderLabel(attrs ...string) template.HTML {
	return template.HTML(fmt.Sprintf("<label for=\"%s\" %s>%s</label>", field.Name, strings.Join(attrs, " "), field.Label))
}

func (field *BaseField) HasErrors() bool {
	return len(field.errors) > 0
}

func (field *BaseField) RenderInput(attrs ...string) template.HTML {
	return template.HTML("")
}

func (field *BaseField) GetName() string {
	return field.Name
}

func (field *BaseField) AddError(err string) {
	field.errors = append(field.errors, err)
}

func (field *BaseField) RenderErrors() template.HTML {
	result := ""
	for _, err := range field.errors {
		result += fmt.Sprintf(`<span class="help-block">%s</span>`, err)
	}

	return template.HTML(result)
}

func (field *BaseField) Validate() bool {
	// 如果有Required并且输入为空,不在进行其他检查
	for _, validator := range field.Validators {
		if _, ok := validator.(Required); ok {
			if ok, message := validator.CleanData(field.GetValue()); !ok {
				field.errors = append(field.errors, message)
				return false
			}
		}
	}

	result := true

	for _, validator := range field.Validators {
		if ok, message := validator.CleanData(field.GetValue()); !ok {
			result = false
			field.errors = append(field.errors, message)
		}
	}

	return result
}

func (field *BaseField) GetValue() string {
	return field.Value
}

func (field *BaseField) SetValue(value string) {
	field.Value = value
}

func (field *BaseField) IsName(name string) bool {
	return field.Name == name
}

func (field *BaseField) RenderFull(attrs []string) template.HTML {
	return template.HTML("")
}

func (field *BaseField) Errors() []string {
	return field.errors
}

type TextField struct {
	BaseField
}

func (field *TextField) RenderInput(attrs ...string) template.HTML {
	attrsStr := ""
	if len(attrs) > 0 {
		for _, attr := range attrs {
			attrsStr += " " + attr
		}
	}
	return template.HTML(fmt.Sprintf(`<input type="text" value="%s" name=%q id=%q%s>`, field.Value, field.Name, field.Name, attrsStr))
}

func NewTextField(name string, label string, value string, validators ...IValidator) *TextField {
	field := TextField{}
	field.Name = name
	field.Label = label
	field.Value = value
	field.Validators = validators

	return &field
}

type PasswordField struct {
	BaseField
}

func (field *PasswordField) RenderInput(attrs ...string) template.HTML {
	attrsStr := ""
	if len(attrs) > 0 {
		for _, attr := range attrs {
			attrsStr += " " + attr
		}
	}
	return template.HTML(fmt.Sprintf(`<input type="password" name=%q id=%q%s>`, field.Name, field.Name, attrsStr))
}

func NewPasswordField(name string, label string, validators ...IValidator) *PasswordField {
	field := PasswordField{}
	field.Name = name
	field.Label = label
	field.Validators = validators

	return &field
}

type TextArea struct {
	BaseField
}

func (field *TextArea) RenderInput(attrs ...string) template.HTML {
	attrsStr := ""
	if len(attrs) > 0 {
		for _, attr := range attrs {
			attrsStr += " " + attr
		}
	}

	return template.HTML(fmt.Sprintf(`<textarea id=%q name=%q%s>%s</textarea>`, field.Name, field.Name, attrsStr, field.Value))
}

func NewTextArea(name string, label string, value string, validators ...IValidator) *TextArea {
	field := TextArea{}
	field.Name = name
	field.Label = label
	field.Value = value
	field.Validators = validators

	return &field
}

type Choice struct {
	Value string
	Label string
}

type SelectField struct {
	BaseField
	Choices []Choice
}

func (field *SelectField) RenderInput(attrs ...string) template.HTML {
	attrsStr := ""
	if len(attrs) > 0 {
		for _, attr := range attrs {
			attrsStr += " " + attr
		}
	}
	options := ""
	for _, choice := range field.Choices {
		selected := ""
		if choice.Value == field.Value {
			selected = " selected"
		}
		options += fmt.Sprintf(`<option value=%q%s>%s</option>`, choice.Value, selected, choice.Label)
	}

	return template.HTML(fmt.Sprintf(`<select id=%q name=%q%s>%s</select>`, field.Name, field.Name, attrsStr, options))
}

func NewSelectField(name string, label string, choices []Choice, defaultValue string, validators ...IValidator) *SelectField {
	field := SelectField{}
	field.Name = name
	field.Label = label
	field.Value = defaultValue
	field.Choices = choices
	field.Validators = validators

	return &field
}

type HiddenField struct {
	BaseField
}

func (field *HiddenField) RenderInput(attrs ...string) template.HTML {
	return template.HTML(fmt.Sprintf(`<input type="hidden" value=%q name=%q id=%q>`, field.Value, field.Name, field.Name))
}

func NewHiddenField(name string, value string) *HiddenField {
	field := HiddenField{}
	field.Name = name
	field.Value = value

	return &field
}
