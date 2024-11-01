#!/bin/sh

/usr/local/bin/transac-parser-svc run service &
sleep 2
/usr/local/bin/transac-parser-svc migrate down &
sleep 2
/usr/local/bin/transac-parser-svc migrate up


wait