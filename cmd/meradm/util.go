package main

import "regexp"

// Simple regex to ensure we have something:port. We rely on merlin to perform proper validation.
var ipPortRegex = regexp.MustCompile(`^([^:]+):(\d+)$`)
