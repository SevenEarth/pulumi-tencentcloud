// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource for an AS (Auto scaling) notification.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const asNotification = new tencentcloud.As.Notification("as_notification", {
 *     notificationTypes: [
 *         "SCALE_OUT_FAILED",
 *         "SCALE_IN_SUCCESSFUL",
 *         "SCALE_IN_FAILED",
 *         "REPLACE_UNHEALTHY_INSTANCE_FAILED",
 *     ],
 *     notificationUserGroupIds: ["76955"],
 *     scalingGroupId: "sg-12af45",
 * });
 * ```
 */
export class Notification extends pulumi.CustomResource {
    /**
     * Get an existing Notification resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NotificationState, opts?: pulumi.CustomResourceOptions): Notification {
        return new Notification(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:As/notification:Notification';

    /**
     * Returns true if the given object is an instance of Notification.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Notification {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Notification.__pulumiType;
    }

    /**
     * A list of Notification Types that trigger notifications. Acceptable values are `SCALE_OUT_FAILED`, `SCALE_IN_SUCCESSFUL`, `SCALE_IN_FAILED`, `REPLACE_UNHEALTHY_INSTANCE_SUCCESSFUL` and `REPLACE_UNHEALTHY_INSTANCE_FAILED`.
     */
    public readonly notificationTypes!: pulumi.Output<string[]>;
    /**
     * A group of user IDs to be notified.
     */
    public readonly notificationUserGroupIds!: pulumi.Output<string[]>;
    /**
     * ID of a scaling group.
     */
    public readonly scalingGroupId!: pulumi.Output<string>;

    /**
     * Create a Notification resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NotificationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NotificationArgs | NotificationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NotificationState | undefined;
            resourceInputs["notificationTypes"] = state ? state.notificationTypes : undefined;
            resourceInputs["notificationUserGroupIds"] = state ? state.notificationUserGroupIds : undefined;
            resourceInputs["scalingGroupId"] = state ? state.scalingGroupId : undefined;
        } else {
            const args = argsOrState as NotificationArgs | undefined;
            if ((!args || args.notificationTypes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'notificationTypes'");
            }
            if ((!args || args.notificationUserGroupIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'notificationUserGroupIds'");
            }
            if ((!args || args.scalingGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scalingGroupId'");
            }
            resourceInputs["notificationTypes"] = args ? args.notificationTypes : undefined;
            resourceInputs["notificationUserGroupIds"] = args ? args.notificationUserGroupIds : undefined;
            resourceInputs["scalingGroupId"] = args ? args.scalingGroupId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Notification.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Notification resources.
 */
export interface NotificationState {
    /**
     * A list of Notification Types that trigger notifications. Acceptable values are `SCALE_OUT_FAILED`, `SCALE_IN_SUCCESSFUL`, `SCALE_IN_FAILED`, `REPLACE_UNHEALTHY_INSTANCE_SUCCESSFUL` and `REPLACE_UNHEALTHY_INSTANCE_FAILED`.
     */
    notificationTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A group of user IDs to be notified.
     */
    notificationUserGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ID of a scaling group.
     */
    scalingGroupId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Notification resource.
 */
export interface NotificationArgs {
    /**
     * A list of Notification Types that trigger notifications. Acceptable values are `SCALE_OUT_FAILED`, `SCALE_IN_SUCCESSFUL`, `SCALE_IN_FAILED`, `REPLACE_UNHEALTHY_INSTANCE_SUCCESSFUL` and `REPLACE_UNHEALTHY_INSTANCE_FAILED`.
     */
    notificationTypes: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A group of user IDs to be notified.
     */
    notificationUserGroupIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ID of a scaling group.
     */
    scalingGroupId: pulumi.Input<string>;
}
