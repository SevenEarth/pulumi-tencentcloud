// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a ses receiver
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const receiver = new tencentcloud.ses.Receiver("receiver", {
 *     datas: [
 *         {
 *             email: "abc@abc.com",
 *         },
 *         {
 *             email: "abcd@abcd.com",
 *         },
 *     ],
 *     desc: "description",
 *     receiversName: "terraform_test",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const receiver = new tencentcloud.ses.Receiver("receiver", {
 *     datas: [
 *         {
 *             email: "abc@abc.com",
 *             templateData: "{\"name\":\"xxx\",\"age\":\"xx\"}",
 *         },
 *         {
 *             email: "abcd@abcd.com",
 *             templateData: "{\"name\":\"xxx\",\"age\":\"xx\"}",
 *         },
 *     ],
 *     desc: "description",
 *     receiversName: "terraform_test",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * ses email_address can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import tencentcloud:Ses/receiver:Receiver receiver receiverId
 * ```
 */
export class Receiver extends pulumi.CustomResource {
    /**
     * Get an existing Receiver resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ReceiverState, opts?: pulumi.CustomResourceOptions): Receiver {
        return new Receiver(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Ses/receiver:Receiver';

    /**
     * Returns true if the given object is an instance of Receiver.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Receiver {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Receiver.__pulumiType;
    }

    /**
     * Recipient email and template parameters in array format. The number of recipients is limited to within 20,000. If there is an object in the `data` list that inputs `templateData`, then other objects are also required.
     */
    public readonly datas!: pulumi.Output<outputs.Ses.ReceiverData[]>;
    /**
     * Recipient group description.
     */
    public readonly desc!: pulumi.Output<string | undefined>;
    /**
     * Recipient group name.
     */
    public readonly receiversName!: pulumi.Output<string>;

    /**
     * Create a Receiver resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReceiverArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ReceiverArgs | ReceiverState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ReceiverState | undefined;
            resourceInputs["datas"] = state ? state.datas : undefined;
            resourceInputs["desc"] = state ? state.desc : undefined;
            resourceInputs["receiversName"] = state ? state.receiversName : undefined;
        } else {
            const args = argsOrState as ReceiverArgs | undefined;
            if ((!args || args.datas === undefined) && !opts.urn) {
                throw new Error("Missing required property 'datas'");
            }
            if ((!args || args.receiversName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'receiversName'");
            }
            resourceInputs["datas"] = args ? args.datas : undefined;
            resourceInputs["desc"] = args ? args.desc : undefined;
            resourceInputs["receiversName"] = args ? args.receiversName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Receiver.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Receiver resources.
 */
export interface ReceiverState {
    /**
     * Recipient email and template parameters in array format. The number of recipients is limited to within 20,000. If there is an object in the `data` list that inputs `templateData`, then other objects are also required.
     */
    datas?: pulumi.Input<pulumi.Input<inputs.Ses.ReceiverData>[]>;
    /**
     * Recipient group description.
     */
    desc?: pulumi.Input<string>;
    /**
     * Recipient group name.
     */
    receiversName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Receiver resource.
 */
export interface ReceiverArgs {
    /**
     * Recipient email and template parameters in array format. The number of recipients is limited to within 20,000. If there is an object in the `data` list that inputs `templateData`, then other objects are also required.
     */
    datas: pulumi.Input<pulumi.Input<inputs.Ses.ReceiverData>[]>;
    /**
     * Recipient group description.
     */
    desc?: pulumi.Input<string>;
    /**
     * Recipient group name.
     */
    receiversName: pulumi.Input<string>;
}
