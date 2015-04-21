#!/usr/local/bin/python
from jnpr.junos import Device
import getpass

hostname = raw_input("Please enter the host you would like to connect to: ")
print "The host has now been set to: %s " % hostname
#Get username and password
user = raw_input("\n\nEnter your username: ")
password = getpass.getpass(prompt="Enter your password: ")

dev = Device(user=user, password=password, host=hostname)

#Connect to device
try:
	dev.open()
except:
	print "You've gone and broken something!"
#Collect BGP summary info from device
bgpsum = dev.rpc.get_bgp_summary_information()

#Print total number of routes in RIB and also the active routes in FIB
print(dev.hostname)
print 35*"="
print "Total Routing Table Entries: ", bgpsum.findtext('bgp-rib/total-prefix-count'), "\nTotal Active Entries: ", bgpsum.findtext('bgp-rib/active-prefix-count'), "\n"

dev.close()