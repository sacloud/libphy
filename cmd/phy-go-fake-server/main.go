// Copyright 2021 The phy-go authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/sacloud/phy-go"

	"github.com/sacloud/phy-go/fake"
	"github.com/sacloud/phy-go/fake/server"
	"github.com/spf13/cobra"
)

var listenAddr string

var cmd = &cobra.Command{
	Use:     "phy-go-fake-server",
	Short:   "Start the web server",
	RunE:    run,
	Version: phy.Version,
}

func init() {
	cmd.Flags().StringVarP(&listenAddr, "addr", "", ":8080", "the address for the server to listen on")
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := cmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	errCh := make(chan error)

	go func() {
		errCh <- startServer(listenAddr)
	}()

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		fmt.Println("shutting down") // nolint
	}
	return ctx.Err()
}

func startServer(addr string) error {
	fakeServer := server.Server{
		Engine: &fake.Engine{},
	}
	httpServer := &http.Server{
		Handler: fakeServer.Handler(),
	}

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer listener.Close()

	return httpServer.Serve(listener)
}
