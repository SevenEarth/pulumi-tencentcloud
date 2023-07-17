// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Sqlserver.Inputs
{

    public sealed class GeneralCloudInstanceResourceTagArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// tag key.
        /// </summary>
        [Input("tagKey")]
        public Input<string>? TagKey { get; set; }

        /// <summary>
        /// tag value.
        /// </summary>
        [Input("tagValue")]
        public Input<string>? TagValue { get; set; }

        public GeneralCloudInstanceResourceTagArgs()
        {
        }
    }
}
