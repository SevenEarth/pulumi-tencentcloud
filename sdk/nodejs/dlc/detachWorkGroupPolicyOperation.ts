// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class DetachWorkGroupPolicyOperation extends pulumi.CustomResource {
    /**
     * Get an existing DetachWorkGroupPolicyOperation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DetachWorkGroupPolicyOperationState, opts?: pulumi.CustomResourceOptions): DetachWorkGroupPolicyOperation {
        return new DetachWorkGroupPolicyOperation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Dlc/detachWorkGroupPolicyOperation:DetachWorkGroupPolicyOperation';

    /**
     * Returns true if the given object is an instance of DetachWorkGroupPolicyOperation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DetachWorkGroupPolicyOperation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DetachWorkGroupPolicyOperation.__pulumiType;
    }

    /**
     * The set of policies to be bound.
     */
    public readonly policySets!: pulumi.Output<outputs.Dlc.DetachWorkGroupPolicyOperationPolicySet[] | undefined>;
    /**
     * Work group id.
     */
    public readonly workGroupId!: pulumi.Output<number>;

    /**
     * Create a DetachWorkGroupPolicyOperation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DetachWorkGroupPolicyOperationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DetachWorkGroupPolicyOperationArgs | DetachWorkGroupPolicyOperationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DetachWorkGroupPolicyOperationState | undefined;
            resourceInputs["policySets"] = state ? state.policySets : undefined;
            resourceInputs["workGroupId"] = state ? state.workGroupId : undefined;
        } else {
            const args = argsOrState as DetachWorkGroupPolicyOperationArgs | undefined;
            if ((!args || args.workGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'workGroupId'");
            }
            resourceInputs["policySets"] = args ? args.policySets : undefined;
            resourceInputs["workGroupId"] = args ? args.workGroupId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DetachWorkGroupPolicyOperation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DetachWorkGroupPolicyOperation resources.
 */
export interface DetachWorkGroupPolicyOperationState {
    /**
     * The set of policies to be bound.
     */
    policySets?: pulumi.Input<pulumi.Input<inputs.Dlc.DetachWorkGroupPolicyOperationPolicySet>[]>;
    /**
     * Work group id.
     */
    workGroupId?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a DetachWorkGroupPolicyOperation resource.
 */
export interface DetachWorkGroupPolicyOperationArgs {
    /**
     * The set of policies to be bound.
     */
    policySets?: pulumi.Input<pulumi.Input<inputs.Dlc.DetachWorkGroupPolicyOperationPolicySet>[]>;
    /**
     * Work group id.
     */
    workGroupId: pulumi.Input<number>;
}
