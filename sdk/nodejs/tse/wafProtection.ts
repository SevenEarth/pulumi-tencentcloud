// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a tse wafProtection
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const wafProtection = new tencentcloud.tse.WafProtection("wafProtection", {
 *     gatewayId: "gateway-ed63e957",
 *     lists: ["7324a769-9d87-48ce-a904-48c3defc4abd"],
 *     operate: "open",
 *     type: "Route",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class WafProtection extends pulumi.CustomResource {
    /**
     * Get an existing WafProtection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WafProtectionState, opts?: pulumi.CustomResourceOptions): WafProtection {
        return new WafProtection(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Tse/wafProtection:WafProtection';

    /**
     * Returns true if the given object is an instance of WafProtection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WafProtection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WafProtection.__pulumiType;
    }

    /**
     * Gateway ID.
     */
    public readonly gatewayId!: pulumi.Output<string>;
    /**
     * Global protection status.
     */
    public /*out*/ readonly globalStatus!: pulumi.Output<string>;
    /**
     * Means the list of services or routes when the resource type `Type` is `Service` or `Route`.
     */
    public readonly lists!: pulumi.Output<string[] | undefined>;
    /**
     * `open`: open the protection, `close`: close the protection.
     */
    public readonly operate!: pulumi.Output<string>;
    /**
     * The type of protection resource. Reference value: `Global`: instance, `Service`: service, `Route`: route, `Object`: obejct (This interface does not currently support this type).
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a WafProtection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WafProtectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WafProtectionArgs | WafProtectionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WafProtectionState | undefined;
            resourceInputs["gatewayId"] = state ? state.gatewayId : undefined;
            resourceInputs["globalStatus"] = state ? state.globalStatus : undefined;
            resourceInputs["lists"] = state ? state.lists : undefined;
            resourceInputs["operate"] = state ? state.operate : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as WafProtectionArgs | undefined;
            if ((!args || args.gatewayId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'gatewayId'");
            }
            if ((!args || args.operate === undefined) && !opts.urn) {
                throw new Error("Missing required property 'operate'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["gatewayId"] = args ? args.gatewayId : undefined;
            resourceInputs["lists"] = args ? args.lists : undefined;
            resourceInputs["operate"] = args ? args.operate : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["globalStatus"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WafProtection.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WafProtection resources.
 */
export interface WafProtectionState {
    /**
     * Gateway ID.
     */
    gatewayId?: pulumi.Input<string>;
    /**
     * Global protection status.
     */
    globalStatus?: pulumi.Input<string>;
    /**
     * Means the list of services or routes when the resource type `Type` is `Service` or `Route`.
     */
    lists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * `open`: open the protection, `close`: close the protection.
     */
    operate?: pulumi.Input<string>;
    /**
     * The type of protection resource. Reference value: `Global`: instance, `Service`: service, `Route`: route, `Object`: obejct (This interface does not currently support this type).
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WafProtection resource.
 */
export interface WafProtectionArgs {
    /**
     * Gateway ID.
     */
    gatewayId: pulumi.Input<string>;
    /**
     * Means the list of services or routes when the resource type `Type` is `Service` or `Route`.
     */
    lists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * `open`: open the protection, `close`: close the protection.
     */
    operate: pulumi.Input<string>;
    /**
     * The type of protection resource. Reference value: `Global`: instance, `Service`: service, `Route`: route, `Object`: obejct (This interface does not currently support this type).
     */
    type: pulumi.Input<string>;
}
