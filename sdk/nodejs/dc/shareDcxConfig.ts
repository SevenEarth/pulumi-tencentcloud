// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a dc shareDcxConfig
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const shareDcxConfig = new tencentcloud.Dc.ShareDcxConfig("share_dcx_config", {
 *     directConnectTunnelId: "dcx-4z49tnws",
 *     enable: false,
 * });
 * ```
 *
 * ## Import
 *
 * dc share_dcx_config can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Dc/shareDcxConfig:ShareDcxConfig share_dcx_config dcx_id
 * ```
 */
export class ShareDcxConfig extends pulumi.CustomResource {
    /**
     * Get an existing ShareDcxConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ShareDcxConfigState, opts?: pulumi.CustomResourceOptions): ShareDcxConfig {
        return new ShareDcxConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Dc/shareDcxConfig:ShareDcxConfig';

    /**
     * Returns true if the given object is an instance of ShareDcxConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ShareDcxConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ShareDcxConfig.__pulumiType;
    }

    /**
     * the direct connect owner accept or reject the apply of direct connect tunnel.
     */
    public readonly directConnectTunnelId!: pulumi.Output<string>;
    /**
     * if accept or reject direct connect tunnel.
     */
    public readonly enable!: pulumi.Output<boolean>;

    /**
     * Create a ShareDcxConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ShareDcxConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ShareDcxConfigArgs | ShareDcxConfigState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ShareDcxConfigState | undefined;
            resourceInputs["directConnectTunnelId"] = state ? state.directConnectTunnelId : undefined;
            resourceInputs["enable"] = state ? state.enable : undefined;
        } else {
            const args = argsOrState as ShareDcxConfigArgs | undefined;
            if ((!args || args.directConnectTunnelId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'directConnectTunnelId'");
            }
            if ((!args || args.enable === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enable'");
            }
            resourceInputs["directConnectTunnelId"] = args ? args.directConnectTunnelId : undefined;
            resourceInputs["enable"] = args ? args.enable : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ShareDcxConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ShareDcxConfig resources.
 */
export interface ShareDcxConfigState {
    /**
     * the direct connect owner accept or reject the apply of direct connect tunnel.
     */
    directConnectTunnelId?: pulumi.Input<string>;
    /**
     * if accept or reject direct connect tunnel.
     */
    enable?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a ShareDcxConfig resource.
 */
export interface ShareDcxConfigArgs {
    /**
     * the direct connect owner accept or reject the apply of direct connect tunnel.
     */
    directConnectTunnelId: pulumi.Input<string>;
    /**
     * if accept or reject direct connect tunnel.
     */
    enable: pulumi.Input<boolean>;
}