package nagioscfg

import (
	"fmt"
	"log"
	"testing"
)

func TestParseString_empty(t *testing.T) {
	config := ""
	result, err := ParseString(config)
	if err != nil || len(result) != 0 {
		t.Fail()
	}
}

func TestParseString_host(t *testing.T) {
	config := `define host{
	host_name			bogus-router
	alias				Bogus Router #1
	address				192.168.1.254
	parents				server-backbone
	check_command			check-host-alive
	check_interval			5
	retry_interval			1
	max_check_attempts		5
	check_period			24x7
	process_perf_data		0
	retain_nonstatus_information	0
	contact_groups			router-admins
	notification_interval		30
	notification_period		24x7
	notification_options		d,u,r
	}`

	results, err := ParseString(config)
	if err != nil {
		log.Printf("Got error: %s", err)
		t.Fail()
	}

	if len(results) != 1 {
		log.Printf("Expected 1 object, got %d", len(results))
		t.Fail()
	}

	fmt.Println(results)
}
