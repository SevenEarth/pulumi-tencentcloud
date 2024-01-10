// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Teo.Inputs
{

    public sealed class ZoneOwnershipVerificationDnsVerificationGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Record type.
        /// </summary>
        [Input("recordType")]
        public Input<string>? RecordType { get; set; }

        /// <summary>
        /// Record the value.
        /// </summary>
        [Input("recordValue")]
        public Input<string>? RecordValue { get; set; }

        /// <summary>
        /// Host record.
        /// </summary>
        [Input("subdomain")]
        public Input<string>? Subdomain { get; set; }

        public ZoneOwnershipVerificationDnsVerificationGetArgs()
        {
        }
    }
}
