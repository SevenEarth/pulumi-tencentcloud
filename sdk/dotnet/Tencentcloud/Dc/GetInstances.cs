// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dc
{
    public static class GetInstances
    {
        /// <summary>
        /// Use this data source to query detailed information of DC instances.
        /// </summary>
        public static Task<GetInstancesResult> InvokeAsync(GetInstancesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInstancesResult>("tencentcloud:Dc/getInstances:getInstances", args ?? new GetInstancesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of DC instances.
        /// </summary>
        public static Output<GetInstancesResult> Invoke(GetInstancesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInstancesResult>("tencentcloud:Dc/getInstances:getInstances", args ?? new GetInstancesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstancesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the DC to be queried.
        /// </summary>
        [Input("dcId")]
        public string? DcId { get; set; }

        /// <summary>
        /// Name of the DC to be queried.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetInstancesArgs()
        {
        }
        public static new GetInstancesArgs Empty => new GetInstancesArgs();
    }

    public sealed class GetInstancesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the DC to be queried.
        /// </summary>
        [Input("dcId")]
        public Input<string>? DcId { get; set; }

        /// <summary>
        /// Name of the DC to be queried.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetInstancesInvokeArgs()
        {
        }
        public static new GetInstancesInvokeArgs Empty => new GetInstancesInvokeArgs();
    }


    [OutputType]
    public sealed class GetInstancesResult
    {
        /// <summary>
        /// ID of the DC.
        /// </summary>
        public readonly string? DcId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Information list of the DC.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstancesInstanceListResult> InstanceLists;
        /// <summary>
        /// Name of the DC.
        /// </summary>
        public readonly string? Name;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetInstancesResult(
            string? dcId,

            string id,

            ImmutableArray<Outputs.GetInstancesInstanceListResult> instanceLists,

            string? name,

            string? resultOutputFile)
        {
            DcId = dcId;
            Id = id;
            InstanceLists = instanceLists;
            Name = name;
            ResultOutputFile = resultOutputFile;
        }
    }
}
