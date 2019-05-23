load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    url = "https://github.com/bazelbuild/rules_go/releases/download/0.17.6/rules_go-0.17.6.tar.gz",
    sha256 = "d322611c090f019ad6788e0f9a96762c1e5cf98a93b2eadc14ca702a9ce27420",
)

http_archive(
    name = "bazel_gazelle",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.17.0/bazel-gazelle-0.17.0.tar.gz"],
    sha256 = "3c681998538231a2d24d0c07ed5a7658cb72bfb5fd4bf9911157c0e9ac6a2687",
)

http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "aed1c249d4ec8f703edddf35cbe9dfaca0b5f5ea6e4cd9e83e99f3b0d1136c3d",
    strip_prefix = "rules_docker-0.7.0",
    urls = ["https://github.com/bazelbuild/rules_docker/archive/v0.7.0.tar.gz"],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

load(
    "@io_bazel_rules_docker//go:image.bzl",
    _go_image_repos = "repositories",
)

_go_image_repos()

go_repository(
    name = "com_github_k0kubun_pp",
    commit = "027a6d1765d673d337e687394dbe780dd64e2a1e",
    importpath = "github.com/k0kubun/pp",
)

go_repository(
    name = "com_github_mattn_go_colorable",
    commit = "167de6bfdfba052fa6b2d3664c8f5272e23c9072",
    importpath = "github.com/mattn/go-colorable",
)

go_repository(
    name = "com_github_mattn_go_isatty",
    commit = "0360b2af4f38e8d38c7fce2a9f4e702702d73a39",
    importpath = "github.com/mattn/go-isatty",
)

go_repository(
    name = "org_golang_x_sys",
    commit = "ac767d655b305d4e9612f5f6e33120b9176c4ad4",
    importpath = "golang.org/x/sys",
)
