// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.ApiGateway.Outputs
{

    [OutputType]
    public sealed class GetUsagePlanEnvironmentsListResult
    {
        /// <summary>
        /// The API ID, this value is empty if attach service.
        /// </summary>
        public readonly string ApiId;
        /// <summary>
        /// The API name, this value is empty if attach service.
        /// </summary>
        public readonly string ApiName;
        /// <summary>
        /// Creation time in the format of `YYYY-MM-DDThh:mm:ssZ` according to ISO 8601 standard. UTC time is used.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The environment name.
        /// </summary>
        public readonly string Environment;
        /// <summary>
        /// The API method, this value is empty if attach service.
        /// </summary>
        public readonly string Method;
        /// <summary>
        /// Last modified time in the format of `YYYY-MM-DDThh:mm:ssZ` according to ISO 8601 standard. UTC time is used.
        /// </summary>
        public readonly string ModifyTime;
        /// <summary>
        /// The API path, this value is empty if attach service.
        /// </summary>
        public readonly string Path;
        /// <summary>
        /// The service ID.
        /// </summary>
        public readonly string ServiceId;
        /// <summary>
        /// The service name.
        /// </summary>
        public readonly string ServiceName;

        [OutputConstructor]
        private GetUsagePlanEnvironmentsListResult(
            string apiId,

            string apiName,

            string createTime,

            string environment,

            string method,

            string modifyTime,

            string path,

            string serviceId,

            string serviceName)
        {
            ApiId = apiId;
            ApiName = apiName;
            CreateTime = createTime;
            Environment = environment;
            Method = method;
            ModifyTime = modifyTime;
            Path = path;
            ServiceId = serviceId;
            ServiceName = serviceName;
        }
    }
}
