// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a monitor policySetDefault
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const policySetDefault = new tencentcloud.monitor.AlarmPolicySetDefault("policySetDefault", {
 *     module: "monitor",
 *     policyId: "policy-u4iykjkt",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class AlarmPolicySetDefault extends pulumi.CustomResource {
    /**
     * Get an existing AlarmPolicySetDefault resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AlarmPolicySetDefaultState, opts?: pulumi.CustomResourceOptions): AlarmPolicySetDefault {
        return new AlarmPolicySetDefault(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Monitor/alarmPolicySetDefault:AlarmPolicySetDefault';

    /**
     * Returns true if the given object is an instance of AlarmPolicySetDefault.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AlarmPolicySetDefault {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AlarmPolicySetDefault.__pulumiType;
    }

    /**
     * Fixed value, as `monitor`.
     */
    public readonly module!: pulumi.Output<string>;
    /**
     * Policy id.
     */
    public readonly policyId!: pulumi.Output<string>;

    /**
     * Create a AlarmPolicySetDefault resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AlarmPolicySetDefaultArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AlarmPolicySetDefaultArgs | AlarmPolicySetDefaultState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AlarmPolicySetDefaultState | undefined;
            resourceInputs["module"] = state ? state.module : undefined;
            resourceInputs["policyId"] = state ? state.policyId : undefined;
        } else {
            const args = argsOrState as AlarmPolicySetDefaultArgs | undefined;
            if ((!args || args.module === undefined) && !opts.urn) {
                throw new Error("Missing required property 'module'");
            }
            if ((!args || args.policyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policyId'");
            }
            resourceInputs["module"] = args ? args.module : undefined;
            resourceInputs["policyId"] = args ? args.policyId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AlarmPolicySetDefault.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AlarmPolicySetDefault resources.
 */
export interface AlarmPolicySetDefaultState {
    /**
     * Fixed value, as `monitor`.
     */
    module?: pulumi.Input<string>;
    /**
     * Policy id.
     */
    policyId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AlarmPolicySetDefault resource.
 */
export interface AlarmPolicySetDefaultArgs {
    /**
     * Fixed value, as `monitor`.
     */
    module: pulumi.Input<string>;
    /**
     * Policy id.
     */
    policyId: pulumi.Input<string>;
}
