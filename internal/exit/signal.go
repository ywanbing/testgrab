package exit

import (
	"os"
	"os/signal"
	"syscall"
)

// WaitSignal 等待退出信号
func WaitSignal() {
	// 退出信号
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
}
