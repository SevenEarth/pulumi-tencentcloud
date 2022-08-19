// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.PrivateDns.Inputs
{

    public sealed class ZoneTagSetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Key of Tag.
        /// </summary>
        [Input("tagKey", required: true)]
        public Input<string> TagKey { get; set; } = null!;

        /// <summary>
        /// Value of Tag.
        /// </summary>
        [Input("tagValue", required: true)]
        public Input<string> TagValue { get; set; } = null!;

        public ZoneTagSetArgs()
        {
        }
    }
}
