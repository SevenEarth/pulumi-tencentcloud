// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a dlc userDataEngineConfig
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const userDataEngineConfig = new tencentcloud.Dlc.UserDataEngineConfig("user_data_engine_config", {
 *     dataEngineConfigPairs: [{
 *         configItem: "qq",
 *         configValue: "ff",
 *     }],
 *     dataEngineId: "DataEngine-cgkvbas6",
 *     sessionResourceTemplate: {
 *         driverSize: "small",
 *         executorMaxNumbers: 1,
 *         executorNums: 1,
 *         executorSize: "small",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * dlc user_data_engine_config can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Dlc/userDataEngineConfig:UserDataEngineConfig user_data_engine_config user_data_engine_config_id
 * ```
 */
export class UserDataEngineConfig extends pulumi.CustomResource {
    /**
     * Get an existing UserDataEngineConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserDataEngineConfigState, opts?: pulumi.CustomResourceOptions): UserDataEngineConfig {
        return new UserDataEngineConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Dlc/userDataEngineConfig:UserDataEngineConfig';

    /**
     * Returns true if the given object is an instance of UserDataEngineConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserDataEngineConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserDataEngineConfig.__pulumiType;
    }

    /**
     * Engine configuration items.
     */
    public readonly dataEngineConfigPairs!: pulumi.Output<outputs.Dlc.UserDataEngineConfigDataEngineConfigPair[] | undefined>;
    /**
     * Engine unique id.
     */
    public readonly dataEngineId!: pulumi.Output<string>;
    /**
     * Job engine resource configuration template.
     */
    public readonly sessionResourceTemplate!: pulumi.Output<outputs.Dlc.UserDataEngineConfigSessionResourceTemplate | undefined>;

    /**
     * Create a UserDataEngineConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserDataEngineConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserDataEngineConfigArgs | UserDataEngineConfigState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserDataEngineConfigState | undefined;
            resourceInputs["dataEngineConfigPairs"] = state ? state.dataEngineConfigPairs : undefined;
            resourceInputs["dataEngineId"] = state ? state.dataEngineId : undefined;
            resourceInputs["sessionResourceTemplate"] = state ? state.sessionResourceTemplate : undefined;
        } else {
            const args = argsOrState as UserDataEngineConfigArgs | undefined;
            if ((!args || args.dataEngineId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataEngineId'");
            }
            resourceInputs["dataEngineConfigPairs"] = args ? args.dataEngineConfigPairs : undefined;
            resourceInputs["dataEngineId"] = args ? args.dataEngineId : undefined;
            resourceInputs["sessionResourceTemplate"] = args ? args.sessionResourceTemplate : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(UserDataEngineConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserDataEngineConfig resources.
 */
export interface UserDataEngineConfigState {
    /**
     * Engine configuration items.
     */
    dataEngineConfigPairs?: pulumi.Input<pulumi.Input<inputs.Dlc.UserDataEngineConfigDataEngineConfigPair>[]>;
    /**
     * Engine unique id.
     */
    dataEngineId?: pulumi.Input<string>;
    /**
     * Job engine resource configuration template.
     */
    sessionResourceTemplate?: pulumi.Input<inputs.Dlc.UserDataEngineConfigSessionResourceTemplate>;
}

/**
 * The set of arguments for constructing a UserDataEngineConfig resource.
 */
export interface UserDataEngineConfigArgs {
    /**
     * Engine configuration items.
     */
    dataEngineConfigPairs?: pulumi.Input<pulumi.Input<inputs.Dlc.UserDataEngineConfigDataEngineConfigPair>[]>;
    /**
     * Engine unique id.
     */
    dataEngineId: pulumi.Input<string>;
    /**
     * Job engine resource configuration template.
     */
    sessionResourceTemplate?: pulumi.Input<inputs.Dlc.UserDataEngineConfigSessionResourceTemplate>;
}
