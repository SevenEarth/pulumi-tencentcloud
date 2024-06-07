// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a CBS snapshot policy attachment resource.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const foo = new tencentcloud.cbs.SnapshotPolicyAttachment("foo", {
 *     storageId: tencentcloud_cbs_storage.foo.id,
 *     snapshotPolicyId: tencentcloud_cbs_snapshot_policy.policy.id,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class SnapshotPolicyAttachment extends pulumi.CustomResource {
    /**
     * Get an existing SnapshotPolicyAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SnapshotPolicyAttachmentState, opts?: pulumi.CustomResourceOptions): SnapshotPolicyAttachment {
        return new SnapshotPolicyAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Cbs/snapshotPolicyAttachment:SnapshotPolicyAttachment';

    /**
     * Returns true if the given object is an instance of SnapshotPolicyAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SnapshotPolicyAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SnapshotPolicyAttachment.__pulumiType;
    }

    /**
     * ID of CBS snapshot policy.
     */
    public readonly snapshotPolicyId!: pulumi.Output<string>;
    /**
     * ID of CBS.
     */
    public readonly storageId!: pulumi.Output<string>;

    /**
     * Create a SnapshotPolicyAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SnapshotPolicyAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SnapshotPolicyAttachmentArgs | SnapshotPolicyAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SnapshotPolicyAttachmentState | undefined;
            resourceInputs["snapshotPolicyId"] = state ? state.snapshotPolicyId : undefined;
            resourceInputs["storageId"] = state ? state.storageId : undefined;
        } else {
            const args = argsOrState as SnapshotPolicyAttachmentArgs | undefined;
            if ((!args || args.snapshotPolicyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'snapshotPolicyId'");
            }
            if ((!args || args.storageId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storageId'");
            }
            resourceInputs["snapshotPolicyId"] = args ? args.snapshotPolicyId : undefined;
            resourceInputs["storageId"] = args ? args.storageId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SnapshotPolicyAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SnapshotPolicyAttachment resources.
 */
export interface SnapshotPolicyAttachmentState {
    /**
     * ID of CBS snapshot policy.
     */
    snapshotPolicyId?: pulumi.Input<string>;
    /**
     * ID of CBS.
     */
    storageId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SnapshotPolicyAttachment resource.
 */
export interface SnapshotPolicyAttachmentArgs {
    /**
     * ID of CBS snapshot policy.
     */
    snapshotPolicyId: pulumi.Input<string>;
    /**
     * ID of CBS.
     */
    storageId: pulumi.Input<string>;
}
