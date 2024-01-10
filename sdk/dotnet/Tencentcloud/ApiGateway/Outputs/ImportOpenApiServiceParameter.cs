// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.ApiGateway.Outputs
{

    [OutputType]
    public sealed class ImportOpenApiServiceParameter
    {
        /// <summary>
        /// The default value for the backend service parameters of the API. This parameter is only used when ServiceType is HTTP.Note: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        public readonly string? DefaultValue;
        /// <summary>
        /// The backend service parameter name of the API. This parameter is only used when ServiceType is HTTP. The front and rear parameter names can be different.Note: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The backend service parameter location of the API, such as head. This parameter is only used when ServiceType is HTTP. The parameter positions at the front and rear ends can be configured differently.Note: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        public readonly string? Position;
        /// <summary>
        /// Remarks on the backend service parameters of the API. This parameter is only used when ServiceType is HTTP.Note: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        public readonly string? RelevantRequestParameterDesc;
        /// <summary>
        /// The name of the front-end parameter corresponding to the backend service parameter of the API. This parameter is only used when ServiceType is HTTP.Note: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        public readonly string? RelevantRequestParameterName;
        /// <summary>
        /// The location of the front-end parameters corresponding to the backend service parameters of the API, such as head. This parameter is only used when ServiceType is HTTP.Note: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        public readonly string? RelevantRequestParameterPosition;
        /// <summary>
        /// The backend service parameter type of the API. This parameter is only used when ServiceType is HTTP.Note: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        public readonly string? RelevantRequestParameterType;

        [OutputConstructor]
        private ImportOpenApiServiceParameter(
            string? defaultValue,

            string? name,

            string? position,

            string? relevantRequestParameterDesc,

            string? relevantRequestParameterName,

            string? relevantRequestParameterPosition,

            string? relevantRequestParameterType)
        {
            DefaultValue = defaultValue;
            Name = name;
            Position = position;
            RelevantRequestParameterDesc = relevantRequestParameterDesc;
            RelevantRequestParameterName = relevantRequestParameterName;
            RelevantRequestParameterPosition = relevantRequestParameterPosition;
            RelevantRequestParameterType = relevantRequestParameterType;
        }
    }
}
