// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to abort multipart upload
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const abortMultipartUpload = new tencentcloud.cos.ObjectAbortMultipartUploadOperation("abortMultipartUpload", {
 *     bucket: "keep-test-xxxxxx",
 *     key: "object",
 *     uploadId: "xxxxxx",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class ObjectAbortMultipartUploadOperation extends pulumi.CustomResource {
    /**
     * Get an existing ObjectAbortMultipartUploadOperation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ObjectAbortMultipartUploadOperationState, opts?: pulumi.CustomResourceOptions): ObjectAbortMultipartUploadOperation {
        return new ObjectAbortMultipartUploadOperation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Cos/objectAbortMultipartUploadOperation:ObjectAbortMultipartUploadOperation';

    /**
     * Returns true if the given object is an instance of ObjectAbortMultipartUploadOperation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ObjectAbortMultipartUploadOperation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ObjectAbortMultipartUploadOperation.__pulumiType;
    }

    /**
     * Bucket.
     */
    public readonly bucket!: pulumi.Output<string>;
    /**
     * Object key.
     */
    public readonly key!: pulumi.Output<string>;
    /**
     * Multipart uploaded id.
     */
    public readonly uploadId!: pulumi.Output<string>;

    /**
     * Create a ObjectAbortMultipartUploadOperation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ObjectAbortMultipartUploadOperationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ObjectAbortMultipartUploadOperationArgs | ObjectAbortMultipartUploadOperationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ObjectAbortMultipartUploadOperationState | undefined;
            resourceInputs["bucket"] = state ? state.bucket : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["uploadId"] = state ? state.uploadId : undefined;
        } else {
            const args = argsOrState as ObjectAbortMultipartUploadOperationArgs | undefined;
            if ((!args || args.bucket === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bucket'");
            }
            if ((!args || args.key === undefined) && !opts.urn) {
                throw new Error("Missing required property 'key'");
            }
            if ((!args || args.uploadId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'uploadId'");
            }
            resourceInputs["bucket"] = args ? args.bucket : undefined;
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["uploadId"] = args ? args.uploadId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ObjectAbortMultipartUploadOperation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ObjectAbortMultipartUploadOperation resources.
 */
export interface ObjectAbortMultipartUploadOperationState {
    /**
     * Bucket.
     */
    bucket?: pulumi.Input<string>;
    /**
     * Object key.
     */
    key?: pulumi.Input<string>;
    /**
     * Multipart uploaded id.
     */
    uploadId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ObjectAbortMultipartUploadOperation resource.
 */
export interface ObjectAbortMultipartUploadOperationArgs {
    /**
     * Bucket.
     */
    bucket: pulumi.Input<string>;
    /**
     * Object key.
     */
    key: pulumi.Input<string>;
    /**
     * Multipart uploaded id.
     */
    uploadId: pulumi.Input<string>;
}
