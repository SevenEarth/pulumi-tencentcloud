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
    public static class GetServices
    {
        /// <summary>
        /// Use this data source to query API gateway services.
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
        ///     var name = Tencentcloud.ApiGateway.GetServices.Invoke(new()
        ///     {
        ///         ServiceName = service.ServiceName,
        ///     });
        /// 
        ///     var id = Tencentcloud.ApiGateway.GetServices.Invoke(new()
        ///     {
        ///         ServiceId = service.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetServicesResult> InvokeAsync(GetServicesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetServicesResult>("tencentcloud:ApiGateway/getServices:getServices", args ?? new GetServicesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query API gateway services.
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
        ///     var name = Tencentcloud.ApiGateway.GetServices.Invoke(new()
        ///     {
        ///         ServiceName = service.ServiceName,
        ///     });
        /// 
        ///     var id = Tencentcloud.ApiGateway.GetServices.Invoke(new()
        ///     {
        ///         ServiceId = service.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetServicesResult> Invoke(GetServicesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetServicesResult>("tencentcloud:ApiGateway/getServices:getServices", args ?? new GetServicesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServicesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// Service ID for query.
        /// </summary>
        [Input("serviceId")]
        public string? ServiceId { get; set; }

        /// <summary>
        /// Service name for query.
        /// </summary>
        [Input("serviceName")]
        public string? ServiceName { get; set; }

        public GetServicesArgs()
        {
        }
        public static new GetServicesArgs Empty => new GetServicesArgs();
    }

    public sealed class GetServicesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// Service ID for query.
        /// </summary>
        [Input("serviceId")]
        public Input<string>? ServiceId { get; set; }

        /// <summary>
        /// Service name for query.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        public GetServicesInvokeArgs()
        {
        }
        public static new GetServicesInvokeArgs Empty => new GetServicesInvokeArgs();
    }


    [OutputType]
    public sealed class GetServicesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of services.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetServicesListResult> Lists;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// Custom service ID.
        /// </summary>
        public readonly string? ServiceId;
        /// <summary>
        /// Custom service name.
        /// </summary>
        public readonly string? ServiceName;

        [OutputConstructor]
        private GetServicesResult(
            string id,

            ImmutableArray<Outputs.GetServicesListResult> lists,

            string? resultOutputFile,

            string? serviceId,

            string? serviceName)
        {
            Id = id;
            Lists = lists;
            ResultOutputFile = resultOutputFile;
            ServiceId = serviceId;
            ServiceName = serviceName;
        }
    }
}
