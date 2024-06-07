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
    public static class GetUnitRules
    {
        /// <summary>
        /// Use this data source to query detailed information of tsf unit_rules
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
        ///     var unitRules = Tencentcloud.Tsf.GetUnitRules.Invoke(new()
        ///     {
        ///         GatewayInstanceId = "gw-ins-lvdypq5k",
        ///         Status = "disabled",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetUnitRulesResult> InvokeAsync(GetUnitRulesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetUnitRulesResult>("tencentcloud:Tsf/getUnitRules:getUnitRules", args ?? new GetUnitRulesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of tsf unit_rules
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
        ///     var unitRules = Tencentcloud.Tsf.GetUnitRules.Invoke(new()
        ///     {
        ///         GatewayInstanceId = "gw-ins-lvdypq5k",
        ///         Status = "disabled",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetUnitRulesResult> Invoke(GetUnitRulesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetUnitRulesResult>("tencentcloud:Tsf/getUnitRules:getUnitRules", args ?? new GetUnitRulesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUnitRulesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// gateway instance id.
        /// </summary>
        [Input("gatewayInstanceId", required: true)]
        public string GatewayInstanceId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// Enabled state, disabled: unpublished, enabled: published.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        public GetUnitRulesArgs()
        {
        }
        public static new GetUnitRulesArgs Empty => new GetUnitRulesArgs();
    }

    public sealed class GetUnitRulesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// gateway instance id.
        /// </summary>
        [Input("gatewayInstanceId", required: true)]
        public Input<string> GatewayInstanceId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// Enabled state, disabled: unpublished, enabled: published.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public GetUnitRulesInvokeArgs()
        {
        }
        public static new GetUnitRulesInvokeArgs Empty => new GetUnitRulesInvokeArgs();
    }


    [OutputType]
    public sealed class GetUnitRulesResult
    {
        /// <summary>
        /// Gateway Entity ID.
        /// </summary>
        public readonly string GatewayInstanceId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// Pagination list information.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetUnitRulesResultResult> Results;
        /// <summary>
        /// Use status: enabled/disabled.
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private GetUnitRulesResult(
            string gatewayInstanceId,

            string id,

            string? resultOutputFile,

            ImmutableArray<Outputs.GetUnitRulesResultResult> results,

            string? status)
        {
            GatewayInstanceId = gatewayInstanceId;
            Id = id;
            ResultOutputFile = resultOutputFile;
            Results = results;
            Status = status;
        }
    }
}
