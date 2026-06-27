package openapi

// Compatibility helpers for consumers migrating from v0.x of sv-tools/openapi/spec.

// NewOperation creates a new extendable Operation.
func NewOperation() *Extendable[Operation] {
	return NewExtendable(&Operation{})
}

// NewResponses creates a new extendable Responses.
func NewResponses() *Extendable[Responses] {
	return NewExtendable(&Responses{})
}

// NewRequestBodySpec creates a new RefOrSpec for Extendable[RequestBody].
func NewRequestBodySpec() *RefOrSpec[Extendable[RequestBody]] {
	return NewRefOrExtSpec[RequestBody](&RequestBody{})
}

// NewMediaType creates a new extendable MediaType.
func NewMediaType() *Extendable[MediaType] {
	return NewExtendable(&MediaType{})
}

// NewSchemaSpec creates a new RefOrSpec for Schema.
func NewSchemaSpec() *RefOrSpec[Schema] {
	return NewRefOrSpec[Schema](&Schema{})
}

// NewServer creates a new extendable Server.
func NewServer() *Extendable[Server] {
	return NewExtendable(&Server{})
}

// NewRefOrSpecCompat creates a RefOrSpec from an optional ref and spec,
// matching the v0.x two-argument signature.
func NewRefOrSpecCompat[T any](ref *Ref, spec *T) *RefOrSpec[T] {
	if ref != nil {
		return NewRefOrSpec[T](ref)
	}
	return NewRefOrSpec[T](spec)
}

// NewBoolOrSchemaCompat creates a BoolOrSchema from a bool and optional schema,
// matching the v0.x two-argument signature.
func NewBoolOrSchemaCompat(b bool, schema *RefOrSpec[Schema]) *BoolOrSchema {
	if schema != nil {
		return NewBoolOrSchema(schema)
	}
	return NewBoolOrSchema(b)
}

// NewSchemaRef creates a RefOrSpec for Schema from a Ref.
func NewSchemaRef(ref *Ref) *RefOrSpec[Schema] {
	return NewRefOrSpec[Schema](ref)
}

// NewRef creates a Ref from a string.
func NewRef(ref string) *Ref {
	return &Ref{Ref: ref}
}

// NewHeaderSpec creates a RefOrSpec for Extendable[Header].
func NewHeaderSpec() *RefOrSpec[Extendable[Header]] {
	return NewRefOrExtSpec[Header](&Header{})
}

// NewResponseSpec creates a RefOrSpec for Extendable[Response].
func NewResponseSpec() *RefOrSpec[Extendable[Response]] {
	return NewRefOrExtSpec[Response](&Response{})
}

// NewDiscriminator creates a new Discriminator.
func NewDiscriminator() *Discriminator {
	return &Discriminator{}
}

// NewContact creates a new extendable Contact.
func NewContact() *Extendable[Contact] {
	return NewExtendable(&Contact{})
}

// NewExternalDocs creates a new extendable ExternalDocs.
func NewExternalDocs() *Extendable[ExternalDocs] {
	return NewExtendable(&ExternalDocs{})
}

// NewInfo creates a new extendable Info.
func NewInfo() *Extendable[Info] {
	return NewExtendable(&Info{})
}

// NewLicense creates a new extendable License.
func NewLicense() *Extendable[License] {
	return NewExtendable(&License{})
}

// NewSecuritySchemeSpec creates a RefOrSpec for Extendable[SecurityScheme].
func NewSecuritySchemeSpec() *RefOrSpec[Extendable[SecurityScheme]] {
	return NewRefOrExtSpec[SecurityScheme](&SecurityScheme{})
}

// NewServerVariable creates a new extendable ServerVariable.
func NewServerVariable() *Extendable[ServerVariable] {
	return NewExtendable(&ServerVariable{})
}

// NewOAuthFlow creates a new extendable OAuthFlow.
func NewOAuthFlow() *Extendable[OAuthFlow] {
	return NewExtendable(&OAuthFlow{})
}

// NewOAuthFlows creates a new extendable OAuthFlows.
func NewOAuthFlows() *Extendable[OAuthFlows] {
	return NewExtendable(&OAuthFlows{})
}
