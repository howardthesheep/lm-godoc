# lm-godoc
lm-godoc is a documentation automation tool. It traverses a Golang project directory and parses 
high level code information such as: function names, structures, and comments. It then renders this 
information into an HTML template. This HTML template is then served from an http server to the user's
web browser. 

## What is the purpose of lm-godoc
lm-godoc is meant to increase the availability of high level project documentation. This will increase
efficiency for both new and seasoned programmers. This takes the pain out of searching the code base for
code documentation.