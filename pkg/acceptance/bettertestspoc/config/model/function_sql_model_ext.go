package model

import (
	tfconfig "github.com/hashicorp/terraform-plugin-testing/config"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance/bettertestspoc/config"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk/datatypes"
)

func FunctionSqlBasicInline(resourceName string, id sdk.SchemaObjectIdentifierWithArguments, functionDefinition string, returnType string) *FunctionSqlModel {
	return FunctionSql(resourceName, id.DatabaseName(), functionDefinition, id.Name(), returnType, id.SchemaName()).
		WithFunctionDefinitionValue(config.MultilineWrapperVariable(functionDefinition))
}

func (f *FunctionSqlModel) WithArgument(argName string, argDataType datatypes.DataType) *FunctionSqlModel {
	return f.WithArgumentsValue(
		tfconfig.ObjectVariable(
			map[string]tfconfig.Variable{
				"arg_name":      tfconfig.StringVariable(argName),
				"arg_data_type": tfconfig.StringVariable(argDataType.ToSql()),
			},
		),
	)
}
