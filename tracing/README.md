# Instruction

## Development

__elasticsearch__
```
[vagrant@localhost tracing]$ docker-compose up -d elasticsearch
Creating network "tracing_default" with the default driver
Creating tracing_elasticsearch_1 ... 
Creating tracing_elasticsearch_1 ... done
```

### Collector

Run
```
[vagrant@localhost tracing]$ tracing jaeger-collector --dependency-storage.type=elasticsearch --span-storage.type=elasticsearch --es.server-urls=http://172.19.0.2:9200 --es.username=elastic --es.password=changeme
{"level":"info","ts":1507443550.7656195,"caller":"healthcheck/handler.go:46","msg":"Health Check server started","http-port":14269}
{"level":"info","ts":1507443551.058716,"caller":"server/collector.go:109","msg":"Starting Jaeger Collector HTTP server","http-port":14268}
{"level":"info","ts":1507443551.0588026,"caller":"healthcheck/handler.go:88","msg":"Health Check state change","http-status":204}
Start gRPC on host [::]:12355
So gRPC is running
Start gRPC Gateway into host :12355
http on host: [::]:12356
```

### Agent

Run
```
[vagrant@localhost tracing]$ tracing jaeger-agent --collector.host-port=localhost:14267
{"level":"info","ts":1507444888.708636,"caller":"tchannel/bulider.go:89","msg":"Enabling service discovery","service":"jaeger-collector"}
{"level":"info","ts":1507444888.7087512,"caller":"peerlistmgr/peer_list_mgr.go:111","msg":"Registering active peer","peer":"localhost:14267"}
{"level":"info","ts":1507444888.710328,"caller":"server/agent.go:52","msg":"Starting agent"}
Start gRPC on host [::]:12365
So gRPC is running
Start gRPC Gateway into host :12365
http on host: [::]:12366
{"level":"info","ts":1507444889.710737,"caller":"peerlistmgr/peer_list_mgr.go:159","msg":"Not enough connected peers","connected":0,"required":1}
{"level":"info","ts":1507444889.7118149,"caller":"peerlistmgr/peer_list_mgr.go:166","msg":"Trying to connect to peer","host:port":"localhost:14267"}
{"level":"info","ts":1507444889.7151206,"caller":"peerlistmgr/peer_list_mgr.go:176","msg":"Connected to peer","host:port":"[::]:14267"}
```

### Client

Java: `git clone --depth=1 https://github.com/jaegertracing/jaeger-client-java jaegertracing/jaeger-client-java` or
```
fanhonglingdeMacBook-Pro:jaeger-client-java fanhongling$ git pull
remote: Counting objects: 35, done.
remote: Compressing objects: 100% (18/18), done.
remote: Total 35 (delta 17), reused 27 (delta 10), pack-reused 0
Unpacking objects: 100% (35/35), done.
From https://github.com/uber/jaeger-client-java
   91108a8..c1241e9  master     -> origin/master
Updating 91108a8..c1241e9
Fast-forward
 .travis.yml                                               |  8 ++++----
 CHANGELOG.rst                                             |  6 ++++++
 README.md                                                 |  8 ++++----
 jaeger-core/src/main/java/com/uber/jaeger/Tracer.java     |  2 +-
 jaeger-core/src/test/java/com/uber/jaeger/TracerTest.java | 14 ++++++++++++++
 jaeger-thrift/build.gradle                                |  8 ++------
 6 files changed, 31 insertions(+), 15 deletions(-)
```

IDL
```
fanhonglingdeMacBook-Pro:jaeger-client-java fanhongling$ git submodule init
Submodule 'idl' (https://github.com/uber/jaeger-idl.git) registered for path 'idl'
fanhonglingdeMacBook-Pro:jaeger-client-java fanhongling$ git submodule sync
Synchronizing submodule url for 'idl'
fanhonglingdeMacBook-Pro:idl fanhongling$ git submodule update
Cloning into 'idl'...
remote: Counting objects: 235, done.
remote: Total 235 (delta 0), reused 0 (delta 0), pack-reused 235
Receiving objects: 100% (235/235), 36.40 KiB | 0 bytes/s, done.
Resolving deltas: 100% (137/137), done.
Checking connectivity... done.
Submodule path '../idl': checked out 'c5adbc98341c5228472af96ea612c54173a3ec2e'
```

Build: Note that you need to install thrift 0.9.2 on your system for this task to work.
```
fanhonglingdeMacBook-Pro:jaeger-client-java fanhongling$ ./gradlew build
Downloading https://services.gradle.org/distributions/gradle-4.0-bin.zip
... snip ...
BUILD SUCCESSFUL in 8m 47s
119 actionable tasks: 92 executed, 27 up-to-date
```

Packages
```
[vagrant@localhost jaeger-client-java]$ ls jaeger-core/build/libs/
jaeger-core-0.21.0-SNAPSHOT.jar          jaeger-core-0.21.0-SNAPSHOT-okhttp381.jar
jaeger-core-0.21.0-SNAPSHOT-javadoc.jar  jaeger-core-0.21.0-SNAPSHOT-sources.jar
[vagrant@localhost jaeger-client-java]$ ls jaeger-thrift/build/libs/
jaeger-thrift-0.21.0-SNAPSHOT.jar          jaeger-thrift-0.21.0-SNAPSHOT-sources.jar
jaeger-thrift-0.21.0-SNAPSHOT-javadoc.jar  jaeger-thrift-0.21.0-SNAPSHOT-thrift92.jar
```

### Thrift

Refer to http://thrift.apache.org/docs/install/centos

`git clone --depth=1 -b 0.9.2 https://github.com/apache/thrift ...`

Install bison

Install _flex_ (Fedora 23)
```
[vagrant@localhost thrift]$ sudo dnf -y flex flex-devel
```

Configure (without no language libs)
```
[vagrant@localhost thrift]$ ./configure --enable-libs=no
checking for a BSD-compatible install... /bin/install -c
checking whether build environment is sane... yes
checking for a thread-safe mkdir -p... /bin/mkdir -p
checking for gawk... gawk
checking whether make sets $(MAKE)... yes
checking whether make supports nested variables... yes
checking whether UID '1000' is supported by ustar format... yes
checking whether GID '1000' is supported by ustar format... yes
checking how to create a ustar tar archive... gnutar
checking for pkg-config... /bin/pkg-config
checking pkg-config is at least version 0.9.0... yes
checking for gcc... gcc
checking whether the C compiler works... yes
checking for C compiler default output file name... a.out
checking for suffix of executables... 
checking whether we are cross compiling... no
checking for suffix of object files... o
checking whether we are using the GNU C compiler... yes
checking whether gcc accepts -g... yes
checking for gcc option to accept ISO C89... none needed
checking whether gcc understands -c and -o together... yes
checking for style of include used by make... GNU
checking dependency style of gcc... gcc3
checking how to run the C preprocessor... gcc -E
checking for g++... g++
checking whether we are using the GNU C++ compiler... yes
checking whether g++ accepts -g... yes
checking dependency style of g++... gcc3
checking build system type... x86_64-unknown-linux-gnu
checking host system type... x86_64-unknown-linux-gnu
checking how to print strings... printf
checking for a sed that does not truncate output... /bin/sed
checking for grep that handles long lines and -e... /bin/grep
checking for egrep... /bin/grep -E
checking for fgrep... /bin/grep -F
checking for ld used by gcc... /bin/ld
checking if the linker (/bin/ld) is GNU ld... yes
checking for BSD- or MS-compatible name lister (nm)... /bin/nm -B
checking the name lister (/bin/nm -B) interface... BSD nm
checking whether ln -s works... yes
checking the maximum length of command line arguments... 1572864
checking how to convert x86_64-unknown-linux-gnu file names to x86_64-unknown-linux-gnu format... func_convert_file_noop
checking how to convert x86_64-unknown-linux-gnu file names to toolchain format... func_convert_file_noop
checking for /bin/ld option to reload object files... -r
checking for objdump... objdump
checking how to recognize dependent libraries... pass_all
checking for dlltool... no
checking how to associate runtime and link libraries... printf %s\n
checking for ar... ar
checking for archiver @FILE support... @
checking for strip... strip
checking for ranlib... ranlib
checking command to parse /bin/nm -B output from gcc object... ok
checking for sysroot... no
checking for a working dd... /bin/dd
checking how to truncate binary pipes... /bin/dd bs=4096 count=1
checking for mt... no
checking if : is a manifest tool... no
checking for ANSI C header files... yes
checking for sys/types.h... yes
checking for sys/stat.h... yes
checking for stdlib.h... yes
checking for string.h... yes
checking for memory.h... yes
checking for strings.h... yes
checking for inttypes.h... yes
checking for stdint.h... yes
checking for unistd.h... yes
checking for dlfcn.h... yes
checking for objdir... .libs
checking if gcc supports -fno-rtti -fno-exceptions... no
checking for gcc option to produce PIC... -fPIC -DPIC
checking if gcc PIC flag -fPIC -DPIC works... yes
checking if gcc static flag -static works... yes
checking if gcc supports -c -o file.o... yes
checking if gcc supports -c -o file.o... (cached) yes
checking whether the gcc linker (/bin/ld -m elf_x86_64) supports shared libraries... yes
checking whether -lc should be explicitly linked in... no
checking dynamic linker characteristics... GNU/Linux ld.so
checking how to hardcode library paths into programs... immediate
checking whether stripping libraries is possible... yes
checking if libtool supports shared libraries... yes
checking whether to build shared libraries... yes
checking whether to build static libraries... yes
checking how to run the C++ preprocessor... g++ -E
checking for ld used by g++... /bin/ld -m elf_x86_64
checking if the linker (/bin/ld -m elf_x86_64) is GNU ld... yes
checking whether the g++ linker (/bin/ld -m elf_x86_64) supports shared libraries... yes
checking for g++ option to produce PIC... -fPIC -DPIC
checking if g++ PIC flag -fPIC -DPIC works... yes
checking if g++ static flag -static works... no
checking if g++ supports -c -o file.o... yes
checking if g++ supports -c -o file.o... (cached) yes
checking whether the g++ linker (/bin/ld -m elf_x86_64) supports shared libraries... yes
checking dynamic linker characteristics... (cached) GNU/Linux ld.so
checking how to hardcode library paths into programs... immediate
checking whether make sets $(MAKE)... (cached) yes
checking for bison... yes
checking for bison version >= 2.5... yes
checking for bison... bison -y
checking for flex... flex
checking lex output file root... lex.yy
checking lex library... -lfl
checking whether yytext is a pointer... yes
checking whether ln -s works... yes
checking for gawk... (cached) gawk
checking for ranlib... (cached) ranlib
checking whether g++ supports C++11 features by default... no
checking whether g++ supports C++11 features with -std=c++11... yes
checking for phpunit... no
checking for library containing strerror... none required
checking for an ANSI C-conforming const... yes
checking for inline... inline
checking for working volatile... yes
checking for stdbool.h that conforms to C99... no
checking for _Bool... no
checking for ANSI C header files... (cached) yes
checking whether time.h and sys/time.h may both be included... yes
checking for sys/wait.h that is POSIX.1 compatible... yes
checking return type of signal handlers... void
checking arpa/inet.h usability... yes
checking arpa/inet.h presence... yes
checking for arpa/inet.h... yes
checking sys/param.h usability... yes
checking sys/param.h presence... yes
checking for sys/param.h... yes
checking fcntl.h usability... yes
checking fcntl.h presence... yes
checking for fcntl.h... yes
checking for inttypes.h... (cached) yes
checking limits.h usability... yes
checking limits.h presence... yes
checking for limits.h... yes
checking netdb.h usability... yes
checking netdb.h presence... yes
checking for netdb.h... yes
checking netinet/in.h usability... yes
checking netinet/in.h presence... yes
checking for netinet/in.h... yes
checking pthread.h usability... yes
checking pthread.h presence... yes
checking for pthread.h... yes
checking stddef.h usability... yes
checking stddef.h presence... yes
checking for stddef.h... yes
checking for stdlib.h... (cached) yes
checking sys/socket.h usability... yes
checking sys/socket.h presence... yes
checking for sys/socket.h... yes
checking sys/time.h usability... yes
checking sys/time.h presence... yes
checking for sys/time.h... yes
checking sys/un.h usability... yes
checking sys/un.h presence... yes
checking for sys/un.h... yes
checking sys/poll.h usability... yes
checking sys/poll.h presence... yes
checking for sys/poll.h... yes
checking sys/resource.h usability... yes
checking sys/resource.h presence... yes
checking for sys/resource.h... yes
checking for unistd.h... (cached) yes
checking libintl.h usability... yes
checking libintl.h presence... yes
checking for libintl.h... yes
checking malloc.h usability... yes
checking malloc.h presence... yes
checking for malloc.h... yes
checking openssl/ssl.h usability... yes
checking openssl/ssl.h presence... yes
checking for openssl/ssl.h... yes
checking openssl/rand.h usability... yes
checking openssl/rand.h presence... yes
checking for openssl/rand.h... yes
checking openssl/x509v3.h usability... yes
checking openssl/x509v3.h presence... yes
checking for openssl/x509v3.h... yes
checking sched.h usability... yes
checking sched.h presence... yes
checking for sched.h... yes
checking wchar.h usability... yes
checking wchar.h presence... yes
checking for wchar.h... yes
checking for pthread_create in -lpthread... yes
checking for clock_gettime in -lrt... yes
checking for setsockopt in -lsocket... no
checking for int16_t... yes
checking for int32_t... yes
checking for int64_t... yes
checking for int8_t... yes
checking for mode_t... yes
checking for off_t... yes
checking for size_t... yes
checking for ssize_t... yes
checking for uint16_t... yes
checking for uint32_t... yes
checking for uint64_t... yes
checking for uint8_t... yes
checking for ptrdiff_t... yes
checking whether struct tm is in sys/time.h or time.h... time.h
checking whether AI_ADDRCONFIG is declared... yes
checking for working alloca.h... yes
checking for alloca... yes
checking for pid_t... yes
checking vfork.h usability... no
checking vfork.h presence... no
checking for vfork.h... no
checking for fork... yes
checking for vfork... yes
checking for working fork... yes
checking for working vfork... (cached) yes
checking for stdlib.h... (cached) yes
checking for GNU libc compatible malloc... yes
checking for working memcmp... yes
checking for stdlib.h... (cached) yes
checking for GNU libc compatible realloc... yes
checking sys/select.h usability... yes
checking sys/select.h presence... yes
checking for sys/select.h... yes
checking for sys/socket.h... (cached) yes
checking types of arguments for select... int,fd_set *,struct timeval *
checking whether lstat correctly handles trailing slash... yes
checking whether stat accepts an empty string... no
checking whether strerror_r is declared... yes
checking for strerror_r... yes
checking whether strerror_r returns char *... yes
checking for strftime... yes
checking for vprintf... yes
checking for _doprnt... no
checking for strtoul... yes
checking for bzero... yes
checking for ftruncate... yes
checking for gethostbyname... yes
checking for gethostbyname_r... yes
checking for gettimeofday... yes
checking for memmove... yes
checking for memset... yes
checking for mkdir... yes
checking for realpath... yes
checking for select... yes
checking for setlocale... yes
checking for socket... yes
checking for strchr... yes
checking for strdup... yes
checking for strerror... yes
checking for strstr... yes
checking for strtol... yes
checking for sqrt... yes
checking for alarm... yes
checking for clock_gettime... yes
checking for sched_get_priority_min... yes
checking for sched_get_priority_max... yes
checking for inet_ntoa... yes
checking for pow... yes
checking the behavior of a signed right shift... arithmetic
checking that generated files are newer than configure... done
configure: creating ./config.status
config.status: creating Makefile
config.status: creating compiler/cpp/Makefile
config.status: creating compiler/cpp/version.h
config.status: creating compiler/cpp/src/windows/version.h
config.status: creating lib/Makefile
...
config.status: creating test/Makefile
...
config.status: creating tutorial/Makefile
...
config.status: creating config.h
config.status: creating lib/cpp/src/thrift/config.h
config.status: executing depfiles commands
config.status: executing libtool commands

thrift 0.9.3

Building C++ Library ......... : no
Building C (GLib) Library .... : no
Building Java Library ........ : no
Building C# Library .......... : no
Building Python Library ...... : no
Building Ruby Library ........ : no
Building Haxe Library ........ : no
Building Haskell Library ..... : no
Building Perl Library ........ : no
Building PHP Library ......... : no
Building Erlang Library ...... : no
Building Go Library .......... : no
Building D Library ........... : no
Building NodeJS Library ...... : no
Building Lua Library ......... : no

If something is missing that you think should be present,
please skim the output of configure to find the missing
component.  Details are present in config.log.
```

Make
```
[vagrant@localhost thrift]$ make
make  all-recursive
make[1]: Entering directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift'
Making all in compiler/cpp
make[2]: Entering directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift/compiler/cpp'
make  all-am
make[3]: Entering directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift/compiler/cpp'
g++ -DHAVE_CONFIG_H -I. -I../.. -I../../lib/cpp/src/thrift  -I./src  -Wall -Wextra -pedantic -g -O2 -std=c++11 -MT src/thrift-main.o -MD -MP -MF src/.deps/thrift-main.Tpo -c -o src/thrift-main.o `test -f 'src/main.cc' || echo './'`src/main.cc
mv -f src/.deps/thrift-main.Tpo src/.deps/thrift-main.Po
...
libtool: link: g++ -Wall -Wextra -pedantic -g -O2 -std=c++11 -o thrift src/thrift-main.o src/thrift-md5.o src/generate/thrift-t_generator.o src/audit/thrift-t_audit.o src/parse/thrift-t_typedef.o src/parse/thrift-parse.o src/generate/thrift-t_c_glib_generator.o src/generate/thrift-t_cpp_generator.o src/generate/thrift-t_java_generator.o src/generate/thrift-t_json_generator.o src/generate/thrift-t_as3_generator.o src/generate/thrift-t_haxe_generator.o src/generate/thrift-t_csharp_generator.o src/generate/thrift-t_py_generator.o src/generate/thrift-t_rb_generator.o src/generate/thrift-t_perl_generator.o src/generate/thrift-t_php_generator.o src/generate/thrift-t_erl_generator.o src/generate/thrift-t_cocoa_generator.o src/generate/thrift-t_st_generator.o src/generate/thrift-t_ocaml_generator.o src/generate/thrift-t_hs_generator.o src/generate/thrift-t_xsd_generator.o src/generate/thrift-t_html_generator.o src/generate/thrift-t_js_generator.o src/generate/thrift-t_javame_generator.o src/generate/thrift-t_delphi_generator.o src/generate/thrift-t_go_generator.o src/generate/thrift-t_gv_generator.o src/generate/thrift-t_d_generator.o src/generate/thrift-t_lua_generator.o  -lfl libparse.a
make[3]: Leaving directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift/compiler/cpp'
make[2]: Leaving directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift/compiler/cpp'
Making all in lib
make[2]: Entering directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift/lib'
make[3]: Entering directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift/lib'
make[3]: Nothing to be done for 'all-am'.
make[3]: Leaving directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift/lib'
make[2]: Leaving directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift/lib'
Making all in test
make[2]: Entering directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift/test'
make[3]: Entering directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift/test'
make[3]: Nothing to be done for 'all-am'.
make[3]: Leaving directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift/test'
make[2]: Leaving directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift/test'
Making all in tutorial
make[2]: Entering directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift/tutorial'
make[3]: Entering directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift/tutorial'
../compiler/cpp/thrift --gen html -r ../tutorial/tutorial.thrift
make[3]: Leaving directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift/tutorial'
make[2]: Leaving directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift/tutorial'
make[2]: Entering directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift'
make[2]: Nothing to be done for 'all-am'.
make[2]: Leaving directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift'
make[1]: Leaving directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift'
```

Install
```
[vagrant@localhost thrift]$ sudo make install
Making install in compiler/cpp
make[1]: Entering directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift/compiler/cpp'
make  install-am
make[2]: Entering directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift/compiler/cpp'
make[3]: Entering directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift/compiler/cpp'
 /bin/mkdir -p '/usr/local/bin'
  /bin/sh ../../libtool   --mode=install /bin/install -c thrift '/usr/local/bin'
libtool: install: /bin/install -c thrift /usr/local/bin/thrift
make[3]: Nothing to be done for 'install-data-am'.
make[3]: Leaving directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift/compiler/cpp'
make[2]: Leaving directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift/compiler/cpp'
make[1]: Leaving directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift/compiler/cpp'
Making install in lib
make[1]: Entering directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift/lib'
make[2]: Entering directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift/lib'
make[3]: Entering directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift/lib'
make[3]: Nothing to be done for 'install-exec-am'.
make[3]: Nothing to be done for 'install-data-am'.
make[3]: Leaving directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift/lib'
make[2]: Leaving directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift/lib'
make[1]: Leaving directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift/lib'
Making install in test
make[1]: Entering directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift/test'
make[2]: Entering directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift/test'
make[3]: Entering directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift/test'
make[3]: Nothing to be done for 'install-exec-am'.
make[3]: Nothing to be done for 'install-data-am'.
make[3]: Leaving directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift/test'
make[2]: Leaving directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift/test'
make[1]: Leaving directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift/test'
Making install in tutorial
make[1]: Entering directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift/tutorial'
make[2]: Entering directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift/tutorial'
../compiler/cpp/thrift --gen html -r ../tutorial/tutorial.thrift
make[3]: Entering directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift/tutorial'
make[3]: Nothing to be done for 'install-exec-am'.
make[3]: Nothing to be done for 'install-data-am'.
make[3]: Leaving directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift/tutorial'
make[2]: Leaving directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift/tutorial'
make[1]: Leaving directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift/tutorial'
make[1]: Entering directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift'
make[2]: Entering directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift'
make[2]: Nothing to be done for 'install-exec-am'.
make[2]: Nothing to be done for 'install-data-am'.
make[2]: Leaving directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift'
make[1]: Leaving directory '/Users/fanhongling/Downloads/workspace/src/github.com/apache/thrift'
```

bin
```
[vagrant@localhost thrift]$ which thrift
/usr/local/bin/thrift
[vagrant@localhost thrift]$ thrift --version
Thrift version 0.9.3
[vagrant@localhost thrift]$ thrift --help
Usage: thrift [options] file
Options:
  -version    Print the compiler version
  -o dir      Set the output directory for gen-* packages
               (default: current directory)
  -out dir    Set the ouput location for generated files.
               (no gen-* folder will be created)
  -I dir      Add a directory to the list of directories
                searched for include directives
  -nowarn     Suppress all compiler warnings (BAD!)
  -strict     Strict compiler warnings on
  -v[erbose]  Verbose mode
  -r[ecurse]  Also generate included files
  -debug      Parse debug trace to stdout
  --allow-neg-keys  Allow negative field keys (Used to preserve protocol
                compatibility with older .thrift files)
  --allow-64bit-consts  Do not print warnings about using 64-bit constants
  --gen STR   Generate code with a dynamically-registered generator.
                STR has the form language[:key1=val1[,key2[,key3=val3]]].
                Keys and values are options passed to the generator.
                Many options will not require values.

Options related to audit operation
   --audit OldFile   Old Thrift file to be audited with 'file'
  -Iold dir    Add a directory to the list of directories
                searched for include directives for old thrift file
  -Inew dir    Add a directory to the list of directories
                searched for include directives for new thrift file

Available generators (and options):
  as3 (AS3):
    bindable:        Add [bindable] metadata to all the struct classes.
  c_glib (C, using GLib):
  cocoa (Cocoa):
    log_unexpected:  Log every time an unexpected field ID or type is encountered.
    validate_required:
                     Throws exception if any required field is not set.
    async_clients:   Generate clients which invoke asynchronously via block syntax.
  cpp (C++):
    cob_style:       Generate "Continuation OBject"-style classes.
    no_client_completion:
                     Omit calls to completion__() in CobClient class.
    no_default_operators:
                     Omits generation of default operators ==, != and <
    templates:       Generate templatized reader/writer methods.
    pure_enums:      Generate pure enums instead of wrapper classes.
    include_prefix:  Use full include paths in generated files.
    moveable_types:  Generate move constructors and assignment operators.
  csharp (C#):
    async:           Adds Async support using Task.Run.
    asyncctp:        Adds Async CTP support using TaskEx.Run.
    wcf:             Adds bindings for WCF to generated classes.
    serial:          Add serialization support to generated classes.
    nullable:        Use nullable types for properties.
    hashcode:        Generate a hashcode and equals implementation for classes.
    union:           Use new union typing, which includes a static read function for union types.
  d (D):
  delphi (delphi):
    ansistr_binary:  Use AnsiString for binary datatype (default is TBytes).
    register_types:  Enable TypeRegistry, allows for creation of struct, union
                     and container instances by interface or TypeInfo()
    constprefix:     Name TConstants classes after IDL to reduce ambiguities
    events:          Enable and use processing events in the generated code.
    xmldoc:          Enable XMLDoc comments for Help Insight etc.
  erl (Erlang):
    legacynames: Output files retain naming conventions of Thrift 0.9.1 and earlier.
    maps:        Generate maps instead of dicts.
    otp16:       Generate non-namespaced dict and set instead of dict:dict and sets:set.
  go (Go):
    package_prefix=  Package prefix for generated files.
    thrift_import=   Override thrift package import path (default:git.apache.org/thrift.git/lib/go/thrift)
    package=         Package name (default: inferred from thrift file name)
    ignore_initialisms
                     Disable automatic spelling correction of initialisms (e.g. "URL")
    read_write_private
                     Make read/write methods private, default is public Read/Write
  gv (Graphviz):
    exceptions:      Whether to draw arrows from functions to exception.
  haxe (Haxe):
    callbacks        Use onError()/onSuccess() callbacks for service methods (like AS3)
    rtti             Enable @:rtti for generated classes and interfaces
    buildmacro=my.macros.Class.method(args)
                     Add @:build macro calls to generated classes and interfaces
  hs (Haskell):
  html (HTML):
    standalone:      Self-contained mode, includes all CSS in the HTML files.
                     Generates no style.css file, but HTML files will be larger.
    noescape:        Do not escape html in doc text.
  java (Java):
    beans:           Members will be private, and setter methods will return void.
    private-members: Members will be private, but setter methods will return 'this' like usual.
    nocamel:         Do not use CamelCase field accessors with beans.
    fullcamel:       Convert underscored_accessor_or_service_names to camelCase.
    android:         Generated structures are Parcelable.
    android_legacy:  Do not use java.io.IOException(throwable) (available for Android 2.3 and above).
    option_type:     Wrap optional fields in an Option type.
    java5:           Generate Java 1.5 compliant code (includes android_legacy flag).
    reuse-objects:   Data objects will not be allocated, but existing instances will be used (read and write).
    sorted_containers:
                     Use TreeSet/TreeMap instead of HashSet/HashMap as a implementation of set/map.
    generated_annotations=[undated|suppress]:
                     undated: suppress the date at @Generated annotations
                     suppress: suppress @Generated annotations entirely
  javame (Java ME):
  js (Javascript):
    jquery:          Generate jQuery compatible code.
    node:            Generate node.js compatible code.
    ts:              Generate TypeScript definition files.
  json (JSON):
    merge:           Generate output with included files merged
  lua (Lua):
  ocaml (OCaml):
  perl (Perl):
  php (PHP):
    inlined:         Generate PHP inlined files
    server:          Generate PHP server stubs
    oop:             Generate PHP with object oriented subclasses
    rest:            Generate PHP REST processors
    nsglobal=NAME:   Set global namespace
    validate:        Generate PHP validator methods
    json:            Generate JsonSerializable classes (requires PHP >= 5.4)
  py (Python):
    new_style:       Generate new-style classes.
    twisted:         Generate Twisted-friendly RPC services.
    tornado:         Generate code for use with Tornado.
    utf8strings:     Encode/decode strings using utf8 in the generated code.
    coding=CODING:   Add file encoding declare in generated file.
    slots:           Generate code using slots for instance members.
    dynamic:         Generate dynamic code, less code generated but slower.
    dynbase=CLS      Derive generated classes from class CLS instead of TBase.
    dynexc=CLS       Derive generated exceptions from CLS instead of TExceptionBase.
    dynimport='from foo.bar import CLS'
                     Add an import line to generated code to find the dynbase class.
  rb (Ruby):
    rubygems:        Add a "require 'rubygems'" line to the top of each generated file.
    namespaced:      Generate files in idiomatic namespaced directories.
  st (Smalltalk):
  xsd (XSD):
```