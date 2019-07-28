workspace(name = "research")

load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "f04d2373bcaf8aa09bccb08a98a57e721306c8f6043a2a0ee610fd6853dcde3d",
    urls = [
        "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/rules_go/releases/download/0.18.6/rules_go-0.18.6.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/0.18.6/rules_go-0.18.6.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()

git_repository(
    name = "bazel_gazelle",
    commit = "81dc0cd8d440cab58dc18c5200708ff774be79c8",
    remote = "https://github.com/bazelbuild/bazel-gazelle",
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

# Imports Docker rules for Bazel (e.g. docker_image)
git_repository(
    name = "io_bazel_rules_docker",
    commit = "3732c9d05315bef6a3dbd195c545d6fea3b86880",
    remote = "https://github.com/bazelbuild/rules_docker.git",
    shallow_since = "1547471117 +0100",
)

# Loads Docker rules for Bazel
load(
    "@io_bazel_rules_docker//repositories:repositories.bzl",
    container_repositories = "repositories",
)

container_repositories()

load(
    "@io_bazel_rules_docker//go:image.bzl",
    _go_image_repos = "repositories",
)

_go_image_repos()

git_repository(
    name = "io_bazel_rules_k8s",
    commit = "fe8a5192d37103cc6de01676ead7fca1f55d5a0e",
    remote = "https://github.com/seibert-media/rules_k8s.git",
    shallow_since = "1561385037 -0700",
)

load("@io_bazel_rules_k8s//k8s:k8s.bzl", "k8s_repositories")

k8s_repositories()

load("@io_bazel_rules_k8s//k8s:with-defaults.bzl", "k8s_defaults")

k8s_defaults(
    name = "k8s_deploy",
    cluster = "c1",
    image_chroot = "eu.gcr.io/argument-analysis-images/",
    namespace = "{STABLE_KUBE_NAMESPACE}",
)

# # Imports Maven for Java/Kotlin
# git_repository(
#     name = "org_pubref_rules_maven",
#     commit = "339c378f856461add63f155d82077de5813e649e",
#     remote = "https://github.com/pubref/rules_maven",
# )

rules_kotlin_version = "4c71740a1b63b785fc90afd8d4d4d5bfda527107"

http_archive(
    name = "io_bazel_rules_kotlin",
    strip_prefix = "rules_kotlin-%s" % rules_kotlin_version,
    type = "zip",
    urls = ["https://github.com/bazelbuild/rules_kotlin/archive/%s.zip" % rules_kotlin_version],
)

load("@io_bazel_rules_kotlin//kotlin:kotlin.bzl", "kotlin_repositories", "kt_register_toolchains")
load("@io_bazel_rules_kotlin//kotlin:kotlin.bzl", "kt_jvm_import")

kotlin_repositories()

kt_register_toolchains()

# load("@org_pubref_rules_maven//maven:rules.bzl", "maven_repositories", "maven_repository")
# maven_repositories()
# load("//3rdparty:workspace.bzl", "maven_dependencies")
# maven_dependencies()
# maven_repository(
#     name = "appdeps",
#     transitive_deps = [
#         "73836e9cf29f978e47817584f9cee86b5e1f4c09:io.javalin:javalin:2.3.0",
#         "3cd63d075497751784b2fa84be59432f4905bf7c:javax.servlet:javax.servlet-api:3.1.0",
#         "1d329d68f31dce13135243c06013aaf6f708f7e7:org.eclipse.jetty:jetty-client:9.4.12.v20180830",
#         "1341796dde4e16df69bca83f3e87688ba2e7d703:org.eclipse.jetty:jetty-http:9.4.12.v20180830",
#         "e93f5adaa35a9a6a85ba130f589c5305c6ecc9e3:org.eclipse.jetty:jetty-io:9.4.12.v20180830",
#         "299e0602a9c0b753ba232cc1c1dda72ddd9addcf:org.eclipse.jetty:jetty-security:9.4.12.v20180830",
#         "b0f25df0d32a445fd07d5f16fff1411c16b888fa:org.eclipse.jetty:jetty-server:9.4.12.v20180830",
#         "4c1149328eda9fa39a274262042420f66d9ffd5f:org.eclipse.jetty:jetty-servlet:9.4.12.v20180830",
#         "cb4ccec9bd1fe4b10a04a0fb25d7053c1050188a:org.eclipse.jetty:jetty-util:9.4.12.v20180830",
#         "a3e119df2da04fcf5aa290c8c35c5b310ce2dcd1:org.eclipse.jetty:jetty-webapp:9.4.12.v20180830",
#         "e9f1874e9b5edd498f2fe7cd0904405da07cc300:org.eclipse.jetty:jetty-xml:9.4.12.v20180830",
#         "97d6376f70ae6c01112325c5254e566af118bc75:org.eclipse.jetty.websocket:websocket-api:9.4.12.v20180830",
#         "75880b6a90a6eda83fdbfc20a42f23eade4b975d:org.eclipse.jetty.websocket:websocket-client:9.4.12.v20180830",
#         "33997cdafbabb3ffd6947a5c33057f967e10535b:org.eclipse.jetty.websocket:websocket-common:9.4.12.v20180830",
#         "fadf609aec6026cb25f25b6bc0b979821f849fd7:org.eclipse.jetty.websocket:websocket-server:9.4.12.v20180830",
#         "8d212616b6ea21b96152ff202c2f53fdca8b8b53:org.eclipse.jetty.websocket:websocket-servlet:9.4.12.v20180830",
#         "919f0dfe192fb4e063e7dacadee7f8bb9a2672a9:org.jetbrains:annotations:13.0",
#         "d9717625bb3c731561251f8dd2c67a1011d6764c:org.jetbrains.kotlin:kotlin-stdlib:1.2.71",
#         "ba18ca1aa0e40eb6f1865b324af2f4cbb691c1ec:org.jetbrains.kotlin:kotlin-stdlib-common:1.2.71",
#         "4ce93f539e2133f172f1167291a911f83400a5d0:org.jetbrains.kotlin:kotlin-stdlib-jdk7:1.2.71",
#         "5470d1f752cd342edb77e1062bac07e838d2cea4:org.jetbrains.kotlin:kotlin-stdlib-jdk8:1.2.71",
#         "da76ca59f6a57ee3102f8f9bd9cee742973efa8a:org.slf4j:slf4j-api:1.7.25",
#         "8dacf9514f0c707cbbcdd6fd699e8940d42fb54e:org.slf4j:slf4j-simple:1.7.25",
#     ],
#     deps = [
#         "io.javalin:javalin:2.3.0",
#         "org.slf4j:slf4j-simple:1.7.25",
#     ],
# )
# load("@appdeps//:rules.bzl", "appdeps_compile")
# appdeps_compile()

RULES_JVM_EXTERNAL_TAG = "2.5"

RULES_JVM_EXTERNAL_SHA = "249e8129914be6d987ca57754516be35a14ea866c616041ff0cd32ea94d2f3a1"

http_archive(
    name = "rules_jvm_external",
    sha256 = RULES_JVM_EXTERNAL_SHA,
    strip_prefix = "rules_jvm_external-%s" % RULES_JVM_EXTERNAL_TAG,
    url = "https://github.com/bazelbuild/rules_jvm_external/archive/%s.zip" % RULES_JVM_EXTERNAL_TAG,
)

load("@rules_jvm_external//:defs.bzl", "maven_install")

maven_install(
    artifacts = [
        "io.javalin:javalin:2.3.0",
        "com.fasterxml.jackson.core:jackson-databind:2.9.6",
        "com.fasterxml.jackson.module:jackson-module-kotlin:2.9.4.1",
        "org.slf4j:slf4j-api:1.7.25",
        "org.slf4j:slf4j-simple:1.7.25",
    ],
    repositories = [
        "https://jcenter.bintray.com/",
        "https://maven.google.com",
        "https://repo1.maven.org/maven2",
    ],
)

load(
    "@io_bazel_rules_docker//repositories:repositories.bzl",
    container_repositories = "repositories",
)
load(
    "@io_bazel_rules_docker//container:container.bzl",
    "container_layer",
)
load(
    "@io_bazel_rules_docker//java:image.bzl",
    _java_image_repos = "repositories",
)

_java_image_repos()

load(
    "@io_bazel_rules_docker//container:container.bzl",
    "container_pull",
)
load(
    "@io_bazel_rules_docker//java:java.bzl",
    _JAVA_DIGESTS = "DIGESTS",
)

container_pull(
    name = "java_image_base",
    digest = _JAVA_DIGESTS["latest"],
    registry = "gcr.io",
    repository = "distroless/java",
)

container_repositories()

http_archive(
    name = "bazel_toolchains",
    sha256 = "07dfbe80638eb1fe681f7c07e61b34b579c6710c691e49ee90ccdc6e9e75ebbb",
    strip_prefix = "bazel-toolchains-9a111bd82161c1fbe8ed17a593ca1023fd941c70",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-toolchains/archive/9a111bd82161c1fbe8ed17a593ca1023fd941c70.tar.gz",
        "https://github.com/bazelbuild/bazel-toolchains/archive/9a111bd82161c1fbe8ed17a593ca1023fd941c70.tar.gz",
    ],
)

load("//:repositories.bzl", "go_repositories")

go_repositories()
