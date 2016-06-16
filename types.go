package main

import (
	"encoding/json"
)

type (
	Repo struct {
		FullName string
	}

	Build struct {
		Link string
	}

	Netrc struct {
		Login string
	}

	Config struct {
		Server       string
		Token        string
		Threshold    float64
		Include      string
		MustIncrease bool
		CACert       string
	}
)
