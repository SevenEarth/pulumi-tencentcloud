// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Elasticsearch.Outputs
{

    [OutputType]
    public sealed class GetDiagnoseDiagnoseResultJobResultMetricDetailMetricDimensionResult
    {
        /// <summary>
        /// Key.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// Value.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private GetDiagnoseDiagnoseResultJobResultMetricDetailMetricDimensionResult(
            string key,

            string value)
        {
            Key = key;
            Value = value;
        }
    }
}
