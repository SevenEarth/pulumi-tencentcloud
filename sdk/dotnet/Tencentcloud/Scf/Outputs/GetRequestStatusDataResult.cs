// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Scf.Outputs
{

    [OutputType]
    public sealed class GetRequestStatusDataResult
    {
        /// <summary>
        /// Time consumed for the request in ms.
        /// </summary>
        public readonly double Duration;
        /// <summary>
        /// Function name.
        /// </summary>
        public readonly string FunctionName;
        /// <summary>
        /// Time consumed by the request in MB.
        /// </summary>
        public readonly double MemUsage;
        /// <summary>
        /// Request ID.
        /// </summary>
        public readonly string RequestId;
        /// <summary>
        /// Result of the request. `0`: succeeded, `1`: running, `-1`: exception.
        /// </summary>
        public readonly int RetCode;
        /// <summary>
        /// Return value after the function is executed.
        /// </summary>
        public readonly string RetMsg;
        /// <summary>
        /// Retry Attempts.
        /// </summary>
        public readonly int RetryNum;
        /// <summary>
        /// Start time of the query, for example `2017-05-16 20:00:00`. If it's left empty, it defaults to 15 minutes before the current time.
        /// </summary>
        public readonly string StartTime;

        [OutputConstructor]
        private GetRequestStatusDataResult(
            double duration,

            string functionName,

            double memUsage,

            string requestId,

            int retCode,

            string retMsg,

            int retryNum,

            string startTime)
        {
            Duration = duration;
            FunctionName = functionName;
            MemUsage = memUsage;
            RequestId = requestId;
            RetCode = retCode;
            RetMsg = retMsg;
            RetryNum = retryNum;
            StartTime = startTime;
        }
    }
}
