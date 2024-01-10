// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ssm
{
    public static class GetServiceStatus
    {
        /// <summary>
        /// Use this data source to query detailed information of ssm service_status
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Tencentcloud.Ssm.GetServiceStatus.InvokeAsync());
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetServiceStatusResult> InvokeAsync(GetServiceStatusArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServiceStatusResult>("tencentcloud:Ssm/getServiceStatus:getServiceStatus", args ?? new GetServiceStatusArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of ssm service_status
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Tencentcloud.Ssm.GetServiceStatus.InvokeAsync());
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetServiceStatusResult> Invoke(GetServiceStatusInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetServiceStatusResult>("tencentcloud:Ssm/getServiceStatus:getServiceStatus", args ?? new GetServiceStatusInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServiceStatusArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetServiceStatusArgs()
        {
        }
    }

    public sealed class GetServiceStatusInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetServiceStatusInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetServiceStatusResult
    {
        /// <summary>
        /// True means that the user can already use the key safe hosting function, false means that the user cannot use the key safe hosting function temporarily.
        /// </summary>
        public readonly bool AccessKeyEscrowEnabled;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Service unavailability type: 0-Not purchased, 1-Normal, 2-Service suspended due to arrears, 3-Resource release.
        /// </summary>
        public readonly int InvalidType;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// True means the service has been activated, false means the service has not been activated yet.
        /// </summary>
        public readonly bool ServiceEnabled;

        [OutputConstructor]
        private GetServiceStatusResult(
            bool accessKeyEscrowEnabled,

            string id,

            int invalidType,

            string? resultOutputFile,

            bool serviceEnabled)
        {
            AccessKeyEscrowEnabled = accessKeyEscrowEnabled;
            Id = id;
            InvalidType = invalidType;
            ResultOutputFile = resultOutputFile;
            ServiceEnabled = serviceEnabled;
        }
    }
}
