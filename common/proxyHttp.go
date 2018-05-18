package common

import (
	"net/http"
	"fmt"
	"golang.org/x/net/proxy"
	"os"
)

func ProxyHttpGet(mediaUrl string) (*http.Response, error) {
	dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:1080", nil, proxy.Direct)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't connect to the proxy:", err)
		os.Exit(1)
	}
	// setup a http client
	httpTransport := &http.Transport{}
	httpClient := &http.Client{Transport: httpTransport}
	// set our socks5 as the dialer
	httpTransport.Dial = dialer.Dial
	resp, err := httpClient.Get(mediaUrl)
	//defer resp.Body.Close()
	fmt.Println(err)
	return resp, err
}
