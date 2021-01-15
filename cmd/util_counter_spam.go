package cmd

import (
	"bufio"
	"context"
	"fmt"
	"math/rand"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/spf13/cobra"

	"github.com/goforbroke1006/lock-free-research/internal/component/shutdowner"
)

func init() {
	var (
		addr       = "0.0.0.0:10000"
		concurrent = 25
	)

	var runStandardCmd = &cobra.Command{
		Use:   "counter-spam",
		Short: "TODO: write me",
		Long:  `TODO: write me`,
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := context.WithCancel(context.Background())

			ccWait := sync.WaitGroup{}
			for i := 0; i < concurrent; i++ {
				ccWait.Add(1)
				go func(ctx context.Context, cid int) {
					defer ccWait.Done()

					conn, err := net.Dial("tcp", addr)
					if err != nil {
						panic(err)
					}

				INFINITE:
					for {
						select {
						case <-ctx.Done():
							break INFINITE
						default:
							time.Sleep(2500 * time.Microsecond)
							num := 10 + rand.Intn(89)
							_, _ = fmt.Fprintf(conn, "%d\n", num)
						}
					}

					fmt.Printf("Clint # %d finished\n", cid)
				}(ctx, i)
			}

			fmt.Println("Press Ctrl+C to terminate...")
			shutdowner.WaitForInterrupt()

			cancel()

			ccWait.Wait()

			// request for result sum
			conn, err := net.Dial("tcp", addr)
			if err != nil {
				panic(err)
			}
			_, _ = fmt.Fprintf(conn, "0\n")
			response, _ := bufio.NewReader(conn).ReadString('\n')
			total := strings.TrimSpace(response)

			fmt.Printf("Concurrency: %d\n", concurrent)
			fmt.Printf("Result:      %s\n", total)
			fmt.Printf("\n")
		},
	}

	runStandardCmd.PersistentFlags().StringVar(&addr, "addr", addr, "TCP server address")
	runStandardCmd.PersistentFlags().IntVar(&concurrent, "concurrent", concurrent, "Concurrent operations count")

	utilCmd.AddCommand(runStandardCmd)
}
