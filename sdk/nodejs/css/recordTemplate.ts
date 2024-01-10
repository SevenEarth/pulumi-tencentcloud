// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export class RecordTemplate extends pulumi.CustomResource {
    /**
     * Get an existing RecordTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RecordTemplateState, opts?: pulumi.CustomResourceOptions): RecordTemplate {
        return new RecordTemplate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Css/recordTemplate:RecordTemplate';

    /**
     * Returns true if the given object is an instance of RecordTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RecordTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RecordTemplate.__pulumiType;
    }

    /**
     * Aac recording parameters are set when Aac recording is enabled.
     */
    public readonly aacParam!: pulumi.Output<outputs.Css.RecordTemplateAacParam | undefined>;
    /**
     * Description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Flv recording parameters are set when Flv recording is enabled.
     */
    public readonly flvParam!: pulumi.Output<outputs.Css.RecordTemplateFlvParam | undefined>;
    /**
     * FLV records special parameters.
     */
    public readonly flvSpecialParam!: pulumi.Output<outputs.Css.RecordTemplateFlvSpecialParam | undefined>;
    /**
     * Hls recording parameters, which are set when hls recording is enabled.
     */
    public readonly hlsParam!: pulumi.Output<outputs.Css.RecordTemplateHlsParam | undefined>;
    /**
     * HLS specific recording parameters.
     */
    public readonly hlsSpecialParam!: pulumi.Output<outputs.Css.RecordTemplateHlsSpecialParam | undefined>;
    /**
     * Live broadcast type, 0 by default. 0: Ordinary live broadcast, 1: Slow broadcast.
     */
    public readonly isDelayLive!: pulumi.Output<number | undefined>;
    /**
     * Mp3 recording parameters are set when Mp3 recording is turned on.
     */
    public readonly mp3Param!: pulumi.Output<outputs.Css.RecordTemplateMp3Param | undefined>;
    /**
     * Mp4 recording parameters are set when Mp4 recording is enabled.
     */
    public readonly mp4Param!: pulumi.Output<outputs.Css.RecordTemplateMp4Param | undefined>;
    /**
     * Whether to remove the watermark. This parameter is invalid when the type is slow live broadcast.
     */
    public readonly removeWatermark!: pulumi.Output<boolean | undefined>;
    /**
     * Template name. Only `Chinese`, `English`, `numbers`, `_`, `-` are supported.
     */
    public readonly templateName!: pulumi.Output<string>;

    /**
     * Create a RecordTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RecordTemplateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RecordTemplateArgs | RecordTemplateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RecordTemplateState | undefined;
            resourceInputs["aacParam"] = state ? state.aacParam : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["flvParam"] = state ? state.flvParam : undefined;
            resourceInputs["flvSpecialParam"] = state ? state.flvSpecialParam : undefined;
            resourceInputs["hlsParam"] = state ? state.hlsParam : undefined;
            resourceInputs["hlsSpecialParam"] = state ? state.hlsSpecialParam : undefined;
            resourceInputs["isDelayLive"] = state ? state.isDelayLive : undefined;
            resourceInputs["mp3Param"] = state ? state.mp3Param : undefined;
            resourceInputs["mp4Param"] = state ? state.mp4Param : undefined;
            resourceInputs["removeWatermark"] = state ? state.removeWatermark : undefined;
            resourceInputs["templateName"] = state ? state.templateName : undefined;
        } else {
            const args = argsOrState as RecordTemplateArgs | undefined;
            if ((!args || args.templateName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'templateName'");
            }
            resourceInputs["aacParam"] = args ? args.aacParam : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["flvParam"] = args ? args.flvParam : undefined;
            resourceInputs["flvSpecialParam"] = args ? args.flvSpecialParam : undefined;
            resourceInputs["hlsParam"] = args ? args.hlsParam : undefined;
            resourceInputs["hlsSpecialParam"] = args ? args.hlsSpecialParam : undefined;
            resourceInputs["isDelayLive"] = args ? args.isDelayLive : undefined;
            resourceInputs["mp3Param"] = args ? args.mp3Param : undefined;
            resourceInputs["mp4Param"] = args ? args.mp4Param : undefined;
            resourceInputs["removeWatermark"] = args ? args.removeWatermark : undefined;
            resourceInputs["templateName"] = args ? args.templateName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RecordTemplate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RecordTemplate resources.
 */
export interface RecordTemplateState {
    /**
     * Aac recording parameters are set when Aac recording is enabled.
     */
    aacParam?: pulumi.Input<inputs.Css.RecordTemplateAacParam>;
    /**
     * Description.
     */
    description?: pulumi.Input<string>;
    /**
     * Flv recording parameters are set when Flv recording is enabled.
     */
    flvParam?: pulumi.Input<inputs.Css.RecordTemplateFlvParam>;
    /**
     * FLV records special parameters.
     */
    flvSpecialParam?: pulumi.Input<inputs.Css.RecordTemplateFlvSpecialParam>;
    /**
     * Hls recording parameters, which are set when hls recording is enabled.
     */
    hlsParam?: pulumi.Input<inputs.Css.RecordTemplateHlsParam>;
    /**
     * HLS specific recording parameters.
     */
    hlsSpecialParam?: pulumi.Input<inputs.Css.RecordTemplateHlsSpecialParam>;
    /**
     * Live broadcast type, 0 by default. 0: Ordinary live broadcast, 1: Slow broadcast.
     */
    isDelayLive?: pulumi.Input<number>;
    /**
     * Mp3 recording parameters are set when Mp3 recording is turned on.
     */
    mp3Param?: pulumi.Input<inputs.Css.RecordTemplateMp3Param>;
    /**
     * Mp4 recording parameters are set when Mp4 recording is enabled.
     */
    mp4Param?: pulumi.Input<inputs.Css.RecordTemplateMp4Param>;
    /**
     * Whether to remove the watermark. This parameter is invalid when the type is slow live broadcast.
     */
    removeWatermark?: pulumi.Input<boolean>;
    /**
     * Template name. Only `Chinese`, `English`, `numbers`, `_`, `-` are supported.
     */
    templateName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RecordTemplate resource.
 */
export interface RecordTemplateArgs {
    /**
     * Aac recording parameters are set when Aac recording is enabled.
     */
    aacParam?: pulumi.Input<inputs.Css.RecordTemplateAacParam>;
    /**
     * Description.
     */
    description?: pulumi.Input<string>;
    /**
     * Flv recording parameters are set when Flv recording is enabled.
     */
    flvParam?: pulumi.Input<inputs.Css.RecordTemplateFlvParam>;
    /**
     * FLV records special parameters.
     */
    flvSpecialParam?: pulumi.Input<inputs.Css.RecordTemplateFlvSpecialParam>;
    /**
     * Hls recording parameters, which are set when hls recording is enabled.
     */
    hlsParam?: pulumi.Input<inputs.Css.RecordTemplateHlsParam>;
    /**
     * HLS specific recording parameters.
     */
    hlsSpecialParam?: pulumi.Input<inputs.Css.RecordTemplateHlsSpecialParam>;
    /**
     * Live broadcast type, 0 by default. 0: Ordinary live broadcast, 1: Slow broadcast.
     */
    isDelayLive?: pulumi.Input<number>;
    /**
     * Mp3 recording parameters are set when Mp3 recording is turned on.
     */
    mp3Param?: pulumi.Input<inputs.Css.RecordTemplateMp3Param>;
    /**
     * Mp4 recording parameters are set when Mp4 recording is enabled.
     */
    mp4Param?: pulumi.Input<inputs.Css.RecordTemplateMp4Param>;
    /**
     * Whether to remove the watermark. This parameter is invalid when the type is slow live broadcast.
     */
    removeWatermark?: pulumi.Input<boolean>;
    /**
     * Template name. Only `Chinese`, `English`, `numbers`, `_`, `-` are supported.
     */
    templateName: pulumi.Input<string>;
}
