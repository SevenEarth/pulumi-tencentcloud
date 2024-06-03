// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tse.Inputs
{

    public sealed class CngwNetworkAccessControlAccessControlGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("cidrBlackLists")]
        private InputList<string>? _cidrBlackLists;

        /// <summary>
        /// Black list.
        /// </summary>
        public InputList<string> CidrBlackLists
        {
            get => _cidrBlackLists ?? (_cidrBlackLists = new InputList<string>());
            set => _cidrBlackLists = value;
        }

        [Input("cidrWhiteLists")]
        private InputList<string>? _cidrWhiteLists;

        /// <summary>
        /// White list.
        /// </summary>
        public InputList<string> CidrWhiteLists
        {
            get => _cidrWhiteLists ?? (_cidrWhiteLists = new InputList<string>());
            set => _cidrWhiteLists = value;
        }

        /// <summary>
        /// Access mode: `Whitelist`, `Blacklist`.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        public CngwNetworkAccessControlAccessControlGetArgs()
        {
        }
        public static new CngwNetworkAccessControlAccessControlGetArgs Empty => new CngwNetworkAccessControlAccessControlGetArgs();
    }
}
