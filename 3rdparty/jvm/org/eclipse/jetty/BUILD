java_library(
    name = "jetty_client",
    visibility = [
        "//3rdparty/jvm:__subpackages__",
    ],
    exports = [
        "//external:jar/org/eclipse/jetty/jetty_client",
    ],
)

java_library(
    name = "jetty_http",
    visibility = [
        "//3rdparty/jvm:__subpackages__",
    ],
    exports = [
        ":jetty_util",
        "//external:jar/org/eclipse/jetty/jetty_http",
    ],
)

java_library(
    name = "jetty_io",
    visibility = [
        "//3rdparty/jvm:__subpackages__",
    ],
    exports = [
        "//external:jar/org/eclipse/jetty/jetty_io",
    ],
)

java_library(
    name = "jetty_security",
    visibility = [
        "//3rdparty/jvm:__subpackages__",
    ],
    exports = [
        "//external:jar/org/eclipse/jetty/jetty_security",
    ],
)

java_library(
    name = "jetty_server",
    visibility = [
        "//3rdparty/jvm:__subpackages__",
    ],
    exports = [
        ":jetty_http",
        ":jetty_io",
        "//3rdparty/jvm/javax/servlet:javax_servlet_api",
        "//external:jar/org/eclipse/jetty/jetty_server",
    ],
)

java_library(
    name = "jetty_servlet",
    visibility = [
        "//3rdparty/jvm:__subpackages__",
    ],
    exports = [
        ":jetty_security",
        "//external:jar/org/eclipse/jetty/jetty_servlet",
    ],
)

java_library(
    name = "jetty_util",
    visibility = [
        "//3rdparty/jvm:__subpackages__",
    ],
    exports = [
        "//external:jar/org/eclipse/jetty/jetty_util",
    ],
)

java_library(
    name = "jetty_webapp",
    visibility = [
        "//3rdparty/jvm:__subpackages__",
    ],
    exports = [
        ":jetty_servlet",
        ":jetty_xml",
        "//external:jar/org/eclipse/jetty/jetty_webapp",
    ],
)

java_library(
    name = "jetty_xml",
    visibility = [
        "//3rdparty/jvm:__subpackages__",
    ],
    exports = [
        "//external:jar/org/eclipse/jetty/jetty_xml",
    ],
)
