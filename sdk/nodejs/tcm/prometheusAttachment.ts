// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a tcm prometheusAttachment
 *
 * > **NOTE:** Instructions for use: 1. Use Tencent Cloud Prometheus to monitor TMP, please enter `vpcId`, `subnetId`, `region` or `instanceId`, it is recommended to use an existing tmp instance; 2. To use the third-party Prometheus service, please enter `customProm`; 3. `tencentcloud.Tcm.PrometheusAttachment` does not support modification; 4. If you use Tencent Cloud Prometheus to monitor TMP, enter `vpcId`, `subnetId`, `region` to create a new Prometheus monitoring instance, destroy will not destroy the Prometheus monitoring instance
 * **NOTE:** If you use the config attribute prometheus in tencentcloud_tcm_mesh, do not use tencentcloud.Tcm.PrometheusAttachment
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const prometheusAttachment = new tencentcloud.tcm.PrometheusAttachment("prometheusAttachment", {
 *     meshId: "mesh-rofjmxxx",
 *     prometheus: {
 *         instanceId: "",
 *         region: "ap-guangzhou",
 *         subnetId: "subnet-driddxxx",
 *         vpcId: "vpc-pewdpxxx",
 *     },
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * tcm prometheus_attachment can be imported using the mesh_id, e.g.
 *
 * ```sh
 * $ pulumi import tencentcloud:Tcm/prometheusAttachment:PrometheusAttachment prometheus_attachment mesh-rofjmxxx
 * ```
 */
export class PrometheusAttachment extends pulumi.CustomResource {
    /**
     * Get an existing PrometheusAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PrometheusAttachmentState, opts?: pulumi.CustomResourceOptions): PrometheusAttachment {
        return new PrometheusAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Tcm/prometheusAttachment:PrometheusAttachment';

    /**
     * Returns true if the given object is an instance of PrometheusAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PrometheusAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PrometheusAttachment.__pulumiType;
    }

    /**
     * Mesh ID.
     */
    public readonly meshId!: pulumi.Output<string>;
    /**
     * Prometheus configuration.
     */
    public readonly prometheus!: pulumi.Output<outputs.Tcm.PrometheusAttachmentPrometheus>;

    /**
     * Create a PrometheusAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PrometheusAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PrometheusAttachmentArgs | PrometheusAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PrometheusAttachmentState | undefined;
            resourceInputs["meshId"] = state ? state.meshId : undefined;
            resourceInputs["prometheus"] = state ? state.prometheus : undefined;
        } else {
            const args = argsOrState as PrometheusAttachmentArgs | undefined;
            if ((!args || args.meshId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'meshId'");
            }
            if ((!args || args.prometheus === undefined) && !opts.urn) {
                throw new Error("Missing required property 'prometheus'");
            }
            resourceInputs["meshId"] = args ? args.meshId : undefined;
            resourceInputs["prometheus"] = args ? args.prometheus : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PrometheusAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PrometheusAttachment resources.
 */
export interface PrometheusAttachmentState {
    /**
     * Mesh ID.
     */
    meshId?: pulumi.Input<string>;
    /**
     * Prometheus configuration.
     */
    prometheus?: pulumi.Input<inputs.Tcm.PrometheusAttachmentPrometheus>;
}

/**
 * The set of arguments for constructing a PrometheusAttachment resource.
 */
export interface PrometheusAttachmentArgs {
    /**
     * Mesh ID.
     */
    meshId: pulumi.Input<string>;
    /**
     * Prometheus configuration.
     */
    prometheus: pulumi.Input<inputs.Tcm.PrometheusAttachmentPrometheus>;
}
