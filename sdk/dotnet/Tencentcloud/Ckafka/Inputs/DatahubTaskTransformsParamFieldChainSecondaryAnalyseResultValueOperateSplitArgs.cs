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

    public sealed class DatahubTaskTransformsParamFieldChainSecondaryAnalyseResultValueOperateSplitArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// delimiter.
        /// </summary>
        [Input("regex", required: true)]
        public Input<string> Regex { get; set; } = null!;

        public DatahubTaskTransformsParamFieldChainSecondaryAnalyseResultValueOperateSplitArgs()
        {
        }
        public static new DatahubTaskTransformsParamFieldChainSecondaryAnalyseResultValueOperateSplitArgs Empty => new DatahubTaskTransformsParamFieldChainSecondaryAnalyseResultValueOperateSplitArgs();
    }
}
