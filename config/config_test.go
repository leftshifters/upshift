package config

import (
	"testing"
)

func TestConfig(t *testing.T) {
	conf, _ := Load("sample.toml")

	if conf.Application.Debug != false {
		t.Fail()
		t.Log("conf.Application.Debug Failed")
	}
	if conf.Runner.RootPassword != "testPassword" {
		t.Fail()
		t.Log("conf.Runner.RootPassword Failed")
	}
	if conf.Build.GitRepoURL != "testRepo" {
		t.Fail()
		t.Log("conf.Build.GitRepoURL Failed")
	}
	if conf.Build.GitRepoRemote != "origin" {
		t.Fail()
		t.Log("conf.Build.GitRepoRemote Failed")
	}
	if conf.Build.CleanBeforeBuild != false {
		t.Fail()
		t.Log("conf.Build.CleanBeforeBuild Failed")
	}
	if conf.Build.UninstallOlderBuilds != false {
		t.Fail()
		t.Log("conf.Build.UninstallOlderBuilds Failed")
	}
	if conf.IOS.ProjectName != "testProject" {
		t.Fail()
		t.Log("conf.IOS.ProjectName Failed")
	}
	if conf.IOS.UseWorkspace != false {
		t.Fail()
		t.Log("conf.IOS.UseWorkspace Failed")
	}
	if conf.IOS.Scheme != "testScheme" {
		t.Fail()
		t.Log("conf.IOS.Scheme Failed")
	}
	if conf.IOS.TestDevice != "iPhone 6" {
		t.Fail()
		t.Log("conf.IOS.TestDevice Failed")
	}
	if conf.IOS.Xcode != "7.3.1" {
		t.Fail()
		t.Log("conf.IOS.Xcode Failed")
	}
	if conf.Android.PackageName != "testPackage" {
		t.Fail()
		t.Log("conf.Android.PackageName Failed")
	}
	if conf.Android.MainActivityName != "testActivity" {
		t.Fail()
		t.Log("conf.Android.MainActivityName Failed")
	}
}
