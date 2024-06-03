// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a monitor grafanaIntegration
 *
 * ## Example Usage
 *
 * ### Create a grafan instance and integrate the configuration
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const config = new pulumi.Config();
 * const availabilityZone = config.get("availabilityZone") || "ap-guangzhou-6";
 * const vpc = new tencentcloud.vpc.Instance("vpc", {cidrBlock: "10.0.0.0/16"});
 * const subnet = new tencentcloud.subnet.Instance("subnet", {
 *     vpcId: vpc.id,
 *     availabilityZone: availabilityZone,
 *     cidrBlock: "10.0.1.0/24",
 * });
 * const foo = new tencentcloud.monitor.GrafanaInstance("foo", {
 *     instanceName: "test-grafana",
 *     vpcId: vpc.id,
 *     subnetIds: [subnet.id],
 *     grafanaInitPassword: "1234567890",
 *     enableInternet: false,
 *     tags: {
 *         createdBy: "test",
 *     },
 * });
 * const grafanaIntegration = new tencentcloud.monitor.GrafanaIntegration("grafanaIntegration", {
 *     instanceId: foo.id,
 *     kind: "tencentcloud-monitor-app",
 *     content: "{\"kind\":\"tencentcloud-monitor-app\",\"spec\":{\"dataSourceSpec\":{\"authProvider\":{\"__anyOf\":\"使用密钥\",\"useRole\":true,\"secretId\":\"arunma@tencent.com\",\"secretKey\":\"12345678\"},\"name\":\"uint-test\"},\"grafanaSpec\":{\"organizationIds\":[]}}}",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class GrafanaIntegration extends pulumi.CustomResource {
    /**
     * Get an existing GrafanaIntegration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GrafanaIntegrationState, opts?: pulumi.CustomResourceOptions): GrafanaIntegration {
        return new GrafanaIntegration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Monitor/grafanaIntegration:GrafanaIntegration';

    /**
     * Returns true if the given object is an instance of GrafanaIntegration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GrafanaIntegration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GrafanaIntegration.__pulumiType;
    }

    /**
     * generated json string of given integration json schema.
     */
    public readonly content!: pulumi.Output<string>;
    /**
     * integration desc.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * grafana instance id.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * integration id.
     */
    public /*out*/ readonly integrationId!: pulumi.Output<string>;
    /**
     * integration json schema kind.
     */
    public readonly kind!: pulumi.Output<string>;

    /**
     * Create a GrafanaIntegration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GrafanaIntegrationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GrafanaIntegrationArgs | GrafanaIntegrationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GrafanaIntegrationState | undefined;
            resourceInputs["content"] = state ? state.content : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["integrationId"] = state ? state.integrationId : undefined;
            resourceInputs["kind"] = state ? state.kind : undefined;
        } else {
            const args = argsOrState as GrafanaIntegrationArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["content"] = args ? args.content : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["kind"] = args ? args.kind : undefined;
            resourceInputs["integrationId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GrafanaIntegration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GrafanaIntegration resources.
 */
export interface GrafanaIntegrationState {
    /**
     * generated json string of given integration json schema.
     */
    content?: pulumi.Input<string>;
    /**
     * integration desc.
     */
    description?: pulumi.Input<string>;
    /**
     * grafana instance id.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * integration id.
     */
    integrationId?: pulumi.Input<string>;
    /**
     * integration json schema kind.
     */
    kind?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GrafanaIntegration resource.
 */
export interface GrafanaIntegrationArgs {
    /**
     * generated json string of given integration json schema.
     */
    content?: pulumi.Input<string>;
    /**
     * integration desc.
     */
    description?: pulumi.Input<string>;
    /**
     * grafana instance id.
     */
    instanceId: pulumi.Input<string>;
    /**
     * integration json schema kind.
     */
    kind?: pulumi.Input<string>;
}
