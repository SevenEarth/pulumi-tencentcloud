// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a ssm productSecret
 *
 * ## Example Usage
 * ### Ssm secret for mysql
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const zones = tencentcloud.Availability.getZonesByProduct({
 *     product: "cdb",
 * });
 * const vpc = new tencentcloud.vpc.Instance("vpc", {cidrBlock: "10.0.0.0/16"});
 * const subnet = new tencentcloud.subnet.Instance("subnet", {
 *     availabilityZone: zones.then(zones => zones.zones?[0]?.name),
 *     vpcId: vpc.id,
 *     cidrBlock: "10.0.0.0/16",
 *     isMulticast: false,
 * });
 * const securityGroup = new tencentcloud.security.Group("securityGroup", {description: "desc."});
 * const exampleInstance = new tencentcloud.mysql.Instance("exampleInstance", {
 *     internetService: 1,
 *     engineVersion: "5.7",
 *     chargeType: "POSTPAID",
 *     rootPassword: "PassWord123",
 *     slaveDeployMode: 0,
 *     availabilityZone: zones.then(zones => zones.zones?[0]?.name),
 *     slaveSyncMode: 1,
 *     instanceName: "tf-example",
 *     memSize: 4000,
 *     volumeSize: 200,
 *     vpcId: vpc.id,
 *     subnetId: subnet.id,
 *     intranetPort: 3306,
 *     securityGroups: [securityGroup.id],
 *     tags: {
 *         createBy: "terraform",
 *     },
 *     parameters: {
 *         character_set_server: "utf8",
 *         max_connections: "1000",
 *     },
 * });
 * const exampleKey = new tencentcloud.kms.Key("exampleKey", {
 *     alias: "tf-example-kms-key",
 *     description: "example of kms key",
 *     keyRotationEnabled: false,
 *     isEnabled: true,
 *     tags: {
 *         createdBy: "terraform",
 *     },
 * });
 * const exampleProductSecret = new tencentcloud.ssm.ProductSecret("exampleProductSecret", {
 *     secretName: "tf-example",
 *     userNamePrefix: "prefix",
 *     productName: "Mysql",
 *     instanceId: exampleInstance.id,
 *     domains: ["10.0.0.0"],
 *     privilegesLists: [{
 *         privilegeName: "GlobalPrivileges",
 *         privileges: ["ALTER ROUTINE"],
 *     }],
 *     description: "for ssm product test",
 *     kmsKeyId: exampleKey.id,
 *     status: "Enabled",
 *     enableRotation: true,
 *     rotationBeginTime: "2023-08-05 20:54:33",
 *     rotationFrequency: 30,
 *     tags: {
 *         createdBy: "terraform",
 *     },
 * });
 * ```
 * ### Ssm secret for tdsql-c-mysql
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 *
 * const example = new tencentcloud.ssm.ProductSecret("example", {
 *     secretName: "tf-tdsql-c-example",
 *     userNamePrefix: "prefix",
 *     productName: "Tdsql_C_Mysql",
 *     instanceId: "cynosdbmysql-xxxxxx",
 *     domains: ["%"],
 *     privilegesLists: [
 *         {
 *             privilegeName: "GlobalPrivileges",
 *             privileges: [
 *                 "ALTER",
 *                 "CREATE",
 *                 "DELETE",
 *             ],
 *         },
 *         {
 *             privilegeName: "DatabasePrivileges",
 *             database: "test",
 *             privileges: [
 *                 "ALTER",
 *                 "CREATE",
 *                 "DELETE",
 *                 "SELECT",
 *             ],
 *         },
 *     ],
 *     description: "test tdsql-c",
 *     kmsKeyId: undefined,
 *     status: "Enabled",
 *     enableRotation: false,
 *     rotationBeginTime: "2023-08-05 20:54:33",
 *     rotationFrequency: 30,
 *     tags: {
 *         createdBy: "terraform",
 *     },
 * });
 * ```
 */
export class ProductSecret extends pulumi.CustomResource {
    /**
     * Get an existing ProductSecret resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProductSecretState, opts?: pulumi.CustomResourceOptions): ProductSecret {
        return new ProductSecret(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Ssm/productSecret:ProductSecret';

    /**
     * Returns true if the given object is an instance of ProductSecret.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProductSecret {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProductSecret.__pulumiType;
    }

    /**
     * Credential creation time in UNIX timestamp format.
     */
    public /*out*/ readonly createTime!: pulumi.Output<number>;
    /**
     * Description, which is used to describe the purpose in detail and can contain up to 2,048 bytes.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Domain name of the account in the form of IP. You can enter `%`.
     */
    public readonly domains!: pulumi.Output<string[]>;
    /**
     * Specifies whether to enable rotation, when secret status is `Disabled`, rotation will be disabled. `True` - enable, `False` - do not enable. If this parameter is not specified, `False` will be used by default.
     */
    public readonly enableRotation!: pulumi.Output<boolean>;
    /**
     * Tencent Cloud service instance ID.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Specifies the KMS CMK that encrypts the credential. If this parameter is left empty, the CMK created by Secrets Manager by default will be used for encryption.You can also specify a custom KMS CMK created in the same region for encryption.
     */
    public readonly kmsKeyId!: pulumi.Output<string | undefined>;
    /**
     * List of permissions that need to be granted when the credential is bound to a Tencent Cloud service.
     */
    public readonly privilegesLists!: pulumi.Output<outputs.Ssm.ProductSecretPrivilegesList[]>;
    /**
     * Name of the Tencent Cloud service bound to the credential, such as `Mysql`, `Tdsql-mysql`, `Tdsql_C_Mysql`. you can use dataSource `tencentcloud.Ssm.getProducts` to query supported products.
     */
    public readonly productName!: pulumi.Output<string>;
    /**
     * User-Defined rotation start time in the format of 2006-01-02 15:04:05.When `EnableRotation` is `True`, this parameter is required.
     */
    public readonly rotationBeginTime!: pulumi.Output<string>;
    /**
     * Rotation frequency in days. Default value: 1 day.
     */
    public readonly rotationFrequency!: pulumi.Output<number>;
    /**
     * Credential name, which must be unique in the same region. It can contain 128 bytes of letters, digits, hyphens, and underscores and must begin with a letter or digit.
     */
    public readonly secretName!: pulumi.Output<string>;
    /**
     * `0`: user-defined secret. `1`: Tencent Cloud services secret. `2`: SSH key secret. `3`: Tencent Cloud API key secret. Note: this field may return `null`, indicating that no valid values can be obtained.
     */
    public /*out*/ readonly secretType!: pulumi.Output<number>;
    /**
     * Enable or Disable Secret. Valid values is `Enabled` or `Disabled`. Default is `Enabled`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Tags of secret.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Prefix of the user account name, which is specified by you and can contain up to 8 characters.Supported character sets include:Digits: [0, 9].Lowercase letters: [a, z].Uppercase letters: [A, Z].Special symbols: underscore.The prefix must begin with a letter.
     */
    public readonly userNamePrefix!: pulumi.Output<string>;

    /**
     * Create a ProductSecret resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProductSecretArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProductSecretArgs | ProductSecretState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProductSecretState | undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["domains"] = state ? state.domains : undefined;
            resourceInputs["enableRotation"] = state ? state.enableRotation : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["kmsKeyId"] = state ? state.kmsKeyId : undefined;
            resourceInputs["privilegesLists"] = state ? state.privilegesLists : undefined;
            resourceInputs["productName"] = state ? state.productName : undefined;
            resourceInputs["rotationBeginTime"] = state ? state.rotationBeginTime : undefined;
            resourceInputs["rotationFrequency"] = state ? state.rotationFrequency : undefined;
            resourceInputs["secretName"] = state ? state.secretName : undefined;
            resourceInputs["secretType"] = state ? state.secretType : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["userNamePrefix"] = state ? state.userNamePrefix : undefined;
        } else {
            const args = argsOrState as ProductSecretArgs | undefined;
            if ((!args || args.domains === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domains'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.privilegesLists === undefined) && !opts.urn) {
                throw new Error("Missing required property 'privilegesLists'");
            }
            if ((!args || args.productName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'productName'");
            }
            if ((!args || args.secretName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'secretName'");
            }
            if ((!args || args.userNamePrefix === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userNamePrefix'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["domains"] = args ? args.domains : undefined;
            resourceInputs["enableRotation"] = args ? args.enableRotation : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            resourceInputs["privilegesLists"] = args ? args.privilegesLists : undefined;
            resourceInputs["productName"] = args ? args.productName : undefined;
            resourceInputs["rotationBeginTime"] = args ? args.rotationBeginTime : undefined;
            resourceInputs["rotationFrequency"] = args ? args.rotationFrequency : undefined;
            resourceInputs["secretName"] = args ? args.secretName : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["userNamePrefix"] = args ? args.userNamePrefix : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["secretType"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ProductSecret.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProductSecret resources.
 */
export interface ProductSecretState {
    /**
     * Credential creation time in UNIX timestamp format.
     */
    createTime?: pulumi.Input<number>;
    /**
     * Description, which is used to describe the purpose in detail and can contain up to 2,048 bytes.
     */
    description?: pulumi.Input<string>;
    /**
     * Domain name of the account in the form of IP. You can enter `%`.
     */
    domains?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies whether to enable rotation, when secret status is `Disabled`, rotation will be disabled. `True` - enable, `False` - do not enable. If this parameter is not specified, `False` will be used by default.
     */
    enableRotation?: pulumi.Input<boolean>;
    /**
     * Tencent Cloud service instance ID.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Specifies the KMS CMK that encrypts the credential. If this parameter is left empty, the CMK created by Secrets Manager by default will be used for encryption.You can also specify a custom KMS CMK created in the same region for encryption.
     */
    kmsKeyId?: pulumi.Input<string>;
    /**
     * List of permissions that need to be granted when the credential is bound to a Tencent Cloud service.
     */
    privilegesLists?: pulumi.Input<pulumi.Input<inputs.Ssm.ProductSecretPrivilegesList>[]>;
    /**
     * Name of the Tencent Cloud service bound to the credential, such as `Mysql`, `Tdsql-mysql`, `Tdsql_C_Mysql`. you can use dataSource `tencentcloud.Ssm.getProducts` to query supported products.
     */
    productName?: pulumi.Input<string>;
    /**
     * User-Defined rotation start time in the format of 2006-01-02 15:04:05.When `EnableRotation` is `True`, this parameter is required.
     */
    rotationBeginTime?: pulumi.Input<string>;
    /**
     * Rotation frequency in days. Default value: 1 day.
     */
    rotationFrequency?: pulumi.Input<number>;
    /**
     * Credential name, which must be unique in the same region. It can contain 128 bytes of letters, digits, hyphens, and underscores and must begin with a letter or digit.
     */
    secretName?: pulumi.Input<string>;
    /**
     * `0`: user-defined secret. `1`: Tencent Cloud services secret. `2`: SSH key secret. `3`: Tencent Cloud API key secret. Note: this field may return `null`, indicating that no valid values can be obtained.
     */
    secretType?: pulumi.Input<number>;
    /**
     * Enable or Disable Secret. Valid values is `Enabled` or `Disabled`. Default is `Enabled`.
     */
    status?: pulumi.Input<string>;
    /**
     * Tags of secret.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Prefix of the user account name, which is specified by you and can contain up to 8 characters.Supported character sets include:Digits: [0, 9].Lowercase letters: [a, z].Uppercase letters: [A, Z].Special symbols: underscore.The prefix must begin with a letter.
     */
    userNamePrefix?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProductSecret resource.
 */
export interface ProductSecretArgs {
    /**
     * Description, which is used to describe the purpose in detail and can contain up to 2,048 bytes.
     */
    description?: pulumi.Input<string>;
    /**
     * Domain name of the account in the form of IP. You can enter `%`.
     */
    domains: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies whether to enable rotation, when secret status is `Disabled`, rotation will be disabled. `True` - enable, `False` - do not enable. If this parameter is not specified, `False` will be used by default.
     */
    enableRotation?: pulumi.Input<boolean>;
    /**
     * Tencent Cloud service instance ID.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Specifies the KMS CMK that encrypts the credential. If this parameter is left empty, the CMK created by Secrets Manager by default will be used for encryption.You can also specify a custom KMS CMK created in the same region for encryption.
     */
    kmsKeyId?: pulumi.Input<string>;
    /**
     * List of permissions that need to be granted when the credential is bound to a Tencent Cloud service.
     */
    privilegesLists: pulumi.Input<pulumi.Input<inputs.Ssm.ProductSecretPrivilegesList>[]>;
    /**
     * Name of the Tencent Cloud service bound to the credential, such as `Mysql`, `Tdsql-mysql`, `Tdsql_C_Mysql`. you can use dataSource `tencentcloud.Ssm.getProducts` to query supported products.
     */
    productName: pulumi.Input<string>;
    /**
     * User-Defined rotation start time in the format of 2006-01-02 15:04:05.When `EnableRotation` is `True`, this parameter is required.
     */
    rotationBeginTime?: pulumi.Input<string>;
    /**
     * Rotation frequency in days. Default value: 1 day.
     */
    rotationFrequency?: pulumi.Input<number>;
    /**
     * Credential name, which must be unique in the same region. It can contain 128 bytes of letters, digits, hyphens, and underscores and must begin with a letter or digit.
     */
    secretName: pulumi.Input<string>;
    /**
     * Enable or Disable Secret. Valid values is `Enabled` or `Disabled`. Default is `Enabled`.
     */
    status?: pulumi.Input<string>;
    /**
     * Tags of secret.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Prefix of the user account name, which is specified by you and can contain up to 8 characters.Supported character sets include:Digits: [0, 9].Lowercase letters: [a, z].Uppercase letters: [A, Z].Special symbols: underscore.The prefix must begin with a letter.
     */
    userNamePrefix: pulumi.Input<string>;
}
