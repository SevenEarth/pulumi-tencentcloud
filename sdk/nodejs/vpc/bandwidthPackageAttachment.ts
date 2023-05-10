// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a vpc bandwidthPackageAttachment
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const bandwidthPackageAttachment = new tencentcloud.Vpc.BandwidthPackageAttachment("bandwidth_package_attachment", {
 *     bandwidthPackageId: "bwp-atmf0p9g",
 *     networkType: "BGP",
 *     protocol: "",
 *     resourceId: "lb-dv1ai6ma",
 *     resourceType: "LoadBalance",
 * });
 * ```
 */
export class BandwidthPackageAttachment extends pulumi.CustomResource {
    /**
     * Get an existing BandwidthPackageAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BandwidthPackageAttachmentState, opts?: pulumi.CustomResourceOptions): BandwidthPackageAttachment {
        return new BandwidthPackageAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Vpc/bandwidthPackageAttachment:BandwidthPackageAttachment';

    /**
     * Returns true if the given object is an instance of BandwidthPackageAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BandwidthPackageAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BandwidthPackageAttachment.__pulumiType;
    }

    /**
     * Bandwidth package unique ID, in the form of `bwp-xxxx`.
     */
    public readonly bandwidthPackageId!: pulumi.Output<string>;
    /**
     * Bandwidth packet type, currently supports `BGP` type, indicating that the internal resource is BGP IP.
     */
    public readonly networkType!: pulumi.Output<string | undefined>;
    /**
     * Bandwidth packet protocol type. Currently `ipv4` and `ipv6` protocol types are supported.
     */
    public readonly protocol!: pulumi.Output<string | undefined>;
    /**
     * The unique ID of the resource, currently supports EIP resources and LB resources, such as `eip-xxxx`, `lb-xxxx`.
     */
    public readonly resourceId!: pulumi.Output<string>;
    /**
     * Resource types, including `Address`, `LoadBalance`.
     */
    public readonly resourceType!: pulumi.Output<string | undefined>;

    /**
     * Create a BandwidthPackageAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BandwidthPackageAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BandwidthPackageAttachmentArgs | BandwidthPackageAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BandwidthPackageAttachmentState | undefined;
            resourceInputs["bandwidthPackageId"] = state ? state.bandwidthPackageId : undefined;
            resourceInputs["networkType"] = state ? state.networkType : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["resourceId"] = state ? state.resourceId : undefined;
            resourceInputs["resourceType"] = state ? state.resourceType : undefined;
        } else {
            const args = argsOrState as BandwidthPackageAttachmentArgs | undefined;
            if ((!args || args.bandwidthPackageId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bandwidthPackageId'");
            }
            if ((!args || args.resourceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceId'");
            }
            resourceInputs["bandwidthPackageId"] = args ? args.bandwidthPackageId : undefined;
            resourceInputs["networkType"] = args ? args.networkType : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["resourceId"] = args ? args.resourceId : undefined;
            resourceInputs["resourceType"] = args ? args.resourceType : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BandwidthPackageAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BandwidthPackageAttachment resources.
 */
export interface BandwidthPackageAttachmentState {
    /**
     * Bandwidth package unique ID, in the form of `bwp-xxxx`.
     */
    bandwidthPackageId?: pulumi.Input<string>;
    /**
     * Bandwidth packet type, currently supports `BGP` type, indicating that the internal resource is BGP IP.
     */
    networkType?: pulumi.Input<string>;
    /**
     * Bandwidth packet protocol type. Currently `ipv4` and `ipv6` protocol types are supported.
     */
    protocol?: pulumi.Input<string>;
    /**
     * The unique ID of the resource, currently supports EIP resources and LB resources, such as `eip-xxxx`, `lb-xxxx`.
     */
    resourceId?: pulumi.Input<string>;
    /**
     * Resource types, including `Address`, `LoadBalance`.
     */
    resourceType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BandwidthPackageAttachment resource.
 */
export interface BandwidthPackageAttachmentArgs {
    /**
     * Bandwidth package unique ID, in the form of `bwp-xxxx`.
     */
    bandwidthPackageId: pulumi.Input<string>;
    /**
     * Bandwidth packet type, currently supports `BGP` type, indicating that the internal resource is BGP IP.
     */
    networkType?: pulumi.Input<string>;
    /**
     * Bandwidth packet protocol type. Currently `ipv4` and `ipv6` protocol types are supported.
     */
    protocol?: pulumi.Input<string>;
    /**
     * The unique ID of the resource, currently supports EIP resources and LB resources, such as `eip-xxxx`, `lb-xxxx`.
     */
    resourceId: pulumi.Input<string>;
    /**
     * Resource types, including `Address`, `LoadBalance`.
     */
    resourceType?: pulumi.Input<string>;
}
