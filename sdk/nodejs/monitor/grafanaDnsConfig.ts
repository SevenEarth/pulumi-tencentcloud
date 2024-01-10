// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a monitor grafanaDnsConfig
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const grafanaDnsConfig = new tencentcloud.Monitor.GrafanaDnsConfig("grafana_dns_config", {
 *     instanceId: "grafana-dp2hnnfa",
 *     nameServers: [
 *         "10.1.2.1",
 *         "10.1.2.2",
 *         "10.1.2.3",
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * monitor grafana_dns_config can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Monitor/grafanaDnsConfig:GrafanaDnsConfig grafana_dns_config instance_id
 * ```
 */
export class GrafanaDnsConfig extends pulumi.CustomResource {
    /**
     * Get an existing GrafanaDnsConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GrafanaDnsConfigState, opts?: pulumi.CustomResourceOptions): GrafanaDnsConfig {
        return new GrafanaDnsConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Monitor/grafanaDnsConfig:GrafanaDnsConfig';

    /**
     * Returns true if the given object is an instance of GrafanaDnsConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GrafanaDnsConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GrafanaDnsConfig.__pulumiType;
    }

    /**
     * Grafana instance ID.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * DNS nameserver list.
     */
    public readonly nameServers!: pulumi.Output<string[] | undefined>;

    /**
     * Create a GrafanaDnsConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GrafanaDnsConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GrafanaDnsConfigArgs | GrafanaDnsConfigState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GrafanaDnsConfigState | undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["nameServers"] = state ? state.nameServers : undefined;
        } else {
            const args = argsOrState as GrafanaDnsConfigArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["nameServers"] = args ? args.nameServers : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GrafanaDnsConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GrafanaDnsConfig resources.
 */
export interface GrafanaDnsConfigState {
    /**
     * Grafana instance ID.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * DNS nameserver list.
     */
    nameServers?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a GrafanaDnsConfig resource.
 */
export interface GrafanaDnsConfigArgs {
    /**
     * Grafana instance ID.
     */
    instanceId: pulumi.Input<string>;
    /**
     * DNS nameserver list.
     */
    nameServers?: pulumi.Input<pulumi.Input<string>[]>;
}
