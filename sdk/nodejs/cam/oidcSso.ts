// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a CAM-OIDC-SSO.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const foo = new tencentcloud.cam.OidcSso("foo", {
 *     authorizationEndpoint: "https://login.microsoftonline.com/.../oauth2/v2.0/authorize",
 *     clientId: "...",
 *     identityKey: "...",
 *     identityUrl: "https://login.microsoftonline.com/.../v2.0",
 *     mappingFiled: "name",
 *     responseMode: "form_post",
 *     responseType: "id_token",
 *     scopes: [
 *         "openid",
 *         "email",
 *     ],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * CAM-OIDC-SSO can be imported using the client_id or any string which can identifier resource, e.g.
 *
 * ```sh
 * $ pulumi import tencentcloud:Cam/oidcSso:OidcSso foo xxxxxxxxxxx
 * ```
 */
export class OidcSso extends pulumi.CustomResource {
    /**
     * Get an existing OidcSso resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OidcSsoState, opts?: pulumi.CustomResourceOptions): OidcSso {
        return new OidcSso(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Cam/oidcSso:OidcSso';

    /**
     * Returns true if the given object is an instance of OidcSso.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OidcSso {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OidcSso.__pulumiType;
    }

    /**
     * Authorization request Endpoint, OpenID Connect identity provider authorization address. Corresponds to the value of the `authorizationEndpoint` field in the Openid-configuration provided by the Enterprise IdP.
     */
    public readonly authorizationEndpoint!: pulumi.Output<string>;
    /**
     * Client ID, the client ID registered with the OpenID Connect identity provider.
     */
    public readonly clientId!: pulumi.Output<string>;
    /**
     * The signature public key requires base64_encode. Verify the public key signed by the OpenID Connect identity provider ID Token. For the security of your account, we recommend that you rotate the signed public key regularly.
     */
    public readonly identityKey!: pulumi.Output<string>;
    /**
     * Identity provider URL. OpenID Connect identity provider identity.Corresponds to the value of the `issuer` field in the Openid-configuration provided by the Enterprise IdP.
     */
    public readonly identityUrl!: pulumi.Output<string>;
    /**
     * Map field names. Which field in the IdP's idToken maps to the user name of the subuser, usually the sub or name field.
     */
    public readonly mappingFiled!: pulumi.Output<string>;
    /**
     * Authorize the request Forsonse mode. Authorization request return mode, formPost and frogment two optional modes, recommended to select formPost mode.
     */
    public readonly responseMode!: pulumi.Output<string>;
    /**
     * Authorization requests The Response type, with a fixed value id_token.
     */
    public readonly responseType!: pulumi.Output<string>;
    /**
     * Authorize the request Scope. openid; email; profile; Authorization request information scope. The default is required openid.
     */
    public readonly scopes!: pulumi.Output<string[] | undefined>;

    /**
     * Create a OidcSso resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OidcSsoArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OidcSsoArgs | OidcSsoState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OidcSsoState | undefined;
            resourceInputs["authorizationEndpoint"] = state ? state.authorizationEndpoint : undefined;
            resourceInputs["clientId"] = state ? state.clientId : undefined;
            resourceInputs["identityKey"] = state ? state.identityKey : undefined;
            resourceInputs["identityUrl"] = state ? state.identityUrl : undefined;
            resourceInputs["mappingFiled"] = state ? state.mappingFiled : undefined;
            resourceInputs["responseMode"] = state ? state.responseMode : undefined;
            resourceInputs["responseType"] = state ? state.responseType : undefined;
            resourceInputs["scopes"] = state ? state.scopes : undefined;
        } else {
            const args = argsOrState as OidcSsoArgs | undefined;
            if ((!args || args.authorizationEndpoint === undefined) && !opts.urn) {
                throw new Error("Missing required property 'authorizationEndpoint'");
            }
            if ((!args || args.clientId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientId'");
            }
            if ((!args || args.identityKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'identityKey'");
            }
            if ((!args || args.identityUrl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'identityUrl'");
            }
            if ((!args || args.mappingFiled === undefined) && !opts.urn) {
                throw new Error("Missing required property 'mappingFiled'");
            }
            if ((!args || args.responseMode === undefined) && !opts.urn) {
                throw new Error("Missing required property 'responseMode'");
            }
            if ((!args || args.responseType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'responseType'");
            }
            resourceInputs["authorizationEndpoint"] = args ? args.authorizationEndpoint : undefined;
            resourceInputs["clientId"] = args ? args.clientId : undefined;
            resourceInputs["identityKey"] = args ? args.identityKey : undefined;
            resourceInputs["identityUrl"] = args ? args.identityUrl : undefined;
            resourceInputs["mappingFiled"] = args ? args.mappingFiled : undefined;
            resourceInputs["responseMode"] = args ? args.responseMode : undefined;
            resourceInputs["responseType"] = args ? args.responseType : undefined;
            resourceInputs["scopes"] = args ? args.scopes : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(OidcSso.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OidcSso resources.
 */
export interface OidcSsoState {
    /**
     * Authorization request Endpoint, OpenID Connect identity provider authorization address. Corresponds to the value of the `authorizationEndpoint` field in the Openid-configuration provided by the Enterprise IdP.
     */
    authorizationEndpoint?: pulumi.Input<string>;
    /**
     * Client ID, the client ID registered with the OpenID Connect identity provider.
     */
    clientId?: pulumi.Input<string>;
    /**
     * The signature public key requires base64_encode. Verify the public key signed by the OpenID Connect identity provider ID Token. For the security of your account, we recommend that you rotate the signed public key regularly.
     */
    identityKey?: pulumi.Input<string>;
    /**
     * Identity provider URL. OpenID Connect identity provider identity.Corresponds to the value of the `issuer` field in the Openid-configuration provided by the Enterprise IdP.
     */
    identityUrl?: pulumi.Input<string>;
    /**
     * Map field names. Which field in the IdP's idToken maps to the user name of the subuser, usually the sub or name field.
     */
    mappingFiled?: pulumi.Input<string>;
    /**
     * Authorize the request Forsonse mode. Authorization request return mode, formPost and frogment two optional modes, recommended to select formPost mode.
     */
    responseMode?: pulumi.Input<string>;
    /**
     * Authorization requests The Response type, with a fixed value id_token.
     */
    responseType?: pulumi.Input<string>;
    /**
     * Authorize the request Scope. openid; email; profile; Authorization request information scope. The default is required openid.
     */
    scopes?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a OidcSso resource.
 */
export interface OidcSsoArgs {
    /**
     * Authorization request Endpoint, OpenID Connect identity provider authorization address. Corresponds to the value of the `authorizationEndpoint` field in the Openid-configuration provided by the Enterprise IdP.
     */
    authorizationEndpoint: pulumi.Input<string>;
    /**
     * Client ID, the client ID registered with the OpenID Connect identity provider.
     */
    clientId: pulumi.Input<string>;
    /**
     * The signature public key requires base64_encode. Verify the public key signed by the OpenID Connect identity provider ID Token. For the security of your account, we recommend that you rotate the signed public key regularly.
     */
    identityKey: pulumi.Input<string>;
    /**
     * Identity provider URL. OpenID Connect identity provider identity.Corresponds to the value of the `issuer` field in the Openid-configuration provided by the Enterprise IdP.
     */
    identityUrl: pulumi.Input<string>;
    /**
     * Map field names. Which field in the IdP's idToken maps to the user name of the subuser, usually the sub or name field.
     */
    mappingFiled: pulumi.Input<string>;
    /**
     * Authorize the request Forsonse mode. Authorization request return mode, formPost and frogment two optional modes, recommended to select formPost mode.
     */
    responseMode: pulumi.Input<string>;
    /**
     * Authorization requests The Response type, with a fixed value id_token.
     */
    responseType: pulumi.Input<string>;
    /**
     * Authorize the request Scope. openid; email; profile; Authorization request information scope. The default is required openid.
     */
    scopes?: pulumi.Input<pulumi.Input<string>[]>;
}
