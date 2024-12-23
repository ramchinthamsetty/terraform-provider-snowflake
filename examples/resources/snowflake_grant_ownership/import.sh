### OnObject
`terraform import snowflake_grant_ownership.example '<role_type>|<role_identifier>|<outbound_privileges_behavior>|OnObject|<object_type>|<object_name>'`

### OnAll (contains inner types: InDatabase | InSchema)

#### InDatabase
`terraform import snowflake_grant_ownership.example '<role_type>|<role_identifier>|<outbound_privileges_behavior>|OnAll|<object_type_plural>|InDatabase|<database_name>'`

#### InSchema
`terraform import snowflake_grant_ownership.example '<role_type>|<role_identifier>|<outbound_privileges_behavior>|OnAll|<object_type_plural>|InSchema|<schema_name>'`

### OnFuture (contains inner types: InDatabase | InSchema)

#### InDatabase
`terraform import snowflake_grant_ownership.example '<role_type>|<role_identifier>|<outbound_privileges_behavior>|OnFuture|<object_type_plural>|InDatabase|<database_name>'`

#### InSchema
`terraform import snowflake_grant_ownership.example '<role_type>|<role_identifier>|<outbound_privileges_behavior>|OnFuture|<object_type_plural>|InSchema|<schema_name>'`

### Import examples

#### OnObject on Schema ToAccountRole
`terraform import snowflake_grant_ownership.example 'ToAccountRole|"account_role"|COPY|OnObject|SCHEMA|"database_name"."schema_name"'`

#### OnObject on Schema ToDatabaseRole
`terraform import snowflake_grant_ownership.example 'ToDatabaseRole|"database_name"."database_role_name"|COPY|OnObject|SCHEMA|"database_name"."schema_name"'`

#### OnObject on Table
`terraform import snowflake_grant_ownership.example 'ToAccountRole|"account_role"|COPY|OnObject|TABLE|"database_name"."schema_name"."table_name"'`

#### OnAll InDatabase
`terraform import snowflake_grant_ownership.example 'ToAccountRole|"account_role"|REVOKE|OnAll|TABLES|InDatabase|"database_name"'`

#### OnAll InSchema
`terraform import snowflake_grant_ownership.example 'ToAccountRole|"account_role"||OnAll|TABLES|InSchema|"database_name"."schema_name"'`

#### OnFuture InDatabase
`terraform import snowflake_grant_ownership.example 'ToAccountRole|"account_role"||OnFuture|TABLES|InDatabase|"database_name"'`

#### OnFuture InSchema
`terraform import snowflake_grant_ownership.example 'ToAccountRole|"account_role"|COPY|OnFuture|TABLES|InSchema|"database_name"."schema_name"'`
