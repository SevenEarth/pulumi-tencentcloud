// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ses.Inputs
{

    public sealed class DomainAttributeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Values that need to be configured.
        /// </summary>
        [Input("expectedValue")]
        public Input<string>? ExpectedValue { get; set; }

        /// <summary>
        /// Domain name.
        /// </summary>
        [Input("sendDomain")]
        public Input<string>? SendDomain { get; set; }

        /// <summary>
        /// Record Type CNAME | A | TXT | MX.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public DomainAttributeArgs()
        {
        }
    }
}
