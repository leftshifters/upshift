#if os(OSX)
    import Darwin
#else
    import Glibc
#endif

// Define the version number of the app
let version = "0.7.3"
print(version)

// Setup the configuration by reading config.toml
var conf = Config()
conf.loadTOML()

var tasks = TaskHandler()