// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tse
{
    public static class GetAccessAddress
    {
        /// <summary>
        /// Use this data source to query detailed information of tse access_address
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
        ///         var accessAddress = Output.Create(Tencentcloud.Tse.GetAccessAddress.InvokeAsync(new Tencentcloud.Tse.GetAccessAddressArgs
        ///         {
        ///             EngineRegion = "ap-guangzhou",
        ///             InstanceId = "ins-7eb7eea7",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAccessAddressResult> InvokeAsync(GetAccessAddressArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAccessAddressResult>("tencentcloud:Tse/getAccessAddress:getAccessAddress", args ?? new GetAccessAddressArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of tse access_address
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
        ///         var accessAddress = Output.Create(Tencentcloud.Tse.GetAccessAddress.InvokeAsync(new Tencentcloud.Tse.GetAccessAddressArgs
        ///         {
        ///             EngineRegion = "ap-guangzhou",
        ///             InstanceId = "ins-7eb7eea7",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetAccessAddressResult> Invoke(GetAccessAddressInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAccessAddressResult>("tencentcloud:Tse/getAccessAddress:getAccessAddress", args ?? new GetAccessAddressInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAccessAddressArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Deploy region.
        /// </summary>
        [Input("engineRegion")]
        public string? EngineRegion { get; set; }

        /// <summary>
        /// engine instance Id.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// Subnet ID, Zookeeper does not need to pass vpcid and subnetid; nacos and Polaris need to pass vpcid and subnetid.
        /// </summary>
        [Input("subnetId")]
        public string? SubnetId { get; set; }

        /// <summary>
        /// VPC ID, Zookeeper does not need to pass vpcid and subnetid; nacos and Polaris need to pass vpcid and subnetid.
        /// </summary>
        [Input("vpcId")]
        public string? VpcId { get; set; }

        /// <summary>
        /// Name of other engine components(pushgateway, polaris-limiter).
        /// </summary>
        [Input("workload")]
        public string? Workload { get; set; }

        public GetAccessAddressArgs()
        {
        }
    }

    public sealed class GetAccessAddressInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Deploy region.
        /// </summary>
        [Input("engineRegion")]
        public Input<string>? EngineRegion { get; set; }

        /// <summary>
        /// engine instance Id.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// Subnet ID, Zookeeper does not need to pass vpcid and subnetid; nacos and Polaris need to pass vpcid and subnetid.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        /// <summary>
        /// VPC ID, Zookeeper does not need to pass vpcid and subnetid; nacos and Polaris need to pass vpcid and subnetid.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// Name of other engine components(pushgateway, polaris-limiter).
        /// </summary>
        [Input("workload")]
        public Input<string>? Workload { get; set; }

        public GetAccessAddressInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAccessAddressResult
    {
        /// <summary>
        /// Console public network access addressNote: This field may return null, indicating that a valid value is not available.
        /// </summary>
        public readonly string ConsoleInternetAddress;
        /// <summary>
        /// Console public network bandwidthNote: This field may return null, indicating that a valid value is not available.
        /// </summary>
        public readonly int ConsoleInternetBandWidth;
        /// <summary>
        /// Console Intranet access addressNote: This field may return null, indicating that a valid value is not available.
        /// </summary>
        public readonly string ConsoleIntranetAddress;
        public readonly string? EngineRegion;
        /// <summary>
        /// Apollo Multi-environment public ip address.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAccessAddressEnvAddressInfoResult> EnvAddressInfos;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        /// <summary>
        /// Public access address.
        /// </summary>
        public readonly string InternetAddress;
        /// <summary>
        /// Client public network bandwidthNote: This field may return null, indicating that a valid value is not available.
        /// </summary>
        public readonly int InternetBandWidth;
        /// <summary>
        /// VPC access IP address listNote: This field may return null, indicating that a valid value is not available.
        /// </summary>
        public readonly string IntranetAddress;
        /// <summary>
        /// Access IP address of the Polaris traffic limiting server nodeNote: This field may return null, indicating that a valid value is not available.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAccessAddressLimiterAddressInfoResult> LimiterAddressInfos;
        public readonly string? ResultOutputFile;
        public readonly string? SubnetId;
        public readonly string? VpcId;
        public readonly string? Workload;

        [OutputConstructor]
        private GetAccessAddressResult(
            string consoleInternetAddress,

            int consoleInternetBandWidth,

            string consoleIntranetAddress,

            string? engineRegion,

            ImmutableArray<Outputs.GetAccessAddressEnvAddressInfoResult> envAddressInfos,

            string id,

            string instanceId,

            string internetAddress,

            int internetBandWidth,

            string intranetAddress,

            ImmutableArray<Outputs.GetAccessAddressLimiterAddressInfoResult> limiterAddressInfos,

            string? resultOutputFile,

            string? subnetId,

            string? vpcId,

            string? workload)
        {
            ConsoleInternetAddress = consoleInternetAddress;
            ConsoleInternetBandWidth = consoleInternetBandWidth;
            ConsoleIntranetAddress = consoleIntranetAddress;
            EngineRegion = engineRegion;
            EnvAddressInfos = envAddressInfos;
            Id = id;
            InstanceId = instanceId;
            InternetAddress = internetAddress;
            InternetBandWidth = internetBandWidth;
            IntranetAddress = intranetAddress;
            LimiterAddressInfos = limiterAddressInfos;
            ResultOutputFile = resultOutputFile;
            SubnetId = subnetId;
            VpcId = vpcId;
            Workload = workload;
        }
    }
}
