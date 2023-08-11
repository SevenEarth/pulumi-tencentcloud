// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a eb eventTarget
 *
 * ## Example Usage
 * ### Create an event target of type scf
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const config = new pulumi.Config();
 * const zone = config.get("zone") || "ap-guangzhou";
 * const namespace = config.get("namespace") || "default";
 * const function = config.get("function") || "keep-1676351130";
 * const functionVersion = config.get("functionVersion") || `$LATEST`;
 * const fooUsers = tencentcloud.Cam.getUsers({});
 * const fooEventBus = new tencentcloud.eb.EventBus("fooEventBus", {
 *     eventBusName: "tf-event_bus",
 *     description: "event bus desc",
 *     enableStore: false,
 *     saveDays: 1,
 *     tags: {
 *         createdBy: "terraform",
 *     },
 * });
 * const fooEventRule = new tencentcloud.eb.EventRule("fooEventRule", {
 *     eventBusId: fooEventBus.id,
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
 * const scfTarget = new tencentcloud.eb.EventTarget("scfTarget", {
 *     eventBusId: fooEventBus.id,
 *     ruleId: fooEventRule.ruleId,
 *     type: "scf",
 *     targetDescription: {
 *         resourceDescription: fooUsers.then(fooUsers => `qcs::scf:${zone}:uin/${fooUsers.userLists?[0]?.uin}:namespace/${namespace}/function/${_function}/${functionVersion}`),
 *         scfParams: {
 *             batchEventCount: 1,
 *             batchTimeout: 1,
 *             enableBatchDelivery: true,
 *         },
 *     },
 * });
 * ```
 * ### Create an event target of type ckafka
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 *
 * const config = new pulumi.Config();
 * const ckafka = config.get("ckafka") || "ckafka-qzoeaqx8";
 * const ckafkaTarget = new tencentcloud.eb.EventTarget("ckafkaTarget", {
 *     eventBusId: tencentcloud_eb_event_bus.foo.id,
 *     ruleId: tencentcloud_eb_event_rule.foo.rule_id,
 *     type: "ckafka",
 *     targetDescription: {
 *         resourceDescription: `qcs::scf:${_var.zone}:uin/${data.tencentcloud_cam_users.foo.user_list[0].uin}:ckafkaId/uin/${data.tencentcloud_cam_users.foo.user_list[0].uin}/${ckafka}`,
 *         ckafkaTargetParams: {
 *             topicName: "dasdasd",
 *             retryPolicy: {
 *                 maxRetryAttempts: 360,
 *                 retryInterval: 60,
 *             },
 *         },
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * eb event_target can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Eb/eventTarget:EventTarget event_target event_target_id
 * ```
 */
export class EventTarget extends pulumi.CustomResource {
    /**
     * Get an existing EventTarget resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EventTargetState, opts?: pulumi.CustomResourceOptions): EventTarget {
        return new EventTarget(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Eb/eventTarget:EventTarget';

    /**
     * Returns true if the given object is an instance of EventTarget.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EventTarget {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EventTarget.__pulumiType;
    }

    /**
     * event bus id.
     */
    public readonly eventBusId!: pulumi.Output<string>;
    /**
     * event rule id.
     */
    public readonly ruleId!: pulumi.Output<string>;
    /**
     * target description.
     */
    public readonly targetDescription!: pulumi.Output<outputs.Eb.EventTargetTargetDescription>;
    /**
     * target type.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a EventTarget resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EventTargetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EventTargetArgs | EventTargetState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EventTargetState | undefined;
            resourceInputs["eventBusId"] = state ? state.eventBusId : undefined;
            resourceInputs["ruleId"] = state ? state.ruleId : undefined;
            resourceInputs["targetDescription"] = state ? state.targetDescription : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as EventTargetArgs | undefined;
            if ((!args || args.eventBusId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'eventBusId'");
            }
            if ((!args || args.ruleId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ruleId'");
            }
            if ((!args || args.targetDescription === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetDescription'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["eventBusId"] = args ? args.eventBusId : undefined;
            resourceInputs["ruleId"] = args ? args.ruleId : undefined;
            resourceInputs["targetDescription"] = args ? args.targetDescription : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EventTarget.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EventTarget resources.
 */
export interface EventTargetState {
    /**
     * event bus id.
     */
    eventBusId?: pulumi.Input<string>;
    /**
     * event rule id.
     */
    ruleId?: pulumi.Input<string>;
    /**
     * target description.
     */
    targetDescription?: pulumi.Input<inputs.Eb.EventTargetTargetDescription>;
    /**
     * target type.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EventTarget resource.
 */
export interface EventTargetArgs {
    /**
     * event bus id.
     */
    eventBusId: pulumi.Input<string>;
    /**
     * event rule id.
     */
    ruleId: pulumi.Input<string>;
    /**
     * target description.
     */
    targetDescription: pulumi.Input<inputs.Eb.EventTargetTargetDescription>;
    /**
     * target type.
     */
    type: pulumi.Input<string>;
}
