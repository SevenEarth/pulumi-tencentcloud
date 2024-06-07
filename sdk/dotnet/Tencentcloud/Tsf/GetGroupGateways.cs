// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tsf
{
    public static class GetGroupGateways
    {
        /// <summary>
        /// Use this data source to query detailed information of tsf group_gateways
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
        ///     var groupGateways = Tencentcloud.Tsf.GetGroupGateways.Invoke(new()
        ///     {
        ///         GatewayDeployGroupId = "group-aeoej4qy",
        ///         SearchWord = "test",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetGroupGatewaysResult> InvokeAsync(GetGroupGatewaysArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGroupGatewaysResult>("tencentcloud:Tsf/getGroupGateways:getGroupGateways", args ?? new GetGroupGatewaysArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of tsf group_gateways
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
        ///     var groupGateways = Tencentcloud.Tsf.GetGroupGateways.Invoke(new()
        ///     {
        ///         GatewayDeployGroupId = "group-aeoej4qy",
        ///         SearchWord = "test",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetGroupGatewaysResult> Invoke(GetGroupGatewaysInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGroupGatewaysResult>("tencentcloud:Tsf/getGroupGateways:getGroupGateways", args ?? new GetGroupGatewaysInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGroupGatewaysArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// gateway group Id.
        /// </summary>
        [Input("gatewayDeployGroupId", required: true)]
        public string GatewayDeployGroupId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// search word.
        /// </summary>
        [Input("searchWord")]
        public string? SearchWord { get; set; }

        public GetGroupGatewaysArgs()
        {
        }
        public static new GetGroupGatewaysArgs Empty => new GetGroupGatewaysArgs();
    }

    public sealed class GetGroupGatewaysInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// gateway group Id.
        /// </summary>
        [Input("gatewayDeployGroupId", required: true)]
        public Input<string> GatewayDeployGroupId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// search word.
        /// </summary>
        [Input("searchWord")]
        public Input<string>? SearchWord { get; set; }

        public GetGroupGatewaysInvokeArgs()
        {
        }
        public static new GetGroupGatewaysInvokeArgs Empty => new GetGroupGatewaysInvokeArgs();
    }


    [OutputType]
    public sealed class GetGroupGatewaysResult
    {
        public readonly string GatewayDeployGroupId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// api group information.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetGroupGatewaysResultResult> Results;
        public readonly string? SearchWord;

        [OutputConstructor]
        private GetGroupGatewaysResult(
            string gatewayDeployGroupId,

            string id,

            string? resultOutputFile,

            ImmutableArray<Outputs.GetGroupGatewaysResultResult> results,

            string? searchWord)
        {
            GatewayDeployGroupId = gatewayDeployGroupId;
            Id = id;
            ResultOutputFile = resultOutputFile;
            Results = results;
            SearchWord = searchWord;
        }
    }
}
