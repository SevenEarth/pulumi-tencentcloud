// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Lighthouse.Inputs
{

    public sealed class GetDisksFilterInputArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Fields to be filtered. Valid names: `disk-id`: Filters by disk id; `instance-id`: Filter by instance id; `disk-name`: Filter by disk name; `zone`: Filter by zone; `disk-usage`: Filter by disk usage(Values: `SYSTEM_DISK` or `DATA_DISK`); `disk-state`: Filter by disk state.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("values", required: true)]
        private InputList<string>? _values;

        /// <summary>
        /// Value of the field.
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public GetDisksFilterInputArgs()
        {
        }
    }
}
