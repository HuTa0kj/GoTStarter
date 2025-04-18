# GoTStarter
GoTStarter is a scaffold for quickly building terminal tools. It is equipped with basic functions such as command line parameter parsing, log output, configuration file loading, and code file compilation.

The project is based on Go 1.22

```bash
git clone git@github.com:HuTa0kj/GoTStarter.git
cd GoTStarter
```

You can execute the initialization script to replace the package name you currently need to develop.

```bash
chmod 777 ./rename.sh
./rename.sh
```

The first time you execute Origin Module, you need to fill in gotstarter.

```
Origin Module Name（Such as gotstarter）: gotstarter
New Module Name（Such as myproject）: <Your Module Name>
```

You can test it by executing the command.

```bash
go run ./cmd/gotstarter/main.go -i input-test -d
```

Compile into a binary file.

```
./build.sh
```

The same effect as you can see.

```
./gotstarter-linux -i input-test -d
```

```

   ______    ___________ __             __           
  / ____/___/_  __/ ___// /_____ ______/ /____  _____
 / / __/ __ \/ /  \__ \/ __/ __ */ ___/ __/ _ \/ ___/
/ /_/ / /_/ / /  ___/ / /_/ /_/ / /  / /_/  __/ /
\____/\____/_/  /____/\__/\__,_/_/   \__/\___/_/     0.0.1

[INF] Debug mode enabled
[INF] Read Config Successful
[DBG] Initialization completed
[INF] input-test
```

You can generate your favorite banner icon, tool name, and more.
