package main

import "errors"

var ErrForbiddenTemplateAndColumnFlags = errors.New("using both flags template and columns is forbidden")
