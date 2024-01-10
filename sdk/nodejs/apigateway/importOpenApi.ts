// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a apiGateway importOpenApi
 *
 * ## Example Usage
 * ### Import open Api by YAML
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const example = new tencentcloud.ApiGateway.ImportOpenApi("example", {
 *     content: `info:
 *   title: keep-service
 *   version: 1.0.1
 * openapi: 3.0.0
 * paths:
 *   /api/test:
 *     get:
 *       description: desc
 *       operationId: test
 *       responses:
 *         '200':
 *           content:
 *             text/html:
 *               example: '200'
 *           description: '200'
 *         default:
 *           content:
 *             text/html:
 *               example: '400'
 *           description: '400'
 *       x-apigw-api-business-type: NORMAL
 *       x-apigw-api-type: NORMAL
 *       x-apigw-backend:
 *         ServiceConfig:
 *           Method: GET
 *           Path: /test
 *           Url: http://domain.com
 *         ServiceType: HTTP
 *       x-apigw-cors: false
 *       x-apigw-protocol: HTTP
 *       x-apigw-service-timeout: 15
 * `,
 *     contentVersion: "openAPI",
 *     encodeType: "YAML",
 *     serviceId: "service-nxz6yync",
 * });
 * ```
 * ### Import open Api by JSON
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const example = new tencentcloud.ApiGateway.ImportOpenApi("example", {
 *     content: "{\"openapi\": \"3.0.0\", \"info\": {\"title\": \"keep-service\", \"version\": \"1.0.1\"}, \"paths\": {\"/api/test\": {\"get\": {\"operationId\": \"test\", \"description\": \"desc\", \"responses\": {\"200\": {\"description\": \"200\", \"content\": {\"text/html\": {\"example\": \"200\"}}}, \"default\": {\"content\": {\"text/html\": {\"example\": \"400\"}}, \"description\": \"400\"}}, \"x-apigw-api-type\": \"NORMAL\", \"x-apigw-api-business-type\": \"NORMAL\", \"x-apigw-protocol\": \"HTTP\", \"x-apigw-cors\": false, \"x-apigw-service-timeout\": 15, \"x-apigw-backend\": {\"ServiceType\": \"HTTP\", \"ServiceConfig\": {\"Url\": \"http://domain.com\", \"Path\": \"/test\", \"Method\": \"GET\"}}}}}}",
 *     contentVersion: "openAPI",
 *     encodeType: "JSON",
 *     serviceId: "service-nxz6yync",
 * });
 * ```
 */
export class ImportOpenApi extends pulumi.CustomResource {
    /**
     * Get an existing ImportOpenApi resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ImportOpenApiState, opts?: pulumi.CustomResourceOptions): ImportOpenApi {
        return new ImportOpenApi(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:ApiGateway/importOpenApi:ImportOpenApi';

    /**
     * Returns true if the given object is an instance of ImportOpenApi.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ImportOpenApi {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ImportOpenApi.__pulumiType;
    }

    /**
     * When `authType` is OAUTH, this field is valid, NORMAL: Business API, OAUTH: Authorization API.
     */
    public /*out*/ readonly apiBusinessType!: pulumi.Output<string>;
    /**
     * Custom API description.
     */
    public /*out*/ readonly apiDesc!: pulumi.Output<string>;
    /**
     * Custom Api Id.
     */
    public /*out*/ readonly apiId!: pulumi.Output<string>;
    /**
     * Custom API name.
     */
    public /*out*/ readonly apiName!: pulumi.Output<string>;
    /**
     * API type, supports NORMAL (regular API) and TSF (microservice API), defaults to NORMAL.
     */
    public /*out*/ readonly apiType!: pulumi.Output<string>;
    /**
     * The unique ID of the associated authorization API takes effect when AuthType is OAUTH and ApiBusinessType is NORMAL. The unique ID of the oauth2.0 authorized API that identifies the business API binding.
     */
    public /*out*/ readonly authRelationApiId!: pulumi.Output<string>;
    /**
     * API authentication type. Support SECRET (Key Pair Authentication), NONE (Authentication Exemption), OAUTH, APP (Application Authentication). The default is NONE.
     */
    public /*out*/ readonly authType!: pulumi.Output<string>;
    /**
     * Constant parameter.
     */
    public /*out*/ readonly constantParameters!: pulumi.Output<outputs.ApiGateway.ImportOpenApiConstantParameter[]>;
    /**
     * OpenAPI body content.
     */
    public readonly content!: pulumi.Output<string>;
    /**
     * The Content version defaults to OpenAPI and currently only supports OpenAPI.
     */
    public readonly contentVersion!: pulumi.Output<string | undefined>;
    /**
     * Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Whether to enable CORS. Default value: `true`.
     */
    public /*out*/ readonly enableCors!: pulumi.Output<boolean>;
    /**
     * The Content format can only be YAML or JSON, and the default is YAML.
     */
    public readonly encodeType!: pulumi.Output<string | undefined>;
    /**
     * Whether to enable Base64 encoding will only take effect when the backend is scf.
     */
    public /*out*/ readonly isBase64Encoded!: pulumi.Output<boolean>;
    /**
     * Charge after starting debugging. (Cloud Market Reserved Fields).
     */
    public /*out*/ readonly isDebugAfterCharge!: pulumi.Output<boolean>;
    /**
     * Do you want to delete the custom response configuration error code? If it is not passed or False is passed, it will not be deleted. If True is passed, all custom response configuration error codes for this API will be deleted.
     */
    public /*out*/ readonly isDeleteResponseErrorCodes!: pulumi.Output<boolean>;
    /**
     * API bound microservice list.
     */
    public /*out*/ readonly microServices!: pulumi.Output<outputs.ApiGateway.ImportOpenApiMicroService[]>;
    /**
     * OAuth configuration. Effective when AuthType is OAUTH.
     */
    public /*out*/ readonly oauthConfigs!: pulumi.Output<outputs.ApiGateway.ImportOpenApiOauthConfig[]>;
    /**
     * API frontend request type. Valid values: `HTTP`, `WEBSOCKET`. Default value: `HTTP`.
     */
    public /*out*/ readonly protocol!: pulumi.Output<string>;
    /**
     * Request frontend method configuration. Valid values: `GET`,`POST`,`PUT`,`DELETE`,`HEAD`,`ANY`. Default value: `GET`.
     */
    public /*out*/ readonly requestConfigMethod!: pulumi.Output<string>;
    /**
     * Request frontend path configuration. Like `/user/getinfo`.
     */
    public /*out*/ readonly requestConfigPath!: pulumi.Output<string>;
    /**
     * Frontend request parameters.
     */
    public /*out*/ readonly requestParameters!: pulumi.Output<outputs.ApiGateway.ImportOpenApiRequestParameter[]>;
    /**
     * Custom error code configuration. Must keep at least one after set.
     */
    public /*out*/ readonly responseErrorCodes!: pulumi.Output<outputs.ApiGateway.ImportOpenApiResponseErrorCode[]>;
    /**
     * Response failure sample of custom response configuration.
     */
    public /*out*/ readonly responseFailExample!: pulumi.Output<string>;
    /**
     * Successful response sample of custom response configuration.
     */
    public /*out*/ readonly responseSuccessExample!: pulumi.Output<string>;
    /**
     * Return type. Valid values: `HTML`, `JSON`, `TEXT`, `BINARY`, `XML`. Default value: `HTML`.
     */
    public /*out*/ readonly responseType!: pulumi.Output<string>;
    /**
     * API backend COS configuration. If ServiceType is COS, then this parameter must be passed.Note: This field may return null, indicating that a valid value cannot be obtained.
     */
    public /*out*/ readonly serviceConfigCosConfigs!: pulumi.Output<outputs.ApiGateway.ImportOpenApiServiceConfigCosConfig[]>;
    /**
     * API backend service request method, such as `GET`. If `serviceConfigType` is `HTTP`, this parameter will be required. The frontend `requestConfigMethod` and backend method `serviceConfigMethod` can be different.
     */
    public /*out*/ readonly serviceConfigMethod!: pulumi.Output<string>;
    /**
     * Returned information of API backend mocking. This parameter is required when `serviceConfigType` is `MOCK`.
     */
    public /*out*/ readonly serviceConfigMockReturnMessage!: pulumi.Output<string>;
    /**
     * API backend service path, such as /path. If `serviceConfigType` is `HTTP`, this parameter will be required. The frontend `requestConfigPath` and backend path `serviceConfigPath` can be different.
     */
    public /*out*/ readonly serviceConfigPath!: pulumi.Output<string>;
    /**
     * Backend type. Effective when enabling vpc, currently supported types are clb, cvm, and upstream.
     */
    public /*out*/ readonly serviceConfigProduct!: pulumi.Output<string>;
    /**
     * SCF function name. This parameter takes effect when `serviceConfigType` is `SCF`.
     */
    public /*out*/ readonly serviceConfigScfFunctionName!: pulumi.Output<string>;
    /**
     * SCF function namespace. This parameter takes effect when `serviceConfigType` is `SCF`.
     */
    public /*out*/ readonly serviceConfigScfFunctionNamespace!: pulumi.Output<string>;
    /**
     * SCF function version. This parameter takes effect when `serviceConfigType` is `SCF`.
     */
    public /*out*/ readonly serviceConfigScfFunctionQualifier!: pulumi.Output<string>;
    /**
     * Scf function type. Effective when the backend type is SCF. Support Event Triggering (EVENT) and HTTP Direct Cloud Function (HTTP).
     */
    public /*out*/ readonly serviceConfigScfFunctionType!: pulumi.Output<string>;
    /**
     * Whether to enable response integration. Effective when the backend type is SCF.
     */
    public /*out*/ readonly serviceConfigScfIsIntegratedResponse!: pulumi.Output<boolean>;
    /**
     * API backend service timeout period in seconds. Default value: `5`.
     */
    public /*out*/ readonly serviceConfigTimeout!: pulumi.Output<number>;
    /**
     * The backend service type of the API. Supports HTTP, MOCK, TSF, SCF, WEBSOCKET, COS, TARGET (internal testing).
     */
    public /*out*/ readonly serviceConfigType!: pulumi.Output<string>;
    /**
     * Only required when binding to VPC channelsNote: This field may return null, indicating that a valid value cannot be obtained.
     */
    public /*out*/ readonly serviceConfigUpstreamId!: pulumi.Output<string>;
    /**
     * The backend service URL of the API. If the ServiceType is HTTP, this parameter must be passed.
     */
    public /*out*/ readonly serviceConfigUrl!: pulumi.Output<string>;
    /**
     * Unique VPC ID.
     */
    public /*out*/ readonly serviceConfigVpcId!: pulumi.Output<string>;
    /**
     * Scf websocket cleaning function. It takes effect when the current end type is WEBSOCKET and the backend type is SCF.
     */
    public /*out*/ readonly serviceConfigWebsocketCleanupFunctionName!: pulumi.Output<string>;
    /**
     * Scf websocket cleans up the function namespace. It takes effect when the current end type is WEBSOCKET and the backend type is SCF.
     */
    public /*out*/ readonly serviceConfigWebsocketCleanupFunctionNamespace!: pulumi.Output<string>;
    /**
     * Scf websocket cleaning function version. It takes effect when the current end type is WEBSOCKET and the backend type is SCF.
     */
    public /*out*/ readonly serviceConfigWebsocketCleanupFunctionQualifier!: pulumi.Output<string>;
    /**
     * Scf websocket registration function. It takes effect when the current end type is WEBSOCKET and the backend type is SCF.
     */
    public /*out*/ readonly serviceConfigWebsocketRegisterFunctionName!: pulumi.Output<string>;
    /**
     * Scf websocket registers function namespaces. It takes effect when the current end type is WEBSOCKET and the backend type is SCF.
     */
    public /*out*/ readonly serviceConfigWebsocketRegisterFunctionNamespace!: pulumi.Output<string>;
    /**
     * Scf websocket transfer function version. It takes effect when the current end type is WEBSOCKET and the backend type is SCF.
     */
    public /*out*/ readonly serviceConfigWebsocketRegisterFunctionQualifier!: pulumi.Output<string>;
    /**
     * Scf websocket transfer function. It takes effect when the current end type is WEBSOCKET and the backend type is SCF.
     */
    public /*out*/ readonly serviceConfigWebsocketTransportFunctionName!: pulumi.Output<string>;
    /**
     * Scf websocket transfer function namespace. It takes effect when the current end type is WEBSOCKET and the backend type is SCF.
     */
    public /*out*/ readonly serviceConfigWebsocketTransportFunctionNamespace!: pulumi.Output<string>;
    /**
     * Scf websocket transfer function version. It takes effect when the current end type is WEBSOCKET and the backend type is SCF.
     */
    public /*out*/ readonly serviceConfigWebsocketTransportFunctionQualifier!: pulumi.Output<string>;
    /**
     * The unique ID of the service where the API is located.
     */
    public readonly serviceId!: pulumi.Output<string>;
    /**
     * The backend service parameters of the API.
     */
    public /*out*/ readonly serviceParameters!: pulumi.Output<outputs.ApiGateway.ImportOpenApiServiceParameter[]>;
    /**
     * Health check configuration for microservices.
     */
    public /*out*/ readonly serviceTsfHealthCheckConfs!: pulumi.Output<outputs.ApiGateway.ImportOpenApiServiceTsfHealthCheckConf[]>;
    /**
     * Load balancing configuration for microservices.
     */
    public /*out*/ readonly serviceTsfLoadBalanceConfs!: pulumi.Output<outputs.ApiGateway.ImportOpenApiServiceTsfLoadBalanceConf[]>;
    /**
     * Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a ImportOpenApi resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ImportOpenApiArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ImportOpenApiArgs | ImportOpenApiState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ImportOpenApiState | undefined;
            resourceInputs["apiBusinessType"] = state ? state.apiBusinessType : undefined;
            resourceInputs["apiDesc"] = state ? state.apiDesc : undefined;
            resourceInputs["apiId"] = state ? state.apiId : undefined;
            resourceInputs["apiName"] = state ? state.apiName : undefined;
            resourceInputs["apiType"] = state ? state.apiType : undefined;
            resourceInputs["authRelationApiId"] = state ? state.authRelationApiId : undefined;
            resourceInputs["authType"] = state ? state.authType : undefined;
            resourceInputs["constantParameters"] = state ? state.constantParameters : undefined;
            resourceInputs["content"] = state ? state.content : undefined;
            resourceInputs["contentVersion"] = state ? state.contentVersion : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["enableCors"] = state ? state.enableCors : undefined;
            resourceInputs["encodeType"] = state ? state.encodeType : undefined;
            resourceInputs["isBase64Encoded"] = state ? state.isBase64Encoded : undefined;
            resourceInputs["isDebugAfterCharge"] = state ? state.isDebugAfterCharge : undefined;
            resourceInputs["isDeleteResponseErrorCodes"] = state ? state.isDeleteResponseErrorCodes : undefined;
            resourceInputs["microServices"] = state ? state.microServices : undefined;
            resourceInputs["oauthConfigs"] = state ? state.oauthConfigs : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["requestConfigMethod"] = state ? state.requestConfigMethod : undefined;
            resourceInputs["requestConfigPath"] = state ? state.requestConfigPath : undefined;
            resourceInputs["requestParameters"] = state ? state.requestParameters : undefined;
            resourceInputs["responseErrorCodes"] = state ? state.responseErrorCodes : undefined;
            resourceInputs["responseFailExample"] = state ? state.responseFailExample : undefined;
            resourceInputs["responseSuccessExample"] = state ? state.responseSuccessExample : undefined;
            resourceInputs["responseType"] = state ? state.responseType : undefined;
            resourceInputs["serviceConfigCosConfigs"] = state ? state.serviceConfigCosConfigs : undefined;
            resourceInputs["serviceConfigMethod"] = state ? state.serviceConfigMethod : undefined;
            resourceInputs["serviceConfigMockReturnMessage"] = state ? state.serviceConfigMockReturnMessage : undefined;
            resourceInputs["serviceConfigPath"] = state ? state.serviceConfigPath : undefined;
            resourceInputs["serviceConfigProduct"] = state ? state.serviceConfigProduct : undefined;
            resourceInputs["serviceConfigScfFunctionName"] = state ? state.serviceConfigScfFunctionName : undefined;
            resourceInputs["serviceConfigScfFunctionNamespace"] = state ? state.serviceConfigScfFunctionNamespace : undefined;
            resourceInputs["serviceConfigScfFunctionQualifier"] = state ? state.serviceConfigScfFunctionQualifier : undefined;
            resourceInputs["serviceConfigScfFunctionType"] = state ? state.serviceConfigScfFunctionType : undefined;
            resourceInputs["serviceConfigScfIsIntegratedResponse"] = state ? state.serviceConfigScfIsIntegratedResponse : undefined;
            resourceInputs["serviceConfigTimeout"] = state ? state.serviceConfigTimeout : undefined;
            resourceInputs["serviceConfigType"] = state ? state.serviceConfigType : undefined;
            resourceInputs["serviceConfigUpstreamId"] = state ? state.serviceConfigUpstreamId : undefined;
            resourceInputs["serviceConfigUrl"] = state ? state.serviceConfigUrl : undefined;
            resourceInputs["serviceConfigVpcId"] = state ? state.serviceConfigVpcId : undefined;
            resourceInputs["serviceConfigWebsocketCleanupFunctionName"] = state ? state.serviceConfigWebsocketCleanupFunctionName : undefined;
            resourceInputs["serviceConfigWebsocketCleanupFunctionNamespace"] = state ? state.serviceConfigWebsocketCleanupFunctionNamespace : undefined;
            resourceInputs["serviceConfigWebsocketCleanupFunctionQualifier"] = state ? state.serviceConfigWebsocketCleanupFunctionQualifier : undefined;
            resourceInputs["serviceConfigWebsocketRegisterFunctionName"] = state ? state.serviceConfigWebsocketRegisterFunctionName : undefined;
            resourceInputs["serviceConfigWebsocketRegisterFunctionNamespace"] = state ? state.serviceConfigWebsocketRegisterFunctionNamespace : undefined;
            resourceInputs["serviceConfigWebsocketRegisterFunctionQualifier"] = state ? state.serviceConfigWebsocketRegisterFunctionQualifier : undefined;
            resourceInputs["serviceConfigWebsocketTransportFunctionName"] = state ? state.serviceConfigWebsocketTransportFunctionName : undefined;
            resourceInputs["serviceConfigWebsocketTransportFunctionNamespace"] = state ? state.serviceConfigWebsocketTransportFunctionNamespace : undefined;
            resourceInputs["serviceConfigWebsocketTransportFunctionQualifier"] = state ? state.serviceConfigWebsocketTransportFunctionQualifier : undefined;
            resourceInputs["serviceId"] = state ? state.serviceId : undefined;
            resourceInputs["serviceParameters"] = state ? state.serviceParameters : undefined;
            resourceInputs["serviceTsfHealthCheckConfs"] = state ? state.serviceTsfHealthCheckConfs : undefined;
            resourceInputs["serviceTsfLoadBalanceConfs"] = state ? state.serviceTsfLoadBalanceConfs : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
        } else {
            const args = argsOrState as ImportOpenApiArgs | undefined;
            if ((!args || args.content === undefined) && !opts.urn) {
                throw new Error("Missing required property 'content'");
            }
            if ((!args || args.serviceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceId'");
            }
            resourceInputs["content"] = args ? args.content : undefined;
            resourceInputs["contentVersion"] = args ? args.contentVersion : undefined;
            resourceInputs["encodeType"] = args ? args.encodeType : undefined;
            resourceInputs["serviceId"] = args ? args.serviceId : undefined;
            resourceInputs["apiBusinessType"] = undefined /*out*/;
            resourceInputs["apiDesc"] = undefined /*out*/;
            resourceInputs["apiId"] = undefined /*out*/;
            resourceInputs["apiName"] = undefined /*out*/;
            resourceInputs["apiType"] = undefined /*out*/;
            resourceInputs["authRelationApiId"] = undefined /*out*/;
            resourceInputs["authType"] = undefined /*out*/;
            resourceInputs["constantParameters"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["enableCors"] = undefined /*out*/;
            resourceInputs["isBase64Encoded"] = undefined /*out*/;
            resourceInputs["isDebugAfterCharge"] = undefined /*out*/;
            resourceInputs["isDeleteResponseErrorCodes"] = undefined /*out*/;
            resourceInputs["microServices"] = undefined /*out*/;
            resourceInputs["oauthConfigs"] = undefined /*out*/;
            resourceInputs["protocol"] = undefined /*out*/;
            resourceInputs["requestConfigMethod"] = undefined /*out*/;
            resourceInputs["requestConfigPath"] = undefined /*out*/;
            resourceInputs["requestParameters"] = undefined /*out*/;
            resourceInputs["responseErrorCodes"] = undefined /*out*/;
            resourceInputs["responseFailExample"] = undefined /*out*/;
            resourceInputs["responseSuccessExample"] = undefined /*out*/;
            resourceInputs["responseType"] = undefined /*out*/;
            resourceInputs["serviceConfigCosConfigs"] = undefined /*out*/;
            resourceInputs["serviceConfigMethod"] = undefined /*out*/;
            resourceInputs["serviceConfigMockReturnMessage"] = undefined /*out*/;
            resourceInputs["serviceConfigPath"] = undefined /*out*/;
            resourceInputs["serviceConfigProduct"] = undefined /*out*/;
            resourceInputs["serviceConfigScfFunctionName"] = undefined /*out*/;
            resourceInputs["serviceConfigScfFunctionNamespace"] = undefined /*out*/;
            resourceInputs["serviceConfigScfFunctionQualifier"] = undefined /*out*/;
            resourceInputs["serviceConfigScfFunctionType"] = undefined /*out*/;
            resourceInputs["serviceConfigScfIsIntegratedResponse"] = undefined /*out*/;
            resourceInputs["serviceConfigTimeout"] = undefined /*out*/;
            resourceInputs["serviceConfigType"] = undefined /*out*/;
            resourceInputs["serviceConfigUpstreamId"] = undefined /*out*/;
            resourceInputs["serviceConfigUrl"] = undefined /*out*/;
            resourceInputs["serviceConfigVpcId"] = undefined /*out*/;
            resourceInputs["serviceConfigWebsocketCleanupFunctionName"] = undefined /*out*/;
            resourceInputs["serviceConfigWebsocketCleanupFunctionNamespace"] = undefined /*out*/;
            resourceInputs["serviceConfigWebsocketCleanupFunctionQualifier"] = undefined /*out*/;
            resourceInputs["serviceConfigWebsocketRegisterFunctionName"] = undefined /*out*/;
            resourceInputs["serviceConfigWebsocketRegisterFunctionNamespace"] = undefined /*out*/;
            resourceInputs["serviceConfigWebsocketRegisterFunctionQualifier"] = undefined /*out*/;
            resourceInputs["serviceConfigWebsocketTransportFunctionName"] = undefined /*out*/;
            resourceInputs["serviceConfigWebsocketTransportFunctionNamespace"] = undefined /*out*/;
            resourceInputs["serviceConfigWebsocketTransportFunctionQualifier"] = undefined /*out*/;
            resourceInputs["serviceParameters"] = undefined /*out*/;
            resourceInputs["serviceTsfHealthCheckConfs"] = undefined /*out*/;
            resourceInputs["serviceTsfLoadBalanceConfs"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ImportOpenApi.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ImportOpenApi resources.
 */
export interface ImportOpenApiState {
    /**
     * When `authType` is OAUTH, this field is valid, NORMAL: Business API, OAUTH: Authorization API.
     */
    apiBusinessType?: pulumi.Input<string>;
    /**
     * Custom API description.
     */
    apiDesc?: pulumi.Input<string>;
    /**
     * Custom Api Id.
     */
    apiId?: pulumi.Input<string>;
    /**
     * Custom API name.
     */
    apiName?: pulumi.Input<string>;
    /**
     * API type, supports NORMAL (regular API) and TSF (microservice API), defaults to NORMAL.
     */
    apiType?: pulumi.Input<string>;
    /**
     * The unique ID of the associated authorization API takes effect when AuthType is OAUTH and ApiBusinessType is NORMAL. The unique ID of the oauth2.0 authorized API that identifies the business API binding.
     */
    authRelationApiId?: pulumi.Input<string>;
    /**
     * API authentication type. Support SECRET (Key Pair Authentication), NONE (Authentication Exemption), OAUTH, APP (Application Authentication). The default is NONE.
     */
    authType?: pulumi.Input<string>;
    /**
     * Constant parameter.
     */
    constantParameters?: pulumi.Input<pulumi.Input<inputs.ApiGateway.ImportOpenApiConstantParameter>[]>;
    /**
     * OpenAPI body content.
     */
    content?: pulumi.Input<string>;
    /**
     * The Content version defaults to OpenAPI and currently only supports OpenAPI.
     */
    contentVersion?: pulumi.Input<string>;
    /**
     * Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Whether to enable CORS. Default value: `true`.
     */
    enableCors?: pulumi.Input<boolean>;
    /**
     * The Content format can only be YAML or JSON, and the default is YAML.
     */
    encodeType?: pulumi.Input<string>;
    /**
     * Whether to enable Base64 encoding will only take effect when the backend is scf.
     */
    isBase64Encoded?: pulumi.Input<boolean>;
    /**
     * Charge after starting debugging. (Cloud Market Reserved Fields).
     */
    isDebugAfterCharge?: pulumi.Input<boolean>;
    /**
     * Do you want to delete the custom response configuration error code? If it is not passed or False is passed, it will not be deleted. If True is passed, all custom response configuration error codes for this API will be deleted.
     */
    isDeleteResponseErrorCodes?: pulumi.Input<boolean>;
    /**
     * API bound microservice list.
     */
    microServices?: pulumi.Input<pulumi.Input<inputs.ApiGateway.ImportOpenApiMicroService>[]>;
    /**
     * OAuth configuration. Effective when AuthType is OAUTH.
     */
    oauthConfigs?: pulumi.Input<pulumi.Input<inputs.ApiGateway.ImportOpenApiOauthConfig>[]>;
    /**
     * API frontend request type. Valid values: `HTTP`, `WEBSOCKET`. Default value: `HTTP`.
     */
    protocol?: pulumi.Input<string>;
    /**
     * Request frontend method configuration. Valid values: `GET`,`POST`,`PUT`,`DELETE`,`HEAD`,`ANY`. Default value: `GET`.
     */
    requestConfigMethod?: pulumi.Input<string>;
    /**
     * Request frontend path configuration. Like `/user/getinfo`.
     */
    requestConfigPath?: pulumi.Input<string>;
    /**
     * Frontend request parameters.
     */
    requestParameters?: pulumi.Input<pulumi.Input<inputs.ApiGateway.ImportOpenApiRequestParameter>[]>;
    /**
     * Custom error code configuration. Must keep at least one after set.
     */
    responseErrorCodes?: pulumi.Input<pulumi.Input<inputs.ApiGateway.ImportOpenApiResponseErrorCode>[]>;
    /**
     * Response failure sample of custom response configuration.
     */
    responseFailExample?: pulumi.Input<string>;
    /**
     * Successful response sample of custom response configuration.
     */
    responseSuccessExample?: pulumi.Input<string>;
    /**
     * Return type. Valid values: `HTML`, `JSON`, `TEXT`, `BINARY`, `XML`. Default value: `HTML`.
     */
    responseType?: pulumi.Input<string>;
    /**
     * API backend COS configuration. If ServiceType is COS, then this parameter must be passed.Note: This field may return null, indicating that a valid value cannot be obtained.
     */
    serviceConfigCosConfigs?: pulumi.Input<pulumi.Input<inputs.ApiGateway.ImportOpenApiServiceConfigCosConfig>[]>;
    /**
     * API backend service request method, such as `GET`. If `serviceConfigType` is `HTTP`, this parameter will be required. The frontend `requestConfigMethod` and backend method `serviceConfigMethod` can be different.
     */
    serviceConfigMethod?: pulumi.Input<string>;
    /**
     * Returned information of API backend mocking. This parameter is required when `serviceConfigType` is `MOCK`.
     */
    serviceConfigMockReturnMessage?: pulumi.Input<string>;
    /**
     * API backend service path, such as /path. If `serviceConfigType` is `HTTP`, this parameter will be required. The frontend `requestConfigPath` and backend path `serviceConfigPath` can be different.
     */
    serviceConfigPath?: pulumi.Input<string>;
    /**
     * Backend type. Effective when enabling vpc, currently supported types are clb, cvm, and upstream.
     */
    serviceConfigProduct?: pulumi.Input<string>;
    /**
     * SCF function name. This parameter takes effect when `serviceConfigType` is `SCF`.
     */
    serviceConfigScfFunctionName?: pulumi.Input<string>;
    /**
     * SCF function namespace. This parameter takes effect when `serviceConfigType` is `SCF`.
     */
    serviceConfigScfFunctionNamespace?: pulumi.Input<string>;
    /**
     * SCF function version. This parameter takes effect when `serviceConfigType` is `SCF`.
     */
    serviceConfigScfFunctionQualifier?: pulumi.Input<string>;
    /**
     * Scf function type. Effective when the backend type is SCF. Support Event Triggering (EVENT) and HTTP Direct Cloud Function (HTTP).
     */
    serviceConfigScfFunctionType?: pulumi.Input<string>;
    /**
     * Whether to enable response integration. Effective when the backend type is SCF.
     */
    serviceConfigScfIsIntegratedResponse?: pulumi.Input<boolean>;
    /**
     * API backend service timeout period in seconds. Default value: `5`.
     */
    serviceConfigTimeout?: pulumi.Input<number>;
    /**
     * The backend service type of the API. Supports HTTP, MOCK, TSF, SCF, WEBSOCKET, COS, TARGET (internal testing).
     */
    serviceConfigType?: pulumi.Input<string>;
    /**
     * Only required when binding to VPC channelsNote: This field may return null, indicating that a valid value cannot be obtained.
     */
    serviceConfigUpstreamId?: pulumi.Input<string>;
    /**
     * The backend service URL of the API. If the ServiceType is HTTP, this parameter must be passed.
     */
    serviceConfigUrl?: pulumi.Input<string>;
    /**
     * Unique VPC ID.
     */
    serviceConfigVpcId?: pulumi.Input<string>;
    /**
     * Scf websocket cleaning function. It takes effect when the current end type is WEBSOCKET and the backend type is SCF.
     */
    serviceConfigWebsocketCleanupFunctionName?: pulumi.Input<string>;
    /**
     * Scf websocket cleans up the function namespace. It takes effect when the current end type is WEBSOCKET and the backend type is SCF.
     */
    serviceConfigWebsocketCleanupFunctionNamespace?: pulumi.Input<string>;
    /**
     * Scf websocket cleaning function version. It takes effect when the current end type is WEBSOCKET and the backend type is SCF.
     */
    serviceConfigWebsocketCleanupFunctionQualifier?: pulumi.Input<string>;
    /**
     * Scf websocket registration function. It takes effect when the current end type is WEBSOCKET and the backend type is SCF.
     */
    serviceConfigWebsocketRegisterFunctionName?: pulumi.Input<string>;
    /**
     * Scf websocket registers function namespaces. It takes effect when the current end type is WEBSOCKET and the backend type is SCF.
     */
    serviceConfigWebsocketRegisterFunctionNamespace?: pulumi.Input<string>;
    /**
     * Scf websocket transfer function version. It takes effect when the current end type is WEBSOCKET and the backend type is SCF.
     */
    serviceConfigWebsocketRegisterFunctionQualifier?: pulumi.Input<string>;
    /**
     * Scf websocket transfer function. It takes effect when the current end type is WEBSOCKET and the backend type is SCF.
     */
    serviceConfigWebsocketTransportFunctionName?: pulumi.Input<string>;
    /**
     * Scf websocket transfer function namespace. It takes effect when the current end type is WEBSOCKET and the backend type is SCF.
     */
    serviceConfigWebsocketTransportFunctionNamespace?: pulumi.Input<string>;
    /**
     * Scf websocket transfer function version. It takes effect when the current end type is WEBSOCKET and the backend type is SCF.
     */
    serviceConfigWebsocketTransportFunctionQualifier?: pulumi.Input<string>;
    /**
     * The unique ID of the service where the API is located.
     */
    serviceId?: pulumi.Input<string>;
    /**
     * The backend service parameters of the API.
     */
    serviceParameters?: pulumi.Input<pulumi.Input<inputs.ApiGateway.ImportOpenApiServiceParameter>[]>;
    /**
     * Health check configuration for microservices.
     */
    serviceTsfHealthCheckConfs?: pulumi.Input<pulumi.Input<inputs.ApiGateway.ImportOpenApiServiceTsfHealthCheckConf>[]>;
    /**
     * Load balancing configuration for microservices.
     */
    serviceTsfLoadBalanceConfs?: pulumi.Input<pulumi.Input<inputs.ApiGateway.ImportOpenApiServiceTsfLoadBalanceConf>[]>;
    /**
     * Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
     */
    updateTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ImportOpenApi resource.
 */
export interface ImportOpenApiArgs {
    /**
     * OpenAPI body content.
     */
    content: pulumi.Input<string>;
    /**
     * The Content version defaults to OpenAPI and currently only supports OpenAPI.
     */
    contentVersion?: pulumi.Input<string>;
    /**
     * The Content format can only be YAML or JSON, and the default is YAML.
     */
    encodeType?: pulumi.Input<string>;
    /**
     * The unique ID of the service where the API is located.
     */
    serviceId: pulumi.Input<string>;
}
