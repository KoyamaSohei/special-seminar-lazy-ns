package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/miekg/dns"
	"go.uber.org/zap"
)

var logger *zap.Logger

func serve() {
	i := os.Getenv("TARGET_IP")
	logger.Info("TARGET_IP is " + i)
	dns.HandleFunc(".", func(w dns.ResponseWriter, r *dns.Msg) {
		m := new(dns.Msg)
		m.SetReply(r)
		for k, q := range r.Question {
			if q.Qtype == dns.TypeA {
				n := r.Question[k].Name
				logger.Info("Q. " + n)
				rr := &dns.A{
					Hdr: dns.RR_Header{Name: r.Question[k].Name, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 3600},
					A:   net.ParseIP(i),
				}
				m.Answer = append(m.Answer, rr)
			}
		}

		w.WriteMsg(m)
	})

	go func() {
		server := &dns.Server{Addr: "[::]:53", Net: "udp", TsigSecret: nil, ReusePort: false}
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
}

func main() {
	ops := zap.NewProductionConfig()
	ops.OutputPaths = []string{"stdout"}
	logger, _ = ops.Build()
	serve()
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	s := <-sig
	log.Printf("Signal (%s) received, stopping\n", s)
}