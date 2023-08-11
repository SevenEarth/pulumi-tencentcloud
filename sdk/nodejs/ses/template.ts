// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a ses template.
 *
 * ## Example Usage
 * ### Create a ses html template
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const example = new tencentcloud.Ses.Template("example", {
 *     templateContent: {
 *         html: `<!DOCTYPE html>
 * <html lang="en">
 * <head>
 *   <meta charset="UTF-8">
 *   <meta name="viewport" content="width=device-width, initial-scale=1.0">
 *   <title>mail title</title>
 * </head>
 * <body>
 * <div class="container">
 *   <h1>Welcome to our service! </h1>
 *   <p>Dear user,</p>
 *   <p>Thank you for using Tencent Cloud:</p>
 *   <p><a href="https://cloud.tencent.com/document/product/1653">https://cloud.tencent.com/document/product/1653</a></p>
 *   <p>If you did not request this email, please ignore it. </p>
 *   <p><strong>from the iac team</strong></p>
 * </div>
 * </body>
 * </html>
 * `,
 *     },
 *     templateName: "tf_example_ses_temp",
 * });
 * ```
 *
 * ## Import
 *
 * ses template can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Ses/template:Template example template_id
 * ```
 */
export class Template extends pulumi.CustomResource {
    /**
     * Get an existing Template resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TemplateState, opts?: pulumi.CustomResourceOptions): Template {
        return new Template(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Ses/template:Template';

    /**
     * Returns true if the given object is an instance of Template.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Template {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Template.__pulumiType;
    }

    /**
     * Sms Template Content.
     */
    public readonly templateContent!: pulumi.Output<outputs.Ses.TemplateTemplateContent>;
    /**
     * smsTemplateName, which must be required.
     */
    public readonly templateName!: pulumi.Output<string>;

    /**
     * Create a Template resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TemplateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TemplateArgs | TemplateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TemplateState | undefined;
            resourceInputs["templateContent"] = state ? state.templateContent : undefined;
            resourceInputs["templateName"] = state ? state.templateName : undefined;
        } else {
            const args = argsOrState as TemplateArgs | undefined;
            if ((!args || args.templateContent === undefined) && !opts.urn) {
                throw new Error("Missing required property 'templateContent'");
            }
            if ((!args || args.templateName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'templateName'");
            }
            resourceInputs["templateContent"] = args ? args.templateContent : undefined;
            resourceInputs["templateName"] = args ? args.templateName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Template.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Template resources.
 */
export interface TemplateState {
    /**
     * Sms Template Content.
     */
    templateContent?: pulumi.Input<inputs.Ses.TemplateTemplateContent>;
    /**
     * smsTemplateName, which must be required.
     */
    templateName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Template resource.
 */
export interface TemplateArgs {
    /**
     * Sms Template Content.
     */
    templateContent: pulumi.Input<inputs.Ses.TemplateTemplateContent>;
    /**
     * smsTemplateName, which must be required.
     */
    templateName: pulumi.Input<string>;
}
