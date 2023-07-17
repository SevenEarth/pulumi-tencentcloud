// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Vpc
{
    public static class GetCvmInstances
    {
        /// <summary>
        /// Use this data source to query detailed information of vpc cvm_instances
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
        ///         var cvmInstances = Output.Create(Tencentcloud.Vpc.GetCvmInstances.InvokeAsync(new Tencentcloud.Vpc.GetCvmInstancesArgs
        ///         {
        ///             Filters = 
        ///             {
        ///                 new Tencentcloud.Vpc.Inputs.GetCvmInstancesFilterArgs
        ///                 {
        ///                     Name = "vpc-id",
        ///                     Values = 
        ///                     {
        ///                         "vpc-lh4nqig9",
        ///                     },
        ///                 },
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetCvmInstancesResult> InvokeAsync(GetCvmInstancesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCvmInstancesResult>("tencentcloud:Vpc/getCvmInstances:getCvmInstances", args ?? new GetCvmInstancesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of vpc cvm_instances
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
        ///         var cvmInstances = Output.Create(Tencentcloud.Vpc.GetCvmInstances.InvokeAsync(new Tencentcloud.Vpc.GetCvmInstancesArgs
        ///         {
        ///             Filters = 
        ///             {
        ///                 new Tencentcloud.Vpc.Inputs.GetCvmInstancesFilterArgs
        ///                 {
        ///                     Name = "vpc-id",
        ///                     Values = 
        ///                     {
        ///                         "vpc-lh4nqig9",
        ///                     },
        ///                 },
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetCvmInstancesResult> Invoke(GetCvmInstancesInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCvmInstancesResult>("tencentcloud:Vpc/getCvmInstances:getCvmInstances", args ?? new GetCvmInstancesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCvmInstancesArgs : Pulumi.InvokeArgs
    {
        [Input("filters", required: true)]
        private List<Inputs.GetCvmInstancesFilterArgs>? _filters;

        /// <summary>
        /// Filter condition. `RouteTableIds` and `Filters` cannot be specified at the same time. vpc-id - String - (Filter condition) VPC instance ID, such as `vpc-f49l6u0z`;instance-type - String - (Filter condition) CVM instance ID;instance-name - String - (Filter condition) CVM name.
        /// </summary>
        public List<Inputs.GetCvmInstancesFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetCvmInstancesFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetCvmInstancesArgs()
        {
        }
    }

    public sealed class GetCvmInstancesInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("filters", required: true)]
        private InputList<Inputs.GetCvmInstancesFilterInputArgs>? _filters;

        /// <summary>
        /// Filter condition. `RouteTableIds` and `Filters` cannot be specified at the same time. vpc-id - String - (Filter condition) VPC instance ID, such as `vpc-f49l6u0z`;instance-type - String - (Filter condition) CVM instance ID;instance-name - String - (Filter condition) CVM name.
        /// </summary>
        public InputList<Inputs.GetCvmInstancesFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetCvmInstancesFilterInputArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetCvmInstancesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCvmInstancesResult
    {
        public readonly ImmutableArray<Outputs.GetCvmInstancesFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// List of CVM instances.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCvmInstancesInstanceSetResult> InstanceSets;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetCvmInstancesResult(
            ImmutableArray<Outputs.GetCvmInstancesFilterResult> filters,

            string id,

            ImmutableArray<Outputs.GetCvmInstancesInstanceSetResult> instanceSets,

            string? resultOutputFile)
        {
            Filters = filters;
            Id = id;
            InstanceSets = instanceSets;
            ResultOutputFile = resultOutputFile;
        }
    }
}
