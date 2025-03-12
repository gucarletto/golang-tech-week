package main

import (
	"context"
	"fmt"
	"gucarletto/golang-tech-week/pkg/workerpool"
	"math/rand"
	"sync"
	"time"
)

type NumeroJob struct {
	Numero int
}

type ResultadoNumero struct {
	Valor     int
	WorkerID  int
	Timestamp time.Time
}

func processarNumero(ctx context.Context, job workerpool.Job) workerpool.Result {
	numeroJob := job.(NumeroJob).Numero
	workerId := numeroJob % 3
	sleepTime := time.Duration(800+rand.Intn(400)) * time.Millisecond
	time.Sleep(sleepTime)
	return ResultadoNumero{
		Valor:     numeroJob,
		WorkerID:  workerId,
		Timestamp: time.Now(),
	}
}

func main() {
	valorMaximo := 20
	bufferSize := 10
	workerCount := 3
	wp := workerpool.New(processarNumero, workerpool.Config{
		WorkerCount: workerCount,
	})

	inputCh := make(chan workerpool.Job, bufferSize)
	ctx := context.Background()
	resultCh, err := wp.Start(ctx, inputCh)
	if err != nil {
		panic(err)
	}
	var wg sync.WaitGroup
	wg.Add(valorMaximo)

	fmt.Println("Iniciando processamento...")

	go func() {
		for i := 0; i < valorMaximo; i++ {
			inputCh <- NumeroJob{Numero: i}
		}
		close(inputCh)
	}()

	go func() {
		for result := range resultCh {
			r := result.(ResultadoNumero)
			fmt.Printf("Resultado: %d (worker %d) - %s\n", r.Valor, r.WorkerID, r.Timestamp.Format(time.RFC3339))
			wg.Done()
		}
	}()

	wg.Wait()
	fmt.Println("Processamento finalizado.")
}
