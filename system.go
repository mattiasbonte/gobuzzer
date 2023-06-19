package gobuzzer

import (
	"fmt"

	"github.com/gen2brain/beeep"
)

var (
	beep   = "beep"
	notify = "notify"
	alert  = "alert"
)

// SystemBuzz notifies the user through triggering a beep/notification/alert on the operating system.
// "beep" plays a subtle beep sound only.
// "notify" displays a notification banner without any sound.
// "alert" displays a notification banner with a notification sound.
func SystemBuzz(message string, notificationType string) error {
	switch notificationType {
	case beep:
		if err := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration); err != nil {
			return err
		}
	case notify:
		if err := beeep.Notify("📢", message, "assets/information.png"); err != nil {
			return err
		}
	case alert:
		if err := beeep.Alert("📢", message, "assets/warning.png"); err != nil {
			return err
		}
	default:
		return fmt.Errorf("invalid notification type: %s", notificationType)
	}

	return nil
}
