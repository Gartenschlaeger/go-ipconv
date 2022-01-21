package main

import "regexp"

var regexp_ipv4 = regexp.MustCompile(`(?m)^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}$`)
var regexp_ipv4cidr = regexp.MustCompile(`(?m)^(\d{1,3})(\.\d{1,3})(\.\d{1,3})?(\.\d{1,3})?(\/\d{1,2})$`)
