// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.As
{
    public static class GetInstances
    {
        /// <summary>
        /// Use this data source to query detailed information of as instances
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var scalingGroup = new Tencentcloud.As.ScalingGroup("scalingGroup", new Tencentcloud.As.ScalingGroupArgs
        ///         {
        ///             ScalingGroupName = "tf-as-group-ds-ins-basic",
        ///             ConfigurationId = "your_launch_configuration_id",
        ///             MaxSize = 1,
        ///             MinSize = 1,
        ///             VpcId = "your_vpc_id",
        ///             SubnetIds = 
        ///             {
        ///                 "your_subnet_id",
        ///             },
        ///             Tags = 
        ///             {
        ///                 { "test", "test" },
        ///             },
        ///         });
        ///         var instances = Tencentcloud.As.GetInstances.Invoke(new Tencentcloud.As.GetInstancesInvokeArgs
        ///         {
        ///             Filters = 
        ///             {
        ///                 new Tencentcloud.As.Inputs.GetInstancesFilterInputArgs
        ///                 {
        ///                     Name = "auto-scaling-group-id",
        ///                     Values = 
        ///                     {
        ///                         scalingGroup.Id,
        ///                     },
        ///                 },
        ///             },
        ///         });
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetInstancesResult> InvokeAsync(GetInstancesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstancesResult>("tencentcloud:As/getInstances:getInstances", args ?? new GetInstancesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of as instances
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var scalingGroup = new Tencentcloud.As.ScalingGroup("scalingGroup", new Tencentcloud.As.ScalingGroupArgs
        ///         {
        ///             ScalingGroupName = "tf-as-group-ds-ins-basic",
        ///             ConfigurationId = "your_launch_configuration_id",
        ///             MaxSize = 1,
        ///             MinSize = 1,
        ///             VpcId = "your_vpc_id",
        ///             SubnetIds = 
        ///             {
        ///                 "your_subnet_id",
        ///             },
        ///             Tags = 
        ///             {
        ///                 { "test", "test" },
        ///             },
        ///         });
        ///         var instances = Tencentcloud.As.GetInstances.Invoke(new Tencentcloud.As.GetInstancesInvokeArgs
        ///         {
        ///             Filters = 
        ///             {
        ///                 new Tencentcloud.As.Inputs.GetInstancesFilterInputArgs
        ///                 {
        ///                     Name = "auto-scaling-group-id",
        ///                     Values = 
        ///                     {
        ///                         scalingGroup.Id,
        ///                     },
        ///                 },
        ///             },
        ///         });
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetInstancesResult> Invoke(GetInstancesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetInstancesResult>("tencentcloud:As/getInstances:getInstances", args ?? new GetInstancesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstancesArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetInstancesFilterArgs>? _filters;

        /// <summary>
        /// Filter conditions. If there are multiple Filters, the relationship between Filters is a logical AND (AND) relationship. If there are multiple Values in the same Filter, the relationship between Values under the same Filter is a logical OR (OR) relationship.
        /// </summary>
        public List<Inputs.GetInstancesFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetInstancesFilterArgs>());
            set => _filters = value;
        }

        [Input("instanceIds")]
        private List<string>? _instanceIds;

        /// <summary>
        /// Instance ID of the cloud server (CVM) to be queried. The limit is 100 per request.
        /// </summary>
        public List<string> InstanceIds
        {
            get => _instanceIds ?? (_instanceIds = new List<string>());
            set => _instanceIds = value;
        }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetInstancesArgs()
        {
        }
    }

    public sealed class GetInstancesInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private InputList<Inputs.GetInstancesFilterInputArgs>? _filters;

        /// <summary>
        /// Filter conditions. If there are multiple Filters, the relationship between Filters is a logical AND (AND) relationship. If there are multiple Values in the same Filter, the relationship between Values under the same Filter is a logical OR (OR) relationship.
        /// </summary>
        public InputList<Inputs.GetInstancesFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetInstancesFilterInputArgs>());
            set => _filters = value;
        }

        [Input("instanceIds")]
        private InputList<string>? _instanceIds;

        /// <summary>
        /// Instance ID of the cloud server (CVM) to be queried. The limit is 100 per request.
        /// </summary>
        public InputList<string> InstanceIds
        {
            get => _instanceIds ?? (_instanceIds = new InputList<string>());
            set => _instanceIds = value;
        }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetInstancesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetInstancesResult
    {
        public readonly ImmutableArray<Outputs.GetInstancesFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> InstanceIds;
        /// <summary>
        /// List of instance details.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstancesInstanceListResult> InstanceLists;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetInstancesResult(
            ImmutableArray<Outputs.GetInstancesFilterResult> filters,

            string id,

            ImmutableArray<string> instanceIds,

            ImmutableArray<Outputs.GetInstancesInstanceListResult> instanceLists,

            string? resultOutputFile)
        {
            Filters = filters;
            Id = id;
            InstanceIds = instanceIds;
            InstanceLists = instanceLists;
            ResultOutputFile = resultOutputFile;
        }
    }
}
