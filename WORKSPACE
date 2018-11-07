workspace(name = "research")

# Imports basic Go rules for Bazel (e.g. go_binary)
git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    commit = "e4d0254fb249a09fb01f052b23d3baddae1b70ec",
)

# Imports the Gazelle tool for Go/Bazel
git_repository(
    name = "bazel_gazelle",
    remote = "https://github.com/bazelbuild/bazel-gazelle",
    commit = "644ec7202aa352b78d65bc66abc2c0616d76cc84",
)

# Imports Maven for Java/Kotlin
git_repository(
    name = "org_pubref_rules_maven",
    commit = "9c3b07a6d9b195a1192aea3cd78afd1f66c80710",
    remote = "https://github.com/pubref/rules_maven",
)

# Imports Kotlin
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
rules_kotlin_version = "87bd13f91d166a8070e9bbfbb0861f6f76435e7a"
http_archive(
    name = "io_bazel_rules_kotlin",
    urls = ["https://github.com/bazelbuild/rules_kotlin/archive/%s.zip" % rules_kotlin_version],
    type = "zip",
    strip_prefix = "rules_kotlin-%s" % rules_kotlin_version
)

# Imports Docker rules for Bazel (e.g. docker_image)
git_repository(
    name = "io_bazel_rules_docker",
    remote = "https://github.com/bazelbuild/rules_docker.git",
    tag = "v0.5.1",
)


# Loads Go rules for Bazel
load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains(
    go_version = "1.10.1",
)

# Loads Docker rules for Bazel
load(
    "@io_bazel_rules_docker//go:image.bzl",
    _go_image_repos = "repositories",
)

_go_image_repos()

# Loads Gazelle tool
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")
gazelle_dependencies()



load("@io_bazel_rules_kotlin//kotlin:kotlin.bzl", "kotlin_repositories", "kt_register_toolchains")
load("@io_bazel_rules_kotlin//kotlin:kotlin.bzl", "kt_jvm_import")
kotlin_repositories()
kt_register_toolchains()


load("@org_pubref_rules_maven//maven:rules.bzl", "maven_repositories", "maven_repository")
maven_repositories()

load("//3rdparty:workspace.bzl", "maven_dependencies")
maven_dependencies()

maven_repository(
    name = 'appdeps',
    deps = [
        'io.javalin:javalin:2.3.0',
        'org.slf4j:slf4j-simple:1.7.25',
    ],
    transitive_deps = [
        '73836e9cf29f978e47817584f9cee86b5e1f4c09:io.javalin:javalin:2.3.0',
        '3cd63d075497751784b2fa84be59432f4905bf7c:javax.servlet:javax.servlet-api:3.1.0',
        '1d329d68f31dce13135243c06013aaf6f708f7e7:org.eclipse.jetty:jetty-client:9.4.12.v20180830',
        '1341796dde4e16df69bca83f3e87688ba2e7d703:org.eclipse.jetty:jetty-http:9.4.12.v20180830',
        'e93f5adaa35a9a6a85ba130f589c5305c6ecc9e3:org.eclipse.jetty:jetty-io:9.4.12.v20180830',
        '299e0602a9c0b753ba232cc1c1dda72ddd9addcf:org.eclipse.jetty:jetty-security:9.4.12.v20180830',
        'b0f25df0d32a445fd07d5f16fff1411c16b888fa:org.eclipse.jetty:jetty-server:9.4.12.v20180830',
        '4c1149328eda9fa39a274262042420f66d9ffd5f:org.eclipse.jetty:jetty-servlet:9.4.12.v20180830',
        'cb4ccec9bd1fe4b10a04a0fb25d7053c1050188a:org.eclipse.jetty:jetty-util:9.4.12.v20180830',
        'a3e119df2da04fcf5aa290c8c35c5b310ce2dcd1:org.eclipse.jetty:jetty-webapp:9.4.12.v20180830',
        'e9f1874e9b5edd498f2fe7cd0904405da07cc300:org.eclipse.jetty:jetty-xml:9.4.12.v20180830',
        '97d6376f70ae6c01112325c5254e566af118bc75:org.eclipse.jetty.websocket:websocket-api:9.4.12.v20180830',
        '75880b6a90a6eda83fdbfc20a42f23eade4b975d:org.eclipse.jetty.websocket:websocket-client:9.4.12.v20180830',
        '33997cdafbabb3ffd6947a5c33057f967e10535b:org.eclipse.jetty.websocket:websocket-common:9.4.12.v20180830',
        'fadf609aec6026cb25f25b6bc0b979821f849fd7:org.eclipse.jetty.websocket:websocket-server:9.4.12.v20180830',
        '8d212616b6ea21b96152ff202c2f53fdca8b8b53:org.eclipse.jetty.websocket:websocket-servlet:9.4.12.v20180830',
        '919f0dfe192fb4e063e7dacadee7f8bb9a2672a9:org.jetbrains:annotations:13.0',
        'd9717625bb3c731561251f8dd2c67a1011d6764c:org.jetbrains.kotlin:kotlin-stdlib:1.2.71',
        'ba18ca1aa0e40eb6f1865b324af2f4cbb691c1ec:org.jetbrains.kotlin:kotlin-stdlib-common:1.2.71',
        '4ce93f539e2133f172f1167291a911f83400a5d0:org.jetbrains.kotlin:kotlin-stdlib-jdk7:1.2.71',
        '5470d1f752cd342edb77e1062bac07e838d2cea4:org.jetbrains.kotlin:kotlin-stdlib-jdk8:1.2.71',
        'da76ca59f6a57ee3102f8f9bd9cee742973efa8a:org.slf4j:slf4j-api:1.7.25',
        '8dacf9514f0c707cbbcdd6fd699e8940d42fb54e:org.slf4j:slf4j-simple:1.7.25',
    ],
)

load("@appdeps//:rules.bzl", "appdeps_compile")
appdeps_compile()

load(
  "@io_bazel_rules_docker//container:container.bzl",
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
    _container = "container",
    _repositories = "repositories",
)

load("@io_bazel_rules_docker//java:java.bzl", _JAVA_DIGESTS = "DIGESTS")
container_pull(
    name = "java_image_base",
    registry = "gcr.io",
    repository = "distroless/java",
    digest = _JAVA_DIGESTS["latest"],
)

container_repositories()

# This requires rules_docker to be fully instantiated before
# it is pulled in.
git_repository(
    name = "io_bazel_rules_k8s",
    commit = "73f60d4cec5eef6ac7563b7cdf76bd1b5b360864",
    remote = "https://github.com/seibert-media/rules_k8s.git",
)

# local testing
# local_repository(
#     name = "io_bazel_rules_k8s",
#     path = "/home/kwiesmueller/git/bazel/rules_k8s",
# )

http_archive(
  name = "bazel_toolchains",
  urls = [
    "https://mirror.bazel.build/github.com/bazelbuild/bazel-toolchains/archive/9a111bd82161c1fbe8ed17a593ca1023fd941c70.tar.gz",
    "https://github.com/bazelbuild/bazel-toolchains/archive/9a111bd82161c1fbe8ed17a593ca1023fd941c70.tar.gz",
  ],
  strip_prefix = "bazel-toolchains-9a111bd82161c1fbe8ed17a593ca1023fd941c70",
  sha256 = "07dfbe80638eb1fe681f7c07e61b34b579c6710c691e49ee90ccdc6e9e75ebbb",
)

load("@io_bazel_rules_k8s//k8s:k8s.bzl", "k8s_repositories")

k8s_repositories()

load("@io_bazel_rules_k8s//k8s:k8s.bzl", "k8s_defaults")

k8s_defaults(
  name = "k8s_deploy",
  kind = "deployment",
  namespace = "{NAMESPACE}",
  cluster = "gke_kwiesmueller-development_us-east1-d_cluster-1",
  repo = "eu.gcr.io/kwiesmueller-development/",
)