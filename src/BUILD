package(default_visibility = ["//visibility:public"])

load("@io_bazel_rules_kotlin//kotlin:kotlin.bzl", "kt_jvm_library")

kt_jvm_library(
    name = "code",
    srcs = glob(["main/kotlin/**/*.kt"]),
    data = ["//config:files"],
    deps = [
        "//lib:adw",
        "@maven//:com_fasterxml_jackson_core_jackson_databind",
        "@maven//:com_fasterxml_jackson_module_jackson_module_kotlin",
        "@maven//:io_javalin_javalin",
        "@maven//:org_slf4j_slf4j_api",
        "@maven//:org_slf4j_slf4j_simple",
    ],
)

load("@io_bazel_rules_docker//java:image.bzl", "java_image")
load("@io_bazel_rules_docker//container:image.bzl", "container_image")

container_image(
    name = "resource_image",
    base = "@java_image_base//image",
    data_path = "/app/src/image.binary.runfiles/research/",
    directory = "/app/src/image.binary.runfiles/research/",
    files = [
        "//resources:files",
    ],
)

container_image(
    name = "config_image",
    base = ":resource_image",
    data_path = "/app/src/image.binary.runfiles/research/",
    directory = "/app/src/image.binary.runfiles/research/",
    files = [
        "//config:files",
    ],
)

java_image(
    name = "image",
    main_class = "com.github.cdl.adw.ServiceKt",
    runtime_deps = [
        ":code",
    ],
)

load("@io_bazel_rules_k8s//k8s:objects.bzl", "k8s_objects")
load("@k8s_deploy//:defaults.bzl", "k8s_deploy")
load(
    "//:helpers/k8s.bzl",
    manifest = "template_manifest",
)

manifest(
    name = "dpl",
    template = "//defaults/k8s:deploy-adw.yaml",
)

manifest(
    name = "svc",
    template = "//defaults/k8s:service-adw.yaml",
)

k8s_deploy(
    name = "service",
    template = ":svc",
)

k8s_deploy(
    name = "deployment",
    images = {
        ":dev": ":image",
    },
    template = ":dpl",
)

k8s_objects(
    name = "dev",
    objects = [
        ":deployment",
        ":service",
    ],
)
