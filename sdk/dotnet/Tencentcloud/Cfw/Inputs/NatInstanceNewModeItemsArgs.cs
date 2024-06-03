// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cfw.Inputs
{

    public sealed class NatInstanceNewModeItemsArgs : global::Pulumi.ResourceArgs
    {
        [Input("eips", required: true)]
        private InputList<string>? _eips;

        /// <summary>
        /// List of egress elastic public network IPs bound in the new mode.
        /// </summary>
        public InputList<string> Eips
        {
            get => _eips ?? (_eips = new InputList<string>());
            set => _eips = value;
        }

        [Input("vpcLists", required: true)]
        private InputList<string>? _vpcLists;

        /// <summary>
        /// List of vpcs connected in new mode.
        /// </summary>
        public InputList<string> VpcLists
        {
            get => _vpcLists ?? (_vpcLists = new InputList<string>());
            set => _vpcLists = value;
        }

        public NatInstanceNewModeItemsArgs()
        {
        }
        public static new NatInstanceNewModeItemsArgs Empty => new NatInstanceNewModeItemsArgs();
    }
}
