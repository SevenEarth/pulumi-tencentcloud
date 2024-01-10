// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a waf ccSession
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const example = new tencentcloud.Waf.CcSession("example", {
 *     category: "match",
 *     domain: "www.demo.com",
 *     edition: "sparta-waf",
 *     endMat: "&",
 *     endOffset: "-1",
 *     keyOrStartMat: "key_a=123",
 *     sessionName: "terraformDemo",
 *     source: "get",
 *     startOffset: "-1",
 * });
 * ```
 *
 * ## Import
 *
 * waf cc_session can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Waf/ccSession:CcSession example www.demo.com#sparta-waf#2000000253
 * ```
 */
export class CcSession extends pulumi.CustomResource {
    /**
     * Get an existing CcSession resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CcSessionState, opts?: pulumi.CustomResourceOptions): CcSession {
        return new CcSession(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Waf/ccSession:CcSession';

    /**
     * Returns true if the given object is an instance of CcSession.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CcSession {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CcSession.__pulumiType;
    }

    /**
     * Session match pattern, Optional patterns are match, location.
     */
    public readonly category!: pulumi.Output<string>;
    /**
     * Domain.
     */
    public readonly domain!: pulumi.Output<string>;
    /**
     * Waf edition. clb-waf means clb-waf, sparta-waf means saas-waf.
     */
    public readonly edition!: pulumi.Output<string>;
    /**
     * Session end identifier, when Category is match.
     */
    public readonly endMat!: pulumi.Output<string>;
    /**
     * End offset position, when Category is location.
     */
    public readonly endOffset!: pulumi.Output<string>;
    /**
     * Session identifier.
     */
    public readonly keyOrStartMat!: pulumi.Output<string>;
    /**
     * Session ID.
     */
    public /*out*/ readonly sessionId!: pulumi.Output<number>;
    /**
     * Session Name.
     */
    public readonly sessionName!: pulumi.Output<string>;
    /**
     * Session matching position, Optional locations are get, post, header, cookie.
     */
    public readonly source!: pulumi.Output<string>;
    /**
     * Starting offset position, when Category is location.
     */
    public readonly startOffset!: pulumi.Output<string>;

    /**
     * Create a CcSession resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CcSessionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CcSessionArgs | CcSessionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CcSessionState | undefined;
            resourceInputs["category"] = state ? state.category : undefined;
            resourceInputs["domain"] = state ? state.domain : undefined;
            resourceInputs["edition"] = state ? state.edition : undefined;
            resourceInputs["endMat"] = state ? state.endMat : undefined;
            resourceInputs["endOffset"] = state ? state.endOffset : undefined;
            resourceInputs["keyOrStartMat"] = state ? state.keyOrStartMat : undefined;
            resourceInputs["sessionId"] = state ? state.sessionId : undefined;
            resourceInputs["sessionName"] = state ? state.sessionName : undefined;
            resourceInputs["source"] = state ? state.source : undefined;
            resourceInputs["startOffset"] = state ? state.startOffset : undefined;
        } else {
            const args = argsOrState as CcSessionArgs | undefined;
            if ((!args || args.category === undefined) && !opts.urn) {
                throw new Error("Missing required property 'category'");
            }
            if ((!args || args.domain === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domain'");
            }
            if ((!args || args.edition === undefined) && !opts.urn) {
                throw new Error("Missing required property 'edition'");
            }
            if ((!args || args.endMat === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endMat'");
            }
            if ((!args || args.endOffset === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endOffset'");
            }
            if ((!args || args.keyOrStartMat === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keyOrStartMat'");
            }
            if ((!args || args.sessionName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sessionName'");
            }
            if ((!args || args.source === undefined) && !opts.urn) {
                throw new Error("Missing required property 'source'");
            }
            if ((!args || args.startOffset === undefined) && !opts.urn) {
                throw new Error("Missing required property 'startOffset'");
            }
            resourceInputs["category"] = args ? args.category : undefined;
            resourceInputs["domain"] = args ? args.domain : undefined;
            resourceInputs["edition"] = args ? args.edition : undefined;
            resourceInputs["endMat"] = args ? args.endMat : undefined;
            resourceInputs["endOffset"] = args ? args.endOffset : undefined;
            resourceInputs["keyOrStartMat"] = args ? args.keyOrStartMat : undefined;
            resourceInputs["sessionName"] = args ? args.sessionName : undefined;
            resourceInputs["source"] = args ? args.source : undefined;
            resourceInputs["startOffset"] = args ? args.startOffset : undefined;
            resourceInputs["sessionId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CcSession.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CcSession resources.
 */
export interface CcSessionState {
    /**
     * Session match pattern, Optional patterns are match, location.
     */
    category?: pulumi.Input<string>;
    /**
     * Domain.
     */
    domain?: pulumi.Input<string>;
    /**
     * Waf edition. clb-waf means clb-waf, sparta-waf means saas-waf.
     */
    edition?: pulumi.Input<string>;
    /**
     * Session end identifier, when Category is match.
     */
    endMat?: pulumi.Input<string>;
    /**
     * End offset position, when Category is location.
     */
    endOffset?: pulumi.Input<string>;
    /**
     * Session identifier.
     */
    keyOrStartMat?: pulumi.Input<string>;
    /**
     * Session ID.
     */
    sessionId?: pulumi.Input<number>;
    /**
     * Session Name.
     */
    sessionName?: pulumi.Input<string>;
    /**
     * Session matching position, Optional locations are get, post, header, cookie.
     */
    source?: pulumi.Input<string>;
    /**
     * Starting offset position, when Category is location.
     */
    startOffset?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CcSession resource.
 */
export interface CcSessionArgs {
    /**
     * Session match pattern, Optional patterns are match, location.
     */
    category: pulumi.Input<string>;
    /**
     * Domain.
     */
    domain: pulumi.Input<string>;
    /**
     * Waf edition. clb-waf means clb-waf, sparta-waf means saas-waf.
     */
    edition: pulumi.Input<string>;
    /**
     * Session end identifier, when Category is match.
     */
    endMat: pulumi.Input<string>;
    /**
     * End offset position, when Category is location.
     */
    endOffset: pulumi.Input<string>;
    /**
     * Session identifier.
     */
    keyOrStartMat: pulumi.Input<string>;
    /**
     * Session Name.
     */
    sessionName: pulumi.Input<string>;
    /**
     * Session matching position, Optional locations are get, post, header, cookie.
     */
    source: pulumi.Input<string>;
    /**
     * Starting offset position, when Category is location.
     */
    startOffset: pulumi.Input<string>;
}
