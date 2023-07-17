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
    public sealed class DatahubTaskTransformsParamFieldChainSecondaryAnalyseResultValueOperateReplace
    {
        /// <summary>
        /// new value.
        /// </summary>
        public readonly string NewValue;
        /// <summary>
        /// been replaced value.
        /// </summary>
        public readonly string OldValue;

        [OutputConstructor]
        private DatahubTaskTransformsParamFieldChainSecondaryAnalyseResultValueOperateReplace(
            string newValue,

            string oldValue)
        {
            NewValue = newValue;
            OldValue = oldValue;
        }
    }
}
