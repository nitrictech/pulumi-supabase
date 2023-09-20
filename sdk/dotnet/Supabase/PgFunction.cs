// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Supabase.Supabase
{
    [SupabaseResourceType("supabase:supabase:PgFunction")]
    public partial class PgFunction : global::Pulumi.CustomResource
    {
        [Output("args")]
        public Output<ImmutableArray<string>> Args { get; private set; } = null!;

        [Output("behaviour")]
        public Output<string> Behaviour { get; private set; } = null!;

        [Output("config_params")]
        public Output<ImmutableDictionary<string, object>?> Config_params { get; private set; } = null!;

        [Output("definition")]
        public Output<string> Definition { get; private set; } = null!;

        [Output("function_name")]
        public Output<string> Function_name { get; private set; } = null!;

        [Output("language")]
        public Output<string> Language { get; private set; } = null!;

        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        [Output("project_ref")]
        public Output<string> Project_ref { get; private set; } = null!;

        [Output("return_type")]
        public Output<string> Return_type { get; private set; } = null!;

        [Output("schema")]
        public Output<string> Schema { get; private set; } = null!;

        [Output("verify_jwt")]
        public Output<bool> Verify_jwt { get; private set; } = null!;


        /// <summary>
        /// Create a PgFunction resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PgFunction(string name, PgFunctionArgs args, CustomResourceOptions? options = null)
            : base("supabase:supabase:PgFunction", name, args ?? new PgFunctionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PgFunction(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("supabase:supabase:PgFunction", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing PgFunction resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PgFunction Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PgFunction(name, id, options);
        }
    }

    public sealed class PgFunctionArgs : global::Pulumi.ResourceArgs
    {
        [Input("args")]
        private InputList<string>? _args;
        public InputList<string> Args
        {
            get => _args ?? (_args = new InputList<string>());
            set => _args = value;
        }

        [Input("behaviour", required: true)]
        public Input<string> Behaviour { get; set; } = null!;

        [Input("config_params")]
        private InputMap<object>? _config_params;
        public InputMap<object> Config_params
        {
            get => _config_params ?? (_config_params = new InputMap<object>());
            set => _config_params = value;
        }

        [Input("definition", required: true)]
        public Input<string> Definition { get; set; } = null!;

        [Input("language", required: true)]
        public Input<string> Language { get; set; } = null!;

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project_ref", required: true)]
        public Input<string> Project_ref { get; set; } = null!;

        [Input("return_type", required: true)]
        public Input<string> Return_type { get; set; } = null!;

        [Input("schema", required: true)]
        public Input<string> Schema { get; set; } = null!;

        [Input("verify_jwt", required: true)]
        public Input<bool> Verify_jwt { get; set; } = null!;

        public PgFunctionArgs()
        {
        }
        public static new PgFunctionArgs Empty => new PgFunctionArgs();
    }
}
