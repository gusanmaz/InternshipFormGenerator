package sireader

import "github.com/gusanmaz/InternshipFormGenerator2/globals"

var g globals.Config

func init(){
	g = globals.Globals
}

const(
	Online = iota
	OnlineVerifiable
	Onsite
)


