#!/bin/sh


/usr/local/bin/transac-parser-svc migrate down &
sleep 2
/usr/local/bin/transac-parser-svc migrate up&
sleep 2
/usr/local/bin/transac-parser-svc run service 




wait