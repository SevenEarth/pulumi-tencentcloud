// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ckafka.Outputs
{

    [OutputType]
    public sealed class DatahubTaskSourceResourceScfParam
    {
        /// <summary>
        /// The maximum number of messages sent in each batch, the default is 1000.
        /// </summary>
        public readonly int? BatchSize;
        /// <summary>
        /// SCF function name.
        /// </summary>
        public readonly string FunctionName;
        /// <summary>
        /// The number of retries after the SCF call fails, the default is 5.
        /// </summary>
        public readonly int? MaxRetries;
        /// <summary>
        /// SCF cloud function namespace, the default is default.
        /// </summary>
        public readonly string? Namespace;
        /// <summary>
        /// SCF cloud function version and alias, the default is DEFAULT.
        /// </summary>
        public readonly string? Qualifier;

        [OutputConstructor]
        private DatahubTaskSourceResourceScfParam(
            int? batchSize,

            string functionName,

            int? maxRetries,

            string? @namespace,

            string? qualifier)
        {
            BatchSize = batchSize;
            FunctionName = functionName;
            MaxRetries = maxRetries;
            Namespace = @namespace;
            Qualifier = qualifier;
        }
    }
}
