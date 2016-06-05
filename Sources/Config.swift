struct ApplicationConfig {
	var Debug: Bool

	init() {
		Debug = false
	}
}

struct RunnerConfig {
	var RootPassword: String

	init() {
		RootPassword = ""
	}
}

struct BuildConfig {
	var GitRepoURL: String
	var GitRepoBranch: String
	var CleanBeforeBuild: Bool
	var UninstallOlderBuilds: Bool

	init() {
		GitRepoURL = ""
		GitRepoBranch = ""
		CleanBeforeBuild = false
		UninstallOlderBuilds = false
	}
}

struct iOSConfig {
	var ProjectName: String
	var UseWorkspace: Bool
	var Scheme: String
	var TestDevice : String
	var Xcode: String

	init() {
		ProjectName = ""
		UseWorkspace = false
		Scheme = ""
		TestDevice = ""
		Xcode = ""
	}
}

struct AndroidConfig {
	var PackageName: String
	var MainActivityName: String

	init() {
		PackageName = ""
		MainActivityName = ""
	}
}

struct Config {
	var Application: ApplicationConfig
	var Runner: RunnerConfig
	var Build: BuildConfig
	var iOS: iOSConfig
	var Android: AndroidConfig

	init() {
		Application = ApplicationConfig()
		Runner = RunnerConfig()
		Build = BuildConfig()
		iOS = iOSConfig()
		Android = AndroidConfig()
	}

	mutating func loadTOML() {
		// TODO : We don't know how to parse TOML yet, so we will just hardcode values and return
		// TODO : Find a fucking way to handle TOML files
		// 	This guy talks about how to do it I think - http://ankit.im/swift/2016/03/26/improving-system-modules-support-in-swiftpm/

		self.iOS.Xcode = "7.3.1"
		return 
	}
}