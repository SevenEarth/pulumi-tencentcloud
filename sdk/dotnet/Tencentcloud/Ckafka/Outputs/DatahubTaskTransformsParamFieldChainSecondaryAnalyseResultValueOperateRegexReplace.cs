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
    public sealed class DatahubTaskTransformsParamFieldChainSecondaryAnalyseResultValueOperateRegexReplace
    {
        /// <summary>
        /// new value.
        /// </summary>
        public readonly string NewValue;
        /// <summary>
        /// Regular.
        /// </summary>
        public readonly string Regex;

        [OutputConstructor]
        private DatahubTaskTransformsParamFieldChainSecondaryAnalyseResultValueOperateRegexReplace(
            string newValue,

            string regex)
        {
            NewValue = newValue;
            Regex = regex;
        }
    }
}
