#!/usr/local/bin/python
from jnpr.junos import Device
import getpass


#Get username and password to log into the device
user = raw_input("\n\nEnter your username: ")
password = getpass.getpass(prompt="Enter your password: ")

#Get router hostnames from a list 
hosts = [line.strip() for line in open("/Users/nick/JunosAutomation/Python/routers.txt",  'r')]

#Loop through the list of routers and connect then print the hostname and serial number
for router in hosts:
	dev = Device(user=user, password=password, host=router)

	#Connect to device
	try:
			dev.open()
	except:
    		print "You've gone and broken something!"

    	print(dev.facts['hostname']) + " :  " + (dev.facts['serialnumber'])


    	dev.close()

