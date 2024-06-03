// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a cfw addressTemplate
 *
 * ## Example Usage
 *
 * ### If type is 1
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const example = new tencentcloud.cfw.AddressTemplate("example", {
 *     detail: "test template",
 *     ipString: "1.1.1.1,2.2.2.2",
 *     type: 1,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ### If type is 5
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const example = new tencentcloud.cfw.AddressTemplate("example", {
 *     detail: "test template",
 *     ipString: "www.qq.com,www.tencent.com",
 *     type: 5,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * cfw address_template can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import tencentcloud:Cfw/addressTemplate:AddressTemplate example mb_1300846651_1695611353900
 * ```
 */
export class AddressTemplate extends pulumi.CustomResource {
    /**
     * Get an existing AddressTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AddressTemplateState, opts?: pulumi.CustomResourceOptions): AddressTemplate {
        return new AddressTemplate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Cfw/addressTemplate:AddressTemplate';

    /**
     * Returns true if the given object is an instance of AddressTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AddressTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AddressTemplate.__pulumiType;
    }

    /**
     * Template Detail.
     */
    public readonly detail!: pulumi.Output<string>;
    /**
     * Type is 1, ip template eg: 1.1.1.1,2.2.2.2; Type is 5, domain name template eg: www.qq.com, www.tencent.com.
     */
    public readonly ipString!: pulumi.Output<string>;
    /**
     * Template name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * 1: ip template; 5: domain name templates.
     */
    public readonly type!: pulumi.Output<number>;

    /**
     * Create a AddressTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AddressTemplateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AddressTemplateArgs | AddressTemplateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AddressTemplateState | undefined;
            resourceInputs["detail"] = state ? state.detail : undefined;
            resourceInputs["ipString"] = state ? state.ipString : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as AddressTemplateArgs | undefined;
            if ((!args || args.detail === undefined) && !opts.urn) {
                throw new Error("Missing required property 'detail'");
            }
            if ((!args || args.ipString === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ipString'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["detail"] = args ? args.detail : undefined;
            resourceInputs["ipString"] = args ? args.ipString : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AddressTemplate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AddressTemplate resources.
 */
export interface AddressTemplateState {
    /**
     * Template Detail.
     */
    detail?: pulumi.Input<string>;
    /**
     * Type is 1, ip template eg: 1.1.1.1,2.2.2.2; Type is 5, domain name template eg: www.qq.com, www.tencent.com.
     */
    ipString?: pulumi.Input<string>;
    /**
     * Template name.
     */
    name?: pulumi.Input<string>;
    /**
     * 1: ip template; 5: domain name templates.
     */
    type?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a AddressTemplate resource.
 */
export interface AddressTemplateArgs {
    /**
     * Template Detail.
     */
    detail: pulumi.Input<string>;
    /**
     * Type is 1, ip template eg: 1.1.1.1,2.2.2.2; Type is 5, domain name template eg: www.qq.com, www.tencent.com.
     */
    ipString: pulumi.Input<string>;
    /**
     * Template name.
     */
    name?: pulumi.Input<string>;
    /**
     * 1: ip template; 5: domain name templates.
     */
    type: pulumi.Input<number>;
}
