// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tcr.Inputs
{

    public sealed class NamespaceCveWhitelistItemArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Vulnerability Whitelist ID.
        /// </summary>
        [Input("cveId")]
        public Input<string>? CveId { get; set; }

        public NamespaceCveWhitelistItemArgs()
        {
        }
    }
}
