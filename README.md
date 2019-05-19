# go-serial-comm-sample
Example of implementation of a program for reading and writing data for serial port

Thios sample use go.bug.st/serial.v1   library.

Configuration:

Install go.bug.st/serial.v1  :

go get go.bug.st/serial.v1


For tests with virtual serial port (linux) install socat

You can create a virtual read and write port with the following command:

socat -d -d pty, raw, echo = 0, b115200 pty, raw, echo = 0, b115200


To read in the virtual port use the command:

cat < port

example:
cat </dev/pts/1

To write to a virtual port:

echo "test ......"> port

example:

echo "test ......"> /dev/pts/2

