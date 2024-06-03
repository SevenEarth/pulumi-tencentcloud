// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provide a resource to create a SSM secret.
 *
 * ## Example Usage
 *
 * ### Create user defined secret
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const example = new tencentcloud.ssm.Secret("example", {
 *     description: "desc.",
 *     isEnabled: true,
 *     recoveryWindowInDays: 0,
 *     secretName: "tf-example",
 *     tags: {
 *         createBy: "terraform",
 *     },
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ### Create redis secret
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const zone = tencentcloud.Redis.getZoneConfig({
 *     typeId: 8,
 * });
 * const vpc = new tencentcloud.vpc.Instance("vpc", {cidrBlock: "10.0.0.0/16"});
 * const subnet = new tencentcloud.subnet.Instance("subnet", {
 *     vpcId: vpc.id,
 *     availabilityZone: zone.then(zone => zone.lists?.[3]?.zone),
 *     cidrBlock: "10.0.0.0/16",
 * });
 * const exampleInstance = new tencentcloud.redis.Instance("exampleInstance", {
 *     availabilityZone: zone.then(zone => zone.lists?.[3]?.zone),
 *     typeId: zone.then(zone => zone.lists?.[3]?.typeId),
 *     password: "Qwer@234",
 *     memSize: zone.then(zone => zone.lists?.[3]?.memSizes?.[0]),
 *     redisShardNum: zone.then(zone => zone.lists?.[3]?.redisShardNums?.[0]),
 *     redisReplicasNum: zone.then(zone => zone.lists?.[3]?.redisReplicasNums?.[0]),
 *     port: 6379,
 *     vpcId: vpc.id,
 *     subnetId: subnet.id,
 * });
 * const exampleSecret = new tencentcloud.ssm.Secret("exampleSecret", {
 *     secretName: "tf-example",
 *     description: "redis desc.",
 *     isEnabled: true,
 *     secretType: 4,
 *     additionalConfig: pulumi.jsonStringify({
 *         Region: "ap-guangzhou",
 *         Privilege: "r",
 *         InstanceId: exampleInstance.id,
 *         ReadonlyPolicy: ["master"],
 *         Remark: "for tf test",
 *     }),
 *     tags: {
 *         createdBy: "terraform",
 *     },
 *     recoveryWindowInDays: 0,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * SSM secret can be imported using the secretName, e.g.
 *
 * ```sh
 * $ pulumi import tencentcloud:Ssm/secret:Secret foo test
 * ```
 */
export class Secret extends pulumi.CustomResource {
    /**
     * Get an existing Secret resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretState, opts?: pulumi.CustomResourceOptions): Secret {
        return new Secret(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Ssm/secret:Secret';

    /**
     * Returns true if the given object is an instance of Secret.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Secret {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Secret.__pulumiType;
    }

    /**
     * Additional config for specific secret types in JSON string format.
     */
    public readonly additionalConfig!: pulumi.Output<string | undefined>;
    /**
     * Description of secret. The maximum is 2048 bytes.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specify whether to enable secret. Default value is `true`.
     */
    public readonly isEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * KMS keyId used to encrypt secret. If it is empty, it means that the CMK created by SSM for you by default is used for encryption. You can also specify the KMS CMK created by yourself in the same region for encryption.
     */
    public readonly kmsKeyId!: pulumi.Output<string>;
    /**
     * Specify the scheduled deletion date. Default value is `0` that means to delete immediately. 1-30 means the number of days reserved, completely deleted after this date.
     */
    public readonly recoveryWindowInDays!: pulumi.Output<number | undefined>;
    /**
     * Name of secret which cannot be repeated in the same region. The maximum length is 128 bytes. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
     */
    public readonly secretName!: pulumi.Output<string>;
    /**
     * Type of secret. `0`: user-defined secret. `4`: redis secret. Default is `0`.
     */
    public readonly secretType!: pulumi.Output<number>;
    /**
     * Status of secret.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Tags of secret.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a Secret resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretArgs | SecretState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecretState | undefined;
            resourceInputs["additionalConfig"] = state ? state.additionalConfig : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["isEnabled"] = state ? state.isEnabled : undefined;
            resourceInputs["kmsKeyId"] = state ? state.kmsKeyId : undefined;
            resourceInputs["recoveryWindowInDays"] = state ? state.recoveryWindowInDays : undefined;
            resourceInputs["secretName"] = state ? state.secretName : undefined;
            resourceInputs["secretType"] = state ? state.secretType : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as SecretArgs | undefined;
            if ((!args || args.secretName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'secretName'");
            }
            resourceInputs["additionalConfig"] = args ? args.additionalConfig : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["isEnabled"] = args ? args.isEnabled : undefined;
            resourceInputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            resourceInputs["recoveryWindowInDays"] = args ? args.recoveryWindowInDays : undefined;
            resourceInputs["secretName"] = args ? args.secretName : undefined;
            resourceInputs["secretType"] = args ? args.secretType : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Secret.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Secret resources.
 */
export interface SecretState {
    /**
     * Additional config for specific secret types in JSON string format.
     */
    additionalConfig?: pulumi.Input<string>;
    /**
     * Description of secret. The maximum is 2048 bytes.
     */
    description?: pulumi.Input<string>;
    /**
     * Specify whether to enable secret. Default value is `true`.
     */
    isEnabled?: pulumi.Input<boolean>;
    /**
     * KMS keyId used to encrypt secret. If it is empty, it means that the CMK created by SSM for you by default is used for encryption. You can also specify the KMS CMK created by yourself in the same region for encryption.
     */
    kmsKeyId?: pulumi.Input<string>;
    /**
     * Specify the scheduled deletion date. Default value is `0` that means to delete immediately. 1-30 means the number of days reserved, completely deleted after this date.
     */
    recoveryWindowInDays?: pulumi.Input<number>;
    /**
     * Name of secret which cannot be repeated in the same region. The maximum length is 128 bytes. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
     */
    secretName?: pulumi.Input<string>;
    /**
     * Type of secret. `0`: user-defined secret. `4`: redis secret. Default is `0`.
     */
    secretType?: pulumi.Input<number>;
    /**
     * Status of secret.
     */
    status?: pulumi.Input<string>;
    /**
     * Tags of secret.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a Secret resource.
 */
export interface SecretArgs {
    /**
     * Additional config for specific secret types in JSON string format.
     */
    additionalConfig?: pulumi.Input<string>;
    /**
     * Description of secret. The maximum is 2048 bytes.
     */
    description?: pulumi.Input<string>;
    /**
     * Specify whether to enable secret. Default value is `true`.
     */
    isEnabled?: pulumi.Input<boolean>;
    /**
     * KMS keyId used to encrypt secret. If it is empty, it means that the CMK created by SSM for you by default is used for encryption. You can also specify the KMS CMK created by yourself in the same region for encryption.
     */
    kmsKeyId?: pulumi.Input<string>;
    /**
     * Specify the scheduled deletion date. Default value is `0` that means to delete immediately. 1-30 means the number of days reserved, completely deleted after this date.
     */
    recoveryWindowInDays?: pulumi.Input<number>;
    /**
     * Name of secret which cannot be repeated in the same region. The maximum length is 128 bytes. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
     */
    secretName: pulumi.Input<string>;
    /**
     * Type of secret. `0`: user-defined secret. `4`: redis secret. Default is `0`.
     */
    secretType?: pulumi.Input<number>;
    /**
     * Tags of secret.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}
