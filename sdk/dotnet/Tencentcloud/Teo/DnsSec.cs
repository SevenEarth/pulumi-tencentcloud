// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Teo
{
    /// <summary>
    /// Provides a resource to create a teo dns_sec
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var dnsSec = new Tencentcloud.Teo.DnsSec("dnsSec", new Tencentcloud.Teo.DnsSecArgs
    ///         {
    ///             Status = "enabled",
    ///             ZoneId = "zone-297z8rf93cfw",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// teo dns_sec can be imported using the zone_id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Teo/dnsSec:DnsSec dns_sec zone-297z8rf93cfw
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Teo/dnsSec:DnsSec")]
    public partial class DnsSec : Pulumi.CustomResource
    {
        /// <summary>
        /// DNSSEC infos.
        /// </summary>
        [Output("dnssec")]
        public Output<Outputs.DnsSecDnssec> Dnssec { get; private set; } = null!;

        /// <summary>
        /// Last modification date.
        /// </summary>
        [Output("modifiedOn")]
        public Output<string> ModifiedOn { get; private set; } = null!;

        /// <summary>
        /// DNSSEC status. Valid values: `enabled`, `disabled`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Site ID.
        /// </summary>
        [Output("zoneId")]
        public Output<string> ZoneId { get; private set; } = null!;


        /// <summary>
        /// Create a DnsSec resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DnsSec(string name, DnsSecArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Teo/dnsSec:DnsSec", name, args ?? new DnsSecArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DnsSec(string name, Input<string> id, DnsSecState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Teo/dnsSec:DnsSec", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/tencentcloudstack",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DnsSec resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DnsSec Get(string name, Input<string> id, DnsSecState? state = null, CustomResourceOptions? options = null)
        {
            return new DnsSec(name, id, state, options);
        }
    }

    public sealed class DnsSecArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// DNSSEC infos.
        /// </summary>
        [Input("dnssec")]
        public Input<Inputs.DnsSecDnssecArgs>? Dnssec { get; set; }

        /// <summary>
        /// DNSSEC status. Valid values: `enabled`, `disabled`.
        /// </summary>
        [Input("status", required: true)]
        public Input<string> Status { get; set; } = null!;

        /// <summary>
        /// Site ID.
        /// </summary>
        [Input("zoneId", required: true)]
        public Input<string> ZoneId { get; set; } = null!;

        public DnsSecArgs()
        {
        }
    }

    public sealed class DnsSecState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// DNSSEC infos.
        /// </summary>
        [Input("dnssec")]
        public Input<Inputs.DnsSecDnssecGetArgs>? Dnssec { get; set; }

        /// <summary>
        /// Last modification date.
        /// </summary>
        [Input("modifiedOn")]
        public Input<string>? ModifiedOn { get; set; }

        /// <summary>
        /// DNSSEC status. Valid values: `enabled`, `disabled`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Site ID.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public DnsSecState()
        {
        }
    }
}
