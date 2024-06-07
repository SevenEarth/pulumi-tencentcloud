// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cam
{
    public static class GetSamlProviders
    {
        /// <summary>
        /// Use this data source to query detailed information of CAM SAML providers
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
        ///     var foo = Tencentcloud.Cam.GetSamlProviders.Invoke(new()
        ///     {
        ///         Name = "cam-test-provider",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetSamlProvidersResult> InvokeAsync(GetSamlProvidersArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSamlProvidersResult>("tencentcloud:Cam/getSamlProviders:getSamlProviders", args ?? new GetSamlProvidersArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of CAM SAML providers
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
        ///     var foo = Tencentcloud.Cam.GetSamlProviders.Invoke(new()
        ///     {
        ///         Name = "cam-test-provider",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetSamlProvidersResult> Invoke(GetSamlProvidersInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSamlProvidersResult>("tencentcloud:Cam/getSamlProviders:getSamlProviders", args ?? new GetSamlProvidersInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSamlProvidersArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The description of the CAM SAML provider.
        /// </summary>
        [Input("description")]
        public string? Description { get; set; }

        /// <summary>
        /// Name of the CAM SAML provider to be queried.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetSamlProvidersArgs()
        {
        }
        public static new GetSamlProvidersArgs Empty => new GetSamlProvidersArgs();
    }

    public sealed class GetSamlProvidersInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The description of the CAM SAML provider.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the CAM SAML provider to be queried.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetSamlProvidersInvokeArgs()
        {
        }
        public static new GetSamlProvidersInvokeArgs Empty => new GetSamlProvidersInvokeArgs();
    }


    [OutputType]
    public sealed class GetSamlProvidersResult
    {
        /// <summary>
        /// Description of CAM SAML provider.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Name of CAM SAML provider.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// A list of CAM SAML providers. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSamlProvidersProviderListResult> ProviderLists;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetSamlProvidersResult(
            string? description,

            string id,

            string? name,

            ImmutableArray<Outputs.GetSamlProvidersProviderListResult> providerLists,

            string? resultOutputFile)
        {
            Description = description;
            Id = id;
            Name = name;
            ProviderLists = providerLists;
            ResultOutputFile = resultOutputFile;
        }
    }
}
