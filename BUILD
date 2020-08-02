load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/twocucao/zelda
gazelle(name = "gazelle")

go_library(
    name = "go_default_library",
    srcs = ["zelda.go"],
    importpath = "github.com/twocucao/zelda",
    visibility = ["//visibility:private"],
    deps = [
        "//apps/www:go_default_library",
        "@com_github_urfave_cli//:go_default_library",
    ],
)

go_binary(
    name = "zelda",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
