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

    public sealed class DatahubTaskTransformsParamFieldChainSecondaryAnalyseResultValueOperateDateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Time format.
        /// </summary>
        [Input("format")]
        public Input<string>? Format { get; set; }

        /// <summary>
        /// input type, string|unix.
        /// </summary>
        [Input("targetType")]
        public Input<string>? TargetType { get; set; }

        /// <summary>
        /// default GMT+8.
        /// </summary>
        [Input("timeZone")]
        public Input<string>? TimeZone { get; set; }

        public DatahubTaskTransformsParamFieldChainSecondaryAnalyseResultValueOperateDateArgs()
        {
        }
        public static new DatahubTaskTransformsParamFieldChainSecondaryAnalyseResultValueOperateDateArgs Empty => new DatahubTaskTransformsParamFieldChainSecondaryAnalyseResultValueOperateDateArgs();
    }
}
