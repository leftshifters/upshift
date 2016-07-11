package actions

import "upshift/basher"

// Android : Construct for all activities related to the android binary
type Android struct {
}

// UpgradeAndroid : Upgrade the android sdk
func (a *Android) UpgradeAndroid(logPath string) error {
	var b basher.Basher
	_, err := b.Run("AndroidUpgradeSDK", []string{logPath})
	if err != nil {
		return err
	}
	return nil
}

// SetupAndroid : Install the android sdk
func (a *Android) SetupAndroid(logPath string) error {
	var b basher.Basher
	_, err := b.Run("AndroidInstallSDK", []string{logPath})
	if err != nil {
		return err
	}
	return nil
}
