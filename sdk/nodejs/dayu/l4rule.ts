// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this resource to create dayu layer 4 rule
 *
 * > **NOTE:** This resource only support resource Anti-DDoS of type `bgpip` and `net`
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const testRule = new tencentcloud.Dayu.L4Rule("test_rule", {
 *     dPort: 60,
 *     healthCheckHealthNum: 5,
 *     healthCheckInterval: 35,
 *     healthCheckSwitch: true,
 *     healthCheckTimeout: 30,
 *     healthCheckUnhealthNum: 10,
 *     protocol: "TCP",
 *     resourceId: "bgpip-00000294",
 *     resourceType: "bgpip",
 *     sPort: 80,
 *     sessionSwitch: false,
 *     sessionTime: 300,
 *     sourceLists: [
 *         {
 *             source: "1.1.1.1",
 *             weight: 100,
 *         },
 *         {
 *             source: "2.2.2.2",
 *             weight: 50,
 *         },
 *     ],
 *     sourceType: 2,
 * });
 * ```
 */
export class L4Rule extends pulumi.CustomResource {
    /**
     * Get an existing L4Rule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: L4RuleState, opts?: pulumi.CustomResourceOptions): L4Rule {
        return new L4Rule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Dayu/l4Rule:L4Rule';

    /**
     * Returns true if the given object is an instance of L4Rule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is L4Rule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === L4Rule.__pulumiType;
    }

    /**
     * The destination port of the L4 rule.
     */
    public readonly dPort!: pulumi.Output<number>;
    /**
     * Health threshold of health check, and the default is 3. If a success result is returned for the health check 3 consecutive times, indicates that the forwarding is normal. The value range is 2-10.
     */
    public readonly healthCheckHealthNum!: pulumi.Output<number>;
    /**
     * Interval time of health check. The value range is 10-60 sec, and the default is 15 sec.
     */
    public readonly healthCheckInterval!: pulumi.Output<number>;
    /**
     * Indicates whether health check is enabled. The default is `false`. Only valid when source list has more than one source item.
     */
    public readonly healthCheckSwitch!: pulumi.Output<boolean>;
    /**
     * HTTP Status Code. The default is 26 and value range is 2-60.
     */
    public readonly healthCheckTimeout!: pulumi.Output<number>;
    /**
     * Unhealthy threshold of health check, and the default is 3. If the unhealthy result is returned 3 consecutive times, indicates that the forwarding is abnormal. The value range is 2-10.
     */
    public readonly healthCheckUnhealthNum!: pulumi.Output<number>;
    /**
     * LB type of the rule. Valid values: `1`, `2`. `1` for weight cycling and `2` for IP hash.
     */
    public /*out*/ readonly lbType!: pulumi.Output<number>;
    /**
     * Name of the rule. When the `resourceType` is `net`, this field should be set with valid domain.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Protocol of the rule. Valid values: `http`, `https`. When `sourceType` is 1(host source), the value of this field can only set with `tcp`.
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * ID of the resource that the layer 4 rule works for.
     */
    public readonly resourceId!: pulumi.Output<string>;
    /**
     * Type of the resource that the layer 4 rule works for. Valid values: `bgpip` and `net`.
     */
    public readonly resourceType!: pulumi.Output<string>;
    /**
     * ID of the layer 4 rule.
     */
    public /*out*/ readonly ruleId!: pulumi.Output<string>;
    /**
     * The source port of the L4 rule.
     */
    public readonly sPort!: pulumi.Output<number>;
    /**
     * Indicate that the session will keep or not, and default value is `false`.
     */
    public readonly sessionSwitch!: pulumi.Output<boolean | undefined>;
    /**
     * Session keep time, only valid when `sessionSwitch` is true, the available value ranges from 1 to 300 and unit is second.
     */
    public readonly sessionTime!: pulumi.Output<number>;
    /**
     * Source list of the rule, it can be a set of ip sources or a set of domain sources. The number of items ranges from 1 to 20.
     */
    public readonly sourceLists!: pulumi.Output<outputs.Dayu.L4RuleSourceList[]>;
    /**
     * Source type, `1` for source of host, `2` for source of IP.
     */
    public readonly sourceType!: pulumi.Output<number>;

    /**
     * Create a L4Rule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: L4RuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: L4RuleArgs | L4RuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as L4RuleState | undefined;
            resourceInputs["dPort"] = state ? state.dPort : undefined;
            resourceInputs["healthCheckHealthNum"] = state ? state.healthCheckHealthNum : undefined;
            resourceInputs["healthCheckInterval"] = state ? state.healthCheckInterval : undefined;
            resourceInputs["healthCheckSwitch"] = state ? state.healthCheckSwitch : undefined;
            resourceInputs["healthCheckTimeout"] = state ? state.healthCheckTimeout : undefined;
            resourceInputs["healthCheckUnhealthNum"] = state ? state.healthCheckUnhealthNum : undefined;
            resourceInputs["lbType"] = state ? state.lbType : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["resourceId"] = state ? state.resourceId : undefined;
            resourceInputs["resourceType"] = state ? state.resourceType : undefined;
            resourceInputs["ruleId"] = state ? state.ruleId : undefined;
            resourceInputs["sPort"] = state ? state.sPort : undefined;
            resourceInputs["sessionSwitch"] = state ? state.sessionSwitch : undefined;
            resourceInputs["sessionTime"] = state ? state.sessionTime : undefined;
            resourceInputs["sourceLists"] = state ? state.sourceLists : undefined;
            resourceInputs["sourceType"] = state ? state.sourceType : undefined;
        } else {
            const args = argsOrState as L4RuleArgs | undefined;
            if ((!args || args.dPort === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dPort'");
            }
            if ((!args || args.protocol === undefined) && !opts.urn) {
                throw new Error("Missing required property 'protocol'");
            }
            if ((!args || args.resourceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceId'");
            }
            if ((!args || args.resourceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceType'");
            }
            if ((!args || args.sPort === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sPort'");
            }
            if ((!args || args.sourceLists === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceLists'");
            }
            if ((!args || args.sourceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceType'");
            }
            resourceInputs["dPort"] = args ? args.dPort : undefined;
            resourceInputs["healthCheckHealthNum"] = args ? args.healthCheckHealthNum : undefined;
            resourceInputs["healthCheckInterval"] = args ? args.healthCheckInterval : undefined;
            resourceInputs["healthCheckSwitch"] = args ? args.healthCheckSwitch : undefined;
            resourceInputs["healthCheckTimeout"] = args ? args.healthCheckTimeout : undefined;
            resourceInputs["healthCheckUnhealthNum"] = args ? args.healthCheckUnhealthNum : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["resourceId"] = args ? args.resourceId : undefined;
            resourceInputs["resourceType"] = args ? args.resourceType : undefined;
            resourceInputs["sPort"] = args ? args.sPort : undefined;
            resourceInputs["sessionSwitch"] = args ? args.sessionSwitch : undefined;
            resourceInputs["sessionTime"] = args ? args.sessionTime : undefined;
            resourceInputs["sourceLists"] = args ? args.sourceLists : undefined;
            resourceInputs["sourceType"] = args ? args.sourceType : undefined;
            resourceInputs["lbType"] = undefined /*out*/;
            resourceInputs["ruleId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(L4Rule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering L4Rule resources.
 */
export interface L4RuleState {
    /**
     * The destination port of the L4 rule.
     */
    dPort?: pulumi.Input<number>;
    /**
     * Health threshold of health check, and the default is 3. If a success result is returned for the health check 3 consecutive times, indicates that the forwarding is normal. The value range is 2-10.
     */
    healthCheckHealthNum?: pulumi.Input<number>;
    /**
     * Interval time of health check. The value range is 10-60 sec, and the default is 15 sec.
     */
    healthCheckInterval?: pulumi.Input<number>;
    /**
     * Indicates whether health check is enabled. The default is `false`. Only valid when source list has more than one source item.
     */
    healthCheckSwitch?: pulumi.Input<boolean>;
    /**
     * HTTP Status Code. The default is 26 and value range is 2-60.
     */
    healthCheckTimeout?: pulumi.Input<number>;
    /**
     * Unhealthy threshold of health check, and the default is 3. If the unhealthy result is returned 3 consecutive times, indicates that the forwarding is abnormal. The value range is 2-10.
     */
    healthCheckUnhealthNum?: pulumi.Input<number>;
    /**
     * LB type of the rule. Valid values: `1`, `2`. `1` for weight cycling and `2` for IP hash.
     */
    lbType?: pulumi.Input<number>;
    /**
     * Name of the rule. When the `resourceType` is `net`, this field should be set with valid domain.
     */
    name?: pulumi.Input<string>;
    /**
     * Protocol of the rule. Valid values: `http`, `https`. When `sourceType` is 1(host source), the value of this field can only set with `tcp`.
     */
    protocol?: pulumi.Input<string>;
    /**
     * ID of the resource that the layer 4 rule works for.
     */
    resourceId?: pulumi.Input<string>;
    /**
     * Type of the resource that the layer 4 rule works for. Valid values: `bgpip` and `net`.
     */
    resourceType?: pulumi.Input<string>;
    /**
     * ID of the layer 4 rule.
     */
    ruleId?: pulumi.Input<string>;
    /**
     * The source port of the L4 rule.
     */
    sPort?: pulumi.Input<number>;
    /**
     * Indicate that the session will keep or not, and default value is `false`.
     */
    sessionSwitch?: pulumi.Input<boolean>;
    /**
     * Session keep time, only valid when `sessionSwitch` is true, the available value ranges from 1 to 300 and unit is second.
     */
    sessionTime?: pulumi.Input<number>;
    /**
     * Source list of the rule, it can be a set of ip sources or a set of domain sources. The number of items ranges from 1 to 20.
     */
    sourceLists?: pulumi.Input<pulumi.Input<inputs.Dayu.L4RuleSourceList>[]>;
    /**
     * Source type, `1` for source of host, `2` for source of IP.
     */
    sourceType?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a L4Rule resource.
 */
export interface L4RuleArgs {
    /**
     * The destination port of the L4 rule.
     */
    dPort: pulumi.Input<number>;
    /**
     * Health threshold of health check, and the default is 3. If a success result is returned for the health check 3 consecutive times, indicates that the forwarding is normal. The value range is 2-10.
     */
    healthCheckHealthNum?: pulumi.Input<number>;
    /**
     * Interval time of health check. The value range is 10-60 sec, and the default is 15 sec.
     */
    healthCheckInterval?: pulumi.Input<number>;
    /**
     * Indicates whether health check is enabled. The default is `false`. Only valid when source list has more than one source item.
     */
    healthCheckSwitch?: pulumi.Input<boolean>;
    /**
     * HTTP Status Code. The default is 26 and value range is 2-60.
     */
    healthCheckTimeout?: pulumi.Input<number>;
    /**
     * Unhealthy threshold of health check, and the default is 3. If the unhealthy result is returned 3 consecutive times, indicates that the forwarding is abnormal. The value range is 2-10.
     */
    healthCheckUnhealthNum?: pulumi.Input<number>;
    /**
     * Name of the rule. When the `resourceType` is `net`, this field should be set with valid domain.
     */
    name?: pulumi.Input<string>;
    /**
     * Protocol of the rule. Valid values: `http`, `https`. When `sourceType` is 1(host source), the value of this field can only set with `tcp`.
     */
    protocol: pulumi.Input<string>;
    /**
     * ID of the resource that the layer 4 rule works for.
     */
    resourceId: pulumi.Input<string>;
    /**
     * Type of the resource that the layer 4 rule works for. Valid values: `bgpip` and `net`.
     */
    resourceType: pulumi.Input<string>;
    /**
     * The source port of the L4 rule.
     */
    sPort: pulumi.Input<number>;
    /**
     * Indicate that the session will keep or not, and default value is `false`.
     */
    sessionSwitch?: pulumi.Input<boolean>;
    /**
     * Session keep time, only valid when `sessionSwitch` is true, the available value ranges from 1 to 300 and unit is second.
     */
    sessionTime?: pulumi.Input<number>;
    /**
     * Source list of the rule, it can be a set of ip sources or a set of domain sources. The number of items ranges from 1 to 20.
     */
    sourceLists: pulumi.Input<pulumi.Input<inputs.Dayu.L4RuleSourceList>[]>;
    /**
     * Source type, `1` for source of host, `2` for source of IP.
     */
    sourceType: pulumi.Input<number>;
}
