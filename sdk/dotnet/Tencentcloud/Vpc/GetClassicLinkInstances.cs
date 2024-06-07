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
    public static class GetClassicLinkInstances
    {
        /// <summary>
        /// Use this data source to query detailed information of vpc classic_link_instances
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
        ///     var classicLinkInstances = Tencentcloud.Vpc.GetClassicLinkInstances.Invoke(new()
        ///     {
        ///         Filters = new[]
        ///         {
        ///             new Tencentcloud.Vpc.Inputs.GetClassicLinkInstancesFilterInputArgs
        ///             {
        ///                 Name = "vpc-id",
        ///                 Values = new[]
        ///                 {
        ///                     "vpc-lh4nqig9",
        ///                 },
        ///             },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetClassicLinkInstancesResult> InvokeAsync(GetClassicLinkInstancesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetClassicLinkInstancesResult>("tencentcloud:Vpc/getClassicLinkInstances:getClassicLinkInstances", args ?? new GetClassicLinkInstancesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of vpc classic_link_instances
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
        ///     var classicLinkInstances = Tencentcloud.Vpc.GetClassicLinkInstances.Invoke(new()
        ///     {
        ///         Filters = new[]
        ///         {
        ///             new Tencentcloud.Vpc.Inputs.GetClassicLinkInstancesFilterInputArgs
        ///             {
        ///                 Name = "vpc-id",
        ///                 Values = new[]
        ///                 {
        ///                     "vpc-lh4nqig9",
        ///                 },
        ///             },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetClassicLinkInstancesResult> Invoke(GetClassicLinkInstancesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetClassicLinkInstancesResult>("tencentcloud:Vpc/getClassicLinkInstances:getClassicLinkInstances", args ?? new GetClassicLinkInstancesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetClassicLinkInstancesArgs : global::Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetClassicLinkInstancesFilterArgs>? _filters;

        /// <summary>
        /// Filter conditions.`vpc-id` - String - (Filter condition) The VPC instance ID. `vm-ip` - String - (Filter condition) The IP address of the CVM on the basic network.
        /// </summary>
        public List<Inputs.GetClassicLinkInstancesFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetClassicLinkInstancesFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetClassicLinkInstancesArgs()
        {
        }
        public static new GetClassicLinkInstancesArgs Empty => new GetClassicLinkInstancesArgs();
    }

    public sealed class GetClassicLinkInstancesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("filters")]
        private InputList<Inputs.GetClassicLinkInstancesFilterInputArgs>? _filters;

        /// <summary>
        /// Filter conditions.`vpc-id` - String - (Filter condition) The VPC instance ID. `vm-ip` - String - (Filter condition) The IP address of the CVM on the basic network.
        /// </summary>
        public InputList<Inputs.GetClassicLinkInstancesFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetClassicLinkInstancesFilterInputArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetClassicLinkInstancesInvokeArgs()
        {
        }
        public static new GetClassicLinkInstancesInvokeArgs Empty => new GetClassicLinkInstancesInvokeArgs();
    }


    [OutputType]
    public sealed class GetClassicLinkInstancesResult
    {
        /// <summary>
        /// Classiclink instance.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetClassicLinkInstancesClassicLinkInstanceSetResult> ClassicLinkInstanceSets;
        public readonly ImmutableArray<Outputs.GetClassicLinkInstancesFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetClassicLinkInstancesResult(
            ImmutableArray<Outputs.GetClassicLinkInstancesClassicLinkInstanceSetResult> classicLinkInstanceSets,

            ImmutableArray<Outputs.GetClassicLinkInstancesFilterResult> filters,

            string id,

            string? resultOutputFile)
        {
            ClassicLinkInstanceSets = classicLinkInstanceSets;
            Filters = filters;
            Id = id;
            ResultOutputFile = resultOutputFile;
        }
    }
}
