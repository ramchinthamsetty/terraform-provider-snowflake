// Code generated by assertions generator; DO NOT EDIT.

package resourceassert

import (
	"testing"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance/bettertestspoc/assert"
)

type FunctionJavascriptResourceAssert struct {
	*assert.ResourceAssert
}

func FunctionJavascriptResource(t *testing.T, name string) *FunctionJavascriptResourceAssert {
	t.Helper()

	return &FunctionJavascriptResourceAssert{
		ResourceAssert: assert.NewResourceAssert(name, "resource"),
	}
}

func ImportedFunctionJavascriptResource(t *testing.T, id string) *FunctionJavascriptResourceAssert {
	t.Helper()

	return &FunctionJavascriptResourceAssert{
		ResourceAssert: assert.NewImportedResourceAssert(id, "imported resource"),
	}
}

///////////////////////////////////
// Attribute value string checks //
///////////////////////////////////

func (f *FunctionJavascriptResourceAssert) HasArgumentsString(expected string) *FunctionJavascriptResourceAssert {
	f.AddAssertion(assert.ValueSet("arguments", expected))
	return f
}

func (f *FunctionJavascriptResourceAssert) HasCommentString(expected string) *FunctionJavascriptResourceAssert {
	f.AddAssertion(assert.ValueSet("comment", expected))
	return f
}

func (f *FunctionJavascriptResourceAssert) HasDatabaseString(expected string) *FunctionJavascriptResourceAssert {
	f.AddAssertion(assert.ValueSet("database", expected))
	return f
}

func (f *FunctionJavascriptResourceAssert) HasEnableConsoleOutputString(expected string) *FunctionJavascriptResourceAssert {
	f.AddAssertion(assert.ValueSet("enable_console_output", expected))
	return f
}

func (f *FunctionJavascriptResourceAssert) HasFullyQualifiedNameString(expected string) *FunctionJavascriptResourceAssert {
	f.AddAssertion(assert.ValueSet("fully_qualified_name", expected))
	return f
}

func (f *FunctionJavascriptResourceAssert) HasFunctionDefinitionString(expected string) *FunctionJavascriptResourceAssert {
	f.AddAssertion(assert.ValueSet("function_definition", expected))
	return f
}

func (f *FunctionJavascriptResourceAssert) HasFunctionLanguageString(expected string) *FunctionJavascriptResourceAssert {
	f.AddAssertion(assert.ValueSet("function_language", expected))
	return f
}

func (f *FunctionJavascriptResourceAssert) HasIsSecureString(expected string) *FunctionJavascriptResourceAssert {
	f.AddAssertion(assert.ValueSet("is_secure", expected))
	return f
}

func (f *FunctionJavascriptResourceAssert) HasLogLevelString(expected string) *FunctionJavascriptResourceAssert {
	f.AddAssertion(assert.ValueSet("log_level", expected))
	return f
}

func (f *FunctionJavascriptResourceAssert) HasMetricLevelString(expected string) *FunctionJavascriptResourceAssert {
	f.AddAssertion(assert.ValueSet("metric_level", expected))
	return f
}

func (f *FunctionJavascriptResourceAssert) HasNameString(expected string) *FunctionJavascriptResourceAssert {
	f.AddAssertion(assert.ValueSet("name", expected))
	return f
}

func (f *FunctionJavascriptResourceAssert) HasNullInputBehaviorString(expected string) *FunctionJavascriptResourceAssert {
	f.AddAssertion(assert.ValueSet("null_input_behavior", expected))
	return f
}

func (f *FunctionJavascriptResourceAssert) HasReturnBehaviorString(expected string) *FunctionJavascriptResourceAssert {
	f.AddAssertion(assert.ValueSet("return_results_behavior", expected))
	return f
}

func (f *FunctionJavascriptResourceAssert) HasReturnTypeString(expected string) *FunctionJavascriptResourceAssert {
	f.AddAssertion(assert.ValueSet("return_type", expected))
	return f
}

func (f *FunctionJavascriptResourceAssert) HasSchemaString(expected string) *FunctionJavascriptResourceAssert {
	f.AddAssertion(assert.ValueSet("schema", expected))
	return f
}

func (f *FunctionJavascriptResourceAssert) HasTraceLevelString(expected string) *FunctionJavascriptResourceAssert {
	f.AddAssertion(assert.ValueSet("trace_level", expected))
	return f
}

////////////////////////////
// Attribute empty checks //
////////////////////////////

func (f *FunctionJavascriptResourceAssert) HasNoArguments() *FunctionJavascriptResourceAssert {
	f.AddAssertion(assert.ValueNotSet("arguments"))
	return f
}

func (f *FunctionJavascriptResourceAssert) HasNoComment() *FunctionJavascriptResourceAssert {
	f.AddAssertion(assert.ValueNotSet("comment"))
	return f
}

func (f *FunctionJavascriptResourceAssert) HasNoDatabase() *FunctionJavascriptResourceAssert {
	f.AddAssertion(assert.ValueNotSet("database"))
	return f
}

func (f *FunctionJavascriptResourceAssert) HasNoEnableConsoleOutput() *FunctionJavascriptResourceAssert {
	f.AddAssertion(assert.ValueNotSet("enable_console_output"))
	return f
}

func (f *FunctionJavascriptResourceAssert) HasNoFullyQualifiedName() *FunctionJavascriptResourceAssert {
	f.AddAssertion(assert.ValueNotSet("fully_qualified_name"))
	return f
}

func (f *FunctionJavascriptResourceAssert) HasNoFunctionDefinition() *FunctionJavascriptResourceAssert {
	f.AddAssertion(assert.ValueNotSet("function_definition"))
	return f
}

func (f *FunctionJavascriptResourceAssert) HasNoFunctionLanguage() *FunctionJavascriptResourceAssert {
	f.AddAssertion(assert.ValueNotSet("function_language"))
	return f
}

func (f *FunctionJavascriptResourceAssert) HasNoIsSecure() *FunctionJavascriptResourceAssert {
	f.AddAssertion(assert.ValueNotSet("is_secure"))
	return f
}

func (f *FunctionJavascriptResourceAssert) HasNoLogLevel() *FunctionJavascriptResourceAssert {
	f.AddAssertion(assert.ValueNotSet("log_level"))
	return f
}

func (f *FunctionJavascriptResourceAssert) HasNoMetricLevel() *FunctionJavascriptResourceAssert {
	f.AddAssertion(assert.ValueNotSet("metric_level"))
	return f
}

func (f *FunctionJavascriptResourceAssert) HasNoName() *FunctionJavascriptResourceAssert {
	f.AddAssertion(assert.ValueNotSet("name"))
	return f
}

func (f *FunctionJavascriptResourceAssert) HasNoNullInputBehavior() *FunctionJavascriptResourceAssert {
	f.AddAssertion(assert.ValueNotSet("null_input_behavior"))
	return f
}

func (f *FunctionJavascriptResourceAssert) HasNoReturnBehavior() *FunctionJavascriptResourceAssert {
	f.AddAssertion(assert.ValueNotSet("return_behavior"))
	return f
}

func (f *FunctionJavascriptResourceAssert) HasNoReturnType() *FunctionJavascriptResourceAssert {
	f.AddAssertion(assert.ValueNotSet("return_type"))
	return f
}

func (f *FunctionJavascriptResourceAssert) HasNoSchema() *FunctionJavascriptResourceAssert {
	f.AddAssertion(assert.ValueNotSet("schema"))
	return f
}

func (f *FunctionJavascriptResourceAssert) HasNoTraceLevel() *FunctionJavascriptResourceAssert {
	f.AddAssertion(assert.ValueNotSet("trace_level"))
	return f
}
