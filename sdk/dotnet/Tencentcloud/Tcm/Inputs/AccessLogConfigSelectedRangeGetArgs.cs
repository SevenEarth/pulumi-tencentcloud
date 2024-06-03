// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tcm.Inputs
{

    public sealed class AccessLogConfigSelectedRangeGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Select all if true, default false.
        /// </summary>
        [Input("all")]
        public Input<bool>? All { get; set; }

        [Input("items")]
        private InputList<Inputs.AccessLogConfigSelectedRangeItemGetArgs>? _items;

        /// <summary>
        /// Items.
        /// </summary>
        public InputList<Inputs.AccessLogConfigSelectedRangeItemGetArgs> Items
        {
            get => _items ?? (_items = new InputList<Inputs.AccessLogConfigSelectedRangeItemGetArgs>());
            set => _items = value;
        }

        public AccessLogConfigSelectedRangeGetArgs()
        {
        }
        public static new AccessLogConfigSelectedRangeGetArgs Empty => new AccessLogConfigSelectedRangeGetArgs();
    }
}
