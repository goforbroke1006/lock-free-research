package cmd

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"

	"github.com/goforbroke1006/lock-free-research/internal"
	tcmb "github.com/goforbroke1006/lock-free-research/internal/api/tcp_counter_mutex_based"
	"github.com/goforbroke1006/lock-free-research/internal/component/shutdowner"
	"github.com/goforbroke1006/lock-free-research/internal/component/tcp_server"
)

func init() {
	var (
		addr = "0.0.0.0:10000"
	)

	var tcpCounterMutexBasedCmd = &cobra.Command{
		Use:   "tcp-counter-mutex-based",
		Short: "TODO: write me",
		Long:  `TODO: write me`,
		Run: func(cmd *cobra.Command, args []string) {
			internal.RegisterNumberCounterMetrics("research", "cmb")

			go func() {
				http.Handle("/metrics", promhttp.Handler())
				if err := http.ListenAndServe(":8080", nil); err != nil {
					panic(err)
				}
			}()

			counter := tcmb.NewMutexBasedNumberCounter()

			server := tcp_server.New(addr)
			server.OnMessage(func(message string, reply *string) {
				startSpan := time.Now()
				message = strings.TrimSpace(message)
				num, _ := strconv.ParseInt(message, 10, 64)
				if num == 0 {
					sum, _ := counter.Get()
					*reply = fmt.Sprintf("%d", sum)
				} else if err := counter.Add(num); err != nil {
					internal.OperationCountFailed.Inc()
				} else {
					internal.OperationCountProcessed.Inc()
					internal.OperationDurationProcessed.Observe(float64(time.Since(startSpan).Nanoseconds()))
				}
			})
			go server.Run()

			shutdowner.WaitForInterrupt()

			_ = server.Stop()
		},
	}

	tcpCounterMutexBasedCmd.PersistentFlags().StringVar(&addr, "addr", addr, "TCP listener address")
	apiCmd.AddCommand(tcpCounterMutexBasedCmd)
}
