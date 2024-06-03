// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this resource to create IP strategy of API gateway.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const service = new tencentcloud.apigateway.Service("service", {
 *     serviceName: "niceservice",
 *     protocol: "http&https",
 *     serviceDesc: "your nice service",
 *     netTypes: [
 *         "INNER",
 *         "OUTER",
 *     ],
 *     ipVersion: "IPv4",
 * });
 * const test = new tencentcloud.apigateway.IpStrategy("test", {
 *     serviceId: service.id,
 *     strategyName: "tf_test",
 *     strategyType: "BLACK",
 *     strategyData: "9.9.9.9",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * IP strategy of API gateway can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import tencentcloud:ApiGateway/ipStrategy:IpStrategy test service-ohxqslqe#IPStrategy-q1lk8ud2
 * ```
 */
export class IpStrategy extends pulumi.CustomResource {
    /**
     * Get an existing IpStrategy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IpStrategyState, opts?: pulumi.CustomResourceOptions): IpStrategy {
        return new IpStrategy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:ApiGateway/ipStrategy:IpStrategy';

    /**
     * Returns true if the given object is an instance of IpStrategy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IpStrategy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IpStrategy.__pulumiType;
    }

    /**
     * Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The ID of the API gateway service.
     */
    public readonly serviceId!: pulumi.Output<string>;
    /**
     * IP address data.
     */
    public readonly strategyData!: pulumi.Output<string>;
    /**
     * IP policy ID.
     */
    public /*out*/ readonly strategyId!: pulumi.Output<string>;
    /**
     * User defined strategy name.
     */
    public readonly strategyName!: pulumi.Output<string>;
    /**
     * Blacklist or whitelist.
     */
    public readonly strategyType!: pulumi.Output<string>;

    /**
     * Create a IpStrategy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IpStrategyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IpStrategyArgs | IpStrategyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IpStrategyState | undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["serviceId"] = state ? state.serviceId : undefined;
            resourceInputs["strategyData"] = state ? state.strategyData : undefined;
            resourceInputs["strategyId"] = state ? state.strategyId : undefined;
            resourceInputs["strategyName"] = state ? state.strategyName : undefined;
            resourceInputs["strategyType"] = state ? state.strategyType : undefined;
        } else {
            const args = argsOrState as IpStrategyArgs | undefined;
            if ((!args || args.serviceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceId'");
            }
            if ((!args || args.strategyData === undefined) && !opts.urn) {
                throw new Error("Missing required property 'strategyData'");
            }
            if ((!args || args.strategyName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'strategyName'");
            }
            if ((!args || args.strategyType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'strategyType'");
            }
            resourceInputs["serviceId"] = args ? args.serviceId : undefined;
            resourceInputs["strategyData"] = args ? args.strategyData : undefined;
            resourceInputs["strategyName"] = args ? args.strategyName : undefined;
            resourceInputs["strategyType"] = args ? args.strategyType : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["strategyId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IpStrategy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IpStrategy resources.
 */
export interface IpStrategyState {
    /**
     * Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
     */
    createTime?: pulumi.Input<string>;
    /**
     * The ID of the API gateway service.
     */
    serviceId?: pulumi.Input<string>;
    /**
     * IP address data.
     */
    strategyData?: pulumi.Input<string>;
    /**
     * IP policy ID.
     */
    strategyId?: pulumi.Input<string>;
    /**
     * User defined strategy name.
     */
    strategyName?: pulumi.Input<string>;
    /**
     * Blacklist or whitelist.
     */
    strategyType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IpStrategy resource.
 */
export interface IpStrategyArgs {
    /**
     * The ID of the API gateway service.
     */
    serviceId: pulumi.Input<string>;
    /**
     * IP address data.
     */
    strategyData: pulumi.Input<string>;
    /**
     * User defined strategy name.
     */
    strategyName: pulumi.Input<string>;
    /**
     * Blacklist or whitelist.
     */
    strategyType: pulumi.Input<string>;
}
