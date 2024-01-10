// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a dasb deviceGroup
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const example = new tencentcloud.Dasb.DeviceGroup("example", {
 *     departmentId: "1.2",
 * });
 * ```
 *
 * ## Import
 *
 * dasb device_group can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Dasb/deviceGroup:DeviceGroup example 36
 * ```
 */
export class DeviceGroup extends pulumi.CustomResource {
    /**
     * Get an existing DeviceGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DeviceGroupState, opts?: pulumi.CustomResourceOptions): DeviceGroup {
        return new DeviceGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Dasb/deviceGroup:DeviceGroup';

    /**
     * Returns true if the given object is an instance of DeviceGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DeviceGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DeviceGroup.__pulumiType;
    }

    /**
     * The ID of the department to which the asset group belongs, such as: 1.2.3 name, with a maximum length of 32 characters.
     */
    public readonly departmentId!: pulumi.Output<string | undefined>;
    /**
     * Device group name, the maximum length is 32 characters.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a DeviceGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DeviceGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DeviceGroupArgs | DeviceGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DeviceGroupState | undefined;
            resourceInputs["departmentId"] = state ? state.departmentId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as DeviceGroupArgs | undefined;
            resourceInputs["departmentId"] = args ? args.departmentId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DeviceGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DeviceGroup resources.
 */
export interface DeviceGroupState {
    /**
     * The ID of the department to which the asset group belongs, such as: 1.2.3 name, with a maximum length of 32 characters.
     */
    departmentId?: pulumi.Input<string>;
    /**
     * Device group name, the maximum length is 32 characters.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DeviceGroup resource.
 */
export interface DeviceGroupArgs {
    /**
     * The ID of the department to which the asset group belongs, such as: 1.2.3 name, with a maximum length of 32 characters.
     */
    departmentId?: pulumi.Input<string>;
    /**
     * Device group name, the maximum length is 32 characters.
     */
    name?: pulumi.Input<string>;
}
