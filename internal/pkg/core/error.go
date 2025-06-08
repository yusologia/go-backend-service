package core

import (
	"errors"
	"fmt"
	"github.com/yusologia/go-core/v2/pkg"
	"github.com/yusologia/go-core/v2/response"
	"net/http"
	"os"
	"service/internal/pkg/grpc/example"
)

func ErrorHandler(fn func() error) error {
	errChan := make(chan error, 1)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				var err error
				if panicData, ok := r.(*logiares.ResponseError); ok {
					status := panicData.Status
					err = errors.New(fmt.Sprintf("Code: %d. Message: %s. InternalMsg: %s", status.Code, status.Message, status.InternalMsg))
				} else if panicData, ok := r.(error); ok {
					err = errors.New(fmt.Sprintf("Code: %d. Message: %v", http.StatusInternalServerError, panicData.Error()))
				} else {
					err = errors.New(fmt.Sprintf("Code: %d. Message: An error Occurred.", http.StatusInternalServerError))
				}

				fmt.Fprintf(os.Stderr, "panic: %v\n", r)
				logiapkg.LogError(r)

				errChan <- err
			}
		}()

		if err := fn(); err != nil {
			logiapkg.LogError(err)
			errChan <- err
		} else {
			close(errChan)
		}
	}()

	return <-errChan
}

func GRPCErrorHandler(fn func() (*example.EXResponse, error)) (res *example.EXResponse, err error) {
	resChan := make(chan *example.EXResponse)
	errChan := make(chan error)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				var err error
				if panicData, ok := r.(*logiares.ResponseError); ok {
					status := panicData.Status
					err = errors.New(fmt.Sprintf("Code: %d. Message: %s. InternalMsg: %s", status.Code, status.Message, status.InternalMsg))
				} else if panicData, ok := r.(error); ok {
					err = errors.New(fmt.Sprintf("Code: %d. Message: %v", http.StatusInternalServerError, panicData.Error()))
				} else {
					err = errors.New(fmt.Sprintf("Code: %d. Message: An error Occurred.", http.StatusInternalServerError))
				}

				fmt.Fprintf(os.Stderr, "panic: %v\n", r)
				logiapkg.LogError(r)

				errChan <- err
			}
		}()

		res, err := fn()
		if err != nil {
			logiapkg.LogError(err)

			errChan <- err
		} else {
			resChan <- res
		}
	}()

	select {
	case res := <-resChan:
		return res, nil
	case err := <-errChan:
		return nil, err
	}
}
