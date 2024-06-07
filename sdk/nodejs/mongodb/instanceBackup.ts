// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a mongodb instanceBackup
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const instanceBackup = new tencentcloud.mongodb.InstanceBackup("instanceBackup", {
 *     backupMethod: 0,
 *     backupRemark: "my backup",
 *     instanceId: "cmgo-9d0p6umb",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class InstanceBackup extends pulumi.CustomResource {
    /**
     * Get an existing InstanceBackup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceBackupState, opts?: pulumi.CustomResourceOptions): InstanceBackup {
        return new InstanceBackup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Mongodb/instanceBackup:InstanceBackup';

    /**
     * Returns true if the given object is an instance of InstanceBackup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InstanceBackup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InstanceBackup.__pulumiType;
    }

    /**
     * 0:logical backup, 1:physical backup.
     */
    public readonly backupMethod!: pulumi.Output<number>;
    /**
     * backup notes.
     */
    public readonly backupRemark!: pulumi.Output<string | undefined>;
    /**
     * Instance ID, the format is: cmgo-9d0p6umb.Same as the instance ID displayed in the cloud database console page.
     */
    public readonly instanceId!: pulumi.Output<string>;

    /**
     * Create a InstanceBackup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceBackupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceBackupArgs | InstanceBackupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceBackupState | undefined;
            resourceInputs["backupMethod"] = state ? state.backupMethod : undefined;
            resourceInputs["backupRemark"] = state ? state.backupRemark : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
        } else {
            const args = argsOrState as InstanceBackupArgs | undefined;
            if ((!args || args.backupMethod === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backupMethod'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["backupMethod"] = args ? args.backupMethod : undefined;
            resourceInputs["backupRemark"] = args ? args.backupRemark : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(InstanceBackup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InstanceBackup resources.
 */
export interface InstanceBackupState {
    /**
     * 0:logical backup, 1:physical backup.
     */
    backupMethod?: pulumi.Input<number>;
    /**
     * backup notes.
     */
    backupRemark?: pulumi.Input<string>;
    /**
     * Instance ID, the format is: cmgo-9d0p6umb.Same as the instance ID displayed in the cloud database console page.
     */
    instanceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InstanceBackup resource.
 */
export interface InstanceBackupArgs {
    /**
     * 0:logical backup, 1:physical backup.
     */
    backupMethod: pulumi.Input<number>;
    /**
     * backup notes.
     */
    backupRemark?: pulumi.Input<string>;
    /**
     * Instance ID, the format is: cmgo-9d0p6umb.Same as the instance ID displayed in the cloud database console page.
     */
    instanceId: pulumi.Input<string>;
}
