// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provide a resource to invoke a Url Push request.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const foo = new tencentcloud.Cdn.UrlPush("foo", {
 *     urls: ["https://www.example.com/b"],
 * });
 * ```
 * ### argument to request new push task with same urls
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const foo = new tencentcloud.Cdn.UrlPush("foo", {
 *     redo: 1,
 *     urls: ["https://www.example.com/a"],
 * });
 * ```
 */
export class UrlPush extends pulumi.CustomResource {
    /**
     * Get an existing UrlPush resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UrlPushState, opts?: pulumi.CustomResourceOptions): UrlPush {
        return new UrlPush(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Cdn/urlPush:UrlPush';

    /**
     * Returns true if the given object is an instance of UrlPush.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UrlPush {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UrlPush.__pulumiType;
    }

    /**
     * Specify push area. NOTE: only push same area cache contents.
     */
    public readonly area!: pulumi.Output<string | undefined>;
    /**
     * Layer to push.
     */
    public readonly layer!: pulumi.Output<string | undefined>;
    /**
     * Whether to recursive parse m3u8 files.
     */
    public readonly parseM3u8!: pulumi.Output<boolean | undefined>;
    /**
     * logs of latest push task.
     */
    public /*out*/ readonly pushHistories!: pulumi.Output<outputs.Cdn.UrlPushPushHistory[]>;
    /**
     * Change to push again. NOTE: this argument only works while resource update, if set to `0` or null will not be triggered.
     */
    public readonly redo!: pulumi.Output<number | undefined>;
    /**
     * Push task id.
     */
    public /*out*/ readonly taskId!: pulumi.Output<string>;
    /**
     * List of url to push. NOTE: urls need include protocol prefix `http://` or `https://`.
     */
    public readonly urls!: pulumi.Output<string[]>;
    /**
     * Specify `User-Agent` HTTP header, default: `TencentCdn`.
     */
    public readonly userAgent!: pulumi.Output<string | undefined>;

    /**
     * Create a UrlPush resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UrlPushArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UrlPushArgs | UrlPushState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UrlPushState | undefined;
            resourceInputs["area"] = state ? state.area : undefined;
            resourceInputs["layer"] = state ? state.layer : undefined;
            resourceInputs["parseM3u8"] = state ? state.parseM3u8 : undefined;
            resourceInputs["pushHistories"] = state ? state.pushHistories : undefined;
            resourceInputs["redo"] = state ? state.redo : undefined;
            resourceInputs["taskId"] = state ? state.taskId : undefined;
            resourceInputs["urls"] = state ? state.urls : undefined;
            resourceInputs["userAgent"] = state ? state.userAgent : undefined;
        } else {
            const args = argsOrState as UrlPushArgs | undefined;
            if ((!args || args.urls === undefined) && !opts.urn) {
                throw new Error("Missing required property 'urls'");
            }
            resourceInputs["area"] = args ? args.area : undefined;
            resourceInputs["layer"] = args ? args.layer : undefined;
            resourceInputs["parseM3u8"] = args ? args.parseM3u8 : undefined;
            resourceInputs["redo"] = args ? args.redo : undefined;
            resourceInputs["urls"] = args ? args.urls : undefined;
            resourceInputs["userAgent"] = args ? args.userAgent : undefined;
            resourceInputs["pushHistories"] = undefined /*out*/;
            resourceInputs["taskId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(UrlPush.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UrlPush resources.
 */
export interface UrlPushState {
    /**
     * Specify push area. NOTE: only push same area cache contents.
     */
    area?: pulumi.Input<string>;
    /**
     * Layer to push.
     */
    layer?: pulumi.Input<string>;
    /**
     * Whether to recursive parse m3u8 files.
     */
    parseM3u8?: pulumi.Input<boolean>;
    /**
     * logs of latest push task.
     */
    pushHistories?: pulumi.Input<pulumi.Input<inputs.Cdn.UrlPushPushHistory>[]>;
    /**
     * Change to push again. NOTE: this argument only works while resource update, if set to `0` or null will not be triggered.
     */
    redo?: pulumi.Input<number>;
    /**
     * Push task id.
     */
    taskId?: pulumi.Input<string>;
    /**
     * List of url to push. NOTE: urls need include protocol prefix `http://` or `https://`.
     */
    urls?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specify `User-Agent` HTTP header, default: `TencentCdn`.
     */
    userAgent?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UrlPush resource.
 */
export interface UrlPushArgs {
    /**
     * Specify push area. NOTE: only push same area cache contents.
     */
    area?: pulumi.Input<string>;
    /**
     * Layer to push.
     */
    layer?: pulumi.Input<string>;
    /**
     * Whether to recursive parse m3u8 files.
     */
    parseM3u8?: pulumi.Input<boolean>;
    /**
     * Change to push again. NOTE: this argument only works while resource update, if set to `0` or null will not be triggered.
     */
    redo?: pulumi.Input<number>;
    /**
     * List of url to push. NOTE: urls need include protocol prefix `http://` or `https://`.
     */
    urls: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specify `User-Agent` HTTP header, default: `TencentCdn`.
     */
    userAgent?: pulumi.Input<string>;
}
