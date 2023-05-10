// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a monitor tmpAlertRule
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const tmpAlertRule = new tencentcloud.Monitor.TmpAlertRule("tmpAlertRule", {
 *     annotations: [{
 *         key: "hello2",
 *         value: "world2",
 *     }],
 *     duration: "4m",
 *     expr: "up{service=\"rig-prometheus-agent\"}>0",
 *     instanceId: "prom-c89b3b3u",
 *     labels: [{
 *         key: "hello1",
 *         value: "world1",
 *     }],
 *     receivers: ["notice-l9ziyxw6"],
 *     ruleName: "test123",
 *     ruleState: 2,
 * });
 * ```
 *
 * ## Import
 *
 * monitor tmpAlertRule can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Monitor/tmpAlertRule:TmpAlertRule tmpAlertRule instanceId#Rule_id
 * ```
 */
export class TmpAlertRule extends pulumi.CustomResource {
    /**
     * Get an existing TmpAlertRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TmpAlertRuleState, opts?: pulumi.CustomResourceOptions): TmpAlertRule {
        return new TmpAlertRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Monitor/tmpAlertRule:TmpAlertRule';

    /**
     * Returns true if the given object is an instance of TmpAlertRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TmpAlertRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TmpAlertRule.__pulumiType;
    }

    /**
     * Rule alarm duration.
     */
    public readonly annotations!: pulumi.Output<outputs.Monitor.TmpAlertRuleAnnotation[] | undefined>;
    /**
     * Rule alarm duration.
     */
    public readonly duration!: pulumi.Output<string | undefined>;
    /**
     * Rule expression, reference documentation: `https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/`.
     */
    public readonly expr!: pulumi.Output<string>;
    /**
     * Instance id.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Rule alarm duration.
     */
    public readonly labels!: pulumi.Output<outputs.Monitor.TmpAlertRuleLabel[] | undefined>;
    /**
     * Alarm notification template id list.
     */
    public readonly receivers!: pulumi.Output<string[]>;
    /**
     * Rule name.
     */
    public readonly ruleName!: pulumi.Output<string>;
    /**
     * Rule state code.
     */
    public readonly ruleState!: pulumi.Output<number | undefined>;
    /**
     * Alarm Policy Template Classification.
     */
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a TmpAlertRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TmpAlertRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TmpAlertRuleArgs | TmpAlertRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TmpAlertRuleState | undefined;
            resourceInputs["annotations"] = state ? state.annotations : undefined;
            resourceInputs["duration"] = state ? state.duration : undefined;
            resourceInputs["expr"] = state ? state.expr : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["receivers"] = state ? state.receivers : undefined;
            resourceInputs["ruleName"] = state ? state.ruleName : undefined;
            resourceInputs["ruleState"] = state ? state.ruleState : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as TmpAlertRuleArgs | undefined;
            if ((!args || args.expr === undefined) && !opts.urn) {
                throw new Error("Missing required property 'expr'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.receivers === undefined) && !opts.urn) {
                throw new Error("Missing required property 'receivers'");
            }
            if ((!args || args.ruleName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ruleName'");
            }
            resourceInputs["annotations"] = args ? args.annotations : undefined;
            resourceInputs["duration"] = args ? args.duration : undefined;
            resourceInputs["expr"] = args ? args.expr : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["receivers"] = args ? args.receivers : undefined;
            resourceInputs["ruleName"] = args ? args.ruleName : undefined;
            resourceInputs["ruleState"] = args ? args.ruleState : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TmpAlertRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TmpAlertRule resources.
 */
export interface TmpAlertRuleState {
    /**
     * Rule alarm duration.
     */
    annotations?: pulumi.Input<pulumi.Input<inputs.Monitor.TmpAlertRuleAnnotation>[]>;
    /**
     * Rule alarm duration.
     */
    duration?: pulumi.Input<string>;
    /**
     * Rule expression, reference documentation: `https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/`.
     */
    expr?: pulumi.Input<string>;
    /**
     * Instance id.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Rule alarm duration.
     */
    labels?: pulumi.Input<pulumi.Input<inputs.Monitor.TmpAlertRuleLabel>[]>;
    /**
     * Alarm notification template id list.
     */
    receivers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Rule name.
     */
    ruleName?: pulumi.Input<string>;
    /**
     * Rule state code.
     */
    ruleState?: pulumi.Input<number>;
    /**
     * Alarm Policy Template Classification.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TmpAlertRule resource.
 */
export interface TmpAlertRuleArgs {
    /**
     * Rule alarm duration.
     */
    annotations?: pulumi.Input<pulumi.Input<inputs.Monitor.TmpAlertRuleAnnotation>[]>;
    /**
     * Rule alarm duration.
     */
    duration?: pulumi.Input<string>;
    /**
     * Rule expression, reference documentation: `https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/`.
     */
    expr: pulumi.Input<string>;
    /**
     * Instance id.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Rule alarm duration.
     */
    labels?: pulumi.Input<pulumi.Input<inputs.Monitor.TmpAlertRuleLabel>[]>;
    /**
     * Alarm notification template id list.
     */
    receivers: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Rule name.
     */
    ruleName: pulumi.Input<string>;
    /**
     * Rule state code.
     */
    ruleState?: pulumi.Input<number>;
    /**
     * Alarm Policy Template Classification.
     */
    type?: pulumi.Input<string>;
}
