#!/usr/bin/python3.8

import cgi
import cgitb
cgitb.enable()

print("Content-Type: text/html\n")
print()

form = cgi.FieldStorage()

for key in form:
    value = form[key].value
    print('<p>%s: %s</p>' % (key, value))
