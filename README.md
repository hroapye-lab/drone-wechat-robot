## drone-wechat-robot

The WeChat for Work Robot notifications for Drone plugin. You can reference [the doc](https://work.weixin.qq.com/help?doc_id=13376) to know how to use it.

## Build

```shell
$ sh build.sh
```

## Environment Reference

You can look at drone pipelines [doc](https://docs.drone.io/pipeline/environment/reference/)

## Usage

```yml
kind: pipeline
type: docker
name: default

steps:
- name: notify
  image: kaynewang/drone-wechat-robot
  settings:
    msgtype: text
    key: your own robot key
    content: "Build status: {{ .CI_BUILD_STATUS }}, Commit: {{ .DRONE_COMMIT }}"
    mentioned_list: @all,kaynewang
    mentioned_mobile_list: @all,kaynewang
```

```yml
kind: pipeline
type: docker
name: default

steps:
- name: notify
  image: kaynewang/drone-wechat-robot
  settings:
    msgtype: markdown
    key: your own robot key
    content: "**Build status**: {{ .CI_BUILD_STATUS }}, **Commit**: {{ .DRONE_COMMIT }}"
```