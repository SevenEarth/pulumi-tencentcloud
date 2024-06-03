// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a cvm launchTemplateDefaultVersion
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const launchTemplateDefaultVersion = new tencentcloud.cvm.LaunchTemplateDefaultVersion("launchTemplateDefaultVersion", {
 *     defaultVersion: 2,
 *     launchTemplateId: "lt-34vaef8fe",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * cvm launch_template_default_version can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import tencentcloud:Cvm/launchTemplateDefaultVersion:LaunchTemplateDefaultVersion launch_template_default_version launch_template_id
 * ```
 */
export class LaunchTemplateDefaultVersion extends pulumi.CustomResource {
    /**
     * Get an existing LaunchTemplateDefaultVersion resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LaunchTemplateDefaultVersionState, opts?: pulumi.CustomResourceOptions): LaunchTemplateDefaultVersion {
        return new LaunchTemplateDefaultVersion(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Cvm/launchTemplateDefaultVersion:LaunchTemplateDefaultVersion';

    /**
     * Returns true if the given object is an instance of LaunchTemplateDefaultVersion.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LaunchTemplateDefaultVersion {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LaunchTemplateDefaultVersion.__pulumiType;
    }

    /**
     * The number of the version that you want to set as the default version.
     */
    public readonly defaultVersion!: pulumi.Output<number>;
    /**
     * Instance launch template ID.
     */
    public readonly launchTemplateId!: pulumi.Output<string>;

    /**
     * Create a LaunchTemplateDefaultVersion resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LaunchTemplateDefaultVersionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LaunchTemplateDefaultVersionArgs | LaunchTemplateDefaultVersionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LaunchTemplateDefaultVersionState | undefined;
            resourceInputs["defaultVersion"] = state ? state.defaultVersion : undefined;
            resourceInputs["launchTemplateId"] = state ? state.launchTemplateId : undefined;
        } else {
            const args = argsOrState as LaunchTemplateDefaultVersionArgs | undefined;
            if ((!args || args.defaultVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'defaultVersion'");
            }
            if ((!args || args.launchTemplateId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'launchTemplateId'");
            }
            resourceInputs["defaultVersion"] = args ? args.defaultVersion : undefined;
            resourceInputs["launchTemplateId"] = args ? args.launchTemplateId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LaunchTemplateDefaultVersion.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LaunchTemplateDefaultVersion resources.
 */
export interface LaunchTemplateDefaultVersionState {
    /**
     * The number of the version that you want to set as the default version.
     */
    defaultVersion?: pulumi.Input<number>;
    /**
     * Instance launch template ID.
     */
    launchTemplateId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LaunchTemplateDefaultVersion resource.
 */
export interface LaunchTemplateDefaultVersionArgs {
    /**
     * The number of the version that you want to set as the default version.
     */
    defaultVersion: pulumi.Input<number>;
    /**
     * Instance launch template ID.
     */
    launchTemplateId: pulumi.Input<string>;
}
