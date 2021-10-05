load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "libwat",
    srcs = glob(["*.go"]),
    importpath = "github.com/gkgoat1/wat/src",
    deps = ["@com_github_llir_llvm//:llvm"],
    visibility = ["//visibility:public"],
)