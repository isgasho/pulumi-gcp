// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Storage.Inputs
{

    public sealed class TransferJobTransferSpecAwsS3DataSourceGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// AWS credentials block.
        /// </summary>
        [Input("awsAccessKey", required: true)]
        public Input<Inputs.TransferJobTransferSpecAwsS3DataSourceAwsAccessKeyGetArgs> AwsAccessKey { get; set; } = null!;

        /// <summary>
        /// S3 Bucket name.
        /// </summary>
        [Input("bucketName", required: true)]
        public Input<string> BucketName { get; set; } = null!;

        public TransferJobTransferSpecAwsS3DataSourceGetArgs()
        {
        }
    }
}
