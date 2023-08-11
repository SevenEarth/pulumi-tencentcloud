// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a vpc flowLog
 *
 * > **NOTE:** The cloud server instance specifications that support stream log collection include: M6ce, M6p, SA3se, S4m, DA3, ITA3, I6t, I6, S5se, SA2, SK1, S4, S5, SN3ne, S3ne, S2ne, SA2a, S3ne, SW3a, SW3b, SW3ne, ITA3, IT5c, IT5, IT5c, IT3, I3, D3, DA2, D2, M6, MA2, M4, C6, IT3a, IT3b, IT3c, C4ne, CN3ne, C3ne, GI1, PNV4, GNV4v, GNV4, GT4, GI3X, GN7, GN7vw.
 *
 * > **NOTE:** The following models no longer support the collection of new stream logs, and the stock stream logs will no longer be reported for data from July 25, 2022: Standard models: S3, SA1, S2, S1;Memory type: M3, M2, M1;Calculation type: C4, CN3, C3, C2;Batch calculation type: BC1, BS1;HPCC: HCCIC5, HCCG5v.
 *
 * ## Import
 *
 * vpc flow_log can be imported using the flow log Id combine vpc Id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Vpc/flowLog:FlowLog flow_log flow_log_id fl-xxxx1234#vpc-yyyy5678
 * ```
 */
export class FlowLog extends pulumi.CustomResource {
    /**
     * Get an existing FlowLog resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FlowLogState, opts?: pulumi.CustomResourceOptions): FlowLog {
        return new FlowLog(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Vpc/flowLog:FlowLog';

    /**
     * Returns true if the given object is an instance of FlowLog.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FlowLog {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FlowLog.__pulumiType;
    }

    /**
     * Specify flow log storage id.
     */
    public readonly cloudLogId!: pulumi.Output<string | undefined>;
    /**
     * Specify flow log storage region, default using current.
     */
    public readonly cloudLogRegion!: pulumi.Output<string>;
    /**
     * Specify flow Log description.
     */
    public readonly flowLogDescription!: pulumi.Output<string | undefined>;
    /**
     * Specify flow log name.
     */
    public readonly flowLogName!: pulumi.Output<string>;
    /**
     * Specify consumer detail, required while `storageType` is `ckafka`.
     */
    public readonly flowLogStorage!: pulumi.Output<outputs.Vpc.FlowLogFlowLogStorage>;
    /**
     * Specify resource unique Id of `resourceType` configured.
     */
    public readonly resourceId!: pulumi.Output<string>;
    /**
     * Specify resource type. NOTE: Only support `NETWORKINTERFACE` for now. Values: `VPC`, `SUBNET`, `NETWORKINTERFACE`, `CCN`, `NAT`, `DCG`.
     */
    public readonly resourceType!: pulumi.Output<string>;
    /**
     * Specify consumer type, values: `cls`, `ckafka`.
     */
    public readonly storageType!: pulumi.Output<string | undefined>;
    /**
     * Tag description list.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Specify log traffic type, values: `ACCEPT`, `REJECT`, `ALL`.
     */
    public readonly trafficType!: pulumi.Output<string>;
    /**
     * Specify vpc Id, ignore while `resourceType` is `CCN` (unsupported) but required while other types.
     */
    public readonly vpcId!: pulumi.Output<string | undefined>;

    /**
     * Create a FlowLog resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FlowLogArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FlowLogArgs | FlowLogState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FlowLogState | undefined;
            resourceInputs["cloudLogId"] = state ? state.cloudLogId : undefined;
            resourceInputs["cloudLogRegion"] = state ? state.cloudLogRegion : undefined;
            resourceInputs["flowLogDescription"] = state ? state.flowLogDescription : undefined;
            resourceInputs["flowLogName"] = state ? state.flowLogName : undefined;
            resourceInputs["flowLogStorage"] = state ? state.flowLogStorage : undefined;
            resourceInputs["resourceId"] = state ? state.resourceId : undefined;
            resourceInputs["resourceType"] = state ? state.resourceType : undefined;
            resourceInputs["storageType"] = state ? state.storageType : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["trafficType"] = state ? state.trafficType : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as FlowLogArgs | undefined;
            if ((!args || args.flowLogName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'flowLogName'");
            }
            if ((!args || args.resourceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceId'");
            }
            if ((!args || args.resourceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceType'");
            }
            if ((!args || args.trafficType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'trafficType'");
            }
            resourceInputs["cloudLogId"] = args ? args.cloudLogId : undefined;
            resourceInputs["cloudLogRegion"] = args ? args.cloudLogRegion : undefined;
            resourceInputs["flowLogDescription"] = args ? args.flowLogDescription : undefined;
            resourceInputs["flowLogName"] = args ? args.flowLogName : undefined;
            resourceInputs["flowLogStorage"] = args ? args.flowLogStorage : undefined;
            resourceInputs["resourceId"] = args ? args.resourceId : undefined;
            resourceInputs["resourceType"] = args ? args.resourceType : undefined;
            resourceInputs["storageType"] = args ? args.storageType : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["trafficType"] = args ? args.trafficType : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FlowLog.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FlowLog resources.
 */
export interface FlowLogState {
    /**
     * Specify flow log storage id.
     */
    cloudLogId?: pulumi.Input<string>;
    /**
     * Specify flow log storage region, default using current.
     */
    cloudLogRegion?: pulumi.Input<string>;
    /**
     * Specify flow Log description.
     */
    flowLogDescription?: pulumi.Input<string>;
    /**
     * Specify flow log name.
     */
    flowLogName?: pulumi.Input<string>;
    /**
     * Specify consumer detail, required while `storageType` is `ckafka`.
     */
    flowLogStorage?: pulumi.Input<inputs.Vpc.FlowLogFlowLogStorage>;
    /**
     * Specify resource unique Id of `resourceType` configured.
     */
    resourceId?: pulumi.Input<string>;
    /**
     * Specify resource type. NOTE: Only support `NETWORKINTERFACE` for now. Values: `VPC`, `SUBNET`, `NETWORKINTERFACE`, `CCN`, `NAT`, `DCG`.
     */
    resourceType?: pulumi.Input<string>;
    /**
     * Specify consumer type, values: `cls`, `ckafka`.
     */
    storageType?: pulumi.Input<string>;
    /**
     * Tag description list.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Specify log traffic type, values: `ACCEPT`, `REJECT`, `ALL`.
     */
    trafficType?: pulumi.Input<string>;
    /**
     * Specify vpc Id, ignore while `resourceType` is `CCN` (unsupported) but required while other types.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FlowLog resource.
 */
export interface FlowLogArgs {
    /**
     * Specify flow log storage id.
     */
    cloudLogId?: pulumi.Input<string>;
    /**
     * Specify flow log storage region, default using current.
     */
    cloudLogRegion?: pulumi.Input<string>;
    /**
     * Specify flow Log description.
     */
    flowLogDescription?: pulumi.Input<string>;
    /**
     * Specify flow log name.
     */
    flowLogName: pulumi.Input<string>;
    /**
     * Specify consumer detail, required while `storageType` is `ckafka`.
     */
    flowLogStorage?: pulumi.Input<inputs.Vpc.FlowLogFlowLogStorage>;
    /**
     * Specify resource unique Id of `resourceType` configured.
     */
    resourceId: pulumi.Input<string>;
    /**
     * Specify resource type. NOTE: Only support `NETWORKINTERFACE` for now. Values: `VPC`, `SUBNET`, `NETWORKINTERFACE`, `CCN`, `NAT`, `DCG`.
     */
    resourceType: pulumi.Input<string>;
    /**
     * Specify consumer type, values: `cls`, `ckafka`.
     */
    storageType?: pulumi.Input<string>;
    /**
     * Tag description list.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Specify log traffic type, values: `ACCEPT`, `REJECT`, `ALL`.
     */
    trafficType: pulumi.Input<string>;
    /**
     * Specify vpc Id, ignore while `resourceType` is `CCN` (unsupported) but required while other types.
     */
    vpcId?: pulumi.Input<string>;
}
