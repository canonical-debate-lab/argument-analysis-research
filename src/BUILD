package(default_visibility = ["//visibility:public"])

load("@io_bazel_rules_kotlin//kotlin:kotlin.bzl", "kt_jvm_library")

kt_jvm_library(
    name = "code",
    srcs = glob(["main/kotlin/**/*.kt"]),
    deps = [
        "//3rdparty/jvm/com/fasterxml/jackson/core:jackson_annotations",
        "//3rdparty/jvm/com/fasterxml/jackson/core:jackson_core",
        "//3rdparty/jvm/com/fasterxml/jackson/core:jackson_databind",
        "//3rdparty/jvm/com/fasterxml/jackson/module:jackson_module_kotlin",
        "//3rdparty/jvm/org/jetbrains/kotlin:kotlin_reflect",
        "//3rdparty/jvm/org/jetbrains/kotlin:kotlin_stdlib",
        "//lib:adw",
        "@appdeps//:compile",
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
    base = ":config_image",
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
    template = "//k8s:deploy-adw.yaml",
)

manifest(
    name = "svc",
    template = "//k8s:service-adw.yaml",
)

k8s_deploy(
    name = "service",
    template = ":svc",
)

k8s_deploy(
    name = "deployment",
    images = {
        "adw:dev": ":image",
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
