// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this resource to create dayu layer 7 rule
 *
 * > **NOTE:** This resource only support resource Anti-DDoS of type `bgpip`
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const testRule = new tencentcloud.dayu.L7Rule("testRule", {
 *     domain: "zhaoshaona.com",
 *     healthCheckCode: 31,
 *     healthCheckHealthNum: 5,
 *     healthCheckInterval: 30,
 *     healthCheckMethod: "GET",
 *     healthCheckPath: "/",
 *     healthCheckSwitch: true,
 *     healthCheckUnhealthNum: 10,
 *     protocol: "https",
 *     resourceId: "bgpip-00000294",
 *     resourceType: "bgpip",
 *     sourceLists: [
 *         "1.1.1.1:80",
 *         "2.2.2.2",
 *     ],
 *     sourceType: 2,
 *     sslId: "%s",
 *     "switch": true,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class L7Rule extends pulumi.CustomResource {
    /**
     * Get an existing L7Rule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: L7RuleState, opts?: pulumi.CustomResourceOptions): L7Rule {
        return new L7Rule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Dayu/l7Rule:L7Rule';

    /**
     * Returns true if the given object is an instance of L7Rule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is L7Rule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === L7Rule.__pulumiType;
    }

    /**
     * Domain that the layer 7 rule works for. Valid string length ranges from 0 to 80.
     */
    public readonly domain!: pulumi.Output<string>;
    /**
     * HTTP Status Code. The default is `26`. Valid value ranges: [1~31]. `1` means the return value '1xx' is health. `2` means the return value '2xx' is health. `4` means the return value '3xx' is health. `8` means the return value '4xx' is health. `16` means the return value '5xx' is health. If you want multiple return codes to indicate health, need to add the corresponding values.
     */
    public readonly healthCheckCode!: pulumi.Output<number>;
    /**
     * Health threshold of health check, and the default is `3`. If a success result is returned for the health check 3 consecutive times, indicates that the forwarding is normal. The value range is [2-10].
     */
    public readonly healthCheckHealthNum!: pulumi.Output<number>;
    /**
     * Interval time of health check. Valid value ranges: [10~60]sec. The default is 15 sec.
     */
    public readonly healthCheckInterval!: pulumi.Output<number>;
    /**
     * Methods of health check. The default is 'HEAD', the available value are 'HEAD' and 'GET'.
     */
    public readonly healthCheckMethod!: pulumi.Output<string>;
    /**
     * Path of health check. The default is `/`.
     */
    public readonly healthCheckPath!: pulumi.Output<string>;
    /**
     * Indicates whether health check is enabled. The default is `false`.
     */
    public readonly healthCheckSwitch!: pulumi.Output<boolean>;
    /**
     * Unhealthy threshold of health check, and the default is `3`. If the unhealthy result is returned 3 consecutive times, indicates that the forwarding is abnormal. The value range is [2-10].
     */
    public readonly healthCheckUnhealthNum!: pulumi.Output<number>;
    /**
     * Name of the rule.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Protocol of the rule. Valid values: `http`, `https`.
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * ID of the resource that the layer 7 rule works for.
     */
    public readonly resourceId!: pulumi.Output<string>;
    /**
     * Type of the resource that the layer 7 rule works for, valid value is `bgpip`.
     */
    public readonly resourceType!: pulumi.Output<string>;
    /**
     * ID of the layer 7 rule.
     */
    public /*out*/ readonly ruleId!: pulumi.Output<string>;
    /**
     * Source list of the rule, it can be a set of ip sources or a set of domain sources. The number of items ranges from 1 to 16.
     */
    public readonly sourceLists!: pulumi.Output<string[]>;
    /**
     * Source type, `1` for source of host, `2` for source of IP.
     */
    public readonly sourceType!: pulumi.Output<number>;
    /**
     * SSL ID, when the `protocol` is `https`, the field should be set with valid SSL id.
     */
    public readonly sslId!: pulumi.Output<string | undefined>;
    /**
     * Status of the rule. `0` for create/modify success, `2` for create/modify fail, `3` for delete success, `5` for delete failed, `6` for waiting to be created/modified, `7` for waiting to be deleted and 8 for waiting to get SSL ID.
     */
    public /*out*/ readonly status!: pulumi.Output<number>;
    /**
     * Indicate the rule will take effect or not.
     */
    public readonly switch!: pulumi.Output<boolean>;

    /**
     * Create a L7Rule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: L7RuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: L7RuleArgs | L7RuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as L7RuleState | undefined;
            resourceInputs["domain"] = state ? state.domain : undefined;
            resourceInputs["healthCheckCode"] = state ? state.healthCheckCode : undefined;
            resourceInputs["healthCheckHealthNum"] = state ? state.healthCheckHealthNum : undefined;
            resourceInputs["healthCheckInterval"] = state ? state.healthCheckInterval : undefined;
            resourceInputs["healthCheckMethod"] = state ? state.healthCheckMethod : undefined;
            resourceInputs["healthCheckPath"] = state ? state.healthCheckPath : undefined;
            resourceInputs["healthCheckSwitch"] = state ? state.healthCheckSwitch : undefined;
            resourceInputs["healthCheckUnhealthNum"] = state ? state.healthCheckUnhealthNum : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["resourceId"] = state ? state.resourceId : undefined;
            resourceInputs["resourceType"] = state ? state.resourceType : undefined;
            resourceInputs["ruleId"] = state ? state.ruleId : undefined;
            resourceInputs["sourceLists"] = state ? state.sourceLists : undefined;
            resourceInputs["sourceType"] = state ? state.sourceType : undefined;
            resourceInputs["sslId"] = state ? state.sslId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["switch"] = state ? state.switch : undefined;
        } else {
            const args = argsOrState as L7RuleArgs | undefined;
            if ((!args || args.domain === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domain'");
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
            if ((!args || args.sourceLists === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceLists'");
            }
            if ((!args || args.sourceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceType'");
            }
            if ((!args || args.switch === undefined) && !opts.urn) {
                throw new Error("Missing required property 'switch'");
            }
            resourceInputs["domain"] = args ? args.domain : undefined;
            resourceInputs["healthCheckCode"] = args ? args.healthCheckCode : undefined;
            resourceInputs["healthCheckHealthNum"] = args ? args.healthCheckHealthNum : undefined;
            resourceInputs["healthCheckInterval"] = args ? args.healthCheckInterval : undefined;
            resourceInputs["healthCheckMethod"] = args ? args.healthCheckMethod : undefined;
            resourceInputs["healthCheckPath"] = args ? args.healthCheckPath : undefined;
            resourceInputs["healthCheckSwitch"] = args ? args.healthCheckSwitch : undefined;
            resourceInputs["healthCheckUnhealthNum"] = args ? args.healthCheckUnhealthNum : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["resourceId"] = args ? args.resourceId : undefined;
            resourceInputs["resourceType"] = args ? args.resourceType : undefined;
            resourceInputs["sourceLists"] = args ? args.sourceLists : undefined;
            resourceInputs["sourceType"] = args ? args.sourceType : undefined;
            resourceInputs["sslId"] = args ? args.sslId : undefined;
            resourceInputs["switch"] = args ? args.switch : undefined;
            resourceInputs["ruleId"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(L7Rule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering L7Rule resources.
 */
export interface L7RuleState {
    /**
     * Domain that the layer 7 rule works for. Valid string length ranges from 0 to 80.
     */
    domain?: pulumi.Input<string>;
    /**
     * HTTP Status Code. The default is `26`. Valid value ranges: [1~31]. `1` means the return value '1xx' is health. `2` means the return value '2xx' is health. `4` means the return value '3xx' is health. `8` means the return value '4xx' is health. `16` means the return value '5xx' is health. If you want multiple return codes to indicate health, need to add the corresponding values.
     */
    healthCheckCode?: pulumi.Input<number>;
    /**
     * Health threshold of health check, and the default is `3`. If a success result is returned for the health check 3 consecutive times, indicates that the forwarding is normal. The value range is [2-10].
     */
    healthCheckHealthNum?: pulumi.Input<number>;
    /**
     * Interval time of health check. Valid value ranges: [10~60]sec. The default is 15 sec.
     */
    healthCheckInterval?: pulumi.Input<number>;
    /**
     * Methods of health check. The default is 'HEAD', the available value are 'HEAD' and 'GET'.
     */
    healthCheckMethod?: pulumi.Input<string>;
    /**
     * Path of health check. The default is `/`.
     */
    healthCheckPath?: pulumi.Input<string>;
    /**
     * Indicates whether health check is enabled. The default is `false`.
     */
    healthCheckSwitch?: pulumi.Input<boolean>;
    /**
     * Unhealthy threshold of health check, and the default is `3`. If the unhealthy result is returned 3 consecutive times, indicates that the forwarding is abnormal. The value range is [2-10].
     */
    healthCheckUnhealthNum?: pulumi.Input<number>;
    /**
     * Name of the rule.
     */
    name?: pulumi.Input<string>;
    /**
     * Protocol of the rule. Valid values: `http`, `https`.
     */
    protocol?: pulumi.Input<string>;
    /**
     * ID of the resource that the layer 7 rule works for.
     */
    resourceId?: pulumi.Input<string>;
    /**
     * Type of the resource that the layer 7 rule works for, valid value is `bgpip`.
     */
    resourceType?: pulumi.Input<string>;
    /**
     * ID of the layer 7 rule.
     */
    ruleId?: pulumi.Input<string>;
    /**
     * Source list of the rule, it can be a set of ip sources or a set of domain sources. The number of items ranges from 1 to 16.
     */
    sourceLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Source type, `1` for source of host, `2` for source of IP.
     */
    sourceType?: pulumi.Input<number>;
    /**
     * SSL ID, when the `protocol` is `https`, the field should be set with valid SSL id.
     */
    sslId?: pulumi.Input<string>;
    /**
     * Status of the rule. `0` for create/modify success, `2` for create/modify fail, `3` for delete success, `5` for delete failed, `6` for waiting to be created/modified, `7` for waiting to be deleted and 8 for waiting to get SSL ID.
     */
    status?: pulumi.Input<number>;
    /**
     * Indicate the rule will take effect or not.
     */
    switch?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a L7Rule resource.
 */
export interface L7RuleArgs {
    /**
     * Domain that the layer 7 rule works for. Valid string length ranges from 0 to 80.
     */
    domain: pulumi.Input<string>;
    /**
     * HTTP Status Code. The default is `26`. Valid value ranges: [1~31]. `1` means the return value '1xx' is health. `2` means the return value '2xx' is health. `4` means the return value '3xx' is health. `8` means the return value '4xx' is health. `16` means the return value '5xx' is health. If you want multiple return codes to indicate health, need to add the corresponding values.
     */
    healthCheckCode?: pulumi.Input<number>;
    /**
     * Health threshold of health check, and the default is `3`. If a success result is returned for the health check 3 consecutive times, indicates that the forwarding is normal. The value range is [2-10].
     */
    healthCheckHealthNum?: pulumi.Input<number>;
    /**
     * Interval time of health check. Valid value ranges: [10~60]sec. The default is 15 sec.
     */
    healthCheckInterval?: pulumi.Input<number>;
    /**
     * Methods of health check. The default is 'HEAD', the available value are 'HEAD' and 'GET'.
     */
    healthCheckMethod?: pulumi.Input<string>;
    /**
     * Path of health check. The default is `/`.
     */
    healthCheckPath?: pulumi.Input<string>;
    /**
     * Indicates whether health check is enabled. The default is `false`.
     */
    healthCheckSwitch?: pulumi.Input<boolean>;
    /**
     * Unhealthy threshold of health check, and the default is `3`. If the unhealthy result is returned 3 consecutive times, indicates that the forwarding is abnormal. The value range is [2-10].
     */
    healthCheckUnhealthNum?: pulumi.Input<number>;
    /**
     * Name of the rule.
     */
    name?: pulumi.Input<string>;
    /**
     * Protocol of the rule. Valid values: `http`, `https`.
     */
    protocol: pulumi.Input<string>;
    /**
     * ID of the resource that the layer 7 rule works for.
     */
    resourceId: pulumi.Input<string>;
    /**
     * Type of the resource that the layer 7 rule works for, valid value is `bgpip`.
     */
    resourceType: pulumi.Input<string>;
    /**
     * Source list of the rule, it can be a set of ip sources or a set of domain sources. The number of items ranges from 1 to 16.
     */
    sourceLists: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Source type, `1` for source of host, `2` for source of IP.
     */
    sourceType: pulumi.Input<number>;
    /**
     * SSL ID, when the `protocol` is `https`, the field should be set with valid SSL id.
     */
    sslId?: pulumi.Input<string>;
    /**
     * Indicate the rule will take effect or not.
     */
    switch: pulumi.Input<boolean>;
}
