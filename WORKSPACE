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

load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

load(
  "@io_bazel_rules_docker//container:container.bzl",
  container_repositories = "repositories",
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