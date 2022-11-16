# Tutorial

### Install command line tools

```shell
~ go install github.com/fengya90/winter-aop/wacli@latest
~ wacli
please run: wacli -h
```

### How to use wacli

#### genarate code

```
wacli gen -d /myhome/mydir1 -d /myhome/mydir2
or
wacli gen -f /myhome/myconfig
```

The configuration file can be JSON, TOML, YAML, HCL and Java properties config files. For example:

```
{
    "source_dir_list":["mydir1","mydir2"]
}
```

#### clean the genarated code

```
wacli clean -d /myhome/mydir1 -d /myhome/mydir2
or
wacli clean -f /myhome/myconfig
```


### Example

refer to the example in this library.

```
~ tree tutorial
tutorial
├── dep
│   ├── dep1.go
│   ├── dep2.go
│   └── init.go
└── main.go
~ wacli gen -d tutorial
generate tutorial/dep/dep1_winter_aop_gen.go
generate tutorial/dep/dep2_winter_aop_gen.go
~ tree tutorial
tutorial
├── dep
│   ├── dep1.go
│   ├── dep1_winter_aop_gen.go
│   ├── dep2.go
│   ├── dep2_winter_aop_gen.go
│   └── init.go
└── main.go
~ cd tutorial
~ go build -tags=winter_aop
~ ./tutorial

highPriorityAround start:addThreeNumber
monitor start:addThreeNumber
monitor finish:addThreeNumber
cost:0
parameters: 2	3	4
result: 9
highPriorityAround finish:addThreeNumber
9
highPriorityAround start:sayAny
monitor start:sayAny
hehe
monitor finish:sayAny
cost:0
parameters: hehe
result:
highPriorityAround finish:sayAny
hello
```

If you want to test the business code, and don't care about 'around functions', you can compile the code normally. Of course, in this case, the command tool wacli is not required.

```
~ wacli clean -d tutorial
~ tree tutorial
tutorial
├── dep
│   ├── dep1.go
│   ├── dep2.go
│   └── init.go
├── main.go
~ cd tutorial
~ go build
~ ./tutorial

9
hehe
hello
```
