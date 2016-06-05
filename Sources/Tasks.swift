struct TaskHandler {
	var job: String, action: String?

	init() {
		if Process.arguments.count == 1 {
			job = "setup"
			action = "showVersion"
		} else if Process.arguments.count == 2 {
			job = Process.arguments[1]
		} else if Process.arguments.count == 3 {
			job = Process.arguments[1]
			action = Process.arguments[2]
		} else {
			job = "version"
		}

		// print("The job is \(job!) and action is \(action!)")
		let tasks = findTaskQueue(forJob: job, andAction: action);
		print(tasks)
	}

	func findTaskQueue(forJob: String, andAction: String?) -> [String] {
		var tasks: [String]
		switch forJob {
		case "ios":
			tasks = findiOSTasks(forAction: action)
		case "android":
			tasks = findAndroidTasks(forAction: action)
		case "setup":
			tasks = findSetupTasks(forAction: action)
		case "install":
			tasks = findInstallTasks(forAction: action)
		case "version":
			tasks = findVersionTasks(forAction: action)
		case "action":
			tasks = findActionTasks(forAction: action)
		default:
			tasks = ["help"]
		}
		return tasks
	}

	func findiOSTasks(forAction: String?) -> [String] {
		var tasks: [String] = ["help"]
		if let action = forAction {
			switch action {
			case "build":
				tasks = ["upgradeScript", "setupXcode", "setupXcpretty", "gitPull", "gitSubmodules", "setupPods", "iosBuild"]
			case "run":
				tasks = ["upgradeScript", "setupXcode", "setupXcpretty", "gitPull", "gitSubmodules", "setupPods", "iosSimulator", "iosRun"]
			case "deploy":
				tasks = ["help"]
			default:
				tasks = ["help"]
			}
		}
		return tasks
	}

	func findAndroidTasks(forAction: String?) -> [String] {
		var tasks: [String] = ["help"]
		if let action = forAction {
			switch action {
			case "build":
				tasks = ["upgradeScript", "gitPull", "gitSubmodules", "setupGradle", "androidBuild"]
			case "run":
				tasks = ["upgradeScript", "gitPull", "gitSubmodules", "setupGradle", "androidEmulator", "androidRun"]
			case "deploy":
				tasks = ["help"]
			default:
				tasks = ["help"]
			}
		}
		return tasks
	}

	func findSetupTasks(forAction: String?) -> [String] {
		var tasks: [String] = ["help"]
		if let action = forAction {
			switch action {
			case "clone":
				tasks = ["gitClone"]
			case "config":
				tasks = ["setupConfig"]
			default:
				// Make an alias "upshift setup" points to "upshift install"
				tasks = ["setupScript"]
			}
		}
		return tasks
	}

	func findInstallTasks(forAction: String?) -> [String] {
		var tasks: [String] = ["help"]
		if let action = forAction {
			switch action {
			case "install":
				tasks = ["setupScript"]
			default:
				tasks = ["help"]
			}
		}
		return tasks
	}

	func findVersionTasks(forAction: String?) -> [String] {
		var tasks: [String] = ["help"]
		if let action = forAction {
			switch action {
			case "version", "-v", "-version":
				tasks = ["showVersion"]
			default:
				tasks = ["help"]
			}
		}
		return tasks
	}

	func findActionTasks(forAction: String?) -> [String] {
		var tasks: [String] = ["help"]
		if let action = forAction {
			switch action {
			case "setupSsh", "setupScript", "setupGradle", "setupPods", "setupXcode", "setupXcpretty", "upgradeScript", "gitPull", "gitClone", "gitSubmodules", "iosSimulator", "androidEmulator", "iosBuild", "androidBuild", "iosRun", "androidRun", "iosDeploy", "androidDeploy":
				tasks = [action]
			default:
				tasks = ["help"]
			}
		}
		return tasks
	}

}