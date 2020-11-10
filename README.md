# GoFINS

[![Build Status](https://travis-ci.org/l1va/gofins.svg?branch=master)](https://travis-ci.org/l1va/gofins)

This is fins command client written by Go.

The library support communication to omron PLC from Go application.

Ideas were taken from https://github.com/hiroeorz/omron-fins-go and https://github.com/patrick--/node-omron-fins

Library was tested with <b>Omron PLC NJ501-1300</b>. Mean time of the cycle request-response is 4ms.
Additional work in the siyka-au repository was tested against a <b>CP1L-EM</b>.

There is simple Omron FINS Server (PLC emulator) in the fins/server.go 

 ### Thanks
github.com/l1va/gofins/fins

Smartwang PS:
I have tested with Omron PLC CP1H + CP1W-CIF41 (Net Module) 
