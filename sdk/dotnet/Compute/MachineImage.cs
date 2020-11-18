// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    /// <summary>
    /// Represents a Machine Image resource. Machine images store all the configuration,
    /// metadata, permissions, and data from one or more disks required to create a
    /// Virtual machine (VM) instance.
    /// 
    /// To get more information about MachineImage, see:
    /// 
    /// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/machineImages)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/compute/docs/machine-images)
    /// 
    /// ## Example Usage
    /// ### Machine Image Basic
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var vm = new Gcp.Compute.Instance("vm", new Gcp.Compute.InstanceArgs
    ///         {
    ///             MachineType = "e2-medium",
    ///             BootDisk = new Gcp.Compute.Inputs.InstanceBootDiskArgs
    ///             {
    ///                 InitializeParams = new Gcp.Compute.Inputs.InstanceBootDiskInitializeParamsArgs
    ///                 {
    ///                     Image = "debian-cloud/debian-9",
    ///                 },
    ///             },
    ///             NetworkInterfaces = 
    ///             {
    ///                 new Gcp.Compute.Inputs.InstanceNetworkInterfaceArgs
    ///                 {
    ///                     Network = "default",
    ///                 },
    ///             },
    ///         }, new CustomResourceOptions
    ///         {
    ///             Provider = google_beta,
    ///         });
    ///         var image = new Gcp.Compute.MachineImage("image", new Gcp.Compute.MachineImageArgs
    ///         {
    ///             SourceInstance = vm.SelfLink,
    ///         }, new CustomResourceOptions
    ///         {
    ///             Provider = google_beta,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// MachineImage can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/machineImage:MachineImage default projects/{{project}}/global/machineImages/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/machineImage:MachineImage default {{project}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/machineImage:MachineImage default {{name}}
    /// ```
    /// </summary>
    public partial class MachineImage : Pulumi.CustomResource
    {
        /// <summary>
        /// A text description of the resource.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Specify this to create an application consistent machine image by informing the OS to prepare for the snapshot process.
        /// Currently only supported on Windows instances using the Volume Shadow Copy Service (VSS).
        /// </summary>
        [Output("guestFlush")]
        public Output<bool?> GuestFlush { get; private set; } = null!;

        /// <summary>
        /// Encrypts the machine image using a customer-supplied encryption key.
        /// After you encrypt a machine image with a customer-supplied key, you must
        /// provide the same key if you use the machine image later (e.g. to create a
        /// instance from the image)
        /// Structure is documented below.
        /// </summary>
        [Output("machineImageEncryptionKey")]
        public Output<Outputs.MachineImageMachineImageEncryptionKey?> MachineImageEncryptionKey { get; private set; } = null!;

        /// <summary>
        /// Name of the resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// The source instance used to create the machine image. You can provide this as a partial or full URL to the resource.
        /// </summary>
        [Output("sourceInstance")]
        public Output<string> SourceInstance { get; private set; } = null!;

        /// <summary>
        /// The regional or multi-regional Cloud Storage bucket location where the machine image is stored.
        /// </summary>
        [Output("storageLocations")]
        public Output<ImmutableArray<string>> StorageLocations { get; private set; } = null!;


        /// <summary>
        /// Create a MachineImage resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MachineImage(string name, MachineImageArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/machineImage:MachineImage", name, args ?? new MachineImageArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MachineImage(string name, Input<string> id, MachineImageState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/machineImage:MachineImage", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing MachineImage resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MachineImage Get(string name, Input<string> id, MachineImageState? state = null, CustomResourceOptions? options = null)
        {
            return new MachineImage(name, id, state, options);
        }
    }

    public sealed class MachineImageArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A text description of the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specify this to create an application consistent machine image by informing the OS to prepare for the snapshot process.
        /// Currently only supported on Windows instances using the Volume Shadow Copy Service (VSS).
        /// </summary>
        [Input("guestFlush")]
        public Input<bool>? GuestFlush { get; set; }

        /// <summary>
        /// Encrypts the machine image using a customer-supplied encryption key.
        /// After you encrypt a machine image with a customer-supplied key, you must
        /// provide the same key if you use the machine image later (e.g. to create a
        /// instance from the image)
        /// Structure is documented below.
        /// </summary>
        [Input("machineImageEncryptionKey")]
        public Input<Inputs.MachineImageMachineImageEncryptionKeyArgs>? MachineImageEncryptionKey { get; set; }

        /// <summary>
        /// Name of the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The source instance used to create the machine image. You can provide this as a partial or full URL to the resource.
        /// </summary>
        [Input("sourceInstance", required: true)]
        public Input<string> SourceInstance { get; set; } = null!;

        public MachineImageArgs()
        {
        }
    }

    public sealed class MachineImageState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A text description of the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specify this to create an application consistent machine image by informing the OS to prepare for the snapshot process.
        /// Currently only supported on Windows instances using the Volume Shadow Copy Service (VSS).
        /// </summary>
        [Input("guestFlush")]
        public Input<bool>? GuestFlush { get; set; }

        /// <summary>
        /// Encrypts the machine image using a customer-supplied encryption key.
        /// After you encrypt a machine image with a customer-supplied key, you must
        /// provide the same key if you use the machine image later (e.g. to create a
        /// instance from the image)
        /// Structure is documented below.
        /// </summary>
        [Input("machineImageEncryptionKey")]
        public Input<Inputs.MachineImageMachineImageEncryptionKeyGetArgs>? MachineImageEncryptionKey { get; set; }

        /// <summary>
        /// Name of the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// The source instance used to create the machine image. You can provide this as a partial or full URL to the resource.
        /// </summary>
        [Input("sourceInstance")]
        public Input<string>? SourceInstance { get; set; }

        [Input("storageLocations")]
        private InputList<string>? _storageLocations;

        /// <summary>
        /// The regional or multi-regional Cloud Storage bucket location where the machine image is stored.
        /// </summary>
        public InputList<string> StorageLocations
        {
            get => _storageLocations ?? (_storageLocations = new InputList<string>());
            set => _storageLocations = value;
        }

        public MachineImageState()
        {
        }
    }
}
