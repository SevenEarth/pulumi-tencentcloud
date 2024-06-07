// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Pts.Inputs
{

    public sealed class JobDomainNameConfigDnsConfigGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("nameservers")]
        private InputList<string>? _nameservers;

        /// <summary>
        /// DNS IP List.
        /// </summary>
        public InputList<string> Nameservers
        {
            get => _nameservers ?? (_nameservers = new InputList<string>());
            set => _nameservers = value;
        }

        public JobDomainNameConfigDnsConfigGetArgs()
        {
        }
        public static new JobDomainNameConfigDnsConfigGetArgs Empty => new JobDomainNameConfigDnsConfigGetArgs();
    }
}
