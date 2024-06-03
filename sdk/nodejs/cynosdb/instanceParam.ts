// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a cynosdb instanceParam
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const instanceParam = new tencentcloud.cynosdb.InstanceParam("instanceParam", {
 *     clusterId: "cynosdbmysql-bws8h88b",
 *     instanceId: "cynosdbmysql-ins-rikr6z4o",
 *     instanceParamLists: [{
 *         currentValue: "0",
 *         paramName: "init_connect",
 *     }],
 *     isInMaintainPeriod: "no",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class InstanceParam extends pulumi.CustomResource {
    /**
     * Get an existing InstanceParam resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceParamState, opts?: pulumi.CustomResourceOptions): InstanceParam {
        return new InstanceParam(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Cynosdb/instanceParam:InstanceParam';

    /**
     * Returns true if the given object is an instance of InstanceParam.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InstanceParam {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InstanceParam.__pulumiType;
    }

    /**
     * Cluster ID.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * Instance ID.
     */
    public readonly instanceId!: pulumi.Output<string | undefined>;
    /**
     * Instance parameter list.
     */
    public readonly instanceParamLists!: pulumi.Output<outputs.Cynosdb.InstanceParamInstanceParamList[] | undefined>;
    /**
     * Yes: modify within the operation and maintenance time window, no: execute immediately (default value).
     */
    public readonly isInMaintainPeriod!: pulumi.Output<string | undefined>;

    /**
     * Create a InstanceParam resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceParamArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceParamArgs | InstanceParamState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceParamState | undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["instanceParamLists"] = state ? state.instanceParamLists : undefined;
            resourceInputs["isInMaintainPeriod"] = state ? state.isInMaintainPeriod : undefined;
        } else {
            const args = argsOrState as InstanceParamArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["instanceParamLists"] = args ? args.instanceParamLists : undefined;
            resourceInputs["isInMaintainPeriod"] = args ? args.isInMaintainPeriod : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(InstanceParam.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InstanceParam resources.
 */
export interface InstanceParamState {
    /**
     * Cluster ID.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * Instance ID.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Instance parameter list.
     */
    instanceParamLists?: pulumi.Input<pulumi.Input<inputs.Cynosdb.InstanceParamInstanceParamList>[]>;
    /**
     * Yes: modify within the operation and maintenance time window, no: execute immediately (default value).
     */
    isInMaintainPeriod?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InstanceParam resource.
 */
export interface InstanceParamArgs {
    /**
     * Cluster ID.
     */
    clusterId: pulumi.Input<string>;
    /**
     * Instance ID.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Instance parameter list.
     */
    instanceParamLists?: pulumi.Input<pulumi.Input<inputs.Cynosdb.InstanceParamInstanceParamList>[]>;
    /**
     * Yes: modify within the operation and maintenance time window, no: execute immediately (default value).
     */
    isInMaintainPeriod?: pulumi.Input<string>;
}
