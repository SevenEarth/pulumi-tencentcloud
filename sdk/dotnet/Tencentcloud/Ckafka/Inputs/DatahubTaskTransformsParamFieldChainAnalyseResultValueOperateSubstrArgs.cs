// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ckafka.Inputs
{

    public sealed class DatahubTaskTransformsParamFieldChainAnalyseResultValueOperateSubstrArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// cut-off position.
        /// </summary>
        [Input("end", required: true)]
        public Input<int> End { get; set; } = null!;

        /// <summary>
        /// interception starting position.
        /// </summary>
        [Input("start", required: true)]
        public Input<int> Start { get; set; } = null!;

        public DatahubTaskTransformsParamFieldChainAnalyseResultValueOperateSubstrArgs()
        {
        }
        public static new DatahubTaskTransformsParamFieldChainAnalyseResultValueOperateSubstrArgs Empty => new DatahubTaskTransformsParamFieldChainAnalyseResultValueOperateSubstrArgs();
    }
}
