// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provide a resource to create a VOD image sprite template.
 *
 * ## Import
 *
 * VOD image sprite template can be imported using the id($subAppId#$templateId), e.g.
 *
 * ```sh
 * $ pulumi import tencentcloud:Vod/imageSpriteTemplate:ImageSpriteTemplate foo $subAppId#$templateId
 * ```
 */
export class ImageSpriteTemplate extends pulumi.CustomResource {
    /**
     * Get an existing ImageSpriteTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ImageSpriteTemplateState, opts?: pulumi.CustomResourceOptions): ImageSpriteTemplate {
        return new ImageSpriteTemplate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Vod/imageSpriteTemplate:ImageSpriteTemplate';

    /**
     * Returns true if the given object is an instance of ImageSpriteTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ImageSpriteTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ImageSpriteTemplate.__pulumiType;
    }

    /**
     * Subimage column count of an image sprite.
     */
    public readonly columnCount!: pulumi.Output<number>;
    /**
     * Template description. Length limit: 256 characters.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * Creation time of template in ISO date format.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Fill refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported: `stretch`: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot shorter or longer; `black`: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks. Default value: `black`.
     */
    public readonly fillType!: pulumi.Output<string | undefined>;
    /**
     * Image format, Valid values:
     * - jpg: jpg format;
     * - png: png format;
     * - webp: webp format;
     * Default value: jpg.
     */
    public readonly format!: pulumi.Output<string>;
    /**
     * Maximum value of the `height` (or short side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both `width` and `height` are `0`, the resolution will be the same as that of the source video; If `width` is `0`, but `height` is not `0`, `width` will be proportionally scaled; If `width` is not `0`, but `height` is `0`, `height` will be proportionally scaled; If both `width` and `height` are not `0`, the custom resolution will be used. Default value: `0`.
     */
    public readonly height!: pulumi.Output<number | undefined>;
    /**
     * Name of a time point screen capturing template. Length limit: 64 characters.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Resolution adaption. Valid values: `true`,`false`. `true`: enabled. In this case, `width` represents the long side of a video, while `height` the short side; `false`: disabled. In this case, `width` represents the width of a video, while `height` the height. Default value: `true`.
     */
    public readonly resolutionAdaptive!: pulumi.Output<boolean | undefined>;
    /**
     * Subimage row count of an image sprite.
     */
    public readonly rowCount!: pulumi.Output<number>;
    /**
     * Sampling interval. If `sampleType` is `Percent`, sampling will be performed at an interval of the specified percentage. If `sampleType` is `Time`, sampling will be performed at the specified time interval in seconds.
     */
    public readonly sampleInterval!: pulumi.Output<number>;
    /**
     * Sampling type. Valid values: `Percent`, `Time`. `Percent`: by percent. `Time`: by time interval.
     */
    public readonly sampleType!: pulumi.Output<string>;
    /**
     * The VOD [application](https://intl.cloud.tencent.com/document/product/266/14574) ID. For customers who activate VOD service from December 25, 2023, if they want to access resources in a VOD application (whether it's the default application or a newly created one), they must fill in this field with the application ID.
     */
    public readonly subAppId!: pulumi.Output<number | undefined>;
    /**
     * Template type, value range:
     * - Preset: system preset template;
     * - Custom: user-defined templates.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Last modified time of template in ISO date format.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * Maximum value of the `width` (or long side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both `width` and `height` are `0`, the resolution will be the same as that of the source video; If `width` is `0`, but `height` is not `0`, width will be proportionally scaled; If `width` is not `0`, but `height` is `0`, `height` will be proportionally scaled; If both `width` and `height` are not `0`, the custom resolution will be used. Default value: `0`.
     */
    public readonly width!: pulumi.Output<number | undefined>;

    /**
     * Create a ImageSpriteTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ImageSpriteTemplateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ImageSpriteTemplateArgs | ImageSpriteTemplateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ImageSpriteTemplateState | undefined;
            resourceInputs["columnCount"] = state ? state.columnCount : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["fillType"] = state ? state.fillType : undefined;
            resourceInputs["format"] = state ? state.format : undefined;
            resourceInputs["height"] = state ? state.height : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["resolutionAdaptive"] = state ? state.resolutionAdaptive : undefined;
            resourceInputs["rowCount"] = state ? state.rowCount : undefined;
            resourceInputs["sampleInterval"] = state ? state.sampleInterval : undefined;
            resourceInputs["sampleType"] = state ? state.sampleType : undefined;
            resourceInputs["subAppId"] = state ? state.subAppId : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
            resourceInputs["width"] = state ? state.width : undefined;
        } else {
            const args = argsOrState as ImageSpriteTemplateArgs | undefined;
            if ((!args || args.columnCount === undefined) && !opts.urn) {
                throw new Error("Missing required property 'columnCount'");
            }
            if ((!args || args.rowCount === undefined) && !opts.urn) {
                throw new Error("Missing required property 'rowCount'");
            }
            if ((!args || args.sampleInterval === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sampleInterval'");
            }
            if ((!args || args.sampleType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sampleType'");
            }
            resourceInputs["columnCount"] = args ? args.columnCount : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["fillType"] = args ? args.fillType : undefined;
            resourceInputs["format"] = args ? args.format : undefined;
            resourceInputs["height"] = args ? args.height : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["resolutionAdaptive"] = args ? args.resolutionAdaptive : undefined;
            resourceInputs["rowCount"] = args ? args.rowCount : undefined;
            resourceInputs["sampleInterval"] = args ? args.sampleInterval : undefined;
            resourceInputs["sampleType"] = args ? args.sampleType : undefined;
            resourceInputs["subAppId"] = args ? args.subAppId : undefined;
            resourceInputs["width"] = args ? args.width : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ImageSpriteTemplate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ImageSpriteTemplate resources.
 */
export interface ImageSpriteTemplateState {
    /**
     * Subimage column count of an image sprite.
     */
    columnCount?: pulumi.Input<number>;
    /**
     * Template description. Length limit: 256 characters.
     */
    comment?: pulumi.Input<string>;
    /**
     * Creation time of template in ISO date format.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Fill refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported: `stretch`: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot shorter or longer; `black`: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks. Default value: `black`.
     */
    fillType?: pulumi.Input<string>;
    /**
     * Image format, Valid values:
     * - jpg: jpg format;
     * - png: png format;
     * - webp: webp format;
     * Default value: jpg.
     */
    format?: pulumi.Input<string>;
    /**
     * Maximum value of the `height` (or short side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both `width` and `height` are `0`, the resolution will be the same as that of the source video; If `width` is `0`, but `height` is not `0`, `width` will be proportionally scaled; If `width` is not `0`, but `height` is `0`, `height` will be proportionally scaled; If both `width` and `height` are not `0`, the custom resolution will be used. Default value: `0`.
     */
    height?: pulumi.Input<number>;
    /**
     * Name of a time point screen capturing template. Length limit: 64 characters.
     */
    name?: pulumi.Input<string>;
    /**
     * Resolution adaption. Valid values: `true`,`false`. `true`: enabled. In this case, `width` represents the long side of a video, while `height` the short side; `false`: disabled. In this case, `width` represents the width of a video, while `height` the height. Default value: `true`.
     */
    resolutionAdaptive?: pulumi.Input<boolean>;
    /**
     * Subimage row count of an image sprite.
     */
    rowCount?: pulumi.Input<number>;
    /**
     * Sampling interval. If `sampleType` is `Percent`, sampling will be performed at an interval of the specified percentage. If `sampleType` is `Time`, sampling will be performed at the specified time interval in seconds.
     */
    sampleInterval?: pulumi.Input<number>;
    /**
     * Sampling type. Valid values: `Percent`, `Time`. `Percent`: by percent. `Time`: by time interval.
     */
    sampleType?: pulumi.Input<string>;
    /**
     * The VOD [application](https://intl.cloud.tencent.com/document/product/266/14574) ID. For customers who activate VOD service from December 25, 2023, if they want to access resources in a VOD application (whether it's the default application or a newly created one), they must fill in this field with the application ID.
     */
    subAppId?: pulumi.Input<number>;
    /**
     * Template type, value range:
     * - Preset: system preset template;
     * - Custom: user-defined templates.
     */
    type?: pulumi.Input<string>;
    /**
     * Last modified time of template in ISO date format.
     */
    updateTime?: pulumi.Input<string>;
    /**
     * Maximum value of the `width` (or long side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both `width` and `height` are `0`, the resolution will be the same as that of the source video; If `width` is `0`, but `height` is not `0`, width will be proportionally scaled; If `width` is not `0`, but `height` is `0`, `height` will be proportionally scaled; If both `width` and `height` are not `0`, the custom resolution will be used. Default value: `0`.
     */
    width?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a ImageSpriteTemplate resource.
 */
export interface ImageSpriteTemplateArgs {
    /**
     * Subimage column count of an image sprite.
     */
    columnCount: pulumi.Input<number>;
    /**
     * Template description. Length limit: 256 characters.
     */
    comment?: pulumi.Input<string>;
    /**
     * Fill refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported: `stretch`: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot shorter or longer; `black`: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks. Default value: `black`.
     */
    fillType?: pulumi.Input<string>;
    /**
     * Image format, Valid values:
     * - jpg: jpg format;
     * - png: png format;
     * - webp: webp format;
     * Default value: jpg.
     */
    format?: pulumi.Input<string>;
    /**
     * Maximum value of the `height` (or short side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both `width` and `height` are `0`, the resolution will be the same as that of the source video; If `width` is `0`, but `height` is not `0`, `width` will be proportionally scaled; If `width` is not `0`, but `height` is `0`, `height` will be proportionally scaled; If both `width` and `height` are not `0`, the custom resolution will be used. Default value: `0`.
     */
    height?: pulumi.Input<number>;
    /**
     * Name of a time point screen capturing template. Length limit: 64 characters.
     */
    name?: pulumi.Input<string>;
    /**
     * Resolution adaption. Valid values: `true`,`false`. `true`: enabled. In this case, `width` represents the long side of a video, while `height` the short side; `false`: disabled. In this case, `width` represents the width of a video, while `height` the height. Default value: `true`.
     */
    resolutionAdaptive?: pulumi.Input<boolean>;
    /**
     * Subimage row count of an image sprite.
     */
    rowCount: pulumi.Input<number>;
    /**
     * Sampling interval. If `sampleType` is `Percent`, sampling will be performed at an interval of the specified percentage. If `sampleType` is `Time`, sampling will be performed at the specified time interval in seconds.
     */
    sampleInterval: pulumi.Input<number>;
    /**
     * Sampling type. Valid values: `Percent`, `Time`. `Percent`: by percent. `Time`: by time interval.
     */
    sampleType: pulumi.Input<string>;
    /**
     * The VOD [application](https://intl.cloud.tencent.com/document/product/266/14574) ID. For customers who activate VOD service from December 25, 2023, if they want to access resources in a VOD application (whether it's the default application or a newly created one), they must fill in this field with the application ID.
     */
    subAppId?: pulumi.Input<number>;
    /**
     * Maximum value of the `width` (or long side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both `width` and `height` are `0`, the resolution will be the same as that of the source video; If `width` is `0`, but `height` is not `0`, width will be proportionally scaled; If `width` is not `0`, but `height` is `0`, `height` will be proportionally scaled; If both `width` and `height` are not `0`, the custom resolution will be used. Default value: `0`.
     */
    width?: pulumi.Input<number>;
}
