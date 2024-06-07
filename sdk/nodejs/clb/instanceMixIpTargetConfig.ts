// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a clb instanceMixIpTargetConfig
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const instanceMixIpTargetConfig = new tencentcloud.clb.InstanceMixIpTargetConfig("instanceMixIpTargetConfig", {
 *     loadBalancerId: "lb-5dnrkgry",
 *     mixIpTarget: false,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * clb instance_mix_ip_target_config can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import tencentcloud:Clb/instanceMixIpTargetConfig:InstanceMixIpTargetConfig instance_mix_ip_target_config instance_id
 * ```
 */
export class InstanceMixIpTargetConfig extends pulumi.CustomResource {
    /**
     * Get an existing InstanceMixIpTargetConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceMixIpTargetConfigState, opts?: pulumi.CustomResourceOptions): InstanceMixIpTargetConfig {
        return new InstanceMixIpTargetConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Clb/instanceMixIpTargetConfig:InstanceMixIpTargetConfig';

    /**
     * Returns true if the given object is an instance of InstanceMixIpTargetConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InstanceMixIpTargetConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InstanceMixIpTargetConfig.__pulumiType;
    }

    /**
     * ID of CLB instances to be queried.
     */
    public readonly loadBalancerId!: pulumi.Output<string>;
    /**
     * False: closed True:open.
     */
    public readonly mixIpTarget!: pulumi.Output<boolean>;

    /**
     * Create a InstanceMixIpTargetConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceMixIpTargetConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceMixIpTargetConfigArgs | InstanceMixIpTargetConfigState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceMixIpTargetConfigState | undefined;
            resourceInputs["loadBalancerId"] = state ? state.loadBalancerId : undefined;
            resourceInputs["mixIpTarget"] = state ? state.mixIpTarget : undefined;
        } else {
            const args = argsOrState as InstanceMixIpTargetConfigArgs | undefined;
            if ((!args || args.loadBalancerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'loadBalancerId'");
            }
            if ((!args || args.mixIpTarget === undefined) && !opts.urn) {
                throw new Error("Missing required property 'mixIpTarget'");
            }
            resourceInputs["loadBalancerId"] = args ? args.loadBalancerId : undefined;
            resourceInputs["mixIpTarget"] = args ? args.mixIpTarget : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(InstanceMixIpTargetConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InstanceMixIpTargetConfig resources.
 */
export interface InstanceMixIpTargetConfigState {
    /**
     * ID of CLB instances to be queried.
     */
    loadBalancerId?: pulumi.Input<string>;
    /**
     * False: closed True:open.
     */
    mixIpTarget?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a InstanceMixIpTargetConfig resource.
 */
export interface InstanceMixIpTargetConfigArgs {
    /**
     * ID of CLB instances to be queried.
     */
    loadBalancerId: pulumi.Input<string>;
    /**
     * False: closed True:open.
     */
    mixIpTarget: pulumi.Input<boolean>;
}
