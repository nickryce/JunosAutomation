#!/usr/bin/env python
from pprint import pprint
from jnpr.junos import Device
from jnpr.junos.op.routes import RouteTable


dev = Device(host='46.226.4.24', user='nick', password='bn6VDSyn' )
dev.open()

tbl = RouteTable(dev)
tbl.get('0.0.0.0/0', protocol='aggregate')
print tbl
for item in tbl:
    print 'protocol:', item.protocol
    print 'age:', item.age
    print 'via:', item.via
    print

dev.close()
