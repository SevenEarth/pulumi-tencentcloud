// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export class CngwRouteRateLimit extends pulumi.CustomResource {
    /**
     * Get an existing CngwRouteRateLimit resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CngwRouteRateLimitState, opts?: pulumi.CustomResourceOptions): CngwRouteRateLimit {
        return new CngwRouteRateLimit(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Tse/cngwRouteRateLimit:CngwRouteRateLimit';

    /**
     * Returns true if the given object is an instance of CngwRouteRateLimit.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CngwRouteRateLimit {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CngwRouteRateLimit.__pulumiType;
    }

    /**
     * gateway ID.
     */
    public readonly gatewayId!: pulumi.Output<string>;
    /**
     * rate limit configuration.
     */
    public readonly limitDetail!: pulumi.Output<outputs.Tse.CngwRouteRateLimitLimitDetail>;
    /**
     * Route id, or route name.
     */
    public readonly routeId!: pulumi.Output<string>;
    /**
     * Tag description list.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a CngwRouteRateLimit resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CngwRouteRateLimitArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CngwRouteRateLimitArgs | CngwRouteRateLimitState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CngwRouteRateLimitState | undefined;
            resourceInputs["gatewayId"] = state ? state.gatewayId : undefined;
            resourceInputs["limitDetail"] = state ? state.limitDetail : undefined;
            resourceInputs["routeId"] = state ? state.routeId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as CngwRouteRateLimitArgs | undefined;
            if ((!args || args.gatewayId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'gatewayId'");
            }
            if ((!args || args.limitDetail === undefined) && !opts.urn) {
                throw new Error("Missing required property 'limitDetail'");
            }
            if ((!args || args.routeId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'routeId'");
            }
            resourceInputs["gatewayId"] = args ? args.gatewayId : undefined;
            resourceInputs["limitDetail"] = args ? args.limitDetail : undefined;
            resourceInputs["routeId"] = args ? args.routeId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CngwRouteRateLimit.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CngwRouteRateLimit resources.
 */
export interface CngwRouteRateLimitState {
    /**
     * gateway ID.
     */
    gatewayId?: pulumi.Input<string>;
    /**
     * rate limit configuration.
     */
    limitDetail?: pulumi.Input<inputs.Tse.CngwRouteRateLimitLimitDetail>;
    /**
     * Route id, or route name.
     */
    routeId?: pulumi.Input<string>;
    /**
     * Tag description list.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a CngwRouteRateLimit resource.
 */
export interface CngwRouteRateLimitArgs {
    /**
     * gateway ID.
     */
    gatewayId: pulumi.Input<string>;
    /**
     * rate limit configuration.
     */
    limitDetail: pulumi.Input<inputs.Tse.CngwRouteRateLimitLimitDetail>;
    /**
     * Route id, or route name.
     */
    routeId: pulumi.Input<string>;
    /**
     * Tag description list.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}