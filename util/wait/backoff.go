package wait

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"k8s.io/apimachinery/pkg/util/wait"
)

// the underlying ExponentialBackoff does not retain the underlying error
// so this addresses this
func Backoff(b wait.Backoff, f func() (bool, error)) error {
	var err error

	retryCount := 0
	waitErr := wait.ExponentialBackoff(b, func() (bool, error) {
		if (retryCount > 1) {
			log.Infof("Retrying ... (%d)", retryCount)
		}
		retryCount++

		var done bool
		done, err = f()
		return done, nil
	})
	if waitErr != nil {
		if err != nil {
			return fmt.Errorf("%v: %v", waitErr, err)
		} else {
			return waitErr
		}
	}
	return err
}
