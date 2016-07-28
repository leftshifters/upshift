package actions

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/leftshifters/upshift/config"
	"github.com/stretchr/testify/assert"
)

func Test_Actions_AndroidBuild(t *testing.T) {
	conf := config.Get()
	currentWD, _ := os.Getwd()
	os.Chdir(filepath.Join("..", "android-test"))
	conf.Settings.CleanBeforeBuild = true

	status := AndroidBuild()
	assert.Equal(t, 0, status)

	os.Chdir(currentWD)
}

func Test_Actions_AndroidDeploy(t *testing.T) {

}

func Test_Actions_AndroidLoadEmulator(t *testing.T) {
	status := AndroidLoadEmulator()
	assert.Equal(t, 0, status)
}

func Test_Actions_AndroidRun(t *testing.T) {

}

func Test_Actions_AndroidTest(t *testing.T) {

}

func Test_Actions_AndroidUpgrade(t *testing.T) {

}

func Test_Actions_GitPull(t *testing.T) {
	status := GitPull()
	assert.Equal(t, 1, status)
}

func Test_Actions_GitSubmodules(t *testing.T) {

}

func Test_Actions_IosArchive(t *testing.T) {

}

func Test_Actions_IosBuild(t *testing.T) {

}

func Test_Actions_IosCertificates(t *testing.T) {

}

func Test_Actions_IosCreateApp(t *testing.T) {

}

func Test_Actions_IosDeploy(t *testing.T) {

}

func Test_Actions_IosExportIPA(t *testing.T) {

}

func Test_Actions_IosPrepare(t *testing.T) {

}

func Test_Actions_IosProvisioning(t *testing.T) {

}

func Test_Actions_IosRun(t *testing.T) {

}

func Test_Actions_IosSimulator(t *testing.T) {

}

func Test_Actions_IosTest(t *testing.T) {

}

func Test_Actions_PodsInstall(t *testing.T) {
	currentWD, _ := os.Getwd()
	os.Chdir(filepath.Join("..", "ios-test-swift"))

	status := SetupPods()
	assert.Equal(t, 0, status)

	os.Chdir(currentWD)
}

func Test_Actions_SetupConfig(t *testing.T) {

}

func Test_Actions_SetupFastlane(t *testing.T) {
	status := SetupFastlane()
	assert.Equal(t, 0, status)
}

func Test_Actions_SetupGradleWrapper(t *testing.T) {
	currentWD, _ := os.Getwd()
	os.Chdir(filepath.Join("..", "android-test"))

	status := SetupGradleWrapper()
	assert.Equal(t, 0, status)

	os.Chdir(currentWD)
}

func Test_Actions_SetupPods(t *testing.T) {
	os.Chdir(filepath.Join("..", "ios-test-swift"))
	status := SetupPods()
	assert.Equal(t, 0, status)
}

func Test_Actions_SetupXcode(t *testing.T) {
	os.Chdir(filepath.Join("..", "ios-test-swift"))
	status := SetupXcode()
	assert.Equal(t, 0, status)
}

func Test_Actions_SetupXcpretty(t *testing.T) {

}

func Test_Actions_SetupXctool(t *testing.T) {

}

func Test_Actions_ShowHelp(t *testing.T) {

}

func Test_Actions_ShowVersion(t *testing.T) {
	status := ShowVersion()
	assert.Equal(t, 0, status)
}

func Test_Actions_UpgradeScript(t *testing.T) {
	status := UpgradeScript()
	assert.Equal(t, 0, status)
}
