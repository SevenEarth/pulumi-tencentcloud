// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a ci mediaSmartCoverTemplate
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const mediaSmartCoverTemplate = new tencentcloud.Ci.MediaSmartCoverTemplate("media_smart_cover_template", {
 *     bucket: "terraform-ci-xxxxxx",
 *     smartCover: {
 *         count: "10",
 *         deleteDuplicates: "true",
 *         format: "jpg",
 *         height: "960",
 *         width: "1280",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * ci media_smart_cover_template can be imported using the bucket#templateId, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Ci/mediaSmartCoverTemplate:MediaSmartCoverTemplate media_smart_cover_template terraform-ci-xxxxxx#t1ede83acc305e423799d638044d859fb7
 * ```
 */
export class MediaSmartCoverTemplate extends pulumi.CustomResource {
    /**
     * Get an existing MediaSmartCoverTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MediaSmartCoverTemplateState, opts?: pulumi.CustomResourceOptions): MediaSmartCoverTemplate {
        return new MediaSmartCoverTemplate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Ci/mediaSmartCoverTemplate:MediaSmartCoverTemplate';

    /**
     * Returns true if the given object is an instance of MediaSmartCoverTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MediaSmartCoverTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MediaSmartCoverTemplate.__pulumiType;
    }

    /**
     * bucket name.
     */
    public readonly bucket!: pulumi.Output<string>;
    /**
     * The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Smart Cover Parameters.
     */
    public readonly smartCover!: pulumi.Output<outputs.Ci.MediaSmartCoverTemplateSmartCover>;

    /**
     * Create a MediaSmartCoverTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MediaSmartCoverTemplateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MediaSmartCoverTemplateArgs | MediaSmartCoverTemplateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MediaSmartCoverTemplateState | undefined;
            resourceInputs["bucket"] = state ? state.bucket : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["smartCover"] = state ? state.smartCover : undefined;
        } else {
            const args = argsOrState as MediaSmartCoverTemplateArgs | undefined;
            if ((!args || args.bucket === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bucket'");
            }
            if ((!args || args.smartCover === undefined) && !opts.urn) {
                throw new Error("Missing required property 'smartCover'");
            }
            resourceInputs["bucket"] = args ? args.bucket : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["smartCover"] = args ? args.smartCover : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(MediaSmartCoverTemplate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MediaSmartCoverTemplate resources.
 */
export interface MediaSmartCoverTemplateState {
    /**
     * bucket name.
     */
    bucket?: pulumi.Input<string>;
    /**
     * The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
     */
    name?: pulumi.Input<string>;
    /**
     * Smart Cover Parameters.
     */
    smartCover?: pulumi.Input<inputs.Ci.MediaSmartCoverTemplateSmartCover>;
}

/**
 * The set of arguments for constructing a MediaSmartCoverTemplate resource.
 */
export interface MediaSmartCoverTemplateArgs {
    /**
     * bucket name.
     */
    bucket: pulumi.Input<string>;
    /**
     * The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
     */
    name?: pulumi.Input<string>;
    /**
     * Smart Cover Parameters.
     */
    smartCover: pulumi.Input<inputs.Ci.MediaSmartCoverTemplateSmartCover>;
}
