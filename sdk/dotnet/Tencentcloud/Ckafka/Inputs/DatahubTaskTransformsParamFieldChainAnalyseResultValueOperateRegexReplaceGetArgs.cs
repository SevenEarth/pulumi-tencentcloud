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

    public sealed class DatahubTaskTransformsParamFieldChainAnalyseResultValueOperateRegexReplaceGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// new value.
        /// </summary>
        [Input("newValue", required: true)]
        public Input<string> NewValue { get; set; } = null!;

        /// <summary>
        /// Regular.
        /// </summary>
        [Input("regex", required: true)]
        public Input<string> Regex { get; set; } = null!;

        public DatahubTaskTransformsParamFieldChainAnalyseResultValueOperateRegexReplaceGetArgs()
        {
        }
    }
}
