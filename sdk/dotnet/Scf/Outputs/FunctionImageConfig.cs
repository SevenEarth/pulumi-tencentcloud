// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Scf.Outputs
{

    [OutputType]
    public sealed class FunctionImageConfig
    {
        /// <summary>
        /// the parameters of command.
        /// </summary>
        public readonly string? Args;
        /// <summary>
        /// The command of entrypoint.
        /// </summary>
        public readonly string? Command;
        /// <summary>
        /// The entrypoint of app.
        /// </summary>
        public readonly string? EntryPoint;
        /// <summary>
        /// The image type. personal or enterprise.
        /// </summary>
        public readonly string ImageType;
        /// <summary>
        /// The uri of image.
        /// </summary>
        public readonly string ImageUri;
        /// <summary>
        /// The registry id of TCR. When image type is enterprise, it must be set.
        /// </summary>
        public readonly string? RegistryId;

        [OutputConstructor]
        private FunctionImageConfig(
            string? args,

            string? command,

            string? entryPoint,

            string imageType,

            string imageUri,

            string? registryId)
        {
            Args = args;
            Command = command;
            EntryPoint = entryPoint;
            ImageType = imageType;
            ImageUri = imageUri;
            RegistryId = registryId;
        }
    }
}
