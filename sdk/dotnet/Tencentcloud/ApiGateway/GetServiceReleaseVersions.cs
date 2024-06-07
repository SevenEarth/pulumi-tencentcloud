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
    public static class GetServiceReleaseVersions
    {
        /// <summary>
        /// Use this data source to query detailed information of apiGateway service_release_versions
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
        ///     var example = Tencentcloud.ApiGateway.GetServiceReleaseVersions.Invoke(new()
        ///     {
        ///         ServiceId = "service-nxz6yync",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetServiceReleaseVersionsResult> InvokeAsync(GetServiceReleaseVersionsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetServiceReleaseVersionsResult>("tencentcloud:ApiGateway/getServiceReleaseVersions:getServiceReleaseVersions", args ?? new GetServiceReleaseVersionsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of apiGateway service_release_versions
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
        ///     var example = Tencentcloud.ApiGateway.GetServiceReleaseVersions.Invoke(new()
        ///     {
        ///         ServiceId = "service-nxz6yync",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetServiceReleaseVersionsResult> Invoke(GetServiceReleaseVersionsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetServiceReleaseVersionsResult>("tencentcloud:ApiGateway/getServiceReleaseVersions:getServiceReleaseVersions", args ?? new GetServiceReleaseVersionsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServiceReleaseVersionsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// The unique ID of the service to be queried.
        /// </summary>
        [Input("serviceId", required: true)]
        public string ServiceId { get; set; } = null!;

        public GetServiceReleaseVersionsArgs()
        {
        }
        public static new GetServiceReleaseVersionsArgs Empty => new GetServiceReleaseVersionsArgs();
    }

    public sealed class GetServiceReleaseVersionsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// The unique ID of the service to be queried.
        /// </summary>
        [Input("serviceId", required: true)]
        public Input<string> ServiceId { get; set; } = null!;

        public GetServiceReleaseVersionsInvokeArgs()
        {
        }
        public static new GetServiceReleaseVersionsInvokeArgs Empty => new GetServiceReleaseVersionsInvokeArgs();
    }


    [OutputType]
    public sealed class GetServiceReleaseVersionsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// List of service releases.Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetServiceReleaseVersionsResultResult> Results;
        public readonly string ServiceId;

        [OutputConstructor]
        private GetServiceReleaseVersionsResult(
            string id,

            string? resultOutputFile,

            ImmutableArray<Outputs.GetServiceReleaseVersionsResultResult> results,

            string serviceId)
        {
            Id = id;
            ResultOutputFile = resultOutputFile;
            Results = results;
            ServiceId = serviceId;
        }
    }
}
