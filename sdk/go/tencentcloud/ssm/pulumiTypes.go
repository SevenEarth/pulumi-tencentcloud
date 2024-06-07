// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssm

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

var _ = internal.GetEnvOrDefault

type ProductSecretPrivilegesList struct {
	// This value takes effect only when `PrivilegeName` is `ColumnPrivileges`, and the following parameters are required in this case:Database: explicitly indicate the database instance.TableName: explicitly indicate the table.
	ColumnName *string `pulumi:"columnName"`
	// This value takes effect only when `PrivilegeName` is `DatabasePrivileges`.
	Database *string `pulumi:"database"`
	// Permission name. Valid values: `GlobalPrivileges`, `DatabasePrivileges`, `TablePrivileges`, `ColumnPrivileges`. When the permission is `DatabasePrivileges`, the database name must be specified by the `Database` parameter; When the permission is `TablePrivileges`, the database name and the table name in the database must be specified by the `Database` and `TableName` parameters; When the permission is `ColumnPrivileges`, the database name, table name in the database, and column name in the table must be specified by the `Database`, `TableName`, and `ColumnName` parameters.
	PrivilegeName string `pulumi:"privilegeName"`
	// Permission list. For the `Mysql` service, optional permission values are: 1. Valid values of `GlobalPrivileges`: SELECT,INSERT,UPDATE,DELETE,CREATE, PROCESS, DROP,REFERENCES,INDEX,ALTER,SHOW DATABASES,CREATE TEMPORARY TABLES,LOCK TABLES,EXECUTE,CREATE VIEW,SHOW VIEW,CREATE ROUTINE,ALTER ROUTINE,EVENT,TRIGGER. Note: if this parameter is not passed in, it means to clear the permission. 2. Valid values of `DatabasePrivileges`: SELECT,INSERT,UPDATE,DELETE,CREATE, DROP,REFERENCES,INDEX,ALTER,CREATE TEMPORARY TABLES,LOCK TABLES,EXECUTE,CREATE VIEW,SHOW VIEW,CREATE ROUTINE,ALTER ROUTINE,EVENT,TRIGGER. Note: if this parameter is not passed in, it means to clear the permission. 3. Valid values of `TablePrivileges`: SELECT,INSERT,UPDATE,DELETE,CREATE, DROP,REFERENCES,INDEX,ALTER,CREATE VIEW,SHOW VIEW, TRIGGER. Note: if this parameter is not passed in, it means to clear the permission. 4. Valid values of `ColumnPrivileges`: SELECT,INSERT,UPDATE,REFERENCES.Note: if this parameter is not passed in, it means to clear the permission.
	Privileges []string `pulumi:"privileges"`
	// This value takes effect only when `PrivilegeName` is `TablePrivileges`, and the `Database` parameter is required in this case to explicitly indicate the database instance.
	TableName *string `pulumi:"tableName"`
}

// ProductSecretPrivilegesListInput is an input type that accepts ProductSecretPrivilegesListArgs and ProductSecretPrivilegesListOutput values.
// You can construct a concrete instance of `ProductSecretPrivilegesListInput` via:
//
//	ProductSecretPrivilegesListArgs{...}
type ProductSecretPrivilegesListInput interface {
	pulumi.Input

	ToProductSecretPrivilegesListOutput() ProductSecretPrivilegesListOutput
	ToProductSecretPrivilegesListOutputWithContext(context.Context) ProductSecretPrivilegesListOutput
}

type ProductSecretPrivilegesListArgs struct {
	// This value takes effect only when `PrivilegeName` is `ColumnPrivileges`, and the following parameters are required in this case:Database: explicitly indicate the database instance.TableName: explicitly indicate the table.
	ColumnName pulumi.StringPtrInput `pulumi:"columnName"`
	// This value takes effect only when `PrivilegeName` is `DatabasePrivileges`.
	Database pulumi.StringPtrInput `pulumi:"database"`
	// Permission name. Valid values: `GlobalPrivileges`, `DatabasePrivileges`, `TablePrivileges`, `ColumnPrivileges`. When the permission is `DatabasePrivileges`, the database name must be specified by the `Database` parameter; When the permission is `TablePrivileges`, the database name and the table name in the database must be specified by the `Database` and `TableName` parameters; When the permission is `ColumnPrivileges`, the database name, table name in the database, and column name in the table must be specified by the `Database`, `TableName`, and `ColumnName` parameters.
	PrivilegeName pulumi.StringInput `pulumi:"privilegeName"`
	// Permission list. For the `Mysql` service, optional permission values are: 1. Valid values of `GlobalPrivileges`: SELECT,INSERT,UPDATE,DELETE,CREATE, PROCESS, DROP,REFERENCES,INDEX,ALTER,SHOW DATABASES,CREATE TEMPORARY TABLES,LOCK TABLES,EXECUTE,CREATE VIEW,SHOW VIEW,CREATE ROUTINE,ALTER ROUTINE,EVENT,TRIGGER. Note: if this parameter is not passed in, it means to clear the permission. 2. Valid values of `DatabasePrivileges`: SELECT,INSERT,UPDATE,DELETE,CREATE, DROP,REFERENCES,INDEX,ALTER,CREATE TEMPORARY TABLES,LOCK TABLES,EXECUTE,CREATE VIEW,SHOW VIEW,CREATE ROUTINE,ALTER ROUTINE,EVENT,TRIGGER. Note: if this parameter is not passed in, it means to clear the permission. 3. Valid values of `TablePrivileges`: SELECT,INSERT,UPDATE,DELETE,CREATE, DROP,REFERENCES,INDEX,ALTER,CREATE VIEW,SHOW VIEW, TRIGGER. Note: if this parameter is not passed in, it means to clear the permission. 4. Valid values of `ColumnPrivileges`: SELECT,INSERT,UPDATE,REFERENCES.Note: if this parameter is not passed in, it means to clear the permission.
	Privileges pulumi.StringArrayInput `pulumi:"privileges"`
	// This value takes effect only when `PrivilegeName` is `TablePrivileges`, and the `Database` parameter is required in this case to explicitly indicate the database instance.
	TableName pulumi.StringPtrInput `pulumi:"tableName"`
}

func (ProductSecretPrivilegesListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProductSecretPrivilegesList)(nil)).Elem()
}

func (i ProductSecretPrivilegesListArgs) ToProductSecretPrivilegesListOutput() ProductSecretPrivilegesListOutput {
	return i.ToProductSecretPrivilegesListOutputWithContext(context.Background())
}

func (i ProductSecretPrivilegesListArgs) ToProductSecretPrivilegesListOutputWithContext(ctx context.Context) ProductSecretPrivilegesListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProductSecretPrivilegesListOutput)
}

// ProductSecretPrivilegesListArrayInput is an input type that accepts ProductSecretPrivilegesListArray and ProductSecretPrivilegesListArrayOutput values.
// You can construct a concrete instance of `ProductSecretPrivilegesListArrayInput` via:
//
//	ProductSecretPrivilegesListArray{ ProductSecretPrivilegesListArgs{...} }
type ProductSecretPrivilegesListArrayInput interface {
	pulumi.Input

	ToProductSecretPrivilegesListArrayOutput() ProductSecretPrivilegesListArrayOutput
	ToProductSecretPrivilegesListArrayOutputWithContext(context.Context) ProductSecretPrivilegesListArrayOutput
}

type ProductSecretPrivilegesListArray []ProductSecretPrivilegesListInput

func (ProductSecretPrivilegesListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProductSecretPrivilegesList)(nil)).Elem()
}

func (i ProductSecretPrivilegesListArray) ToProductSecretPrivilegesListArrayOutput() ProductSecretPrivilegesListArrayOutput {
	return i.ToProductSecretPrivilegesListArrayOutputWithContext(context.Background())
}

func (i ProductSecretPrivilegesListArray) ToProductSecretPrivilegesListArrayOutputWithContext(ctx context.Context) ProductSecretPrivilegesListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProductSecretPrivilegesListArrayOutput)
}

type ProductSecretPrivilegesListOutput struct{ *pulumi.OutputState }

func (ProductSecretPrivilegesListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProductSecretPrivilegesList)(nil)).Elem()
}

func (o ProductSecretPrivilegesListOutput) ToProductSecretPrivilegesListOutput() ProductSecretPrivilegesListOutput {
	return o
}

func (o ProductSecretPrivilegesListOutput) ToProductSecretPrivilegesListOutputWithContext(ctx context.Context) ProductSecretPrivilegesListOutput {
	return o
}

// This value takes effect only when `PrivilegeName` is `ColumnPrivileges`, and the following parameters are required in this case:Database: explicitly indicate the database instance.TableName: explicitly indicate the table.
func (o ProductSecretPrivilegesListOutput) ColumnName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProductSecretPrivilegesList) *string { return v.ColumnName }).(pulumi.StringPtrOutput)
}

// This value takes effect only when `PrivilegeName` is `DatabasePrivileges`.
func (o ProductSecretPrivilegesListOutput) Database() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProductSecretPrivilegesList) *string { return v.Database }).(pulumi.StringPtrOutput)
}

// Permission name. Valid values: `GlobalPrivileges`, `DatabasePrivileges`, `TablePrivileges`, `ColumnPrivileges`. When the permission is `DatabasePrivileges`, the database name must be specified by the `Database` parameter; When the permission is `TablePrivileges`, the database name and the table name in the database must be specified by the `Database` and `TableName` parameters; When the permission is `ColumnPrivileges`, the database name, table name in the database, and column name in the table must be specified by the `Database`, `TableName`, and `ColumnName` parameters.
func (o ProductSecretPrivilegesListOutput) PrivilegeName() pulumi.StringOutput {
	return o.ApplyT(func(v ProductSecretPrivilegesList) string { return v.PrivilegeName }).(pulumi.StringOutput)
}

// Permission list. For the `Mysql` service, optional permission values are: 1. Valid values of `GlobalPrivileges`: SELECT,INSERT,UPDATE,DELETE,CREATE, PROCESS, DROP,REFERENCES,INDEX,ALTER,SHOW DATABASES,CREATE TEMPORARY TABLES,LOCK TABLES,EXECUTE,CREATE VIEW,SHOW VIEW,CREATE ROUTINE,ALTER ROUTINE,EVENT,TRIGGER. Note: if this parameter is not passed in, it means to clear the permission. 2. Valid values of `DatabasePrivileges`: SELECT,INSERT,UPDATE,DELETE,CREATE, DROP,REFERENCES,INDEX,ALTER,CREATE TEMPORARY TABLES,LOCK TABLES,EXECUTE,CREATE VIEW,SHOW VIEW,CREATE ROUTINE,ALTER ROUTINE,EVENT,TRIGGER. Note: if this parameter is not passed in, it means to clear the permission. 3. Valid values of `TablePrivileges`: SELECT,INSERT,UPDATE,DELETE,CREATE, DROP,REFERENCES,INDEX,ALTER,CREATE VIEW,SHOW VIEW, TRIGGER. Note: if this parameter is not passed in, it means to clear the permission. 4. Valid values of `ColumnPrivileges`: SELECT,INSERT,UPDATE,REFERENCES.Note: if this parameter is not passed in, it means to clear the permission.
func (o ProductSecretPrivilegesListOutput) Privileges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ProductSecretPrivilegesList) []string { return v.Privileges }).(pulumi.StringArrayOutput)
}

// This value takes effect only when `PrivilegeName` is `TablePrivileges`, and the `Database` parameter is required in this case to explicitly indicate the database instance.
func (o ProductSecretPrivilegesListOutput) TableName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProductSecretPrivilegesList) *string { return v.TableName }).(pulumi.StringPtrOutput)
}

type ProductSecretPrivilegesListArrayOutput struct{ *pulumi.OutputState }

func (ProductSecretPrivilegesListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProductSecretPrivilegesList)(nil)).Elem()
}

func (o ProductSecretPrivilegesListArrayOutput) ToProductSecretPrivilegesListArrayOutput() ProductSecretPrivilegesListArrayOutput {
	return o
}

func (o ProductSecretPrivilegesListArrayOutput) ToProductSecretPrivilegesListArrayOutputWithContext(ctx context.Context) ProductSecretPrivilegesListArrayOutput {
	return o
}

func (o ProductSecretPrivilegesListArrayOutput) Index(i pulumi.IntInput) ProductSecretPrivilegesListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProductSecretPrivilegesList {
		return vs[0].([]ProductSecretPrivilegesList)[vs[1].(int)]
	}).(ProductSecretPrivilegesListOutput)
}

type GetSecretVersionsSecretVersionList struct {
	// The base64-encoded binary secret.
	SecretBinary string `pulumi:"secretBinary"`
	// The string text of secret.
	SecretString string `pulumi:"secretString"`
	// VersionId used to filter result.
	VersionId string `pulumi:"versionId"`
}

// GetSecretVersionsSecretVersionListInput is an input type that accepts GetSecretVersionsSecretVersionListArgs and GetSecretVersionsSecretVersionListOutput values.
// You can construct a concrete instance of `GetSecretVersionsSecretVersionListInput` via:
//
//	GetSecretVersionsSecretVersionListArgs{...}
type GetSecretVersionsSecretVersionListInput interface {
	pulumi.Input

	ToGetSecretVersionsSecretVersionListOutput() GetSecretVersionsSecretVersionListOutput
	ToGetSecretVersionsSecretVersionListOutputWithContext(context.Context) GetSecretVersionsSecretVersionListOutput
}

type GetSecretVersionsSecretVersionListArgs struct {
	// The base64-encoded binary secret.
	SecretBinary pulumi.StringInput `pulumi:"secretBinary"`
	// The string text of secret.
	SecretString pulumi.StringInput `pulumi:"secretString"`
	// VersionId used to filter result.
	VersionId pulumi.StringInput `pulumi:"versionId"`
}

func (GetSecretVersionsSecretVersionListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecretVersionsSecretVersionList)(nil)).Elem()
}

func (i GetSecretVersionsSecretVersionListArgs) ToGetSecretVersionsSecretVersionListOutput() GetSecretVersionsSecretVersionListOutput {
	return i.ToGetSecretVersionsSecretVersionListOutputWithContext(context.Background())
}

func (i GetSecretVersionsSecretVersionListArgs) ToGetSecretVersionsSecretVersionListOutputWithContext(ctx context.Context) GetSecretVersionsSecretVersionListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetSecretVersionsSecretVersionListOutput)
}

// GetSecretVersionsSecretVersionListArrayInput is an input type that accepts GetSecretVersionsSecretVersionListArray and GetSecretVersionsSecretVersionListArrayOutput values.
// You can construct a concrete instance of `GetSecretVersionsSecretVersionListArrayInput` via:
//
//	GetSecretVersionsSecretVersionListArray{ GetSecretVersionsSecretVersionListArgs{...} }
type GetSecretVersionsSecretVersionListArrayInput interface {
	pulumi.Input

	ToGetSecretVersionsSecretVersionListArrayOutput() GetSecretVersionsSecretVersionListArrayOutput
	ToGetSecretVersionsSecretVersionListArrayOutputWithContext(context.Context) GetSecretVersionsSecretVersionListArrayOutput
}

type GetSecretVersionsSecretVersionListArray []GetSecretVersionsSecretVersionListInput

func (GetSecretVersionsSecretVersionListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetSecretVersionsSecretVersionList)(nil)).Elem()
}

func (i GetSecretVersionsSecretVersionListArray) ToGetSecretVersionsSecretVersionListArrayOutput() GetSecretVersionsSecretVersionListArrayOutput {
	return i.ToGetSecretVersionsSecretVersionListArrayOutputWithContext(context.Background())
}

func (i GetSecretVersionsSecretVersionListArray) ToGetSecretVersionsSecretVersionListArrayOutputWithContext(ctx context.Context) GetSecretVersionsSecretVersionListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetSecretVersionsSecretVersionListArrayOutput)
}

type GetSecretVersionsSecretVersionListOutput struct{ *pulumi.OutputState }

func (GetSecretVersionsSecretVersionListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecretVersionsSecretVersionList)(nil)).Elem()
}

func (o GetSecretVersionsSecretVersionListOutput) ToGetSecretVersionsSecretVersionListOutput() GetSecretVersionsSecretVersionListOutput {
	return o
}

func (o GetSecretVersionsSecretVersionListOutput) ToGetSecretVersionsSecretVersionListOutputWithContext(ctx context.Context) GetSecretVersionsSecretVersionListOutput {
	return o
}

// The base64-encoded binary secret.
func (o GetSecretVersionsSecretVersionListOutput) SecretBinary() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecretVersionsSecretVersionList) string { return v.SecretBinary }).(pulumi.StringOutput)
}

// The string text of secret.
func (o GetSecretVersionsSecretVersionListOutput) SecretString() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecretVersionsSecretVersionList) string { return v.SecretString }).(pulumi.StringOutput)
}

// VersionId used to filter result.
func (o GetSecretVersionsSecretVersionListOutput) VersionId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecretVersionsSecretVersionList) string { return v.VersionId }).(pulumi.StringOutput)
}

type GetSecretVersionsSecretVersionListArrayOutput struct{ *pulumi.OutputState }

func (GetSecretVersionsSecretVersionListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetSecretVersionsSecretVersionList)(nil)).Elem()
}

func (o GetSecretVersionsSecretVersionListArrayOutput) ToGetSecretVersionsSecretVersionListArrayOutput() GetSecretVersionsSecretVersionListArrayOutput {
	return o
}

func (o GetSecretVersionsSecretVersionListArrayOutput) ToGetSecretVersionsSecretVersionListArrayOutputWithContext(ctx context.Context) GetSecretVersionsSecretVersionListArrayOutput {
	return o
}

func (o GetSecretVersionsSecretVersionListArrayOutput) Index(i pulumi.IntInput) GetSecretVersionsSecretVersionListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetSecretVersionsSecretVersionList {
		return vs[0].([]GetSecretVersionsSecretVersionList)[vs[1].(int)]
	}).(GetSecretVersionsSecretVersionListOutput)
}

type GetSecretsSecretList struct {
	// When the credential type is SSH key pair credential, this field is valid and is used to represent the CVM instance ID associated with the SSH key pair.
	AssociatedInstanceIds []string `pulumi:"associatedInstanceIds"`
	// Create time of secret.
	CreateTime int `pulumi:"createTime"`
	// Uin of Creator.
	CreateUin int `pulumi:"createUin"`
	// Delete time of CMK.
	DeleteTime int `pulumi:"deleteTime"`
	// Description of secret.
	Description string `pulumi:"description"`
	// KMS keyId used to encrypt secret.
	KmsKeyId string `pulumi:"kmsKeyId"`
	// KMS CMK type used to encrypt credentials, DEFAULT represents the default key created by SecretsManager, and CUSTOMER represents the user specified key.
	KmsKeyType string `pulumi:"kmsKeyType"`
	// Next rotation start time, uinx timestamp.
	NextRotationTime int `pulumi:"nextRotationTime"`
	// This parameter only takes effect when the SecretType parameter value is 1. When the SecretType value is 1, if the Product Name value is empty, it means to query all types of cloud product credentials. If the Product Name value is MySQL, it means to query MySQL database credentials. If the Product Name value is Tdsql mysql, it means to query Tdsql (MySQL version) credentials.
	ProductName string `pulumi:"productName"`
	// When the credential type is SSH key pair credential, this field is valid and represents the item ID to which the SSH key pair belongs.
	ProjectId int `pulumi:"projectId"`
	// The cloud product instance ID number corresponding to the cloud product credentials.
	ResourceId string `pulumi:"resourceId"`
	// When the credential type is SSH key pair credential, this field is valid and is used to represent the name of the SSH key pair credential.
	ResourceName string `pulumi:"resourceName"`
	// The user specified rotation start time.
	RotationBeginTime string `pulumi:"rotationBeginTime"`
	// The frequency of rotation, in days, takes effect when rotation is on.
	RotationFrequency int `pulumi:"rotationFrequency"`
	// 1: - Turn on the rotation; 0- No rotation Note: This field may return null, indicating that a valid value cannot be obtained.
	RotationStatus int `pulumi:"rotationStatus"`
	// Secret name used to filter result.
	SecretName string `pulumi:"secretName"`
	// 0- represents user-defined credentials, defaults to 0. 1- represents the user's cloud product credentials. 2- represents SSH key pair credentials. 3- represents cloud API key pair credentials.
	SecretType int `pulumi:"secretType"`
	// Status of secret.
	Status string `pulumi:"status"`
	// When the credential type is a cloud API key pair credential, this field is valid and is used to represent the user UIN to which the cloud API key pair belongs.
	TargetUin int `pulumi:"targetUin"`
}

// GetSecretsSecretListInput is an input type that accepts GetSecretsSecretListArgs and GetSecretsSecretListOutput values.
// You can construct a concrete instance of `GetSecretsSecretListInput` via:
//
//	GetSecretsSecretListArgs{...}
type GetSecretsSecretListInput interface {
	pulumi.Input

	ToGetSecretsSecretListOutput() GetSecretsSecretListOutput
	ToGetSecretsSecretListOutputWithContext(context.Context) GetSecretsSecretListOutput
}

type GetSecretsSecretListArgs struct {
	// When the credential type is SSH key pair credential, this field is valid and is used to represent the CVM instance ID associated with the SSH key pair.
	AssociatedInstanceIds pulumi.StringArrayInput `pulumi:"associatedInstanceIds"`
	// Create time of secret.
	CreateTime pulumi.IntInput `pulumi:"createTime"`
	// Uin of Creator.
	CreateUin pulumi.IntInput `pulumi:"createUin"`
	// Delete time of CMK.
	DeleteTime pulumi.IntInput `pulumi:"deleteTime"`
	// Description of secret.
	Description pulumi.StringInput `pulumi:"description"`
	// KMS keyId used to encrypt secret.
	KmsKeyId pulumi.StringInput `pulumi:"kmsKeyId"`
	// KMS CMK type used to encrypt credentials, DEFAULT represents the default key created by SecretsManager, and CUSTOMER represents the user specified key.
	KmsKeyType pulumi.StringInput `pulumi:"kmsKeyType"`
	// Next rotation start time, uinx timestamp.
	NextRotationTime pulumi.IntInput `pulumi:"nextRotationTime"`
	// This parameter only takes effect when the SecretType parameter value is 1. When the SecretType value is 1, if the Product Name value is empty, it means to query all types of cloud product credentials. If the Product Name value is MySQL, it means to query MySQL database credentials. If the Product Name value is Tdsql mysql, it means to query Tdsql (MySQL version) credentials.
	ProductName pulumi.StringInput `pulumi:"productName"`
	// When the credential type is SSH key pair credential, this field is valid and represents the item ID to which the SSH key pair belongs.
	ProjectId pulumi.IntInput `pulumi:"projectId"`
	// The cloud product instance ID number corresponding to the cloud product credentials.
	ResourceId pulumi.StringInput `pulumi:"resourceId"`
	// When the credential type is SSH key pair credential, this field is valid and is used to represent the name of the SSH key pair credential.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
	// The user specified rotation start time.
	RotationBeginTime pulumi.StringInput `pulumi:"rotationBeginTime"`
	// The frequency of rotation, in days, takes effect when rotation is on.
	RotationFrequency pulumi.IntInput `pulumi:"rotationFrequency"`
	// 1: - Turn on the rotation; 0- No rotation Note: This field may return null, indicating that a valid value cannot be obtained.
	RotationStatus pulumi.IntInput `pulumi:"rotationStatus"`
	// Secret name used to filter result.
	SecretName pulumi.StringInput `pulumi:"secretName"`
	// 0- represents user-defined credentials, defaults to 0. 1- represents the user's cloud product credentials. 2- represents SSH key pair credentials. 3- represents cloud API key pair credentials.
	SecretType pulumi.IntInput `pulumi:"secretType"`
	// Status of secret.
	Status pulumi.StringInput `pulumi:"status"`
	// When the credential type is a cloud API key pair credential, this field is valid and is used to represent the user UIN to which the cloud API key pair belongs.
	TargetUin pulumi.IntInput `pulumi:"targetUin"`
}

func (GetSecretsSecretListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecretsSecretList)(nil)).Elem()
}

func (i GetSecretsSecretListArgs) ToGetSecretsSecretListOutput() GetSecretsSecretListOutput {
	return i.ToGetSecretsSecretListOutputWithContext(context.Background())
}

func (i GetSecretsSecretListArgs) ToGetSecretsSecretListOutputWithContext(ctx context.Context) GetSecretsSecretListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetSecretsSecretListOutput)
}

// GetSecretsSecretListArrayInput is an input type that accepts GetSecretsSecretListArray and GetSecretsSecretListArrayOutput values.
// You can construct a concrete instance of `GetSecretsSecretListArrayInput` via:
//
//	GetSecretsSecretListArray{ GetSecretsSecretListArgs{...} }
type GetSecretsSecretListArrayInput interface {
	pulumi.Input

	ToGetSecretsSecretListArrayOutput() GetSecretsSecretListArrayOutput
	ToGetSecretsSecretListArrayOutputWithContext(context.Context) GetSecretsSecretListArrayOutput
}

type GetSecretsSecretListArray []GetSecretsSecretListInput

func (GetSecretsSecretListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetSecretsSecretList)(nil)).Elem()
}

func (i GetSecretsSecretListArray) ToGetSecretsSecretListArrayOutput() GetSecretsSecretListArrayOutput {
	return i.ToGetSecretsSecretListArrayOutputWithContext(context.Background())
}

func (i GetSecretsSecretListArray) ToGetSecretsSecretListArrayOutputWithContext(ctx context.Context) GetSecretsSecretListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetSecretsSecretListArrayOutput)
}

type GetSecretsSecretListOutput struct{ *pulumi.OutputState }

func (GetSecretsSecretListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecretsSecretList)(nil)).Elem()
}

func (o GetSecretsSecretListOutput) ToGetSecretsSecretListOutput() GetSecretsSecretListOutput {
	return o
}

func (o GetSecretsSecretListOutput) ToGetSecretsSecretListOutputWithContext(ctx context.Context) GetSecretsSecretListOutput {
	return o
}

// When the credential type is SSH key pair credential, this field is valid and is used to represent the CVM instance ID associated with the SSH key pair.
func (o GetSecretsSecretListOutput) AssociatedInstanceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSecretsSecretList) []string { return v.AssociatedInstanceIds }).(pulumi.StringArrayOutput)
}

// Create time of secret.
func (o GetSecretsSecretListOutput) CreateTime() pulumi.IntOutput {
	return o.ApplyT(func(v GetSecretsSecretList) int { return v.CreateTime }).(pulumi.IntOutput)
}

// Uin of Creator.
func (o GetSecretsSecretListOutput) CreateUin() pulumi.IntOutput {
	return o.ApplyT(func(v GetSecretsSecretList) int { return v.CreateUin }).(pulumi.IntOutput)
}

// Delete time of CMK.
func (o GetSecretsSecretListOutput) DeleteTime() pulumi.IntOutput {
	return o.ApplyT(func(v GetSecretsSecretList) int { return v.DeleteTime }).(pulumi.IntOutput)
}

// Description of secret.
func (o GetSecretsSecretListOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecretsSecretList) string { return v.Description }).(pulumi.StringOutput)
}

// KMS keyId used to encrypt secret.
func (o GetSecretsSecretListOutput) KmsKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecretsSecretList) string { return v.KmsKeyId }).(pulumi.StringOutput)
}

// KMS CMK type used to encrypt credentials, DEFAULT represents the default key created by SecretsManager, and CUSTOMER represents the user specified key.
func (o GetSecretsSecretListOutput) KmsKeyType() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecretsSecretList) string { return v.KmsKeyType }).(pulumi.StringOutput)
}

// Next rotation start time, uinx timestamp.
func (o GetSecretsSecretListOutput) NextRotationTime() pulumi.IntOutput {
	return o.ApplyT(func(v GetSecretsSecretList) int { return v.NextRotationTime }).(pulumi.IntOutput)
}

// This parameter only takes effect when the SecretType parameter value is 1. When the SecretType value is 1, if the Product Name value is empty, it means to query all types of cloud product credentials. If the Product Name value is MySQL, it means to query MySQL database credentials. If the Product Name value is Tdsql mysql, it means to query Tdsql (MySQL version) credentials.
func (o GetSecretsSecretListOutput) ProductName() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecretsSecretList) string { return v.ProductName }).(pulumi.StringOutput)
}

// When the credential type is SSH key pair credential, this field is valid and represents the item ID to which the SSH key pair belongs.
func (o GetSecretsSecretListOutput) ProjectId() pulumi.IntOutput {
	return o.ApplyT(func(v GetSecretsSecretList) int { return v.ProjectId }).(pulumi.IntOutput)
}

// The cloud product instance ID number corresponding to the cloud product credentials.
func (o GetSecretsSecretListOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecretsSecretList) string { return v.ResourceId }).(pulumi.StringOutput)
}

// When the credential type is SSH key pair credential, this field is valid and is used to represent the name of the SSH key pair credential.
func (o GetSecretsSecretListOutput) ResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecretsSecretList) string { return v.ResourceName }).(pulumi.StringOutput)
}

// The user specified rotation start time.
func (o GetSecretsSecretListOutput) RotationBeginTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecretsSecretList) string { return v.RotationBeginTime }).(pulumi.StringOutput)
}

// The frequency of rotation, in days, takes effect when rotation is on.
func (o GetSecretsSecretListOutput) RotationFrequency() pulumi.IntOutput {
	return o.ApplyT(func(v GetSecretsSecretList) int { return v.RotationFrequency }).(pulumi.IntOutput)
}

// 1: - Turn on the rotation; 0- No rotation Note: This field may return null, indicating that a valid value cannot be obtained.
func (o GetSecretsSecretListOutput) RotationStatus() pulumi.IntOutput {
	return o.ApplyT(func(v GetSecretsSecretList) int { return v.RotationStatus }).(pulumi.IntOutput)
}

// Secret name used to filter result.
func (o GetSecretsSecretListOutput) SecretName() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecretsSecretList) string { return v.SecretName }).(pulumi.StringOutput)
}

// 0- represents user-defined credentials, defaults to 0. 1- represents the user's cloud product credentials. 2- represents SSH key pair credentials. 3- represents cloud API key pair credentials.
func (o GetSecretsSecretListOutput) SecretType() pulumi.IntOutput {
	return o.ApplyT(func(v GetSecretsSecretList) int { return v.SecretType }).(pulumi.IntOutput)
}

// Status of secret.
func (o GetSecretsSecretListOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecretsSecretList) string { return v.Status }).(pulumi.StringOutput)
}

// When the credential type is a cloud API key pair credential, this field is valid and is used to represent the user UIN to which the cloud API key pair belongs.
func (o GetSecretsSecretListOutput) TargetUin() pulumi.IntOutput {
	return o.ApplyT(func(v GetSecretsSecretList) int { return v.TargetUin }).(pulumi.IntOutput)
}

type GetSecretsSecretListArrayOutput struct{ *pulumi.OutputState }

func (GetSecretsSecretListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetSecretsSecretList)(nil)).Elem()
}

func (o GetSecretsSecretListArrayOutput) ToGetSecretsSecretListArrayOutput() GetSecretsSecretListArrayOutput {
	return o
}

func (o GetSecretsSecretListArrayOutput) ToGetSecretsSecretListArrayOutputWithContext(ctx context.Context) GetSecretsSecretListArrayOutput {
	return o
}

func (o GetSecretsSecretListArrayOutput) Index(i pulumi.IntInput) GetSecretsSecretListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetSecretsSecretList {
		return vs[0].([]GetSecretsSecretList)[vs[1].(int)]
	}).(GetSecretsSecretListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProductSecretPrivilegesListInput)(nil)).Elem(), ProductSecretPrivilegesListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProductSecretPrivilegesListArrayInput)(nil)).Elem(), ProductSecretPrivilegesListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetSecretVersionsSecretVersionListInput)(nil)).Elem(), GetSecretVersionsSecretVersionListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetSecretVersionsSecretVersionListArrayInput)(nil)).Elem(), GetSecretVersionsSecretVersionListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetSecretsSecretListInput)(nil)).Elem(), GetSecretsSecretListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetSecretsSecretListArrayInput)(nil)).Elem(), GetSecretsSecretListArray{})
	pulumi.RegisterOutputType(ProductSecretPrivilegesListOutput{})
	pulumi.RegisterOutputType(ProductSecretPrivilegesListArrayOutput{})
	pulumi.RegisterOutputType(GetSecretVersionsSecretVersionListOutput{})
	pulumi.RegisterOutputType(GetSecretVersionsSecretVersionListArrayOutput{})
	pulumi.RegisterOutputType(GetSecretsSecretListOutput{})
	pulumi.RegisterOutputType(GetSecretsSecretListArrayOutput{})
}
