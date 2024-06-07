// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Scf
{
    public static class GetAsyncEventStatus
    {
        /// <summary>
        /// Use this data source to query detailed information of scf async_event_status
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
        ///     var asyncEventStatus = Tencentcloud.Scf.GetAsyncEventStatus.Invoke(new()
        ///     {
        ///         InvokeRequestId = "9de9405a-e33a-498d-bb59-e80b7bed1191",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetAsyncEventStatusResult> InvokeAsync(GetAsyncEventStatusArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAsyncEventStatusResult>("tencentcloud:Scf/getAsyncEventStatus:getAsyncEventStatus", args ?? new GetAsyncEventStatusArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of scf async_event_status
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
        ///     var asyncEventStatus = Tencentcloud.Scf.GetAsyncEventStatus.Invoke(new()
        ///     {
        ///         InvokeRequestId = "9de9405a-e33a-498d-bb59-e80b7bed1191",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetAsyncEventStatusResult> Invoke(GetAsyncEventStatusInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAsyncEventStatusResult>("tencentcloud:Scf/getAsyncEventStatus:getAsyncEventStatus", args ?? new GetAsyncEventStatusInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAsyncEventStatusArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the async execution request.
        /// </summary>
        [Input("invokeRequestId", required: true)]
        public string InvokeRequestId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetAsyncEventStatusArgs()
        {
        }
        public static new GetAsyncEventStatusArgs Empty => new GetAsyncEventStatusArgs();
    }

    public sealed class GetAsyncEventStatusInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the async execution request.
        /// </summary>
        [Input("invokeRequestId", required: true)]
        public Input<string> InvokeRequestId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetAsyncEventStatusInvokeArgs()
        {
        }
        public static new GetAsyncEventStatusInvokeArgs Empty => new GetAsyncEventStatusInvokeArgs();
    }


    [OutputType]
    public sealed class GetAsyncEventStatusResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Async execution request ID.
        /// </summary>
        public readonly string InvokeRequestId;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// Async event status.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAsyncEventStatusResultResult> Results;

        [OutputConstructor]
        private GetAsyncEventStatusResult(
            string id,

            string invokeRequestId,

            string? resultOutputFile,

            ImmutableArray<Outputs.GetAsyncEventStatusResultResult> results)
        {
            Id = id;
            InvokeRequestId = invokeRequestId;
            ResultOutputFile = resultOutputFile;
            Results = results;
        }
    }
}
