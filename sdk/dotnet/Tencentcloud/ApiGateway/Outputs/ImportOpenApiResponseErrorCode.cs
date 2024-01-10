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
    public sealed class ImportOpenApiResponseErrorCode
    {
        /// <summary>
        /// Custom response configuration error code.
        /// </summary>
        public readonly int? Code;
        /// <summary>
        /// Custom error code conversion.
        /// </summary>
        public readonly int? ConvertedCode;
        /// <summary>
        /// Parameter description.
        /// </summary>
        public readonly string? Desc;
        /// <summary>
        /// Custom response configuration error message.
        /// </summary>
        public readonly string? Msg;
        /// <summary>
        /// Whether to enable error code conversion. Default value: `false`.
        /// </summary>
        public readonly bool? NeedConvert;

        [OutputConstructor]
        private ImportOpenApiResponseErrorCode(
            int? code,

            int? convertedCode,

            string? desc,

            string? msg,

            bool? needConvert)
        {
            Code = code;
            ConvertedCode = convertedCode;
            Desc = desc;
            Msg = msg;
            NeedConvert = needConvert;
        }
    }
}
