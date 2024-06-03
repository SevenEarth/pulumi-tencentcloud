// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Eb.Outputs
{

    [OutputType]
    public sealed class GetSearchFilterFilterResult
    {
        /// <summary>
        /// filter field name.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// operator, congruent eq, not equal neq, similar like, exclude similar not like, less than lt, less than and equal to lte, greater than gt, greater than and equal to gte, within range range, not within range norange.
        /// </summary>
        public readonly string Operator;
        /// <summary>
        /// Filter values, range operations need to enter two values at the same time, separated by commas.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private GetSearchFilterFilterResult(
            string key,

            string @operator,

            string value)
        {
            Key = key;
            Operator = @operator;
            Value = value;
        }
    }
}
