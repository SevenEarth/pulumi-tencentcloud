// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a waf clb instance
 *
 * > **NOTE:** Region only supports `ap-guangzhou` and `ap-seoul`.
 *
 * ## Example Usage
 *
 * ### Create a basic waf premium clb instance
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const example = new tencentcloud.waf.ClbInstance("example", {
 *     goodsCategory: "premium_clb",
 *     instanceName: "tf-example-clb-waf",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ### Create a complete waf ultimateClb instance
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const example = new tencentcloud.waf.ClbInstance("example", {
 *     apiSecurity: 1,
 *     autoRenewFlag: 1,
 *     botManagement: 1,
 *     elasticMode: 1,
 *     goodsCategory: "ultimate_clb",
 *     instanceName: "tf-example-clb-waf",
 *     timeSpan: 1,
 *     timeUnit: "m",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ### Set waf ultimateClb instance qps limit
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const example = new tencentcloud.waf.ClbInstance("example", {
 *     apiSecurity: 1,
 *     autoRenewFlag: 1,
 *     botManagement: 1,
 *     elasticMode: 1,
 *     goodsCategory: "ultimate_clb",
 *     instanceName: "tf-example-clb-waf",
 *     qpsLimit: 200000,
 *     timeSpan: 1,
 *     timeUnit: "m",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class ClbInstance extends pulumi.CustomResource {
    /**
     * Get an existing ClbInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClbInstanceState, opts?: pulumi.CustomResourceOptions): ClbInstance {
        return new ClbInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Waf/clbInstance:ClbInstance';

    /**
     * Returns true if the given object is an instance of ClbInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ClbInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ClbInstance.__pulumiType;
    }

    /**
     * Whether to purchase API Security, 1: yes, 0: no. Default is 0.
     */
    public readonly apiSecurity!: pulumi.Output<number | undefined>;
    /**
     * Auto renew flag, 1: enable, 0: disable.
     */
    public readonly autoRenewFlag!: pulumi.Output<number | undefined>;
    /**
     * waf instance start time.
     */
    public /*out*/ readonly beginTime!: pulumi.Output<string>;
    /**
     * Whether to purchase Bot management, 1: yes, 0: no. Default is 0.
     */
    public readonly botManagement!: pulumi.Output<number | undefined>;
    /**
     * waf instance edition, clb or saas.
     */
    public /*out*/ readonly edition!: pulumi.Output<string>;
    /**
     * Is elastic billing enabled, 1: enable, 0: disable.
     */
    public readonly elasticMode!: pulumi.Output<number | undefined>;
    /**
     * Billing order parameters. support: premium_clb, enterprise_clb, ultimate_clb.
     */
    public readonly goodsCategory!: pulumi.Output<string>;
    /**
     * waf instance id.
     */
    public /*out*/ readonly instanceId!: pulumi.Output<string>;
    /**
     * Waf instance name.
     */
    public readonly instanceName!: pulumi.Output<string>;
    /**
     * QPS Limit, Minimum setting 10000. Only `elasticMode` is 1, can be set.
     */
    public readonly qpsLimit!: pulumi.Output<number>;
    /**
     * waf instance status.
     */
    public /*out*/ readonly status!: pulumi.Output<number>;
    /**
     * Time interval.
     */
    public readonly timeSpan!: pulumi.Output<number | undefined>;
    /**
     * Time unit, support d, m, y. d: day, m: month, y: year.
     */
    public readonly timeUnit!: pulumi.Output<string | undefined>;
    /**
     * waf instance valid time.
     */
    public /*out*/ readonly validTime!: pulumi.Output<string>;

    /**
     * Create a ClbInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClbInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClbInstanceArgs | ClbInstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ClbInstanceState | undefined;
            resourceInputs["apiSecurity"] = state ? state.apiSecurity : undefined;
            resourceInputs["autoRenewFlag"] = state ? state.autoRenewFlag : undefined;
            resourceInputs["beginTime"] = state ? state.beginTime : undefined;
            resourceInputs["botManagement"] = state ? state.botManagement : undefined;
            resourceInputs["edition"] = state ? state.edition : undefined;
            resourceInputs["elasticMode"] = state ? state.elasticMode : undefined;
            resourceInputs["goodsCategory"] = state ? state.goodsCategory : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["instanceName"] = state ? state.instanceName : undefined;
            resourceInputs["qpsLimit"] = state ? state.qpsLimit : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["timeSpan"] = state ? state.timeSpan : undefined;
            resourceInputs["timeUnit"] = state ? state.timeUnit : undefined;
            resourceInputs["validTime"] = state ? state.validTime : undefined;
        } else {
            const args = argsOrState as ClbInstanceArgs | undefined;
            if ((!args || args.goodsCategory === undefined) && !opts.urn) {
                throw new Error("Missing required property 'goodsCategory'");
            }
            resourceInputs["apiSecurity"] = args ? args.apiSecurity : undefined;
            resourceInputs["autoRenewFlag"] = args ? args.autoRenewFlag : undefined;
            resourceInputs["botManagement"] = args ? args.botManagement : undefined;
            resourceInputs["elasticMode"] = args ? args.elasticMode : undefined;
            resourceInputs["goodsCategory"] = args ? args.goodsCategory : undefined;
            resourceInputs["instanceName"] = args ? args.instanceName : undefined;
            resourceInputs["qpsLimit"] = args ? args.qpsLimit : undefined;
            resourceInputs["timeSpan"] = args ? args.timeSpan : undefined;
            resourceInputs["timeUnit"] = args ? args.timeUnit : undefined;
            resourceInputs["beginTime"] = undefined /*out*/;
            resourceInputs["edition"] = undefined /*out*/;
            resourceInputs["instanceId"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["validTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ClbInstance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ClbInstance resources.
 */
export interface ClbInstanceState {
    /**
     * Whether to purchase API Security, 1: yes, 0: no. Default is 0.
     */
    apiSecurity?: pulumi.Input<number>;
    /**
     * Auto renew flag, 1: enable, 0: disable.
     */
    autoRenewFlag?: pulumi.Input<number>;
    /**
     * waf instance start time.
     */
    beginTime?: pulumi.Input<string>;
    /**
     * Whether to purchase Bot management, 1: yes, 0: no. Default is 0.
     */
    botManagement?: pulumi.Input<number>;
    /**
     * waf instance edition, clb or saas.
     */
    edition?: pulumi.Input<string>;
    /**
     * Is elastic billing enabled, 1: enable, 0: disable.
     */
    elasticMode?: pulumi.Input<number>;
    /**
     * Billing order parameters. support: premium_clb, enterprise_clb, ultimate_clb.
     */
    goodsCategory?: pulumi.Input<string>;
    /**
     * waf instance id.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Waf instance name.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * QPS Limit, Minimum setting 10000. Only `elasticMode` is 1, can be set.
     */
    qpsLimit?: pulumi.Input<number>;
    /**
     * waf instance status.
     */
    status?: pulumi.Input<number>;
    /**
     * Time interval.
     */
    timeSpan?: pulumi.Input<number>;
    /**
     * Time unit, support d, m, y. d: day, m: month, y: year.
     */
    timeUnit?: pulumi.Input<string>;
    /**
     * waf instance valid time.
     */
    validTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ClbInstance resource.
 */
export interface ClbInstanceArgs {
    /**
     * Whether to purchase API Security, 1: yes, 0: no. Default is 0.
     */
    apiSecurity?: pulumi.Input<number>;
    /**
     * Auto renew flag, 1: enable, 0: disable.
     */
    autoRenewFlag?: pulumi.Input<number>;
    /**
     * Whether to purchase Bot management, 1: yes, 0: no. Default is 0.
     */
    botManagement?: pulumi.Input<number>;
    /**
     * Is elastic billing enabled, 1: enable, 0: disable.
     */
    elasticMode?: pulumi.Input<number>;
    /**
     * Billing order parameters. support: premium_clb, enterprise_clb, ultimate_clb.
     */
    goodsCategory: pulumi.Input<string>;
    /**
     * Waf instance name.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * QPS Limit, Minimum setting 10000. Only `elasticMode` is 1, can be set.
     */
    qpsLimit?: pulumi.Input<number>;
    /**
     * Time interval.
     */
    timeSpan?: pulumi.Input<number>;
    /**
     * Time unit, support d, m, y. d: day, m: month, y: year.
     */
    timeUnit?: pulumi.Input<string>;
}
