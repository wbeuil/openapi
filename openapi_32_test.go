package openapi

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOpenAPI32_Self(t *testing.T) {
	t.Parallel()

	b := NewOpenAPIBuilder().
		OpenAPI("3.2.0").
		Self("https://example.com/api.json").
		Info(NewInfoBuilder().Title("Test").Version("1.0.0").Build()).
		Paths(NewPaths())

	spec := b.Build()
	assert.Equal(t, "3.2.0", spec.Spec.OpenAPI)
	assert.Equal(t, "https://example.com/api.json", spec.Spec.Self)

	data, err := json.Marshal(spec)
	require.NoError(t, err)
	assert.Contains(t, string(data), `"$self":"https://example.com/api.json"`)
}

func TestPathItem32_QueryAndAdditionalOperations(t *testing.T) {
	t.Parallel()

	b := NewPathItemBuilder().
		Query(NewOperationBuilder().Summary("Query pets").Build()).
		AddAdditionalOperation("CUSTOM", NewOperationBuilder().Summary("Custom op").Build())

	item := b.Build()
	assert.NotNil(t, item.Spec.Spec.Query)
	assert.Equal(t, "Query pets", item.Spec.Spec.Query.Spec.Summary)
	assert.Len(t, item.Spec.Spec.AdditionalOperations, 1)
	assert.Equal(t, "Custom op", item.Spec.Spec.AdditionalOperations["CUSTOM"].Spec.Summary)

	data, err := json.Marshal(item)
	require.NoError(t, err)
	assert.Contains(t, string(data), `"query"`)
	assert.Contains(t, string(data), `"additionalOperations"`)
}

func TestTag32_SummaryParentKind(t *testing.T) {
	t.Parallel()

	b := NewTagBuilder().
		Name("pets").
		Summary("Pet operations").
		Parent("animals").
		Kind("nav")

	tag := b.Build()
	assert.Equal(t, "Pet operations", tag.Spec.Summary)
	assert.Equal(t, "animals", tag.Spec.Parent)
	assert.Equal(t, "nav", tag.Spec.Kind)

	data, err := json.Marshal(tag)
	require.NoError(t, err)
	assert.Contains(t, string(data), `"summary":"Pet operations"`)
	assert.Contains(t, string(data), `"parent":"animals"`)
	assert.Contains(t, string(data), `"kind":"nav"`)
}

func TestMediaType32_ItemSchemaPrefixEncodingItemEncoding(t *testing.T) {
	t.Parallel()

	b := NewMediaTypeBuilder().
		ItemSchema(NewSchemaBuilder().AddType("string").Build()).
		AddPrefixEncoding(NewEncodingBuilder().ContentType("text/plain").Build()).
		ItemEncoding(NewEncodingBuilder().ContentType("application/json").Build())

	mt := b.Build()
	assert.NotNil(t, mt.Spec.ItemSchema)
	assert.Len(t, mt.Spec.PrefixEncoding, 1)
	assert.NotNil(t, mt.Spec.ItemEncoding)

	data, err := json.Marshal(mt)
	require.NoError(t, err)
	assert.Contains(t, string(data), `"itemSchema"`)
	assert.Contains(t, string(data), `"prefixEncoding"`)
	assert.Contains(t, string(data), `"itemEncoding"`)
}

func TestParameter32_QueryString(t *testing.T) {
	t.Parallel()

	b := NewParameterBuilder().
		Name("filter").
		In(InQueryString).
		Content(map[string]*Extendable[MediaType]{
			"application/x-www-form-urlencoded": NewMediaTypeBuilder().Build(),
		})

	param := b.Build()
	assert.Equal(t, InQueryString, param.Spec.Spec.In)

	data, err := json.Marshal(param)
	require.NoError(t, err)
	assert.Contains(t, string(data), `"in":"querystring"`)
}
