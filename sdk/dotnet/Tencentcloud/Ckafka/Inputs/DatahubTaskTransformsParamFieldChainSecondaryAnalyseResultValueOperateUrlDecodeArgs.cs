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

    public sealed class DatahubTaskTransformsParamFieldChainSecondaryAnalyseResultValueOperateUrlDecodeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// code.
        /// </summary>
        [Input("charsetName")]
        public Input<string>? CharsetName { get; set; }

        public DatahubTaskTransformsParamFieldChainSecondaryAnalyseResultValueOperateUrlDecodeArgs()
        {
        }
        public static new DatahubTaskTransformsParamFieldChainSecondaryAnalyseResultValueOperateUrlDecodeArgs Empty => new DatahubTaskTransformsParamFieldChainSecondaryAnalyseResultValueOperateUrlDecodeArgs();
    }
}
