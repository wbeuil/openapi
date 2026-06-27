package openapi

// Tag adds metadata to a single tag that is used by the Operation Object.
// It is not mandatory to have a Tag Object per tag defined in the Operation Object instances.
//
// https://spec.openapis.org/oas/v3.1.1#tag-object
//
// Example:
//
//	name: pet
//	description: Pets operations
type Tag struct {
	// Additional external documentation for this tag.
	ExternalDocs *Extendable[ExternalDocs] `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	// REQUIRED.
	// The name of the tag.
	Name string `json:"name" yaml:"name"`
	// A short summary of the tag, used for display purposes.
	Summary string `json:"summary,omitempty" yaml:"summary,omitempty"`
	// A description for the tag.
	// CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// The name of a tag that this tag is nested under.
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`
	// A machine-readable string to categorize what sort of tag it is.
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
}

func (o *Tag) validateSpec(location string, validator *Validator) []*validationError {
	var errs []*validationError
	if o.Name == "" {
		errs = append(errs, newValidationError(joinLoc(location, "name"), ErrRequired))
	}
	if o.ExternalDocs != nil {
		errs = append(errs, o.ExternalDocs.validateSpec(joinLoc(location, "externalDocs"), validator)...)
	}
	validator.visited[joinLoc("tags", o.Name)] = true
	return errs
}

type TagBuilder struct {
	spec *Extendable[Tag]
}

func NewTagBuilder() *TagBuilder {
	return &TagBuilder{
		spec: NewExtendable[Tag](&Tag{}),
	}
}

func (b *TagBuilder) Build() *Extendable[Tag] {
	return b.spec
}

func (b *TagBuilder) Extensions(v map[string]any) *TagBuilder {
	b.spec.Extensions = v
	return b
}

func (b *TagBuilder) AddExt(name string, value any) *TagBuilder {
	b.spec.AddExt(name, value)
	return b
}

func (b *TagBuilder) ExternalDocs(v *Extendable[ExternalDocs]) *TagBuilder {
	b.spec.Spec.ExternalDocs = v
	return b
}

func (b *TagBuilder) Name(v string) *TagBuilder {
	b.spec.Spec.Name = v
	return b
}

func (b *TagBuilder) Description(v string) *TagBuilder {
	b.spec.Spec.Description = v
	return b
}

func (b *TagBuilder) Summary(v string) *TagBuilder {
	b.spec.Spec.Summary = v
	return b
}

func (b *TagBuilder) Parent(v string) *TagBuilder {
	b.spec.Spec.Parent = v
	return b
}

func (b *TagBuilder) Kind(v string) *TagBuilder {
	b.spec.Spec.Kind = v
	return b
}
