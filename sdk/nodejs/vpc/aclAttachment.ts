// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provide a resource to attach an existing subnet to Network ACL.
 *
 * ## Import
 *
 * Acl attachment can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import tencentcloud:Vpc/aclAttachment:AclAttachment attachment acl-eotx5qsg#subnet-91x0geu6
 * ```
 */
export class AclAttachment extends pulumi.CustomResource {
    /**
     * Get an existing AclAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AclAttachmentState, opts?: pulumi.CustomResourceOptions): AclAttachment {
        return new AclAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Vpc/aclAttachment:AclAttachment';

    /**
     * Returns true if the given object is an instance of AclAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AclAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AclAttachment.__pulumiType;
    }

    /**
     * ID of the attached ACL.
     */
    public readonly aclId!: pulumi.Output<string>;
    /**
     * The Subnet instance ID.
     */
    public readonly subnetId!: pulumi.Output<string>;

    /**
     * Create a AclAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AclAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AclAttachmentArgs | AclAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AclAttachmentState | undefined;
            resourceInputs["aclId"] = state ? state.aclId : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
        } else {
            const args = argsOrState as AclAttachmentArgs | undefined;
            if ((!args || args.aclId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'aclId'");
            }
            if ((!args || args.subnetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetId'");
            }
            resourceInputs["aclId"] = args ? args.aclId : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AclAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AclAttachment resources.
 */
export interface AclAttachmentState {
    /**
     * ID of the attached ACL.
     */
    aclId?: pulumi.Input<string>;
    /**
     * The Subnet instance ID.
     */
    subnetId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AclAttachment resource.
 */
export interface AclAttachmentArgs {
    /**
     * ID of the attached ACL.
     */
    aclId: pulumi.Input<string>;
    /**
     * The Subnet instance ID.
     */
    subnetId: pulumi.Input<string>;
}
