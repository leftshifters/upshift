import Foundation


// // http://iswift.org/cookbook/execute-a-shell-command
// // Create a Task instance

// let task = NSTask()
// task.launchPath = "/usr/bin/env"

// let pipe = NSPipe()
// task.standardOutput = pipe
// task.standardError = pipe

// task.launch()
// task.waitUntilExit()
// let data = pipe.fileHandleForReading.readDataToEndOfFile()
// let output = NSString(data: data, encoding: NSUTF8StringEncoding)
// print(output!)


let task = NSTask()

// // Set the task parameters
task.launchPath = "/bin/bash"
// task.arguments = ["TERM_PROGRAM=Apple_Terminal", "SWIFTENV_ROOT=/Users/sudhanshuraheja/.swiftenv", "ANDROID_HOME=/Users/sudhanshuraheja/Library/Android/sdk", "SHELL=/bin/bash", "TERM=xterm-256color", "TMPDIR=/var/folders/d4/zc5jxtfj7dz2c8qml1dwlqcw0000gn/T/", "Apple_PubSub_Socket_Render=/private/tmp/com.apple.launchd.9MhEmbBwpX/Render", "TERM_PROGRAM_VERSION=361.1", "OLDPWD=/Users/sudhanshuraheja/code/upshift", "TERM_SESSION_ID=BA572524-4CD4-417E-A33C-B00CF92A063C", "N_PREFIX=/Users/sudhanshuraheja/n", "USER=sudhanshuraheja", "SSH_AUTH_SOCK=/private/tmp/com.apple.launchd.UsawnP6Tv5/Listeners", "__CF_USER_TEXT_ENCODING=0x1F5:0x0:0x0", "PATH=/Users/sudhanshuraheja/.swiftenv/shims:/Users/sudhanshuraheja/.swiftenv/bin:/Users/sudhanshuraheja/.swiftenv/shims:/Applications/Sublime Text.app/Contents/SharedSupport/bin:/Library/Java/JavaVirtualMachines/jdk1.8.0_73.jdk/Contents/Home/bin:/Users/sudhanshuraheja/torch/install/bin:/Users/sudhanshuraheja/.swiftenv/shims:/Users/sudhanshuraheja/.swiftenv/bin:/Users/sudhanshuraheja/.swiftenv/shims:/Applications/Sublime Text.app/Contents/SharedSupport/bin:/Library/Java/JavaVirtualMachines/jdk1.8.0_73.jdk/Contents/Home/bin:/Users/sudhanshuraheja/torch/install/bin:/Users/sudhanshuraheja/.swiftenv/shims:/Users/sudhanshuraheja/.swiftenv/bin:/Users/sudhanshuraheja/.swiftenv/shims:/Applications/Sublime Text.app/Contents/SharedSupport/bin:/Library/Java/JavaVirtualMachines/jdk1.8.0_73.jdk/Contents/Home/bin:/Users/sudhanshuraheja/torch/install/bin:/Users/sudhanshuraheja/.swiftenv/shims:/Users/sudhanshuraheja/.swiftenv/bin:/Users/sudhanshuraheja/.swiftenv/shims:/Applications/Sublime Text.app/Contents/SharedSupport/bin:/Library/Java/JavaVirtualMachines/jdk1.8.0_73.jdk/Contents/Home/bin:/Users/sudhanshuraheja/torch/install/bin:/Users/sudhanshuraheja/.swiftenv/shims:/Users/sudhanshuraheja/.swiftenv/bin:/Users/sudhanshuraheja/.swiftenv/shims:/Applications/Sublime Text.app/Contents/SharedSupport/bin:/Library/Java/JavaVirtualMachines/jdk1.8.0_73.jdk/Contents/Home/bin:/Users/sudhanshuraheja/torch/install/bin:/Users/sudhanshuraheja/.swiftenv/shims:/Users/sudhanshuraheja/.swiftenv/bin:/Users/sudhanshuraheja/.swiftenv/shims:/Applications/Sublime Text.app/Contents/SharedSupport/bin:/Library/Java/JavaVirtualMachines/jdk1.8.0_73.jdk/Contents/Home/bin:/Users/sudhanshuraheja/torch/install/bin:/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin:/opt/X11/bin:/usr/local/go/bin:/Users/sudhanshuraheja/n/bin:/Users/sudhanshuraheja/Library/Android/sdk/tools:/Users/sudhanshuraheja/Library/Android/sdk/platform-tools:/Users/sudhanshuraheja/gocode/bin:/Users/sudhanshuraheja/Library/Android/sdk/tools:/Users/sudhanshuraheja/Library/Android/sdk/platform-tools:/Users/sudhanshuraheja/gocode/bin:/Users/sudhanshuraheja/Library/Android/sdk/tools:/Users/sudhanshuraheja/Library/Android/sdk/platform-tools:/Users/sudhanshuraheja/gocode/bin:/Users/sudhanshuraheja/Library/Android/sdk/tools:/Users/sudhanshuraheja/Library/Android/sdk/platform-tools:/Users/sudhanshuraheja/gocode/bin:/Users/sudhanshuraheja/Library/Android/sdk/tools:/Users/sudhanshuraheja/Library/Android/sdk/platform-tools:/Users/sudhanshuraheja/gocode/bin:/Users/sudhanshuraheja/Library/Android/sdk/tools:/Users/sudhanshuraheja/Library/Android/sdk/platform-tools:/Users/sudhanshuraheja/gocode/bin", "PWD=/Users/sudhanshuraheja/code/Task", "JAVA_HOME=/Library/Java/JavaVirtualMachines/jdk1.8.0_73.jdk/Contents/Home", "XPC_FLAGS=0x0", "XPC_SERVICE_NAME=0", "SHLVL=1", "HOME=/Users/sudhanshuraheja", "LOGNAME=sudhanshuraheja", "LC_CTYPE=UTF-8", "DISPLAY=/private/tmp/com.apple.launchd.kvnIICz71s/org.macosforge.xquartz:0", "gradle", "tasks"]
task.arguments = ["-c", "gradle tasks"]
 
// // Create a Pipe and make the task
// // put all the output there
// let pipe = NSPipe()
// task.standardOutput = pipe
// task.standardError = pipe

// // Launch the task
task.launch()
task.waitUntilExit()
 
// // Get the data
// // repeat {
// 	let data = pipe.fileHandleForReading.readDataToEndOfFile()
// 	let output = NSString(data: data, encoding: NSUTF8StringEncoding)
// 	print(output!)
// // } while(task.isRunning)


print(task.terminationStatus)


// import Darwin

// let fp = popen("ping -c 4 localhost", "r")
// var buf = Array<CChar>(repeating: 0, count: 128)

// while fgets(&buf, CInt(buf.count), fp) != nil,
//       let str = String.fromCString(buf) {
//     print(str)
// }

// fclose(fp)







