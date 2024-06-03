// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a vpc vpnGatewaySslClientCert
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const vpnGatewaySslClientCert = new tencentcloud.vpn.GatewaySslClientCert("vpnGatewaySslClientCert", {
 *     sslVpnClientId: "vpnc-123456",
 *     "switch": "off",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * vpc vpn_gateway_ssl_client_cert can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import tencentcloud:Vpn/gatewaySslClientCert:GatewaySslClientCert vpn_gateway_ssl_client_cert ssl_client_id
 * ```
 */
export class GatewaySslClientCert extends pulumi.CustomResource {
    /**
     * Get an existing GatewaySslClientCert resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GatewaySslClientCertState, opts?: pulumi.CustomResourceOptions): GatewaySslClientCert {
        return new GatewaySslClientCert(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Vpn/gatewaySslClientCert:GatewaySslClientCert';

    /**
     * Returns true if the given object is an instance of GatewaySslClientCert.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GatewaySslClientCert {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GatewaySslClientCert.__pulumiType;
    }

    /**
     * SSL-VPN-CLIENT Instance ID.
     */
    public readonly sslVpnClientId!: pulumi.Output<string>;
    /**
     * `on`: Enable, `off`: Disable.
     */
    public readonly switch!: pulumi.Output<string | undefined>;

    /**
     * Create a GatewaySslClientCert resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GatewaySslClientCertArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GatewaySslClientCertArgs | GatewaySslClientCertState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GatewaySslClientCertState | undefined;
            resourceInputs["sslVpnClientId"] = state ? state.sslVpnClientId : undefined;
            resourceInputs["switch"] = state ? state.switch : undefined;
        } else {
            const args = argsOrState as GatewaySslClientCertArgs | undefined;
            if ((!args || args.sslVpnClientId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sslVpnClientId'");
            }
            resourceInputs["sslVpnClientId"] = args ? args.sslVpnClientId : undefined;
            resourceInputs["switch"] = args ? args.switch : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GatewaySslClientCert.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GatewaySslClientCert resources.
 */
export interface GatewaySslClientCertState {
    /**
     * SSL-VPN-CLIENT Instance ID.
     */
    sslVpnClientId?: pulumi.Input<string>;
    /**
     * `on`: Enable, `off`: Disable.
     */
    switch?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GatewaySslClientCert resource.
 */
export interface GatewaySslClientCertArgs {
    /**
     * SSL-VPN-CLIENT Instance ID.
     */
    sslVpnClientId: pulumi.Input<string>;
    /**
     * `on`: Enable, `off`: Disable.
     */
    switch?: pulumi.Input<string>;
}
