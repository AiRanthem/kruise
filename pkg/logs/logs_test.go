package logs

import (
	"errors"
	"k8s.io/klog/v2"
	"testing"
)

func TestSetLoggingFormat(t *testing.T) {
	klog.InfoS("text", "key", 1)
	SetLoggingFormat(LoggingFormatJson)
	klog.InfoS("json", "key", 1)
	klog.V(4).InfoS("level 4", "key", 4)
	klog.ErrorS(errors.New("error"), "error", "key", 3.14)
	klog.Info("json too")
}
