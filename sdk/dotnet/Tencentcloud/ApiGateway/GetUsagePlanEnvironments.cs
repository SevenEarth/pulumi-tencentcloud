// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.ApiGateway
{
    public static class GetUsagePlanEnvironments
    {
        /// <summary>
        /// Used to query the environment list bound by the plan.
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
        ///     var plan = new Tencentcloud.ApiGateway.UsagePlan("plan", new()
        ///     {
        ///         UsagePlanName = "my_plan",
        ///         UsagePlanDesc = "nice plan",
        ///         MaxRequestNum = 100,
        ///         MaxRequestNumPreSec = 10,
        ///     });
        /// 
        ///     var service = new Tencentcloud.ApiGateway.Service("service", new()
        ///     {
        ///         ServiceName = "niceservice",
        ///         Protocol = "http&amp;https",
        ///         ServiceDesc = "your nice service",
        ///         NetTypes = new[]
        ///         {
        ///             "INNER",
        ///             "OUTER",
        ///         },
        ///         IpVersion = "IPv4",
        ///     });
        /// 
        ///     var attachService = new Tencentcloud.ApiGateway.UsagePlanAttachment("attachService", new()
        ///     {
        ///         UsagePlanId = plan.Id,
        ///         ServiceId = service.Id,
        ///         Environment = "test",
        ///         BindType = "SERVICE",
        ///     });
        /// 
        ///     var environmentTest = Tencentcloud.ApiGateway.GetUsagePlanEnvironments.Invoke(new()
        ///     {
        ///         UsagePlanId = attachService.UsagePlanId,
        ///         BindType = "SERVICE",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetUsagePlanEnvironmentsResult> InvokeAsync(GetUsagePlanEnvironmentsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetUsagePlanEnvironmentsResult>("tencentcloud:ApiGateway/getUsagePlanEnvironments:getUsagePlanEnvironments", args ?? new GetUsagePlanEnvironmentsArgs(), options.WithDefaults());

        /// <summary>
        /// Used to query the environment list bound by the plan.
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
        ///     var plan = new Tencentcloud.ApiGateway.UsagePlan("plan", new()
        ///     {
        ///         UsagePlanName = "my_plan",
        ///         UsagePlanDesc = "nice plan",
        ///         MaxRequestNum = 100,
        ///         MaxRequestNumPreSec = 10,
        ///     });
        /// 
        ///     var service = new Tencentcloud.ApiGateway.Service("service", new()
        ///     {
        ///         ServiceName = "niceservice",
        ///         Protocol = "http&amp;https",
        ///         ServiceDesc = "your nice service",
        ///         NetTypes = new[]
        ///         {
        ///             "INNER",
        ///             "OUTER",
        ///         },
        ///         IpVersion = "IPv4",
        ///     });
        /// 
        ///     var attachService = new Tencentcloud.ApiGateway.UsagePlanAttachment("attachService", new()
        ///     {
        ///         UsagePlanId = plan.Id,
        ///         ServiceId = service.Id,
        ///         Environment = "test",
        ///         BindType = "SERVICE",
        ///     });
        /// 
        ///     var environmentTest = Tencentcloud.ApiGateway.GetUsagePlanEnvironments.Invoke(new()
        ///     {
        ///         UsagePlanId = attachService.UsagePlanId,
        ///         BindType = "SERVICE",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetUsagePlanEnvironmentsResult> Invoke(GetUsagePlanEnvironmentsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetUsagePlanEnvironmentsResult>("tencentcloud:ApiGateway/getUsagePlanEnvironments:getUsagePlanEnvironments", args ?? new GetUsagePlanEnvironmentsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUsagePlanEnvironmentsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Binding type. Valid values: `API`, `SERVICE`. Default value: `SERVICE`.
        /// </summary>
        [Input("bindType")]
        public string? BindType { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// ID of the usage plan to be queried.
        /// </summary>
        [Input("usagePlanId", required: true)]
        public string UsagePlanId { get; set; } = null!;

        public GetUsagePlanEnvironmentsArgs()
        {
        }
        public static new GetUsagePlanEnvironmentsArgs Empty => new GetUsagePlanEnvironmentsArgs();
    }

    public sealed class GetUsagePlanEnvironmentsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Binding type. Valid values: `API`, `SERVICE`. Default value: `SERVICE`.
        /// </summary>
        [Input("bindType")]
        public Input<string>? BindType { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// ID of the usage plan to be queried.
        /// </summary>
        [Input("usagePlanId", required: true)]
        public Input<string> UsagePlanId { get; set; } = null!;

        public GetUsagePlanEnvironmentsInvokeArgs()
        {
        }
        public static new GetUsagePlanEnvironmentsInvokeArgs Empty => new GetUsagePlanEnvironmentsInvokeArgs();
    }


    [OutputType]
    public sealed class GetUsagePlanEnvironmentsResult
    {
        public readonly string? BindType;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of usage plan binding details.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetUsagePlanEnvironmentsListResult> Lists;
        public readonly string? ResultOutputFile;
        public readonly string UsagePlanId;

        [OutputConstructor]
        private GetUsagePlanEnvironmentsResult(
            string? bindType,

            string id,

            ImmutableArray<Outputs.GetUsagePlanEnvironmentsListResult> lists,

            string? resultOutputFile,

            string usagePlanId)
        {
            BindType = bindType;
            Id = id;
            Lists = lists;
            ResultOutputFile = resultOutputFile;
            UsagePlanId = usagePlanId;
        }
    }
}
