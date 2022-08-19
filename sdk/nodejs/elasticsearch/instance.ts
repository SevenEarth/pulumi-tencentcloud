// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides an elasticsearch instance resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const foo = new tencentcloud.elasticsearch.Instance("foo", {
 *     instanceName: "tf-test",
 *     availabilityZone: "ap-guangzhou-3",
 *     version: "7.5.1",
 *     vpcId: _var.vpc_id,
 *     subnetId: _var.subnet_id,
 *     password: "Test12345",
 *     licenseType: "oss",
 *     webNodeTypeInfos: [{
 *         nodeNum: 1,
 *         nodeType: "ES.S1.MEDIUM4",
 *     }],
 *     nodeInfoLists: [{
 *         nodeNum: 2,
 *         nodeType: "ES.S1.MEDIUM4",
 *         encrypt: false,
 *     }],
 *     tags: {
 *         test: "test",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Elasticsearch instance can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Elasticsearch/instance:Instance foo es-17634f05
 * ```
 */
export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceState, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Elasticsearch/instance:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }

    /**
     * Availability zone. When create multi-az es, this parameter must be omitted.
     */
    public readonly availabilityZone!: pulumi.Output<string | undefined>;
    /**
     * Whether to enable X-Pack security authentication in Basic Edition 6.8 and above. Valid values are `1` and `2`. `1` is disabled, `2` is enabled, and default value is `1`.
     */
    public readonly basicSecurityType!: pulumi.Output<number | undefined>;
    /**
     * The tenancy of the prepaid instance, and uint is month. NOTE: it only works when chargeType is set to `PREPAID`.
     */
    public readonly chargePeriod!: pulumi.Output<number | undefined>;
    /**
     * The charge type of instance. Valid values are `PREPAID` and `POSTPAID_BY_HOUR`.
     */
    public readonly chargeType!: pulumi.Output<string | undefined>;
    /**
     * Instance creation time.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Cluster deployment mode. Valid values are `0` and `1`. `0` is single-AZ deployment, and `1` is multi-AZ deployment. Default value is `0`.
     */
    public readonly deployMode!: pulumi.Output<number | undefined>;
    /**
     * Elasticsearch domain name.
     */
    public /*out*/ readonly elasticsearchDomain!: pulumi.Output<string>;
    /**
     * Elasticsearch port.
     */
    public /*out*/ readonly elasticsearchPort!: pulumi.Output<number>;
    /**
     * Elasticsearch VIP.
     */
    public /*out*/ readonly elasticsearchVip!: pulumi.Output<string>;
    /**
     * Name of the instance, which can contain 1 to 50 English letters, Chinese characters, digits, dashes(-), or underscores(_).
     */
    public readonly instanceName!: pulumi.Output<string | undefined>;
    /**
     * Kibana access URL.
     */
    public /*out*/ readonly kibanaUrl!: pulumi.Output<string>;
    /**
     * License type. Valid values are `oss`, `basic` and `platinum`. The default value is `platinum`.
     */
    public readonly licenseType!: pulumi.Output<string | undefined>;
    /**
     * Details of AZs in multi-AZ deployment mode (which is required when deployMode is `1`).
     */
    public readonly multiZoneInfos!: pulumi.Output<outputs.Elasticsearch.InstanceMultiZoneInfo[]>;
    /**
     * Node information list, which is used to describe the specification information of various types of nodes in the cluster, such as node type, node quantity, node specification, disk type, and disk size.
     */
    public readonly nodeInfoLists!: pulumi.Output<outputs.Elasticsearch.InstanceNodeInfoList[]>;
    /**
     * Password to an instance.
     */
    public readonly password!: pulumi.Output<string>;
    /**
     * When enabled, the instance will be renew automatically when it reach the end of the prepaid tenancy. Valid values are `RENEW_FLAG_AUTO` and `RENEW_FLAG_MANUAL`. NOTE: it only works when chargeType is set to `PREPAID`.
     */
    public readonly renewFlag!: pulumi.Output<string | undefined>;
    /**
     * The ID of a VPC subnetwork. When create multi-az es, this parameter must be omitted.
     */
    public readonly subnetId!: pulumi.Output<string | undefined>;
    /**
     * A mapping of tags to assign to the instance. For tag limits, please refer to [Use Limits](https://intl.cloud.tencent.com/document/product/651/13354).
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Version of the instance. Valid values are `5.6.4`, `6.4.3`, `6.8.2` and `7.5.1`.
     */
    public readonly version!: pulumi.Output<string>;
    /**
     * The ID of a VPC network.
     */
    public readonly vpcId!: pulumi.Output<string>;
    /**
     * Visual node configuration.
     */
    public readonly webNodeTypeInfos!: pulumi.Output<outputs.Elasticsearch.InstanceWebNodeTypeInfo[] | undefined>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceArgs | InstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceState | undefined;
            resourceInputs["availabilityZone"] = state ? state.availabilityZone : undefined;
            resourceInputs["basicSecurityType"] = state ? state.basicSecurityType : undefined;
            resourceInputs["chargePeriod"] = state ? state.chargePeriod : undefined;
            resourceInputs["chargeType"] = state ? state.chargeType : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["deployMode"] = state ? state.deployMode : undefined;
            resourceInputs["elasticsearchDomain"] = state ? state.elasticsearchDomain : undefined;
            resourceInputs["elasticsearchPort"] = state ? state.elasticsearchPort : undefined;
            resourceInputs["elasticsearchVip"] = state ? state.elasticsearchVip : undefined;
            resourceInputs["instanceName"] = state ? state.instanceName : undefined;
            resourceInputs["kibanaUrl"] = state ? state.kibanaUrl : undefined;
            resourceInputs["licenseType"] = state ? state.licenseType : undefined;
            resourceInputs["multiZoneInfos"] = state ? state.multiZoneInfos : undefined;
            resourceInputs["nodeInfoLists"] = state ? state.nodeInfoLists : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["renewFlag"] = state ? state.renewFlag : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["webNodeTypeInfos"] = state ? state.webNodeTypeInfos : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if ((!args || args.nodeInfoLists === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nodeInfoLists'");
            }
            if ((!args || args.password === undefined) && !opts.urn) {
                throw new Error("Missing required property 'password'");
            }
            if ((!args || args.version === undefined) && !opts.urn) {
                throw new Error("Missing required property 'version'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            resourceInputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            resourceInputs["basicSecurityType"] = args ? args.basicSecurityType : undefined;
            resourceInputs["chargePeriod"] = args ? args.chargePeriod : undefined;
            resourceInputs["chargeType"] = args ? args.chargeType : undefined;
            resourceInputs["deployMode"] = args ? args.deployMode : undefined;
            resourceInputs["instanceName"] = args ? args.instanceName : undefined;
            resourceInputs["licenseType"] = args ? args.licenseType : undefined;
            resourceInputs["multiZoneInfos"] = args ? args.multiZoneInfos : undefined;
            resourceInputs["nodeInfoLists"] = args ? args.nodeInfoLists : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["renewFlag"] = args ? args.renewFlag : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["webNodeTypeInfos"] = args ? args.webNodeTypeInfos : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["elasticsearchDomain"] = undefined /*out*/;
            resourceInputs["elasticsearchPort"] = undefined /*out*/;
            resourceInputs["elasticsearchVip"] = undefined /*out*/;
            resourceInputs["kibanaUrl"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Instance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    /**
     * Availability zone. When create multi-az es, this parameter must be omitted.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * Whether to enable X-Pack security authentication in Basic Edition 6.8 and above. Valid values are `1` and `2`. `1` is disabled, `2` is enabled, and default value is `1`.
     */
    basicSecurityType?: pulumi.Input<number>;
    /**
     * The tenancy of the prepaid instance, and uint is month. NOTE: it only works when chargeType is set to `PREPAID`.
     */
    chargePeriod?: pulumi.Input<number>;
    /**
     * The charge type of instance. Valid values are `PREPAID` and `POSTPAID_BY_HOUR`.
     */
    chargeType?: pulumi.Input<string>;
    /**
     * Instance creation time.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Cluster deployment mode. Valid values are `0` and `1`. `0` is single-AZ deployment, and `1` is multi-AZ deployment. Default value is `0`.
     */
    deployMode?: pulumi.Input<number>;
    /**
     * Elasticsearch domain name.
     */
    elasticsearchDomain?: pulumi.Input<string>;
    /**
     * Elasticsearch port.
     */
    elasticsearchPort?: pulumi.Input<number>;
    /**
     * Elasticsearch VIP.
     */
    elasticsearchVip?: pulumi.Input<string>;
    /**
     * Name of the instance, which can contain 1 to 50 English letters, Chinese characters, digits, dashes(-), or underscores(_).
     */
    instanceName?: pulumi.Input<string>;
    /**
     * Kibana access URL.
     */
    kibanaUrl?: pulumi.Input<string>;
    /**
     * License type. Valid values are `oss`, `basic` and `platinum`. The default value is `platinum`.
     */
    licenseType?: pulumi.Input<string>;
    /**
     * Details of AZs in multi-AZ deployment mode (which is required when deployMode is `1`).
     */
    multiZoneInfos?: pulumi.Input<pulumi.Input<inputs.Elasticsearch.InstanceMultiZoneInfo>[]>;
    /**
     * Node information list, which is used to describe the specification information of various types of nodes in the cluster, such as node type, node quantity, node specification, disk type, and disk size.
     */
    nodeInfoLists?: pulumi.Input<pulumi.Input<inputs.Elasticsearch.InstanceNodeInfoList>[]>;
    /**
     * Password to an instance.
     */
    password?: pulumi.Input<string>;
    /**
     * When enabled, the instance will be renew automatically when it reach the end of the prepaid tenancy. Valid values are `RENEW_FLAG_AUTO` and `RENEW_FLAG_MANUAL`. NOTE: it only works when chargeType is set to `PREPAID`.
     */
    renewFlag?: pulumi.Input<string>;
    /**
     * The ID of a VPC subnetwork. When create multi-az es, this parameter must be omitted.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the instance. For tag limits, please refer to [Use Limits](https://intl.cloud.tencent.com/document/product/651/13354).
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Version of the instance. Valid values are `5.6.4`, `6.4.3`, `6.8.2` and `7.5.1`.
     */
    version?: pulumi.Input<string>;
    /**
     * The ID of a VPC network.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * Visual node configuration.
     */
    webNodeTypeInfos?: pulumi.Input<pulumi.Input<inputs.Elasticsearch.InstanceWebNodeTypeInfo>[]>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * Availability zone. When create multi-az es, this parameter must be omitted.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * Whether to enable X-Pack security authentication in Basic Edition 6.8 and above. Valid values are `1` and `2`. `1` is disabled, `2` is enabled, and default value is `1`.
     */
    basicSecurityType?: pulumi.Input<number>;
    /**
     * The tenancy of the prepaid instance, and uint is month. NOTE: it only works when chargeType is set to `PREPAID`.
     */
    chargePeriod?: pulumi.Input<number>;
    /**
     * The charge type of instance. Valid values are `PREPAID` and `POSTPAID_BY_HOUR`.
     */
    chargeType?: pulumi.Input<string>;
    /**
     * Cluster deployment mode. Valid values are `0` and `1`. `0` is single-AZ deployment, and `1` is multi-AZ deployment. Default value is `0`.
     */
    deployMode?: pulumi.Input<number>;
    /**
     * Name of the instance, which can contain 1 to 50 English letters, Chinese characters, digits, dashes(-), or underscores(_).
     */
    instanceName?: pulumi.Input<string>;
    /**
     * License type. Valid values are `oss`, `basic` and `platinum`. The default value is `platinum`.
     */
    licenseType?: pulumi.Input<string>;
    /**
     * Details of AZs in multi-AZ deployment mode (which is required when deployMode is `1`).
     */
    multiZoneInfos?: pulumi.Input<pulumi.Input<inputs.Elasticsearch.InstanceMultiZoneInfo>[]>;
    /**
     * Node information list, which is used to describe the specification information of various types of nodes in the cluster, such as node type, node quantity, node specification, disk type, and disk size.
     */
    nodeInfoLists: pulumi.Input<pulumi.Input<inputs.Elasticsearch.InstanceNodeInfoList>[]>;
    /**
     * Password to an instance.
     */
    password: pulumi.Input<string>;
    /**
     * When enabled, the instance will be renew automatically when it reach the end of the prepaid tenancy. Valid values are `RENEW_FLAG_AUTO` and `RENEW_FLAG_MANUAL`. NOTE: it only works when chargeType is set to `PREPAID`.
     */
    renewFlag?: pulumi.Input<string>;
    /**
     * The ID of a VPC subnetwork. When create multi-az es, this parameter must be omitted.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the instance. For tag limits, please refer to [Use Limits](https://intl.cloud.tencent.com/document/product/651/13354).
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Version of the instance. Valid values are `5.6.4`, `6.4.3`, `6.8.2` and `7.5.1`.
     */
    version: pulumi.Input<string>;
    /**
     * The ID of a VPC network.
     */
    vpcId: pulumi.Input<string>;
    /**
     * Visual node configuration.
     */
    webNodeTypeInfos?: pulumi.Input<pulumi.Input<inputs.Elasticsearch.InstanceWebNodeTypeInfo>[]>;
}
