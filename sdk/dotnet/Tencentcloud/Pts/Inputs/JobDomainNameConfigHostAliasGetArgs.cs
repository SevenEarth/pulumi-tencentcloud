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

    public sealed class JobDomainNameConfigHostAliasGetArgs : Pulumi.ResourceArgs
    {
        [Input("hostNames")]
        private InputList<string>? _hostNames;
        public InputList<string> HostNames
        {
            get => _hostNames ?? (_hostNames = new InputList<string>());
            set => _hostNames = value;
        }

        [Input("ip")]
        public Input<string>? Ip { get; set; }

        public JobDomainNameConfigHostAliasGetArgs()
        {
        }
    }
}