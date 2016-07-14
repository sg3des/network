#NETWORK - golang udp wrapper over *net* package 

Support two operating methods:
	
1. network return connection and use this connection for send messages

	```
		conn,err := network.Client(":12345")
		...

		conn.Reply([]bytes)
	```

2. internal variable contains current connection, and use it for send messages

	```
		_,err := network.Client(":12345")
		...

		network.Reply([]bytes)
	```

internal variable(second method) allow send reply from different packages