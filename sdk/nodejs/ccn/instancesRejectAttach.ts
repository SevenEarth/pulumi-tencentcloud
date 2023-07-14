// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a vpc ccn_instances_reject_attach, you can use this resource to approve cross-region attachment.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const ccnInstancesRejectAttach = new tencentcloud.Ccn.InstancesRejectAttach("ccn_instances_reject_attach", {
 *     ccnId: "ccn-39lqkygf",
 *     instances: [{
 *         instanceId: "vpc-j9yhbzpn",
 *         instanceRegion: "ap-guangzhou",
 *         instanceType: "VPC",
 *     }],
 * });
 * ```
 */
export class InstancesRejectAttach extends pulumi.CustomResource {
    /**
     * Get an existing InstancesRejectAttach resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstancesRejectAttachState, opts?: pulumi.CustomResourceOptions): InstancesRejectAttach {
        return new InstancesRejectAttach(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Ccn/instancesRejectAttach:InstancesRejectAttach';

    /**
     * Returns true if the given object is an instance of InstancesRejectAttach.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InstancesRejectAttach {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InstancesRejectAttach.__pulumiType;
    }

    /**
     * CCN Instance ID.
     */
    public readonly ccnId!: pulumi.Output<string>;
    /**
     * Reject List Of Attachment Instances.
     */
    public readonly instances!: pulumi.Output<outputs.Ccn.InstancesRejectAttachInstance[]>;

    /**
     * Create a InstancesRejectAttach resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstancesRejectAttachArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstancesRejectAttachArgs | InstancesRejectAttachState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstancesRejectAttachState | undefined;
            resourceInputs["ccnId"] = state ? state.ccnId : undefined;
            resourceInputs["instances"] = state ? state.instances : undefined;
        } else {
            const args = argsOrState as InstancesRejectAttachArgs | undefined;
            if ((!args || args.ccnId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ccnId'");
            }
            if ((!args || args.instances === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instances'");
            }
            resourceInputs["ccnId"] = args ? args.ccnId : undefined;
            resourceInputs["instances"] = args ? args.instances : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(InstancesRejectAttach.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InstancesRejectAttach resources.
 */
export interface InstancesRejectAttachState {
    /**
     * CCN Instance ID.
     */
    ccnId?: pulumi.Input<string>;
    /**
     * Reject List Of Attachment Instances.
     */
    instances?: pulumi.Input<pulumi.Input<inputs.Ccn.InstancesRejectAttachInstance>[]>;
}

/**
 * The set of arguments for constructing a InstancesRejectAttach resource.
 */
export interface InstancesRejectAttachArgs {
    /**
     * CCN Instance ID.
     */
    ccnId: pulumi.Input<string>;
    /**
     * Reject List Of Attachment Instances.
     */
    instances: pulumi.Input<pulumi.Input<inputs.Ccn.InstancesRejectAttachInstance>[]>;
}
