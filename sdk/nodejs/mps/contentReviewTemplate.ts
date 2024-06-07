// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a mps contentReviewTemplate
 *
 * ## Import
 *
 * mps content_review_template can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import tencentcloud:Mps/contentReviewTemplate:ContentReviewTemplate content_review_template definition
 * ```
 */
export class ContentReviewTemplate extends pulumi.CustomResource {
    /**
     * Get an existing ContentReviewTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ContentReviewTemplateState, opts?: pulumi.CustomResourceOptions): ContentReviewTemplate {
        return new ContentReviewTemplate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Mps/contentReviewTemplate:ContentReviewTemplate';

    /**
     * Returns true if the given object is an instance of ContentReviewTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ContentReviewTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ContentReviewTemplate.__pulumiType;
    }

    /**
     * Content review template description information, length limit: 256 characters.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * Content review template name, length limit: 64 characters.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Political control parameters.
     */
    public readonly politicalConfigure!: pulumi.Output<outputs.Mps.ContentReviewTemplatePoliticalConfigure | undefined>;
    /**
     * Control parameters for porn image.
     */
    public readonly pornConfigure!: pulumi.Output<outputs.Mps.ContentReviewTemplatePornConfigure | undefined>;
    /**
     * Prohibited control parameters. Prohibited content includes:abuse, drug-related violations.Note: this parameter is not yet supported.
     */
    public readonly prohibitedConfigure!: pulumi.Output<outputs.Mps.ContentReviewTemplateProhibitedConfigure | undefined>;
    /**
     * Control parameters for unsafe information.
     */
    public readonly terrorismConfigure!: pulumi.Output<outputs.Mps.ContentReviewTemplateTerrorismConfigure | undefined>;
    /**
     * User-Defined Content Moderation Control Parameters.
     */
    public readonly userDefineConfigure!: pulumi.Output<outputs.Mps.ContentReviewTemplateUserDefineConfigure | undefined>;

    /**
     * Create a ContentReviewTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ContentReviewTemplateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ContentReviewTemplateArgs | ContentReviewTemplateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ContentReviewTemplateState | undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["politicalConfigure"] = state ? state.politicalConfigure : undefined;
            resourceInputs["pornConfigure"] = state ? state.pornConfigure : undefined;
            resourceInputs["prohibitedConfigure"] = state ? state.prohibitedConfigure : undefined;
            resourceInputs["terrorismConfigure"] = state ? state.terrorismConfigure : undefined;
            resourceInputs["userDefineConfigure"] = state ? state.userDefineConfigure : undefined;
        } else {
            const args = argsOrState as ContentReviewTemplateArgs | undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["politicalConfigure"] = args ? args.politicalConfigure : undefined;
            resourceInputs["pornConfigure"] = args ? args.pornConfigure : undefined;
            resourceInputs["prohibitedConfigure"] = args ? args.prohibitedConfigure : undefined;
            resourceInputs["terrorismConfigure"] = args ? args.terrorismConfigure : undefined;
            resourceInputs["userDefineConfigure"] = args ? args.userDefineConfigure : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ContentReviewTemplate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ContentReviewTemplate resources.
 */
export interface ContentReviewTemplateState {
    /**
     * Content review template description information, length limit: 256 characters.
     */
    comment?: pulumi.Input<string>;
    /**
     * Content review template name, length limit: 64 characters.
     */
    name?: pulumi.Input<string>;
    /**
     * Political control parameters.
     */
    politicalConfigure?: pulumi.Input<inputs.Mps.ContentReviewTemplatePoliticalConfigure>;
    /**
     * Control parameters for porn image.
     */
    pornConfigure?: pulumi.Input<inputs.Mps.ContentReviewTemplatePornConfigure>;
    /**
     * Prohibited control parameters. Prohibited content includes:abuse, drug-related violations.Note: this parameter is not yet supported.
     */
    prohibitedConfigure?: pulumi.Input<inputs.Mps.ContentReviewTemplateProhibitedConfigure>;
    /**
     * Control parameters for unsafe information.
     */
    terrorismConfigure?: pulumi.Input<inputs.Mps.ContentReviewTemplateTerrorismConfigure>;
    /**
     * User-Defined Content Moderation Control Parameters.
     */
    userDefineConfigure?: pulumi.Input<inputs.Mps.ContentReviewTemplateUserDefineConfigure>;
}

/**
 * The set of arguments for constructing a ContentReviewTemplate resource.
 */
export interface ContentReviewTemplateArgs {
    /**
     * Content review template description information, length limit: 256 characters.
     */
    comment?: pulumi.Input<string>;
    /**
     * Content review template name, length limit: 64 characters.
     */
    name?: pulumi.Input<string>;
    /**
     * Political control parameters.
     */
    politicalConfigure?: pulumi.Input<inputs.Mps.ContentReviewTemplatePoliticalConfigure>;
    /**
     * Control parameters for porn image.
     */
    pornConfigure?: pulumi.Input<inputs.Mps.ContentReviewTemplatePornConfigure>;
    /**
     * Prohibited control parameters. Prohibited content includes:abuse, drug-related violations.Note: this parameter is not yet supported.
     */
    prohibitedConfigure?: pulumi.Input<inputs.Mps.ContentReviewTemplateProhibitedConfigure>;
    /**
     * Control parameters for unsafe information.
     */
    terrorismConfigure?: pulumi.Input<inputs.Mps.ContentReviewTemplateTerrorismConfigure>;
    /**
     * User-Defined Content Moderation Control Parameters.
     */
    userDefineConfigure?: pulumi.Input<inputs.Mps.ContentReviewTemplateUserDefineConfigure>;
}
