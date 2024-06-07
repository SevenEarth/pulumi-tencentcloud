// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a cvm imageSharePermission
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const imageSharePermission = new tencentcloud.cvm.ImageSharePermission("imageSharePermission", {
 *     accountIds: ["xxxxxx"],
 *     imageId: "img-xxxxxx",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * cvm image_share_permission can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import tencentcloud:Cvm/imageSharePermission:ImageSharePermission image_share_permission image_share_permission_id
 * ```
 */
export class ImageSharePermission extends pulumi.CustomResource {
    /**
     * Get an existing ImageSharePermission resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ImageSharePermissionState, opts?: pulumi.CustomResourceOptions): ImageSharePermission {
        return new ImageSharePermission(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Cvm/imageSharePermission:ImageSharePermission';

    /**
     * Returns true if the given object is an instance of ImageSharePermission.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ImageSharePermission {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ImageSharePermission.__pulumiType;
    }

    /**
     * List of account IDs with which an image is shared.
     */
    public readonly accountIds!: pulumi.Output<string[]>;
    /**
     * Image ID such as `img-gvbnzy6f`. You can only specify an image in the NORMAL state.
     */
    public readonly imageId!: pulumi.Output<string>;

    /**
     * Create a ImageSharePermission resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ImageSharePermissionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ImageSharePermissionArgs | ImageSharePermissionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ImageSharePermissionState | undefined;
            resourceInputs["accountIds"] = state ? state.accountIds : undefined;
            resourceInputs["imageId"] = state ? state.imageId : undefined;
        } else {
            const args = argsOrState as ImageSharePermissionArgs | undefined;
            if ((!args || args.accountIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountIds'");
            }
            if ((!args || args.imageId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'imageId'");
            }
            resourceInputs["accountIds"] = args ? args.accountIds : undefined;
            resourceInputs["imageId"] = args ? args.imageId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ImageSharePermission.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ImageSharePermission resources.
 */
export interface ImageSharePermissionState {
    /**
     * List of account IDs with which an image is shared.
     */
    accountIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Image ID such as `img-gvbnzy6f`. You can only specify an image in the NORMAL state.
     */
    imageId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ImageSharePermission resource.
 */
export interface ImageSharePermissionArgs {
    /**
     * List of account IDs with which an image is shared.
     */
    accountIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Image ID such as `img-gvbnzy6f`. You can only specify an image in the NORMAL state.
     */
    imageId: pulumi.Input<string>;
}
