// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.ApiGateway.Inputs
{

    public sealed class ApiServiceParameterGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The default value for the backend service parameters of the API. This parameter is only used when ServiceType is HTTP.Note: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        [Input("defaultValue")]
        public Input<string>? DefaultValue { get; set; }

        /// <summary>
        /// The backend service parameter name of the API. This parameter is only used when ServiceType is HTTP. The front and rear parameter names can be different.Note: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The backend service parameter location of the API, such as head. This parameter is only used when ServiceType is HTTP. The parameter positions at the front and rear ends can be configured differently.Note: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        [Input("position")]
        public Input<string>? Position { get; set; }

        /// <summary>
        /// Remarks on the backend service parameters of the API. This parameter is only used when ServiceType is HTTP.Note: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        [Input("relevantRequestParameterDesc")]
        public Input<string>? RelevantRequestParameterDesc { get; set; }

        /// <summary>
        /// The name of the front-end parameter corresponding to the backend service parameter of the API. This parameter is only used when ServiceType is HTTP.Note: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        [Input("relevantRequestParameterName")]
        public Input<string>? RelevantRequestParameterName { get; set; }

        /// <summary>
        /// The location of the front-end parameters corresponding to the backend service parameters of the API, such as head. This parameter is only used when ServiceType is HTTP.Note: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        [Input("relevantRequestParameterPosition")]
        public Input<string>? RelevantRequestParameterPosition { get; set; }

        /// <summary>
        /// The backend service parameter type of the API. This parameter is only used when ServiceType is HTTP.Note: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        [Input("relevantRequestParameterType")]
        public Input<string>? RelevantRequestParameterType { get; set; }

        public ApiServiceParameterGetArgs()
        {
        }
        public static new ApiServiceParameterGetArgs Empty => new ApiServiceParameterGetArgs();
    }
}
