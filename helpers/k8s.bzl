_TEMPLATE = "//defaults/k8s:deploy.yaml"

def _template_manifest_impl(ctx):
    name = "{}".format(ctx.label).replace("//", "").replace("cmd/", "").replace("/", "-").split(":", 1)[0]

    ctx.actions.expand_template(
        template = ctx.file.template,
        output = ctx.outputs.source_file,
        substitutions = {
            "{NAME}": name,
            "{STAGE}": ctx.attr.stage,
        },
    )

template_manifest = rule(
    attrs = {
        "template": attr.label(
            default = Label(_TEMPLATE),
            allow_single_file = True,
        ),
        "stage": attr.string(default = "dev"),
    },
    outputs = {"source_file": "%{name}.yaml"},
    implementation = _template_manifest_impl,
)

def template_image(ctx, *args, **kwargs):
    print(ctx, args, kwargs)
