// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a mysql drInstanceToMater
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const drInstanceToMater = new tencentcloud.Mysql.DrInstanceToMater("dr_instance_to_mater", {
 *     instanceId: "cdb-c1nl9rpv",
 * });
 * ```
 *
 * ## Import
 *
 * mysql dr_instance_to_mater can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Mysql/drInstanceToMater:DrInstanceToMater dr_instance_to_mater dr_instance_to_mater_id
 * ```
 */
export class DrInstanceToMater extends pulumi.CustomResource {
    /**
     * Get an existing DrInstanceToMater resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DrInstanceToMaterState, opts?: pulumi.CustomResourceOptions): DrInstanceToMater {
        return new DrInstanceToMater(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Mysql/drInstanceToMater:DrInstanceToMater';

    /**
     * Returns true if the given object is an instance of DrInstanceToMater.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DrInstanceToMater {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DrInstanceToMater.__pulumiType;
    }

    /**
     * Disaster recovery instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
     */
    public readonly instanceId!: pulumi.Output<string>;

    /**
     * Create a DrInstanceToMater resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DrInstanceToMaterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DrInstanceToMaterArgs | DrInstanceToMaterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DrInstanceToMaterState | undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
        } else {
            const args = argsOrState as DrInstanceToMaterArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DrInstanceToMater.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DrInstanceToMater resources.
 */
export interface DrInstanceToMaterState {
    /**
     * Disaster recovery instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
     */
    instanceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DrInstanceToMater resource.
 */
export interface DrInstanceToMaterArgs {
    /**
     * Disaster recovery instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
     */
    instanceId: pulumi.Input<string>;
}
