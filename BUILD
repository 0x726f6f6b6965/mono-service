load("@gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/0x726f6f6b6965/mono-service
# gazelle:proto disable_global
gazelle(name = "gazelle")

# gazelle(
#     name = "gazelle-update-repos",
#     args = [
#         "-from_file=go.mod",
#         "-to_macro=deps.bzl%go_dependencies",
#         "-prune",
#         "-build_file_proto_mode=disable_global",
#     ],
#     command = "update-repos",
# )

filegroup(
    name = "env-data",
    srcs = [
        ".env",
        "policy/auth.rego",
        "rule-data/allow_uri.json",
    ],
    visibility = ["//visibility:public"],
)
