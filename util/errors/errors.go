package errors

import (
	"net"
	"net/url"
	"strings"

	log "github.com/sirupsen/logrus"

	apierr "k8s.io/apimachinery/pkg/api/errors"

	argoerrs "github.com/argoproj/argo/v2/errors"
)

func IsTransientErr(err error) bool {
	if err == nil {
		return false
	}
	err = argoerrs.Cause(err)

	log.Infof("IsTransientErr? %T", err)

	if (isExceededQuotaErr(err)) {
		log.Infof("Quota Exceeded")
		return true
	}

	if (apierr.IsTooManyRequests(err)) {
		log.Infof("Too Many Requests")
		return true
	}

	if (isResourceQuotaConflictErr(err)) {
		log.Infof("Resource Quota Conflict")
		return true
	}

	if (isTransientNetworkErr(err)) {
		log.Infof("Transient Network Error")
		return true
	}

	log.Errorf("Non Transient Error: %v", err)

	return false
}

func isExceededQuotaErr(err error) bool {
	return apierr.IsForbidden(err) && strings.Contains(err.Error(), "exceeded quota")
}

func isResourceQuotaConflictErr(err error) bool {
	return apierr.IsConflict(err) && strings.Contains(err.Error(), "Operation cannot be fulfilled on resourcequota")
}

func isTransientNetworkErr(err error) bool {
	switch err.(type) {
	case net.Error:
		switch err.(type) {
		case *net.DNSError, *net.OpError, net.UnknownNetworkError:
			log.Infof("Network error")
			return true
		case *url.Error:
			if (strings.Contains(err.Error(), "Connection closed by foreign host")) {
				// For a URL error, where it replies back "connection closed"
				// retry again.
				log.Infof("Connection closed by foreign host")
				return true
			} else if strings.Contains(err.Error(), "network is unreachable") {
				log.Infof("network is unreachable")
				return true
			} else if strings.Contains(err.Error(), "connection reset by peer") {
				log.Infof("connection reset by peer")
				return true
			}
		default:
			if strings.Contains(err.Error(), "net/http: TLS handshake timeout") {
				// If error is - tlsHandshakeTimeoutError, retry.
				log.Infof("TLS handshake timeout")
				return true
			} else if strings.Contains(err.Error(), "i/o timeout") {
				// If error is - tcp timeoutError, retry.
				log.Infof("i/o timeout")
				return true
			} else if strings.Contains(err.Error(), "connection timed out") {
				// If err is a net.Dial timeout, retry.
				log.Infof("connection timed out")
				return true
			}

			log.Errorf("Non Transient Network Error: %v", err)
		}
	}

	return false
}
