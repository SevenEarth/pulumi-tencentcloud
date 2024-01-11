// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a cfw syncRoute
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const example = new tencentcloud.Cfw.SyncRoute("example", {
 *     fwType: "nat",
 *     syncType: "Route",
 * });
 * ```
 */
export class SyncRoute extends pulumi.CustomResource {
    /**
     * Get an existing SyncRoute resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SyncRouteState, opts?: pulumi.CustomResourceOptions): SyncRoute {
        return new SyncRoute(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Cfw/syncRoute:SyncRoute';

    /**
     * Returns true if the given object is an instance of SyncRoute.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SyncRoute {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SyncRoute.__pulumiType;
    }

    /**
     * Firewall type; nat: nat firewall; ew: inter-vpc firewall.
     */
    public readonly fwType!: pulumi.Output<string | undefined>;
    /**
     * Synchronization operation type: Route, synchronize firewall routing.
     */
    public readonly syncType!: pulumi.Output<string>;

    /**
     * Create a SyncRoute resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SyncRouteArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SyncRouteArgs | SyncRouteState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SyncRouteState | undefined;
            resourceInputs["fwType"] = state ? state.fwType : undefined;
            resourceInputs["syncType"] = state ? state.syncType : undefined;
        } else {
            const args = argsOrState as SyncRouteArgs | undefined;
            if ((!args || args.syncType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'syncType'");
            }
            resourceInputs["fwType"] = args ? args.fwType : undefined;
            resourceInputs["syncType"] = args ? args.syncType : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SyncRoute.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SyncRoute resources.
 */
export interface SyncRouteState {
    /**
     * Firewall type; nat: nat firewall; ew: inter-vpc firewall.
     */
    fwType?: pulumi.Input<string>;
    /**
     * Synchronization operation type: Route, synchronize firewall routing.
     */
    syncType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SyncRoute resource.
 */
export interface SyncRouteArgs {
    /**
     * Firewall type; nat: nat firewall; ew: inter-vpc firewall.
     */
    fwType?: pulumi.Input<string>;
    /**
     * Synchronization operation type: Route, synchronize firewall routing.
     */
    syncType: pulumi.Input<string>;
}