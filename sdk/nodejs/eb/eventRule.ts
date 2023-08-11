// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a eb eventRule
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 *
 * const foo = new tencentcloud.eb.EventBus("foo", {
 *     eventBusName: "tf-event_bus",
 *     description: "event bus desc",
 *     enableStore: false,
 *     saveDays: 1,
 *     tags: {
 *         createdBy: "terraform",
 *     },
 * });
 * const eventRule = new tencentcloud.eb.EventRule("eventRule", {
 *     eventBusId: foo.id,
 *     ruleName: "tf-event_rule",
 *     description: "event rule desc",
 *     enable: true,
 *     eventPattern: JSON.stringify({
 *         source: "apigw.cloud.tencent",
 *         type: ["connector:apigw"],
 *     }),
 *     tags: {
 *         createdBy: "terraform",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * eb event_rule can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Eb/eventRule:EventRule event_rule event_rule_id
 * ```
 */
export class EventRule extends pulumi.CustomResource {
    /**
     * Get an existing EventRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EventRuleState, opts?: pulumi.CustomResourceOptions): EventRule {
        return new EventRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Eb/eventRule:EventRule';

    /**
     * Returns true if the given object is an instance of EventRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EventRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EventRule.__pulumiType;
    }

    /**
     * Event set description, unlimited character type, description within 200 characters.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Enable switch.
     */
    public readonly enable!: pulumi.Output<boolean | undefined>;
    /**
     * event bus Id.
     */
    public readonly eventBusId!: pulumi.Output<string>;
    /**
     * Reference: [Event Mode](https://cloud.tencent.com/document/product/1359/56084).
     */
    public readonly eventPattern!: pulumi.Output<string>;
    /**
     * event rule id.
     */
    public /*out*/ readonly ruleId!: pulumi.Output<string>;
    /**
     * Event rule name, which can only contain letters, numbers, underscores, hyphens, starts with a letter and ends with a number or letter, 2~60 characters.
     */
    public readonly ruleName!: pulumi.Output<string>;
    /**
     * Tag description list.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a EventRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EventRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EventRuleArgs | EventRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EventRuleState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["enable"] = state ? state.enable : undefined;
            resourceInputs["eventBusId"] = state ? state.eventBusId : undefined;
            resourceInputs["eventPattern"] = state ? state.eventPattern : undefined;
            resourceInputs["ruleId"] = state ? state.ruleId : undefined;
            resourceInputs["ruleName"] = state ? state.ruleName : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as EventRuleArgs | undefined;
            if ((!args || args.eventBusId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'eventBusId'");
            }
            if ((!args || args.eventPattern === undefined) && !opts.urn) {
                throw new Error("Missing required property 'eventPattern'");
            }
            if ((!args || args.ruleName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ruleName'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["enable"] = args ? args.enable : undefined;
            resourceInputs["eventBusId"] = args ? args.eventBusId : undefined;
            resourceInputs["eventPattern"] = args ? args.eventPattern : undefined;
            resourceInputs["ruleName"] = args ? args.ruleName : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["ruleId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EventRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EventRule resources.
 */
export interface EventRuleState {
    /**
     * Event set description, unlimited character type, description within 200 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Enable switch.
     */
    enable?: pulumi.Input<boolean>;
    /**
     * event bus Id.
     */
    eventBusId?: pulumi.Input<string>;
    /**
     * Reference: [Event Mode](https://cloud.tencent.com/document/product/1359/56084).
     */
    eventPattern?: pulumi.Input<string>;
    /**
     * event rule id.
     */
    ruleId?: pulumi.Input<string>;
    /**
     * Event rule name, which can only contain letters, numbers, underscores, hyphens, starts with a letter and ends with a number or letter, 2~60 characters.
     */
    ruleName?: pulumi.Input<string>;
    /**
     * Tag description list.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a EventRule resource.
 */
export interface EventRuleArgs {
    /**
     * Event set description, unlimited character type, description within 200 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Enable switch.
     */
    enable?: pulumi.Input<boolean>;
    /**
     * event bus Id.
     */
    eventBusId: pulumi.Input<string>;
    /**
     * Reference: [Event Mode](https://cloud.tencent.com/document/product/1359/56084).
     */
    eventPattern: pulumi.Input<string>;
    /**
     * Event rule name, which can only contain letters, numbers, underscores, hyphens, starts with a letter and ends with a number or letter, 2~60 characters.
     */
    ruleName: pulumi.Input<string>;
    /**
     * Tag description list.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}
