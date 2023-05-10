// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manage Guetzli compression functionality
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const foo = new tencentcloud.Ci.Guetzli("foo", {
 *     bucket: "examplebucket-1250000000",
 *     status: "on",
 * });
 * ```
 *
 * ## Import
 *
 * Resource guetzli can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Ci/guetzli:Guetzli example examplebucket-1250000000
 * ```
 */
export class Guetzli extends pulumi.CustomResource {
    /**
     * Get an existing Guetzli resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GuetzliState, opts?: pulumi.CustomResourceOptions): Guetzli {
        return new Guetzli(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Ci/guetzli:Guetzli';

    /**
     * Returns true if the given object is an instance of Guetzli.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Guetzli {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Guetzli.__pulumiType;
    }

    /**
     * The name of a bucket, the format should be [custom name]-[appid], for example `mycos-1258798060`.
     */
    public readonly bucket!: pulumi.Output<string>;
    /**
     * Whether Guetzli is set, options: on/off.
     */
    public readonly status!: pulumi.Output<string>;

    /**
     * Create a Guetzli resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GuetzliArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GuetzliArgs | GuetzliState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GuetzliState | undefined;
            resourceInputs["bucket"] = state ? state.bucket : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as GuetzliArgs | undefined;
            if ((!args || args.bucket === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bucket'");
            }
            if ((!args || args.status === undefined) && !opts.urn) {
                throw new Error("Missing required property 'status'");
            }
            resourceInputs["bucket"] = args ? args.bucket : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Guetzli.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Guetzli resources.
 */
export interface GuetzliState {
    /**
     * The name of a bucket, the format should be [custom name]-[appid], for example `mycos-1258798060`.
     */
    bucket?: pulumi.Input<string>;
    /**
     * Whether Guetzli is set, options: on/off.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Guetzli resource.
 */
export interface GuetzliArgs {
    /**
     * The name of a bucket, the format should be [custom name]-[appid], for example `mycos-1258798060`.
     */
    bucket: pulumi.Input<string>;
    /**
     * Whether Guetzli is set, options: on/off.
     */
    status: pulumi.Input<string>;
}
