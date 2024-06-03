// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cynosdb.Inputs
{

    public sealed class ProxyProxyZoneGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of proxy nodes.
        /// </summary>
        [Input("proxyNodeCount")]
        public Input<int>? ProxyNodeCount { get; set; }

        /// <summary>
        /// Proxy node availability zone.
        /// </summary>
        [Input("proxyNodeZone")]
        public Input<string>? ProxyNodeZone { get; set; }

        public ProxyProxyZoneGetArgs()
        {
        }
        public static new ProxyProxyZoneGetArgs Empty => new ProxyProxyZoneGetArgs();
    }
}
