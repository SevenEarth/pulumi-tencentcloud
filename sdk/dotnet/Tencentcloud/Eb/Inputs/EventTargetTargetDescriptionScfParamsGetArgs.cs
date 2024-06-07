// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Eb.Inputs
{

    public sealed class EventTargetTargetDescriptionScfParamsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Maximum number of events for batch delivery.
        /// </summary>
        [Input("batchEventCount")]
        public Input<int>? BatchEventCount { get; set; }

        /// <summary>
        /// Maximum waiting time for bulk delivery.
        /// </summary>
        [Input("batchTimeout")]
        public Input<int>? BatchTimeout { get; set; }

        /// <summary>
        /// Enable batch delivery.
        /// </summary>
        [Input("enableBatchDelivery")]
        public Input<bool>? EnableBatchDelivery { get; set; }

        public EventTargetTargetDescriptionScfParamsGetArgs()
        {
        }
        public static new EventTargetTargetDescriptionScfParamsGetArgs Empty => new EventTargetTargetDescriptionScfParamsGetArgs();
    }
}
