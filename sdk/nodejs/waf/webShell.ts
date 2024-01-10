// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a waf webShell
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const example = new tencentcloud.Waf.WebShell("example", {
 *     domain: "demo.waf.com",
 *     status: 0,
 * });
 * ```
 *
 * ## Import
 *
 * waf web_shell can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Waf/webShell:WebShell example demo.waf.com
 * ```
 */
export class WebShell extends pulumi.CustomResource {
    /**
     * Get an existing WebShell resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WebShellState, opts?: pulumi.CustomResourceOptions): WebShell {
        return new WebShell(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Waf/webShell:WebShell';

    /**
     * Returns true if the given object is an instance of WebShell.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WebShell {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WebShell.__pulumiType;
    }

    /**
     * Domain.
     */
    public readonly domain!: pulumi.Output<string>;
    /**
     * Webshell status, 1: open; 0: closed; 2: log.
     */
    public readonly status!: pulumi.Output<number>;

    /**
     * Create a WebShell resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WebShellArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WebShellArgs | WebShellState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WebShellState | undefined;
            resourceInputs["domain"] = state ? state.domain : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as WebShellArgs | undefined;
            if ((!args || args.domain === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domain'");
            }
            if ((!args || args.status === undefined) && !opts.urn) {
                throw new Error("Missing required property 'status'");
            }
            resourceInputs["domain"] = args ? args.domain : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WebShell.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WebShell resources.
 */
export interface WebShellState {
    /**
     * Domain.
     */
    domain?: pulumi.Input<string>;
    /**
     * Webshell status, 1: open; 0: closed; 2: log.
     */
    status?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a WebShell resource.
 */
export interface WebShellArgs {
    /**
     * Domain.
     */
    domain: pulumi.Input<string>;
    /**
     * Webshell status, 1: open; 0: closed; 2: log.
     */
    status: pulumi.Input<number>;
}
