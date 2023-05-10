// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a ci bucket
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const bucketAttachment = new tencentcloud.Ci.BucketAttachment("bucket_attachment", {
 *     bucket: "terraform-ci-xxxxxx",
 * });
 * ```
 *
 * ## Import
 *
 * ci bucket can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Ci/bucketAttachment:BucketAttachment bucket_attachment terraform-ci-xxxxxx
 * ```
 */
export class BucketAttachment extends pulumi.CustomResource {
    /**
     * Get an existing BucketAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BucketAttachmentState, opts?: pulumi.CustomResourceOptions): BucketAttachment {
        return new BucketAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Ci/bucketAttachment:BucketAttachment';

    /**
     * Returns true if the given object is an instance of BucketAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BucketAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BucketAttachment.__pulumiType;
    }

    /**
     * bucket name.
     */
    public readonly bucket!: pulumi.Output<string>;
    /**
     * Binding object storage state, `on`: bound, `off`: unbound, `unbinding`: unbinding.
     */
    public /*out*/ readonly ciStatus!: pulumi.Output<string>;

    /**
     * Create a BucketAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BucketAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BucketAttachmentArgs | BucketAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BucketAttachmentState | undefined;
            resourceInputs["bucket"] = state ? state.bucket : undefined;
            resourceInputs["ciStatus"] = state ? state.ciStatus : undefined;
        } else {
            const args = argsOrState as BucketAttachmentArgs | undefined;
            if ((!args || args.bucket === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bucket'");
            }
            resourceInputs["bucket"] = args ? args.bucket : undefined;
            resourceInputs["ciStatus"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BucketAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BucketAttachment resources.
 */
export interface BucketAttachmentState {
    /**
     * bucket name.
     */
    bucket?: pulumi.Input<string>;
    /**
     * Binding object storage state, `on`: bound, `off`: unbound, `unbinding`: unbinding.
     */
    ciStatus?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BucketAttachment resource.
 */
export interface BucketAttachmentArgs {
    /**
     * bucket name.
     */
    bucket: pulumi.Input<string>;
}
