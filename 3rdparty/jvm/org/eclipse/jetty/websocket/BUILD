java_library(
    name = "websocket_api",
    visibility = [
        "//3rdparty/jvm:__subpackages__",
    ],
    exports = [
        "//external:jar/org/eclipse/jetty/websocket/websocket_api",
    ],
)

java_library(
    name = "websocket_client",
    visibility = [
        "//3rdparty/jvm:__subpackages__",
    ],
    exports = [
        "//3rdparty/jvm/org/eclipse/jetty:jetty_client",
        "//external:jar/org/eclipse/jetty/websocket/websocket_client",
    ],
)

java_library(
    name = "websocket_common",
    visibility = [
        "//3rdparty/jvm:__subpackages__",
    ],
    exports = [
        ":websocket_api",
        "//external:jar/org/eclipse/jetty/websocket/websocket_common",
    ],
)

java_library(
    name = "websocket_server",
    visibility = [
        "//3rdparty/jvm:__subpackages__",
    ],
    exports = [
        ":websocket_client",
        ":websocket_common",
        ":websocket_servlet",
        "//external:jar/org/eclipse/jetty/websocket/websocket_server",
    ],
)

java_library(
    name = "websocket_servlet",
    visibility = [
        "//3rdparty/jvm:__subpackages__",
    ],
    exports = [
        "//external:jar/org/eclipse/jetty/websocket/websocket_servlet",
    ],
)
