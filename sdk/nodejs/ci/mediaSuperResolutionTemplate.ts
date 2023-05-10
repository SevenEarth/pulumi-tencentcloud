// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a ci mediaSuperResolutionTemplate
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const mediaSuperResolutionTemplate = new tencentcloud.Ci.MediaSuperResolutionTemplate("media_super_resolution_template", {
 *     bucket: "terraform-ci-1308919341",
 *     enableScaleUp: "true",
 *     resolution: "sdtohd",
 *     version: "Enhance",
 * });
 * ```
 *
 * ## Import
 *
 * ci media_super_resolution_template can be imported using the bucket#templateId, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Ci/mediaSuperResolutionTemplate:MediaSuperResolutionTemplate media_super_resolution_template terraform-ci-xxxxxx#t1d707eb2be3294e22b47123894f85cb8f
 * ```
 */
export class MediaSuperResolutionTemplate extends pulumi.CustomResource {
    /**
     * Get an existing MediaSuperResolutionTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MediaSuperResolutionTemplateState, opts?: pulumi.CustomResourceOptions): MediaSuperResolutionTemplate {
        return new MediaSuperResolutionTemplate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Ci/mediaSuperResolutionTemplate:MediaSuperResolutionTemplate';

    /**
     * Returns true if the given object is an instance of MediaSuperResolutionTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MediaSuperResolutionTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MediaSuperResolutionTemplate.__pulumiType;
    }

    /**
     * bucket name.
     */
    public readonly bucket!: pulumi.Output<string>;
    /**
     * Auto scaling switch, off by default.
     */
    public readonly enableScaleUp!: pulumi.Output<string | undefined>;
    /**
     * The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Resolution Options sdtohd: Standard Definition to Ultra Definition, hdto4k: HD to 4K.
     */
    public readonly resolution!: pulumi.Output<string>;
    /**
     * version, default value Base, Base: basic version, Enhance: enhanced version.
     */
    public readonly version!: pulumi.Output<string | undefined>;

    /**
     * Create a MediaSuperResolutionTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MediaSuperResolutionTemplateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MediaSuperResolutionTemplateArgs | MediaSuperResolutionTemplateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MediaSuperResolutionTemplateState | undefined;
            resourceInputs["bucket"] = state ? state.bucket : undefined;
            resourceInputs["enableScaleUp"] = state ? state.enableScaleUp : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["resolution"] = state ? state.resolution : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as MediaSuperResolutionTemplateArgs | undefined;
            if ((!args || args.bucket === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bucket'");
            }
            if ((!args || args.resolution === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resolution'");
            }
            resourceInputs["bucket"] = args ? args.bucket : undefined;
            resourceInputs["enableScaleUp"] = args ? args.enableScaleUp : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["resolution"] = args ? args.resolution : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(MediaSuperResolutionTemplate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MediaSuperResolutionTemplate resources.
 */
export interface MediaSuperResolutionTemplateState {
    /**
     * bucket name.
     */
    bucket?: pulumi.Input<string>;
    /**
     * Auto scaling switch, off by default.
     */
    enableScaleUp?: pulumi.Input<string>;
    /**
     * The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
     */
    name?: pulumi.Input<string>;
    /**
     * Resolution Options sdtohd: Standard Definition to Ultra Definition, hdto4k: HD to 4K.
     */
    resolution?: pulumi.Input<string>;
    /**
     * version, default value Base, Base: basic version, Enhance: enhanced version.
     */
    version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MediaSuperResolutionTemplate resource.
 */
export interface MediaSuperResolutionTemplateArgs {
    /**
     * bucket name.
     */
    bucket: pulumi.Input<string>;
    /**
     * Auto scaling switch, off by default.
     */
    enableScaleUp?: pulumi.Input<string>;
    /**
     * The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
     */
    name?: pulumi.Input<string>;
    /**
     * Resolution Options sdtohd: Standard Definition to Ultra Definition, hdto4k: HD to 4K.
     */
    resolution: pulumi.Input<string>;
    /**
     * version, default value Base, Base: basic version, Enhance: enhanced version.
     */
    version?: pulumi.Input<string>;
}
