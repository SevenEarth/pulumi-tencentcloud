// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a CLB instance.
 *
 * ## Example Usage
 *
 * INTERNAL CLB
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const internalClb = new tencentcloud.Clb.Instance("internal_clb", {
 *     clbName: "myclb",
 *     networkType: "INTERNAL",
 *     projectId: 0,
 *     subnetId: "subnet-12rastkr",
 *     tags: {
 *         test: "tf",
 *     },
 *     vpcId: "vpc-7007ll7q",
 * });
 * ```
 *
 * OPEN CLB
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const openClb = new tencentcloud.Clb.Instance("open_clb", {
 *     clbName: "myclb",
 *     networkType: "OPEN",
 *     projectId: 0,
 *     securityGroups: ["sg-o0ek7r93"],
 *     tags: {
 *         test: "tf",
 *     },
 *     targetRegionInfoRegion: "ap-guangzhou",
 *     targetRegionInfoVpcId: "vpc-da7ffa61",
 *     vpcId: "vpc-da7ffa61",
 * });
 * ```
 *
 * Default enable
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const foo = new tencentcloud.vpc.Instance("foo", {
 *     cidrBlock: "10.0.0.0/16",
 *     tags: {
 *         test: "mytest",
 *     },
 * });
 * const subnet = new tencentcloud.subnet.Instance("subnet", {
 *     availabilityZone: "ap-guangzhou-1",
 *     vpcId: foo.id,
 *     cidrBlock: "10.0.20.0/28",
 *     isMulticast: false,
 * });
 * const sglab = new tencentcloud.security.Group("sglab", {
 *     description: "favourite sg",
 *     projectId: 0,
 * });
 * const openClb = new tencentcloud.clb.Instance("openClb", {
 *     networkType: "OPEN",
 *     clbName: "my-open-clb",
 *     projectId: 0,
 *     vpcId: foo.id,
 *     loadBalancerPassToTarget: true,
 *     securityGroups: [sglab.id],
 *     targetRegionInfoRegion: "ap-guangzhou",
 *     targetRegionInfoVpcId: foo.id,
 *     tags: {
 *         test: "open",
 *     },
 * });
 * ```
 *
 * CREATE multiple instance
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const openClb1 = new tencentcloud.Clb.Instance("open_clb1", {
 *     clbName: "hello",
 *     masterZoneId: "ap-guangzhou-3",
 *     networkType: "OPEN",
 * });
 * ```
 *
 * CREATE instance with log
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const vpcTest = new tencentcloud.Vpc.Instance("vpc_test", {
 *     cidrBlock: "10.0.0.0/16",
 * });
 * const rtbTest = new tencentcloud.Route.Table("rtb_test", {
 *     vpcId: vpcTest.id,
 * });
 * const subnetTest = new tencentcloud.Subnet.Instance("subnet_test", {
 *     availabilityZone: "ap-guangzhou-3",
 *     cidrBlock: "10.0.1.0/24",
 *     routeTableId: rtbTest.id,
 *     vpcId: vpcTest.id,
 * });
 * const set = new tencentcloud.Clb.LogSet("set", {
 *     period: 7,
 * });
 * const topic = new tencentcloud.Clb.LogTopic("topic", {
 *     logSetId: set.id,
 *     topicName: "clb-topic",
 * });
 * const internalClb = new tencentcloud.Clb.Instance("internal_clb", {
 *     clbName: "myclb",
 *     loadBalancerPassToTarget: true,
 *     logSetId: set.id,
 *     logTopicId: topic.id,
 *     networkType: "INTERNAL",
 *     projectId: 0,
 *     subnetId: subnetTest.id,
 *     tags: {
 *         test: "tf",
 *     },
 *     vpcId: vpcTest.id,
 * });
 * ```
 *
 * ## Import
 *
 * CLB instance can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Clb/instance:Instance foo lb-7a0t6zqb
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
    public static readonly __pulumiType = 'tencentcloud:Clb/instance:Instance';

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
     * IP version, only applicable to open CLB. Valid values are `ipv4`, `ipv6` and `IPv6FullChain`.
     */
    public readonly addressIpVersion!: pulumi.Output<string>;
    /**
     * Bandwidth package id. If set, the `internetChargeType` must be `BANDWIDTH_PACKAGE`.
     */
    public readonly bandwidthPackageId!: pulumi.Output<string | undefined>;
    /**
     * Name of the CLB. The name can only contain Chinese characters, English letters, numbers, underscore and hyphen '-'.
     */
    public readonly clbName!: pulumi.Output<string>;
    /**
     * The virtual service address table of the CLB.
     */
    public /*out*/ readonly clbVips!: pulumi.Output<string[]>;
    /**
     * Max bandwidth out, only applicable to open CLB. Valid value ranges is [1, 2048]. Unit is MB.
     */
    public readonly internetBandwidthMaxOut!: pulumi.Output<number>;
    /**
     * Internet charge type, only applicable to open CLB. Valid values are `TRAFFIC_POSTPAID_BY_HOUR`, `BANDWIDTH_POSTPAID_BY_HOUR` and `BANDWIDTH_PACKAGE`.
     */
    public readonly internetChargeType!: pulumi.Output<string>;
    /**
     * Whether the target allow flow come from clb. If value is true, only check security group of clb, or check both clb and backend instance security group.
     */
    public readonly loadBalancerPassToTarget!: pulumi.Output<boolean | undefined>;
    /**
     * The id of log set.
     */
    public readonly logSetId!: pulumi.Output<string | undefined>;
    /**
     * The id of log topic.
     */
    public readonly logTopicId!: pulumi.Output<string | undefined>;
    /**
     * Setting master zone id of cross available zone disaster recovery, only applicable to open CLB.
     */
    public readonly masterZoneId!: pulumi.Output<string | undefined>;
    /**
     * Type of CLB instance. Valid values: `OPEN` and `INTERNAL`.
     */
    public readonly networkType!: pulumi.Output<string>;
    /**
     * ID of the project within the CLB instance, `0` - Default Project.
     */
    public readonly projectId!: pulumi.Output<number | undefined>;
    /**
     * Security groups of the CLB instance. Supports both `OPEN` and `INTERNAL` CLBs.
     */
    public readonly securityGroups!: pulumi.Output<string[] | undefined>;
    /**
     * Setting slave zone id of cross available zone disaster recovery, only applicable to open CLB. this zone will undertake traffic when the master is down.
     */
    public readonly slaveZoneId!: pulumi.Output<string | undefined>;
    /**
     * Snat Ip List, required with `snat_pro=true`. NOTE: This argument cannot be read and modified here because dynamic ip is untraceable, please import resource `tencentcloud.Clb.SnatIp` to handle fixed ips.
     */
    public readonly snatIps!: pulumi.Output<outputs.Clb.InstanceSnatIp[] | undefined>;
    /**
     * Indicates whether Binding IPs of other VPCs feature switch.
     */
    public readonly snatPro!: pulumi.Output<boolean | undefined>;
    /**
     * Subnet ID of the CLB. Effective only for CLB within the VPC. Only supports `INTERNAL` CLBs. Default is `ipv4`.
     */
    public readonly subnetId!: pulumi.Output<string | undefined>;
    /**
     * The available tags within this CLB.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Region information of backend services are attached the CLB instance. Only supports `OPEN` CLBs.
     */
    public readonly targetRegionInfoRegion!: pulumi.Output<string>;
    /**
     * Vpc information of backend services are attached the CLB instance. Only supports `OPEN` CLBs.
     */
    public readonly targetRegionInfoVpcId!: pulumi.Output<string>;
    /**
     * Network operator, only applicable to open CLB. Valid values are `CMCC`(China Mobile), `CTCC`(Telecom), `CUCC`(China Unicom) and `BGP`. If this ISP is specified, network billing method can only use the bandwidth package billing (BANDWIDTH_PACKAGE).
     */
    public /*out*/ readonly vipIsp!: pulumi.Output<string>;
    /**
     * VPC ID of the CLB.
     */
    public readonly vpcId!: pulumi.Output<string>;
    /**
     * Available zone id, only applicable to open CLB.
     */
    public readonly zoneId!: pulumi.Output<string | undefined>;

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
            resourceInputs["addressIpVersion"] = state ? state.addressIpVersion : undefined;
            resourceInputs["bandwidthPackageId"] = state ? state.bandwidthPackageId : undefined;
            resourceInputs["clbName"] = state ? state.clbName : undefined;
            resourceInputs["clbVips"] = state ? state.clbVips : undefined;
            resourceInputs["internetBandwidthMaxOut"] = state ? state.internetBandwidthMaxOut : undefined;
            resourceInputs["internetChargeType"] = state ? state.internetChargeType : undefined;
            resourceInputs["loadBalancerPassToTarget"] = state ? state.loadBalancerPassToTarget : undefined;
            resourceInputs["logSetId"] = state ? state.logSetId : undefined;
            resourceInputs["logTopicId"] = state ? state.logTopicId : undefined;
            resourceInputs["masterZoneId"] = state ? state.masterZoneId : undefined;
            resourceInputs["networkType"] = state ? state.networkType : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["securityGroups"] = state ? state.securityGroups : undefined;
            resourceInputs["slaveZoneId"] = state ? state.slaveZoneId : undefined;
            resourceInputs["snatIps"] = state ? state.snatIps : undefined;
            resourceInputs["snatPro"] = state ? state.snatPro : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["targetRegionInfoRegion"] = state ? state.targetRegionInfoRegion : undefined;
            resourceInputs["targetRegionInfoVpcId"] = state ? state.targetRegionInfoVpcId : undefined;
            resourceInputs["vipIsp"] = state ? state.vipIsp : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if ((!args || args.clbName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clbName'");
            }
            if ((!args || args.networkType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkType'");
            }
            resourceInputs["addressIpVersion"] = args ? args.addressIpVersion : undefined;
            resourceInputs["bandwidthPackageId"] = args ? args.bandwidthPackageId : undefined;
            resourceInputs["clbName"] = args ? args.clbName : undefined;
            resourceInputs["internetBandwidthMaxOut"] = args ? args.internetBandwidthMaxOut : undefined;
            resourceInputs["internetChargeType"] = args ? args.internetChargeType : undefined;
            resourceInputs["loadBalancerPassToTarget"] = args ? args.loadBalancerPassToTarget : undefined;
            resourceInputs["logSetId"] = args ? args.logSetId : undefined;
            resourceInputs["logTopicId"] = args ? args.logTopicId : undefined;
            resourceInputs["masterZoneId"] = args ? args.masterZoneId : undefined;
            resourceInputs["networkType"] = args ? args.networkType : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["securityGroups"] = args ? args.securityGroups : undefined;
            resourceInputs["slaveZoneId"] = args ? args.slaveZoneId : undefined;
            resourceInputs["snatIps"] = args ? args.snatIps : undefined;
            resourceInputs["snatPro"] = args ? args.snatPro : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["targetRegionInfoRegion"] = args ? args.targetRegionInfoRegion : undefined;
            resourceInputs["targetRegionInfoVpcId"] = args ? args.targetRegionInfoVpcId : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["zoneId"] = args ? args.zoneId : undefined;
            resourceInputs["clbVips"] = undefined /*out*/;
            resourceInputs["vipIsp"] = undefined /*out*/;
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
     * IP version, only applicable to open CLB. Valid values are `ipv4`, `ipv6` and `IPv6FullChain`.
     */
    addressIpVersion?: pulumi.Input<string>;
    /**
     * Bandwidth package id. If set, the `internetChargeType` must be `BANDWIDTH_PACKAGE`.
     */
    bandwidthPackageId?: pulumi.Input<string>;
    /**
     * Name of the CLB. The name can only contain Chinese characters, English letters, numbers, underscore and hyphen '-'.
     */
    clbName?: pulumi.Input<string>;
    /**
     * The virtual service address table of the CLB.
     */
    clbVips?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Max bandwidth out, only applicable to open CLB. Valid value ranges is [1, 2048]. Unit is MB.
     */
    internetBandwidthMaxOut?: pulumi.Input<number>;
    /**
     * Internet charge type, only applicable to open CLB. Valid values are `TRAFFIC_POSTPAID_BY_HOUR`, `BANDWIDTH_POSTPAID_BY_HOUR` and `BANDWIDTH_PACKAGE`.
     */
    internetChargeType?: pulumi.Input<string>;
    /**
     * Whether the target allow flow come from clb. If value is true, only check security group of clb, or check both clb and backend instance security group.
     */
    loadBalancerPassToTarget?: pulumi.Input<boolean>;
    /**
     * The id of log set.
     */
    logSetId?: pulumi.Input<string>;
    /**
     * The id of log topic.
     */
    logTopicId?: pulumi.Input<string>;
    /**
     * Setting master zone id of cross available zone disaster recovery, only applicable to open CLB.
     */
    masterZoneId?: pulumi.Input<string>;
    /**
     * Type of CLB instance. Valid values: `OPEN` and `INTERNAL`.
     */
    networkType?: pulumi.Input<string>;
    /**
     * ID of the project within the CLB instance, `0` - Default Project.
     */
    projectId?: pulumi.Input<number>;
    /**
     * Security groups of the CLB instance. Supports both `OPEN` and `INTERNAL` CLBs.
     */
    securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Setting slave zone id of cross available zone disaster recovery, only applicable to open CLB. this zone will undertake traffic when the master is down.
     */
    slaveZoneId?: pulumi.Input<string>;
    /**
     * Snat Ip List, required with `snat_pro=true`. NOTE: This argument cannot be read and modified here because dynamic ip is untraceable, please import resource `tencentcloud.Clb.SnatIp` to handle fixed ips.
     */
    snatIps?: pulumi.Input<pulumi.Input<inputs.Clb.InstanceSnatIp>[]>;
    /**
     * Indicates whether Binding IPs of other VPCs feature switch.
     */
    snatPro?: pulumi.Input<boolean>;
    /**
     * Subnet ID of the CLB. Effective only for CLB within the VPC. Only supports `INTERNAL` CLBs. Default is `ipv4`.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * The available tags within this CLB.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Region information of backend services are attached the CLB instance. Only supports `OPEN` CLBs.
     */
    targetRegionInfoRegion?: pulumi.Input<string>;
    /**
     * Vpc information of backend services are attached the CLB instance. Only supports `OPEN` CLBs.
     */
    targetRegionInfoVpcId?: pulumi.Input<string>;
    /**
     * Network operator, only applicable to open CLB. Valid values are `CMCC`(China Mobile), `CTCC`(Telecom), `CUCC`(China Unicom) and `BGP`. If this ISP is specified, network billing method can only use the bandwidth package billing (BANDWIDTH_PACKAGE).
     */
    vipIsp?: pulumi.Input<string>;
    /**
     * VPC ID of the CLB.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * Available zone id, only applicable to open CLB.
     */
    zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * IP version, only applicable to open CLB. Valid values are `ipv4`, `ipv6` and `IPv6FullChain`.
     */
    addressIpVersion?: pulumi.Input<string>;
    /**
     * Bandwidth package id. If set, the `internetChargeType` must be `BANDWIDTH_PACKAGE`.
     */
    bandwidthPackageId?: pulumi.Input<string>;
    /**
     * Name of the CLB. The name can only contain Chinese characters, English letters, numbers, underscore and hyphen '-'.
     */
    clbName: pulumi.Input<string>;
    /**
     * Max bandwidth out, only applicable to open CLB. Valid value ranges is [1, 2048]. Unit is MB.
     */
    internetBandwidthMaxOut?: pulumi.Input<number>;
    /**
     * Internet charge type, only applicable to open CLB. Valid values are `TRAFFIC_POSTPAID_BY_HOUR`, `BANDWIDTH_POSTPAID_BY_HOUR` and `BANDWIDTH_PACKAGE`.
     */
    internetChargeType?: pulumi.Input<string>;
    /**
     * Whether the target allow flow come from clb. If value is true, only check security group of clb, or check both clb and backend instance security group.
     */
    loadBalancerPassToTarget?: pulumi.Input<boolean>;
    /**
     * The id of log set.
     */
    logSetId?: pulumi.Input<string>;
    /**
     * The id of log topic.
     */
    logTopicId?: pulumi.Input<string>;
    /**
     * Setting master zone id of cross available zone disaster recovery, only applicable to open CLB.
     */
    masterZoneId?: pulumi.Input<string>;
    /**
     * Type of CLB instance. Valid values: `OPEN` and `INTERNAL`.
     */
    networkType: pulumi.Input<string>;
    /**
     * ID of the project within the CLB instance, `0` - Default Project.
     */
    projectId?: pulumi.Input<number>;
    /**
     * Security groups of the CLB instance. Supports both `OPEN` and `INTERNAL` CLBs.
     */
    securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Setting slave zone id of cross available zone disaster recovery, only applicable to open CLB. this zone will undertake traffic when the master is down.
     */
    slaveZoneId?: pulumi.Input<string>;
    /**
     * Snat Ip List, required with `snat_pro=true`. NOTE: This argument cannot be read and modified here because dynamic ip is untraceable, please import resource `tencentcloud.Clb.SnatIp` to handle fixed ips.
     */
    snatIps?: pulumi.Input<pulumi.Input<inputs.Clb.InstanceSnatIp>[]>;
    /**
     * Indicates whether Binding IPs of other VPCs feature switch.
     */
    snatPro?: pulumi.Input<boolean>;
    /**
     * Subnet ID of the CLB. Effective only for CLB within the VPC. Only supports `INTERNAL` CLBs. Default is `ipv4`.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * The available tags within this CLB.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Region information of backend services are attached the CLB instance. Only supports `OPEN` CLBs.
     */
    targetRegionInfoRegion?: pulumi.Input<string>;
    /**
     * Vpc information of backend services are attached the CLB instance. Only supports `OPEN` CLBs.
     */
    targetRegionInfoVpcId?: pulumi.Input<string>;
    /**
     * VPC ID of the CLB.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * Available zone id, only applicable to open CLB.
     */
    zoneId?: pulumi.Input<string>;
}
