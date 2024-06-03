// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a ci bucketPicStyle
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const bucketPicStyle = new tencentcloud.ci.BucketPicStyle("bucketPicStyle", {
 *     bucket: "terraform-ci-xxxxxx",
 *     styleBody: "imageMogr2/thumbnail/20x/crop/20x20/gravity/center/interlace/0/quality/100",
 *     styleName: "rayscale_2",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * ci bucket_pic_style can be imported using the bucket#styleName, e.g.
 *
 * ```sh
 * $ pulumi import tencentcloud:Ci/bucketPicStyle:BucketPicStyle bucket_pic_style terraform-ci-xxxxxx#rayscale_2
 * ```
 */
export class BucketPicStyle extends pulumi.CustomResource {
    /**
     * Get an existing BucketPicStyle resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BucketPicStyleState, opts?: pulumi.CustomResourceOptions): BucketPicStyle {
        return new BucketPicStyle(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Ci/bucketPicStyle:BucketPicStyle';

    /**
     * Returns true if the given object is an instance of BucketPicStyle.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BucketPicStyle {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BucketPicStyle.__pulumiType;
    }

    /**
     * bucket name.
     */
    public readonly bucket!: pulumi.Output<string>;
    /**
     * style details, example: mageMogr2/grayscale/1.
     */
    public readonly styleBody!: pulumi.Output<string>;
    /**
     * style name, style names are case-sensitive, and a combination of uppercase and lowercase letters, numbers, and `$ + _ ( )` is supported.
     */
    public readonly styleName!: pulumi.Output<string>;

    /**
     * Create a BucketPicStyle resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BucketPicStyleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BucketPicStyleArgs | BucketPicStyleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BucketPicStyleState | undefined;
            resourceInputs["bucket"] = state ? state.bucket : undefined;
            resourceInputs["styleBody"] = state ? state.styleBody : undefined;
            resourceInputs["styleName"] = state ? state.styleName : undefined;
        } else {
            const args = argsOrState as BucketPicStyleArgs | undefined;
            if ((!args || args.bucket === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bucket'");
            }
            if ((!args || args.styleBody === undefined) && !opts.urn) {
                throw new Error("Missing required property 'styleBody'");
            }
            if ((!args || args.styleName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'styleName'");
            }
            resourceInputs["bucket"] = args ? args.bucket : undefined;
            resourceInputs["styleBody"] = args ? args.styleBody : undefined;
            resourceInputs["styleName"] = args ? args.styleName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BucketPicStyle.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BucketPicStyle resources.
 */
export interface BucketPicStyleState {
    /**
     * bucket name.
     */
    bucket?: pulumi.Input<string>;
    /**
     * style details, example: mageMogr2/grayscale/1.
     */
    styleBody?: pulumi.Input<string>;
    /**
     * style name, style names are case-sensitive, and a combination of uppercase and lowercase letters, numbers, and `$ + _ ( )` is supported.
     */
    styleName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BucketPicStyle resource.
 */
export interface BucketPicStyleArgs {
    /**
     * bucket name.
     */
    bucket: pulumi.Input<string>;
    /**
     * style details, example: mageMogr2/grayscale/1.
     */
    styleBody: pulumi.Input<string>;
    /**
     * style name, style names are case-sensitive, and a combination of uppercase and lowercase letters, numbers, and `$ + _ ( )` is supported.
     */
    styleName: pulumi.Input<string>;
}
