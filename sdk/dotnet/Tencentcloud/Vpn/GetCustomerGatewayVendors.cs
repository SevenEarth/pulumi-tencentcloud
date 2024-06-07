// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Vpn
{
    public static class GetCustomerGatewayVendors
    {
        /// <summary>
        /// Use this data source to query detailed information of vpc vpn_customer_gateway_vendors
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
        ///     var vpnCustomerGatewayVendors = Tencentcloud.Vpn.GetCustomerGatewayVendors.Invoke();
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetCustomerGatewayVendorsResult> InvokeAsync(GetCustomerGatewayVendorsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCustomerGatewayVendorsResult>("tencentcloud:Vpn/getCustomerGatewayVendors:getCustomerGatewayVendors", args ?? new GetCustomerGatewayVendorsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of vpc vpn_customer_gateway_vendors
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
        ///     var vpnCustomerGatewayVendors = Tencentcloud.Vpn.GetCustomerGatewayVendors.Invoke();
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetCustomerGatewayVendorsResult> Invoke(GetCustomerGatewayVendorsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCustomerGatewayVendorsResult>("tencentcloud:Vpn/getCustomerGatewayVendors:getCustomerGatewayVendors", args ?? new GetCustomerGatewayVendorsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCustomerGatewayVendorsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetCustomerGatewayVendorsArgs()
        {
        }
        public static new GetCustomerGatewayVendorsArgs Empty => new GetCustomerGatewayVendorsArgs();
    }

    public sealed class GetCustomerGatewayVendorsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetCustomerGatewayVendorsInvokeArgs()
        {
        }
        public static new GetCustomerGatewayVendorsInvokeArgs Empty => new GetCustomerGatewayVendorsInvokeArgs();
    }


    [OutputType]
    public sealed class GetCustomerGatewayVendorsResult
    {
        /// <summary>
        /// Customer Gateway Vendor Set.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCustomerGatewayVendorsCustomerGatewayVendorSetResult> CustomerGatewayVendorSets;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetCustomerGatewayVendorsResult(
            ImmutableArray<Outputs.GetCustomerGatewayVendorsCustomerGatewayVendorSetResult> customerGatewayVendorSets,

            string id,

            string? resultOutputFile)
        {
            CustomerGatewayVendorSets = customerGatewayVendorSets;
            Id = id;
            ResultOutputFile = resultOutputFile;
        }
    }
}
