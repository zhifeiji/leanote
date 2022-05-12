#!/bin/bash

pids=`ps aux|grep leanote|grep -v grep|awk '{print $2}'`
echo '杀掉历史进程 ' $pids

for i in ${pids[@]}
do
        kill $i
done

cp -rf conf/app.conf bin/src/github.com/zhifeiji/leanote/conf/app.conf

echo '启动 leanote 服务'
nohup /bin/bash /root/leanote/bin/run.sh >> /var/log/leanote.log 2>&1 &
