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
    public static class GetScalingConfigs
    {
        /// <summary>
        /// Use this data source to query scaling configuration information.
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
        ///     var asConfigs = Tencentcloud.As.GetScalingConfigs.Invoke(new()
        ///     {
        ///         ConfigurationId = "asc-oqio4yyj",
        ///         ResultOutputFile = "my_test_path",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetScalingConfigsResult> InvokeAsync(GetScalingConfigsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetScalingConfigsResult>("tencentcloud:As/getScalingConfigs:getScalingConfigs", args ?? new GetScalingConfigsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query scaling configuration information.
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
        ///     var asConfigs = Tencentcloud.As.GetScalingConfigs.Invoke(new()
        ///     {
        ///         ConfigurationId = "asc-oqio4yyj",
        ///         ResultOutputFile = "my_test_path",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetScalingConfigsResult> Invoke(GetScalingConfigsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetScalingConfigsResult>("tencentcloud:As/getScalingConfigs:getScalingConfigs", args ?? new GetScalingConfigsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetScalingConfigsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Launch configuration ID.
        /// </summary>
        [Input("configurationId")]
        public string? ConfigurationId { get; set; }

        /// <summary>
        /// Launch configuration name.
        /// </summary>
        [Input("configurationName")]
        public string? ConfigurationName { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetScalingConfigsArgs()
        {
        }
        public static new GetScalingConfigsArgs Empty => new GetScalingConfigsArgs();
    }

    public sealed class GetScalingConfigsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Launch configuration ID.
        /// </summary>
        [Input("configurationId")]
        public Input<string>? ConfigurationId { get; set; }

        /// <summary>
        /// Launch configuration name.
        /// </summary>
        [Input("configurationName")]
        public Input<string>? ConfigurationName { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetScalingConfigsInvokeArgs()
        {
        }
        public static new GetScalingConfigsInvokeArgs Empty => new GetScalingConfigsInvokeArgs();
    }


    [OutputType]
    public sealed class GetScalingConfigsResult
    {
        /// <summary>
        /// Launch configuration ID.
        /// </summary>
        public readonly string? ConfigurationId;
        /// <summary>
        /// A list of configuration. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetScalingConfigsConfigurationListResult> ConfigurationLists;
        /// <summary>
        /// Launch configuration name.
        /// </summary>
        public readonly string? ConfigurationName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetScalingConfigsResult(
            string? configurationId,

            ImmutableArray<Outputs.GetScalingConfigsConfigurationListResult> configurationLists,

            string? configurationName,

            string id,

            string? resultOutputFile)
        {
            ConfigurationId = configurationId;
            ConfigurationLists = configurationLists;
            ConfigurationName = configurationName;
            Id = id;
            ResultOutputFile = resultOutputFile;
        }
    }
}
