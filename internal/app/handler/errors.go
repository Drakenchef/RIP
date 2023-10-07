package handler

import "errors"

var (
	idNotFound                = errors.New("param `id` not found")
	idMustBeEmpty             = errors.New("param `id` must be empty")
	planetCannotBeEmpty       = errors.New("planet name cannot be empty")
	headerNotFound            = errors.New("no file uploaded")
	fridOrPlanetIsEmpty       = errors.New("flight request or planet cannot be empty")
	flightNumberCannotBeEmpty = errors.New("param `flight_number` cannot be empty")
)
