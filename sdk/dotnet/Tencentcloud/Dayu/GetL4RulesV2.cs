// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dayu
{
    public static class GetL4RulesV2
    {
        /// <summary>
        /// Use this data source to query dayu new layer 4 rules
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var tencentcloudDayuL4RulesV2 = Tencentcloud.Dayu.GetL4RulesV2.Invoke(new()
        ///     {
        ///         Business = "bgpip",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetL4RulesV2Result> InvokeAsync(GetL4RulesV2Args args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetL4RulesV2Result>("tencentcloud:Dayu/getL4RulesV2:getL4RulesV2", args ?? new GetL4RulesV2Args(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query dayu new layer 4 rules
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var tencentcloudDayuL4RulesV2 = Tencentcloud.Dayu.GetL4RulesV2.Invoke(new()
        ///     {
        ///         Business = "bgpip",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetL4RulesV2Result> Invoke(GetL4RulesV2InvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetL4RulesV2Result>("tencentcloud:Dayu/getL4RulesV2:getL4RulesV2", args ?? new GetL4RulesV2InvokeArgs(), options.WithDefaults());
    }


    public sealed class GetL4RulesV2Args : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Type of the resource that the layer 4 rule works for, valid values are `bgpip`, `bgp`, `bgp-multip` and `net`.
        /// </summary>
        [Input("business", required: true)]
        public string Business { get; set; } = null!;

        /// <summary>
        /// Ip of the resource.
        /// </summary>
        [Input("ip")]
        public string? Ip { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// Virtual port of resource.
        /// </summary>
        [Input("virtualPort")]
        public int? VirtualPort { get; set; }

        public GetL4RulesV2Args()
        {
        }
        public static new GetL4RulesV2Args Empty => new GetL4RulesV2Args();
    }

    public sealed class GetL4RulesV2InvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Type of the resource that the layer 4 rule works for, valid values are `bgpip`, `bgp`, `bgp-multip` and `net`.
        /// </summary>
        [Input("business", required: true)]
        public Input<string> Business { get; set; } = null!;

        /// <summary>
        /// Ip of the resource.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// Virtual port of resource.
        /// </summary>
        [Input("virtualPort")]
        public Input<int>? VirtualPort { get; set; }

        public GetL4RulesV2InvokeArgs()
        {
        }
        public static new GetL4RulesV2InvokeArgs Empty => new GetL4RulesV2InvokeArgs();
    }


    [OutputType]
    public sealed class GetL4RulesV2Result
    {
        public readonly string Business;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Bind the resource IP information.
        /// </summary>
        public readonly string? Ip;
        /// <summary>
        /// A list of layer 4 rules. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetL4RulesV2ListResult> Lists;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// The virtual port of the layer 4 rule.
        /// </summary>
        public readonly int? VirtualPort;

        [OutputConstructor]
        private GetL4RulesV2Result(
            string business,

            string id,

            string? ip,

            ImmutableArray<Outputs.GetL4RulesV2ListResult> lists,

            string? resultOutputFile,

            int? virtualPort)
        {
            Business = business;
            Id = id;
            Ip = ip;
            Lists = lists;
            ResultOutputFile = resultOutputFile;
            VirtualPort = virtualPort;
        }
    }
}
