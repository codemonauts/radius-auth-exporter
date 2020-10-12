package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
	"layeh.com/radius/rfc2868"
)

func tryRadiusAuth(username string, password string) (*radius.Packet, bool) {
	packet := radius.New(radius.CodeAccessRequest, []byte(cfg.Secret))
	rfc2865.UserName_SetString(packet, username)
	rfc2865.UserPassword_SetString(packet, password)
	address := fmt.Sprintf("%s:%d", cfg.Address, cfg.Port)

	ctx, done := context.WithTimeout(context.Background(), multiplyDuration(cfg.Timeout, time.Second))
	defer done()

	response, err := radius.Exchange(ctx, packet, address)
	if err != nil {
		log.Println(err)
		return nil, false
	}

	if response.Code == radius.CodeAccessAccept {
		return response, true
	}

	return nil, false
}

// Try to extract attribute 81 which is
// 'Tunnel-Private-Group-ID' according to RFC2868
func getVLANfromResponse(paket *radius.Packet) int {
	_, value := rfc2868.TunnelPrivateGroupID_Get(paket)
	i, err := strconv.Atoi(string(value))
	if err != nil {
		log.Printf("Got unparseable VLAN ID: %s", value)
		return 0
	}
	return i
}
