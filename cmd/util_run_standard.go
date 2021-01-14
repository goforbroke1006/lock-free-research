package cmd

import (
	"fmt"
	"sync"
	"time"

	"github.com/spf13/cobra"

	"github.com/goforbroke1006/lock-free-research/internal"
	"github.com/goforbroke1006/lock-free-research/internal/number_counter"
)

func init() {
	var (
		samplesFilename string
		concurrent      = 100
	)

	var runStandardCmd = &cobra.Command{
		Use:   "run-standard",
		Short: "TODO: write me",
		Long:  `TODO: write me`,
		Run: func(cmd *cobra.Command, args []string) {
			samples, err := internal.ReadSamples(samplesFilename)
			if err != nil {
				panic(err)
			}

			nc := number_counter.NewMutexBasedNumberCounter()
			startSpan := time.Now()
			ccWait := sync.WaitGroup{}

			for i := 0; i < concurrent; i++ {
				ccWait.Add(1)
				go func() {
					defer ccWait.Done()
					for _, n := range samples {
						if err := nc.Add(n); err != nil {
							panic(err)
						}
					}
				}()
			}

			ccWait.Wait()
			nanos := time.Since(startSpan).Nanoseconds()
			total, err := nc.Get()
			if err != nil {
				panic(err)
			}

			fmt.Printf("Type:        WITH MUTEX\n")
			fmt.Printf("Samples:     %d\n", len(samples))
			fmt.Printf("Concurrency: %d\n", concurrent)
			fmt.Printf("Result:      %d\n", total)
			fmt.Printf("Duration:    %d nanoseconds\n", nanos)
			fmt.Printf("\n")
		},
	}

	runStandardCmd.PersistentFlags().StringVar(&samplesFilename, "filename", samplesFilename, "File with number location")
	runStandardCmd.PersistentFlags().IntVar(&concurrent, "concurrent", concurrent, "Concurrent operations count")

	utilCmd.AddCommand(runStandardCmd)
}
