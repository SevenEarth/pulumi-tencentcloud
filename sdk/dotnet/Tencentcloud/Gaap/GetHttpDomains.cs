// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Gaap
{
    public static class GetHttpDomains
    {
        /// <summary>
        /// Use this data source to query forward domain of layer7 listeners.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var fooProxy = new Tencentcloud.Gaap.Proxy("fooProxy", new()
        ///     {
        ///         Bandwidth = 10,
        ///         Concurrent = 2,
        ///         AccessRegion = "SouthChina",
        ///         RealserverRegion = "NorthChina",
        ///     });
        /// 
        ///     var fooLayer7Listener = new Tencentcloud.Gaap.Layer7Listener("fooLayer7Listener", new()
        ///     {
        ///         Protocol = "HTTP",
        ///         Port = 80,
        ///         ProxyId = fooProxy.Id,
        ///     });
        /// 
        ///     var fooHttpDomain = new Tencentcloud.Gaap.HttpDomain("fooHttpDomain", new()
        ///     {
        ///         ListenerId = fooLayer7Listener.Id,
        ///         Domain = "www.qq.com",
        ///     });
        /// 
        ///     var fooHttpDomains = Tencentcloud.Gaap.GetHttpDomains.Invoke(new()
        ///     {
        ///         ListenerId = fooLayer7Listener.Id,
        ///         Domain = fooHttpDomain.Domain,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetHttpDomainsResult> InvokeAsync(GetHttpDomainsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetHttpDomainsResult>("tencentcloud:Gaap/getHttpDomains:getHttpDomains", args ?? new GetHttpDomainsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query forward domain of layer7 listeners.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var fooProxy = new Tencentcloud.Gaap.Proxy("fooProxy", new()
        ///     {
        ///         Bandwidth = 10,
        ///         Concurrent = 2,
        ///         AccessRegion = "SouthChina",
        ///         RealserverRegion = "NorthChina",
        ///     });
        /// 
        ///     var fooLayer7Listener = new Tencentcloud.Gaap.Layer7Listener("fooLayer7Listener", new()
        ///     {
        ///         Protocol = "HTTP",
        ///         Port = 80,
        ///         ProxyId = fooProxy.Id,
        ///     });
        /// 
        ///     var fooHttpDomain = new Tencentcloud.Gaap.HttpDomain("fooHttpDomain", new()
        ///     {
        ///         ListenerId = fooLayer7Listener.Id,
        ///         Domain = "www.qq.com",
        ///     });
        /// 
        ///     var fooHttpDomains = Tencentcloud.Gaap.GetHttpDomains.Invoke(new()
        ///     {
        ///         ListenerId = fooLayer7Listener.Id,
        ///         Domain = fooHttpDomain.Domain,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetHttpDomainsResult> Invoke(GetHttpDomainsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetHttpDomainsResult>("tencentcloud:Gaap/getHttpDomains:getHttpDomains", args ?? new GetHttpDomainsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetHttpDomainsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Forward domain of the layer7 listener to be queried.
        /// </summary>
        [Input("domain", required: true)]
        public string Domain { get; set; } = null!;

        /// <summary>
        /// ID of the layer7 listener to be queried.
        /// </summary>
        [Input("listenerId", required: true)]
        public string ListenerId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetHttpDomainsArgs()
        {
        }
        public static new GetHttpDomainsArgs Empty => new GetHttpDomainsArgs();
    }

    public sealed class GetHttpDomainsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Forward domain of the layer7 listener to be queried.
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// ID of the layer7 listener to be queried.
        /// </summary>
        [Input("listenerId", required: true)]
        public Input<string> ListenerId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetHttpDomainsInvokeArgs()
        {
        }
        public static new GetHttpDomainsInvokeArgs Empty => new GetHttpDomainsInvokeArgs();
    }


    [OutputType]
    public sealed class GetHttpDomainsResult
    {
        /// <summary>
        /// Forward domain of the layer7 listener.
        /// </summary>
        public readonly string Domain;
        /// <summary>
        /// An information list of forward domain of the layer7 listeners. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetHttpDomainsDomainResult> Domains;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string ListenerId;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetHttpDomainsResult(
            string domain,

            ImmutableArray<Outputs.GetHttpDomainsDomainResult> domains,

            string id,

            string listenerId,

            string? resultOutputFile)
        {
            Domain = domain;
            Domains = domains;
            Id = id;
            ListenerId = listenerId;
            ResultOutputFile = resultOutputFile;
        }
    }
}
