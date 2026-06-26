package infa

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// WaitExitSignal 阻塞当前 goroutine，直到收到 Ctrl+C 退出信号。
func WaitExitSignal() os.Signal {
	ch := make(chan os.Signal, 1)

	// os.Interrupt 表示 Ctrl+C
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(ch)

	return <-ch
}

// WaitUntil 等待到目标时间。
func WaitUntil(target time.Time)  {
	wait := time.Until(target)
	if wait <= 0 {
		return
	}
	// 等待
	log.Printf("⏳ 等待到目标时间: %s, 剩余 %s",
		target.Format("2006-01-02 15:04:05"), wait,
	)
	time.Sleep(wait)
	log.Println("WaitUntil Done")
}

