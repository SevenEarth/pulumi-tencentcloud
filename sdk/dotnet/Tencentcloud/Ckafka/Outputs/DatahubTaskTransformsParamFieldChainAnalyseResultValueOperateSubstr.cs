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
    public sealed class DatahubTaskTransformsParamFieldChainAnalyseResultValueOperateSubstr
    {
        /// <summary>
        /// cut-off position.
        /// </summary>
        public readonly int End;
        /// <summary>
        /// interception starting position.
        /// </summary>
        public readonly int Start;

        [OutputConstructor]
        private DatahubTaskTransformsParamFieldChainAnalyseResultValueOperateSubstr(
            int end,

            int start)
        {
            End = end;
            Start = start;
        }
    }
}
