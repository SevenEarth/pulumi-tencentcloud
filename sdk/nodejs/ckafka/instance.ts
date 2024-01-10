// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this resource to create ckafka instance.
 *
 * ## Example Usage
 * ### Basic Instance
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const config = new pulumi.Config();
 * const vpcId = config.get("vpcId") || "vpc-68vi2d3h";
 * const subnetId = config.get("subnetId") || "subnet-ob6clqwk";
 * const gz = tencentcloud.Availability.getZonesByProduct({
 *     name: "ap-guangzhou-3",
 *     product: "ckafka",
 * });
 * const kafkaInstancePrepaid = new tencentcloud.ckafka.Instance("kafkaInstancePrepaid", {
 *     instanceName: "ckafka-instance-prepaid",
 *     zoneId: gz.then(gz => gz.zones?[0]?.id),
 *     period: 1,
 *     vpcId: vpcId,
 *     subnetId: subnetId,
 *     msgRetentionTime: 1300,
 *     renewFlag: 0,
 *     kafkaVersion: "2.4.1",
 *     diskSize: 200,
 *     diskType: "CLOUD_BASIC",
 *     bandWidth: 20,
 *     partition: 400,
 *     specificationsType: "standard",
 *     instanceType: 2,
 *     config: {
 *         autoCreateTopicEnable: true,
 *         defaultNumPartitions: 3,
 *         defaultReplicationFactor: 3,
 *     },
 *     dynamicRetentionConfig: {
 *         enable: 1,
 *     },
 * });
 * const kafkaInstancePostpaid = new tencentcloud.ckafka.Instance("kafkaInstancePostpaid", {
 *     instanceName: "ckafka-instance-postpaid",
 *     zoneId: gz.then(gz => gz.zones?[0]?.id),
 *     vpcId: vpcId,
 *     subnetId: subnetId,
 *     msgRetentionTime: 1300,
 *     kafkaVersion: "1.1.1",
 *     diskSize: 200,
 *     bandWidth: 20,
 *     diskType: "CLOUD_BASIC",
 *     partition: 400,
 *     chargeType: "POSTPAID_BY_HOUR",
 *     config: {
 *         autoCreateTopicEnable: true,
 *         defaultNumPartitions: 3,
 *         defaultReplicationFactor: 3,
 *     },
 *     dynamicRetentionConfig: {
 *         enable: 1,
 *     },
 * });
 * ```
 * ### Multi zone Instance
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const config = new pulumi.Config();
 * const vpcId = config.get("vpcId") || "vpc-68vi2d3h";
 * const subnetId = config.get("subnetId") || "subnet-ob6clqwk";
 * const gz3 = tencentcloud.Availability.getZonesByProduct({
 *     name: "ap-guangzhou-3",
 *     product: "ckafka",
 * });
 * const gz6 = tencentcloud.Availability.getZonesByProduct({
 *     name: "ap-guangzhou-6",
 *     product: "ckafka",
 * });
 * const kafkaInstance = new tencentcloud.ckafka.Instance("kafkaInstance", {
 *     instanceName: "ckafka-instance-maz-tf-test",
 *     zoneId: gz3.then(gz3 => gz3.zones?[0]?.id),
 *     multiZoneFlag: true,
 *     zoneIds: [
 *         gz3.then(gz3 => gz3.zones?[0]?.id),
 *         gz6.then(gz6 => gz6.zones?[0]?.id),
 *     ],
 *     period: 1,
 *     vpcId: vpcId,
 *     subnetId: subnetId,
 *     msgRetentionTime: 1300,
 *     renewFlag: 0,
 *     kafkaVersion: "1.1.1",
 *     diskSize: 500,
 *     diskType: "CLOUD_BASIC",
 *     config: {
 *         autoCreateTopicEnable: true,
 *         defaultNumPartitions: 3,
 *         defaultReplicationFactor: 3,
 *     },
 *     dynamicRetentionConfig: {
 *         enable: 1,
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * ckafka instance can be imported using the instance_id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Ckafka/instance:Instance foo ckafka-f9ife4zz
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
    public static readonly __pulumiType = 'tencentcloud:Ckafka/instance:Instance';

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
     * Instance bandwidth in MBps.
     */
    public readonly bandWidth!: pulumi.Output<number>;
    /**
     * The charge type of instance. Valid values are `PREPAID` and `POSTPAID_BY_HOUR`. Default value is `PREPAID`.
     */
    public readonly chargeType!: pulumi.Output<string | undefined>;
    /**
     * Instance configuration.
     */
    public readonly config!: pulumi.Output<outputs.Ckafka.InstanceConfig | undefined>;
    /**
     * Disk Size. Its interval varies with bandwidth, and the input must be within the interval, which can be viewed through the control. If it is not within the interval, the plan will cause a change when first created.
     */
    public readonly diskSize!: pulumi.Output<number>;
    /**
     * Type of disk.
     */
    public readonly diskType!: pulumi.Output<string>;
    /**
     * Dynamic message retention policy configuration.
     */
    public readonly dynamicRetentionConfig!: pulumi.Output<outputs.Ckafka.InstanceDynamicRetentionConfig>;
    /**
     * Instance name.
     */
    public readonly instanceName!: pulumi.Output<string>;
    /**
     * Description of instance type. `profession`: 1, `standard`:  1(general), 2(standard), 3(advanced), 4(capacity), 5(specialized-1), 6(specialized-2), 7(specialized-3), 8(specialized-4), 9(exclusive).
     */
    public readonly instanceType!: pulumi.Output<number>;
    /**
     * Kafka version (0.10.2/1.1.1/2.4.1).
     */
    public readonly kafkaVersion!: pulumi.Output<string>;
    /**
     * The size of a single message in bytes at the instance level. Value range: `1024 - 12*1024*1024 bytes (i.e., 1KB-12MB).
     */
    public readonly maxMessageByte!: pulumi.Output<number>;
    /**
     * The maximum retention time of instance logs, in minutes. the default is 10080 (7 days), the maximum is 30 days, and the default 0 is not filled, which means that the log retention time recovery policy is not enabled.
     */
    public readonly msgRetentionTime!: pulumi.Output<number>;
    /**
     * Indicates whether the instance is multi zones. NOTE: if set to `true`, `zoneIds` must set together.
     */
    public readonly multiZoneFlag!: pulumi.Output<boolean | undefined>;
    /**
     * Partition Size. Its interval varies with bandwidth, and the input must be within the interval, which can be viewed through the control. If it is not within the interval, the plan will cause a change when first created.
     */
    public readonly partition!: pulumi.Output<number>;
    /**
     * Prepaid purchase time, such as 1, is one month.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * It has been deprecated from version 1.81.6. If set public network value, it will cause error. Bandwidth of the public network.
     *
     * @deprecated It has been deprecated from version 1.81.6. If set public network value, it will cause error.
     */
    public readonly publicNetwork!: pulumi.Output<number>;
    /**
     * Modification of the rebalancing time after upgrade.
     */
    public readonly rebalanceTime!: pulumi.Output<number | undefined>;
    /**
     * Prepaid automatic renewal mark, 0 means the default state, the initial state, 1 means automatic renewal, 2 means clear no automatic renewal (user setting).
     */
    public readonly renewFlag!: pulumi.Output<number>;
    /**
     * Specifications type of instance. Allowed values are `standard`, `profession`. Default is `profession`.
     */
    public readonly specificationsType!: pulumi.Output<string | undefined>;
    /**
     * Subnet id, it will be basic network if not set.
     */
    public readonly subnetId!: pulumi.Output<string | undefined>;
    /**
     * Tag set of instance.
     */
    public readonly tagSet!: pulumi.Output<{[key: string]: any}>;
    /**
     * It has been deprecated from version 1.78.5, because it do not support change. Use `tagSet` instead. Tags of instance. Partition size, the professional version does not need tag.
     *
     * @deprecated It has been deprecated from version 1.78.5, because it do not support change. Use `tag_set` instead.
     */
    public readonly tags!: pulumi.Output<outputs.Ckafka.InstanceTag[]>;
    /**
     * POSTPAID_BY_HOUR scale-down mode
     * - 1: stable transformation;
     * - 2: High-speed transformer.
     */
    public readonly upgradeStrategy!: pulumi.Output<number | undefined>;
    /**
     * Vip of instance.
     */
    public /*out*/ readonly vip!: pulumi.Output<string>;
    /**
     * Vpc id, it will be basic network if not set.
     */
    public readonly vpcId!: pulumi.Output<string | undefined>;
    /**
     * Type of instance.
     */
    public /*out*/ readonly vport!: pulumi.Output<string>;
    /**
     * Available zone id.
     */
    public readonly zoneId!: pulumi.Output<number>;
    /**
     * List of available zone id. NOTE: this argument must set together with `multiZoneFlag`.
     */
    public readonly zoneIds!: pulumi.Output<number[] | undefined>;

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
            resourceInputs["bandWidth"] = state ? state.bandWidth : undefined;
            resourceInputs["chargeType"] = state ? state.chargeType : undefined;
            resourceInputs["config"] = state ? state.config : undefined;
            resourceInputs["diskSize"] = state ? state.diskSize : undefined;
            resourceInputs["diskType"] = state ? state.diskType : undefined;
            resourceInputs["dynamicRetentionConfig"] = state ? state.dynamicRetentionConfig : undefined;
            resourceInputs["instanceName"] = state ? state.instanceName : undefined;
            resourceInputs["instanceType"] = state ? state.instanceType : undefined;
            resourceInputs["kafkaVersion"] = state ? state.kafkaVersion : undefined;
            resourceInputs["maxMessageByte"] = state ? state.maxMessageByte : undefined;
            resourceInputs["msgRetentionTime"] = state ? state.msgRetentionTime : undefined;
            resourceInputs["multiZoneFlag"] = state ? state.multiZoneFlag : undefined;
            resourceInputs["partition"] = state ? state.partition : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["publicNetwork"] = state ? state.publicNetwork : undefined;
            resourceInputs["rebalanceTime"] = state ? state.rebalanceTime : undefined;
            resourceInputs["renewFlag"] = state ? state.renewFlag : undefined;
            resourceInputs["specificationsType"] = state ? state.specificationsType : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["tagSet"] = state ? state.tagSet : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["upgradeStrategy"] = state ? state.upgradeStrategy : undefined;
            resourceInputs["vip"] = state ? state.vip : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["vport"] = state ? state.vport : undefined;
            resourceInputs["zoneId"] = state ? state.zoneId : undefined;
            resourceInputs["zoneIds"] = state ? state.zoneIds : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if ((!args || args.instanceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceName'");
            }
            if ((!args || args.zoneId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zoneId'");
            }
            resourceInputs["bandWidth"] = args ? args.bandWidth : undefined;
            resourceInputs["chargeType"] = args ? args.chargeType : undefined;
            resourceInputs["config"] = args ? args.config : undefined;
            resourceInputs["diskSize"] = args ? args.diskSize : undefined;
            resourceInputs["diskType"] = args ? args.diskType : undefined;
            resourceInputs["dynamicRetentionConfig"] = args ? args.dynamicRetentionConfig : undefined;
            resourceInputs["instanceName"] = args ? args.instanceName : undefined;
            resourceInputs["instanceType"] = args ? args.instanceType : undefined;
            resourceInputs["kafkaVersion"] = args ? args.kafkaVersion : undefined;
            resourceInputs["maxMessageByte"] = args ? args.maxMessageByte : undefined;
            resourceInputs["msgRetentionTime"] = args ? args.msgRetentionTime : undefined;
            resourceInputs["multiZoneFlag"] = args ? args.multiZoneFlag : undefined;
            resourceInputs["partition"] = args ? args.partition : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["publicNetwork"] = args ? args.publicNetwork : undefined;
            resourceInputs["rebalanceTime"] = args ? args.rebalanceTime : undefined;
            resourceInputs["renewFlag"] = args ? args.renewFlag : undefined;
            resourceInputs["specificationsType"] = args ? args.specificationsType : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["tagSet"] = args ? args.tagSet : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["upgradeStrategy"] = args ? args.upgradeStrategy : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["zoneId"] = args ? args.zoneId : undefined;
            resourceInputs["zoneIds"] = args ? args.zoneIds : undefined;
            resourceInputs["vip"] = undefined /*out*/;
            resourceInputs["vport"] = undefined /*out*/;
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
     * Instance bandwidth in MBps.
     */
    bandWidth?: pulumi.Input<number>;
    /**
     * The charge type of instance. Valid values are `PREPAID` and `POSTPAID_BY_HOUR`. Default value is `PREPAID`.
     */
    chargeType?: pulumi.Input<string>;
    /**
     * Instance configuration.
     */
    config?: pulumi.Input<inputs.Ckafka.InstanceConfig>;
    /**
     * Disk Size. Its interval varies with bandwidth, and the input must be within the interval, which can be viewed through the control. If it is not within the interval, the plan will cause a change when first created.
     */
    diskSize?: pulumi.Input<number>;
    /**
     * Type of disk.
     */
    diskType?: pulumi.Input<string>;
    /**
     * Dynamic message retention policy configuration.
     */
    dynamicRetentionConfig?: pulumi.Input<inputs.Ckafka.InstanceDynamicRetentionConfig>;
    /**
     * Instance name.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * Description of instance type. `profession`: 1, `standard`:  1(general), 2(standard), 3(advanced), 4(capacity), 5(specialized-1), 6(specialized-2), 7(specialized-3), 8(specialized-4), 9(exclusive).
     */
    instanceType?: pulumi.Input<number>;
    /**
     * Kafka version (0.10.2/1.1.1/2.4.1).
     */
    kafkaVersion?: pulumi.Input<string>;
    /**
     * The size of a single message in bytes at the instance level. Value range: `1024 - 12*1024*1024 bytes (i.e., 1KB-12MB).
     */
    maxMessageByte?: pulumi.Input<number>;
    /**
     * The maximum retention time of instance logs, in minutes. the default is 10080 (7 days), the maximum is 30 days, and the default 0 is not filled, which means that the log retention time recovery policy is not enabled.
     */
    msgRetentionTime?: pulumi.Input<number>;
    /**
     * Indicates whether the instance is multi zones. NOTE: if set to `true`, `zoneIds` must set together.
     */
    multiZoneFlag?: pulumi.Input<boolean>;
    /**
     * Partition Size. Its interval varies with bandwidth, and the input must be within the interval, which can be viewed through the control. If it is not within the interval, the plan will cause a change when first created.
     */
    partition?: pulumi.Input<number>;
    /**
     * Prepaid purchase time, such as 1, is one month.
     */
    period?: pulumi.Input<number>;
    /**
     * It has been deprecated from version 1.81.6. If set public network value, it will cause error. Bandwidth of the public network.
     *
     * @deprecated It has been deprecated from version 1.81.6. If set public network value, it will cause error.
     */
    publicNetwork?: pulumi.Input<number>;
    /**
     * Modification of the rebalancing time after upgrade.
     */
    rebalanceTime?: pulumi.Input<number>;
    /**
     * Prepaid automatic renewal mark, 0 means the default state, the initial state, 1 means automatic renewal, 2 means clear no automatic renewal (user setting).
     */
    renewFlag?: pulumi.Input<number>;
    /**
     * Specifications type of instance. Allowed values are `standard`, `profession`. Default is `profession`.
     */
    specificationsType?: pulumi.Input<string>;
    /**
     * Subnet id, it will be basic network if not set.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * Tag set of instance.
     */
    tagSet?: pulumi.Input<{[key: string]: any}>;
    /**
     * It has been deprecated from version 1.78.5, because it do not support change. Use `tagSet` instead. Tags of instance. Partition size, the professional version does not need tag.
     *
     * @deprecated It has been deprecated from version 1.78.5, because it do not support change. Use `tag_set` instead.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.Ckafka.InstanceTag>[]>;
    /**
     * POSTPAID_BY_HOUR scale-down mode
     * - 1: stable transformation;
     * - 2: High-speed transformer.
     */
    upgradeStrategy?: pulumi.Input<number>;
    /**
     * Vip of instance.
     */
    vip?: pulumi.Input<string>;
    /**
     * Vpc id, it will be basic network if not set.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * Type of instance.
     */
    vport?: pulumi.Input<string>;
    /**
     * Available zone id.
     */
    zoneId?: pulumi.Input<number>;
    /**
     * List of available zone id. NOTE: this argument must set together with `multiZoneFlag`.
     */
    zoneIds?: pulumi.Input<pulumi.Input<number>[]>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * Instance bandwidth in MBps.
     */
    bandWidth?: pulumi.Input<number>;
    /**
     * The charge type of instance. Valid values are `PREPAID` and `POSTPAID_BY_HOUR`. Default value is `PREPAID`.
     */
    chargeType?: pulumi.Input<string>;
    /**
     * Instance configuration.
     */
    config?: pulumi.Input<inputs.Ckafka.InstanceConfig>;
    /**
     * Disk Size. Its interval varies with bandwidth, and the input must be within the interval, which can be viewed through the control. If it is not within the interval, the plan will cause a change when first created.
     */
    diskSize?: pulumi.Input<number>;
    /**
     * Type of disk.
     */
    diskType?: pulumi.Input<string>;
    /**
     * Dynamic message retention policy configuration.
     */
    dynamicRetentionConfig?: pulumi.Input<inputs.Ckafka.InstanceDynamicRetentionConfig>;
    /**
     * Instance name.
     */
    instanceName: pulumi.Input<string>;
    /**
     * Description of instance type. `profession`: 1, `standard`:  1(general), 2(standard), 3(advanced), 4(capacity), 5(specialized-1), 6(specialized-2), 7(specialized-3), 8(specialized-4), 9(exclusive).
     */
    instanceType?: pulumi.Input<number>;
    /**
     * Kafka version (0.10.2/1.1.1/2.4.1).
     */
    kafkaVersion?: pulumi.Input<string>;
    /**
     * The size of a single message in bytes at the instance level. Value range: `1024 - 12*1024*1024 bytes (i.e., 1KB-12MB).
     */
    maxMessageByte?: pulumi.Input<number>;
    /**
     * The maximum retention time of instance logs, in minutes. the default is 10080 (7 days), the maximum is 30 days, and the default 0 is not filled, which means that the log retention time recovery policy is not enabled.
     */
    msgRetentionTime?: pulumi.Input<number>;
    /**
     * Indicates whether the instance is multi zones. NOTE: if set to `true`, `zoneIds` must set together.
     */
    multiZoneFlag?: pulumi.Input<boolean>;
    /**
     * Partition Size. Its interval varies with bandwidth, and the input must be within the interval, which can be viewed through the control. If it is not within the interval, the plan will cause a change when first created.
     */
    partition?: pulumi.Input<number>;
    /**
     * Prepaid purchase time, such as 1, is one month.
     */
    period?: pulumi.Input<number>;
    /**
     * It has been deprecated from version 1.81.6. If set public network value, it will cause error. Bandwidth of the public network.
     *
     * @deprecated It has been deprecated from version 1.81.6. If set public network value, it will cause error.
     */
    publicNetwork?: pulumi.Input<number>;
    /**
     * Modification of the rebalancing time after upgrade.
     */
    rebalanceTime?: pulumi.Input<number>;
    /**
     * Prepaid automatic renewal mark, 0 means the default state, the initial state, 1 means automatic renewal, 2 means clear no automatic renewal (user setting).
     */
    renewFlag?: pulumi.Input<number>;
    /**
     * Specifications type of instance. Allowed values are `standard`, `profession`. Default is `profession`.
     */
    specificationsType?: pulumi.Input<string>;
    /**
     * Subnet id, it will be basic network if not set.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * Tag set of instance.
     */
    tagSet?: pulumi.Input<{[key: string]: any}>;
    /**
     * It has been deprecated from version 1.78.5, because it do not support change. Use `tagSet` instead. Tags of instance. Partition size, the professional version does not need tag.
     *
     * @deprecated It has been deprecated from version 1.78.5, because it do not support change. Use `tag_set` instead.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.Ckafka.InstanceTag>[]>;
    /**
     * POSTPAID_BY_HOUR scale-down mode
     * - 1: stable transformation;
     * - 2: High-speed transformer.
     */
    upgradeStrategy?: pulumi.Input<number>;
    /**
     * Vpc id, it will be basic network if not set.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * Available zone id.
     */
    zoneId: pulumi.Input<number>;
    /**
     * List of available zone id. NOTE: this argument must set together with `multiZoneFlag`.
     */
    zoneIds?: pulumi.Input<pulumi.Input<number>[]>;
}
